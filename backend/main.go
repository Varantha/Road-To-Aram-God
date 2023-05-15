package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

type RiotGamesAPI struct {
	BaseURL string
	APIKey  string
}

type CustomHandler func(api *RiotGamesAPI, c *gin.Context)

type APIRoute struct {
	Path    string
	Method  string
	Handler CustomHandler
}

type inputChallengeCategory struct {
	CategoryName   string
	ChallengeName  string
	ChallengeNames []string
}

// StringInSlice checks if `list` contains `s`
func StringInSlice(s string, list []string) bool {
	for _, x := range list {
		if s == x {
			return true
		}
	}
	return false
}

func LoadAuth() {
	err := godotenv.Load("local.env")
	if err != nil {
		log.Printf("Unable to load local.env: %s", err)
	}
}

func main() {
	LoadAuth()

	apiKey := os.Getenv("API_KEY")
	if len(apiKey) == 0 {
		log.Fatal("No API Key Env variable set, unable to start.")
	}

	api := &RiotGamesAPI{
		BaseURL: "https://%s.api.riotgames.com",
		APIKey:  apiKey,
	}

	routes := []APIRoute{
		{
			Path:    "getPlayerChallenges/:summonerName",
			Method:  "GET",
			Handler: getPlayerChallenges,
		},
		{
			Path:    "getChallengeInfo",
			Method:  "GET",
			Handler: getChallengeInfo,
		},
		{
			Path:    "getCombinedChallengeInfo/:summonerName",
			Method:  "GET",
			Handler: getCombinedChallengeInfo,
		},
		// Add more routes here...
	}

	router := gin.Default()

	routePrefix := "/:region/"
	for _, route := range routes {
		switch route.Method {
		case "GET":
			router.GET(routePrefix+route.Path, routeHandlerWrapper(api, route.Handler))
			// Add more cases for other HTTP methods if needed
		}
	}

	router.Run("0.0.0.0:8080")
}

func routeHandlerWrapper(api *RiotGamesAPI, handler CustomHandler) gin.HandlerFunc {
	return func(c *gin.Context) {
		handler(api, c)
	}
}

func (api *RiotGamesAPI) ritoGet(endpoint string, region string) ([]byte, error) {
	regionedBaseUrl := fmt.Sprintf(api.BaseURL, region)
	url := fmt.Sprintf("%s%s?api_key=%s", regionedBaseUrl, endpoint, api.APIKey)

	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("error %d: %s", resp.StatusCode, string(body))
	}

	return body, nil
}

// getPlayerChallenges
func getPlayerChallenges(api *RiotGamesAPI, c *gin.Context) {
	summonerName := c.Param("summonerName")
	region := c.Param("region")
	challengeData, err := api.getPlayerChallengeData(summonerName, region)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, challengeData)
}

func (api *RiotGamesAPI) getPlayerChallengeData(summonerName string, region string) (*GetChallengesReturn, error) {
	summonerId, err := api.getSummonerID(region, summonerName)
	if err != nil {
		return nil, err
	}
	endpoint := fmt.Sprintf("/lol/challenges/v1/player-data/%s", summonerId)

	body, err := api.ritoGet(endpoint, region)
	if err != nil {
		return nil, err
	}

	var challengeData GetChallengesReturn
	if err := json.Unmarshal(body, &challengeData); err != nil {
		return nil, err
	}

	return &challengeData, nil
}

//getChallengeInfo
func getChallengeInfo(api *RiotGamesAPI, c *gin.Context) {
	region := c.Param("region")
	playerData, err := api.getChallengeInfoData(region)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, playerData)
}

func (api *RiotGamesAPI) getChallengeInfoData(region string) ([]ChallengesInfoReturn, error) {
	endpoint := "/lol/challenges/v1/challenges/config"

	body, err := api.ritoGet(endpoint, region)
	if err != nil {
		return nil, err
	}

	var challengeInfo []ChallengesInfoReturn
	if err := json.Unmarshal(body, &challengeInfo); err != nil {
		return nil, err
	}

	return challengeInfo, nil
}

//getCombinedChallengeInfo
func getCombinedChallengeInfo(api *RiotGamesAPI, c *gin.Context) {
	summonerName := c.Param("summonerName")
	region := c.Param("region")
	playerData, err := api.getPlayerChallengeData(summonerName, region)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	challengeData, err := api.getChallengeInfoData(region)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	challengeCategories := []inputChallengeCategory{
		{
			CategoryName:  "ARAM God",
			ChallengeName: "ARAM Authority",
			ChallengeNames: []string{
				"ARAM Warrior",
				"ARAM Finesse",
				"ARAM Champion",
			},
		},
		{
			CategoryName:  "ARAM Warrior",
			ChallengeName: "ARAM Warrior",
			ChallengeNames: []string{
				"No Hiding",
				"Bad Medicine",
				"DPS Threat",
				"ARAM Legend",
				"Double Decimation",
				"Solo Carry",
				"Farm Champs Not Minions",
				"ARAM Eradication",
			},
		},
		{
			CategoryName:  "ARAM Finesse",
			ChallengeName: "ARAM Finesse",
			ChallengeNames: []string{
				"Snow Day",
				"Another Day, Another Bullseye",
				"Pop Goes The Poro",
				"Free Ticket to Base",
				"It was a... Near-Hit",
				"Free Money",
			},
		},
		{
			CategoryName:  "ARAM Champion",
			ChallengeName: "ARAM Champion",
			ChallengeNames: []string{
				"All Random All Champions",
				"Rapid Demolition",
				"All Random All Flawless",
				"Active Participant",
				"Lightning Round",
				"NA-RAM",
				"Can't Touch This",
			},
		},
	}

	langCode := "en_US"

	CategorisedChallenges := sortChallengesIntoCategories(challengeData, *playerData, challengeCategories, langCode)

	c.JSON(http.StatusOK, CategorisedChallenges)
}

func sortChallengesIntoCategories(challengesInfo []ChallengesInfoReturn, playerChallenges GetChallengesReturn, challengeInputCategories []inputChallengeCategory, langCode string) []ChallengeCategory {

	var challengeOutputCategories []ChallengeCategory

	for _, category := range challengeInputCategories {
		newChallengeCategory := ChallengeCategory{
			CategoryName: category.CategoryName,
		}
		for _, challengeInfo := range challengesInfo {
			matchFound := false
			for _, challenge := range playerChallenges.Challenges {
				if challenge.ChallengeID == challengeInfo.ID {
					matchFound = true
					if challengeInfo.LocalizedNames[langCode].Name == category.ChallengeName {
						challengeDetail := populateChallengeDetails(challengeInfo, challenge, langCode)
						newChallengeCategory.CategoryChallenge = challengeDetail
					} else if StringInSlice(challengeInfo.LocalizedNames[langCode].Name, category.ChallengeNames) {
						challengeDetail := populateChallengeDetails(challengeInfo, challenge, langCode)
						newChallengeCategory.Challenges = append(newChallengeCategory.Challenges, challengeDetail)
					}
					break
				}
			}
			if !matchFound {
				if StringInSlice(challengeInfo.LocalizedNames[langCode].Name, category.ChallengeNames) {
					challengeDetail := populateBlankChallenge(challengeInfo, langCode)
					newChallengeCategory.Challenges = append(newChallengeCategory.Challenges, challengeDetail)
				}
			}
		}
		challengeOutputCategories = append(challengeOutputCategories, newChallengeCategory)
	}

	return challengeOutputCategories
}

func populateChallengeDetails(challengeInfo ChallengesInfoReturn, challengeData Challenge, langCode string) ChallengeDetail {
	return ChallengeDetail{
		ChallengeID:               challengeInfo.ID,
		ChallengeName:             challengeInfo.LocalizedNames[langCode].Name,
		ChallengeDescription:      challengeInfo.LocalizedNames[langCode].Description,
		ChallengeShortDescription: challengeInfo.LocalizedNames[langCode].ShortDescription,
		Percentile:                challengeData.Percentile,
		Level:                     challengeData.Level,
		Value:                     challengeData.Value,
		AchievedTime:              challengeData.AchievedTime,
		Position:                  challengeData.Position,
		PlayersInLevel:            challengeData.PlayersInLevel,
		Thresholds:                challengeInfo.Thresholds,
	}
}

func populateBlankChallenge(challengeInfo ChallengesInfoReturn, langCode string) ChallengeDetail {
	return ChallengeDetail{
		ChallengeID:               challengeInfo.ID,
		ChallengeName:             challengeInfo.LocalizedNames[langCode].Name,
		ChallengeDescription:      challengeInfo.LocalizedNames[langCode].Description,
		ChallengeShortDescription: challengeInfo.LocalizedNames[langCode].ShortDescription,
		Percentile:                0,
		Level:                     "NONE",
		Value:                     0,
		AchievedTime:              0,
		Position:                  0,
		PlayersInLevel:            0,
		Thresholds:                challengeInfo.Thresholds,
	}
}

//GetSummoner
func (api *RiotGamesAPI) getSummonerID(region string, summonerName string) (string, error) {
	endpoint := fmt.Sprintf("/lol/summoner/v4/summoners/by-name/%s", summonerName)

	body, err := api.ritoGet(endpoint, region)
	if err != nil {
		return "", err
	}

	summonerReturn := GetSummonerReturn{}
	err = json.Unmarshal(body, &summonerReturn)
	if err != nil {
		return "", err
	}

	return summonerReturn.Puuid, nil
}
