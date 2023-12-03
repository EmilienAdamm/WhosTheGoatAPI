package main

type Player struct {
	ID             int    `json:"ID"`
	FirstName      string `json:"firstname"`
	LastName       string `json:"lastname"`
	DateOfBirth    string `json:"dateOfBirth"`
	Goals          int    `json:"goals"`
	Country        string `json:"country"`
	Games          int    `json:"games"`
	GoalsSelection int    `json:"goalsSelection"`
	GamesSelection int    `json:"gamesSelection"`
	Assists        int    `json:"assists"`
}

type ScoreData struct {
	Score int `json:"score"`
}

type AnalyticsData struct {
	QuestionID  int `json:"questionID"`
	RightPlayID int `json:"rightPlayID"`
	Player1ID   int `json:"player1ID"`
	Player2ID   int `json:"player2ID"`
	Score       int `json:"score"`
}
