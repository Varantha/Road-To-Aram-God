package main

type GetSummonerReturn struct {
	ID            string `json:"id"`
	AccountID     string `json:"accountId"`
	Puuid         string `json:"puuid"`
	Name          string `json:"name"`
	ProfileIconID int    `json:"profileIconId"`
	RevisionDate  int64  `json:"revisionDate"`
	SummonerLevel int    `json:"summonerLevel"`
}

type GetChallengesReturn struct {
	TotalPoints struct {
		Level      string  `json:"level"`
		Current    int     `json:"current"`
		Max        int     `json:"max"`
		Percentile float64 `json:"percentile"`
	} `json:"totalPoints"`
	CategoryPoints struct {
		Veterancy struct {
			Level      string  `json:"level"`
			Current    int     `json:"current"`
			Max        int     `json:"max"`
			Percentile float64 `json:"percentile"`
		} `json:"VETERANCY"`
		Imagination struct {
			Level      string  `json:"level"`
			Current    int     `json:"current"`
			Max        int     `json:"max"`
			Percentile float64 `json:"percentile"`
		} `json:"IMAGINATION"`
		Collection struct {
			Level      string  `json:"level"`
			Current    int     `json:"current"`
			Max        int     `json:"max"`
			Percentile float64 `json:"percentile"`
		} `json:"COLLECTION"`
		Expertise struct {
			Level      string  `json:"level"`
			Current    int     `json:"current"`
			Max        int     `json:"max"`
			Percentile float64 `json:"percentile"`
		} `json:"EXPERTISE"`
		Teamwork struct {
			Level      string  `json:"level"`
			Current    int     `json:"current"`
			Max        int     `json:"max"`
			Percentile float64 `json:"percentile"`
		} `json:"TEAMWORK"`
	} `json:"categoryPoints"`
	Challenges  []Challenge `json:"challenges"`
	Preferences struct {
		BannerAccent             string `json:"bannerAccent"`
		Title                    string `json:"title"`
		ChallengeIds             []int  `json:"challengeIds"`
		CrestBorder              string `json:"crestBorder"`
		PrestigeCrestBorderLevel int    `json:"prestigeCrestBorderLevel"`
	} `json:"preferences"`
}

type Challenge struct {
	ChallengeID    float64 `json:"challengeId"`
	Percentile     float64 `json:"percentile"`
	Level          string  `json:"level"`
	Value          float64 `json:"value"`
	AchievedTime   int64   `json:"achievedTime,omitempty"`
	Position       float64 `json:"position,omitempty"`
	PlayersInLevel float64 `json:"playersInLevel,omitempty"`
}

type LocalizedName struct {
	Description      string `json:"description"`
	Name             string `json:"name"`
	ShortDescription string `json:"shortDescription"`
}

type Thresholds struct {
	Bronze      float64 `json:"BRONZE,omitempty"`
	Challenger  float64 `json:"CHALLENGER,omitempty"`
	Diamond     float64 `json:"DIAMOND,omitempty"`
	Gold        float64 `json:"GOLD,omitempty"`
	Grandmaster float64 `json:"GRANDMASTER,omitempty"`
	Iron        float64 `json:"IRON,omitempty"`
	Master      float64 `json:"MASTER,omitempty"`
	Platinum    float64 `json:"PLATINUM,omitempty"`
	Silver      float64 `json:"SILVER,omitempty"`
}

type ChallengesInfoReturn struct {
	ID             float64                  `json:"id"`
	Leaderboard    bool                     `json:"leaderboard"`
	LocalizedNames map[string]LocalizedName `json:"localizedNames"`
	State          string                   `json:"state"`
	Thresholds     Thresholds               `json:"thresholds,omitempty"`
}

type ChallengeDetail struct {
	ChallengeID               float64    `json:"challengeId"`
	ChallengeName             string     `json:"challengeName"`
	ChallengeDescription      string     `json:"challengeDescription"`
	ChallengeShortDescription string     `json:"shortDescription"`
	Percentile                float64    `json:"percentile"`
	Level                     string     `json:"level"`
	Value                     float64    `json:"value"`
	AchievedTime              int64      `json:"achievedTime,omitempty"`
	Position                  float64    `json:"position,omitempty"`
	PlayersInLevel            float64    `json:"playersInLevel,omitempty"`
	Thresholds                Thresholds `json:"thresholds,omitempty"`
}

type ChallengeCategory struct {
	CategoryName      string            `json:"categoryName"`
	Challenges        []ChallengeDetail `json:"challenges"`
	CategoryChallenge ChallengeDetail   `json:"categoryChallenge"`
}
