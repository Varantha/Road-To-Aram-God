package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func LoadAuth() {
	err := godotenv.Load("local.env")
	if err != nil {
		log.Printf("Unable to load local.env: %s", err)
	}
}

func main() {
	LoadAuth()

	router := gin.Default()
	router.GET("/getThing/:summoner", getThing)

	router.Run("0.0.0.0:8080")
}

func getThing(c *gin.Context) {
	summoner := c.Param("summoner")

	c.IndentedJSON(http.StatusOK, json.RawMessage(summoner))
}

func getSummonerID(summonerName string) (string, error) {
	API_KEY := os.Getenv("API_KEY")

	SummonerUrl := fmt.Sprintf("https://euw1.api.riotgames.com/lol/summoner/v4/summoners/by-name/%s?api_key=%s", summonerName, API_KEY)

	resp, err := http.Get(SummonerUrl)
	if err != nil {
		return "", err
	}

	if resp.StatusCode == 404 {
		err_404 := errors.New("404 - Summoner Not Found")
		return "", err_404
	}

	body, readErr := ioutil.ReadAll(resp.Body)
	if readErr != nil {
		return "", err
	}

	summonerReturn := GetSummonerReturn{}
	err = json.Unmarshal(body, &summonerReturn)
	if err != nil {
		return "", err
	}

	return summonerReturn.ID, nil
}
