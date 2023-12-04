package main

import (
	"database/sql"
	"math/rand"
	"time"
)

type DAO struct {
	db *sql.DB
}

// NewDAO
//   - @name NewDAO
//   - @description Create a new DAO
//   - @param db *sql.DB
//   - @return *DAO
func NewDAO(db *sql.DB) *DAO {
	return &DAO{db: db}
}

// GetRanPlayer
//   - @name GetRanPlayer
//   - @description Get a random player from the database
//   - @return Player
func (dao *DAO) GetRanPlayer() (Player, error) {
	randomNumber := dao.getRanNum(1, 730)
	query := "SELECT * FROM players WHERE players.ID = ?"
	var player Player

	err := dao.db.QueryRow(query, randomNumber).Scan(
		&player.ID,
		&player.FirstName,
		&player.LastName,
		&player.DateOfBirth,
		&player.Goals,
		&player.Country,
		&player.Games,
		&player.GoalsSelection,
		&player.GamesSelection,
		&player.Assists,
	)
	if err != nil {
		return Player{}, err
	}
	return player, nil
}

// UpdatePlayer
//   - @name UpdatePlayer
//   - @description Update a player in the database Only used in development.
//   - @param data Player
//   - @return error
func (dao *DAO) UpdatePlayer(data Player) error {
	query := `UPDATE players
               SET assists = ?
               WHERE firstname = ?
               AND lastname = ?`

	_, err := dao.db.Exec(query, data.Assists, data.FirstName, data.LastName)
	if err != nil {
		return err
	}

	return nil
}

// AddResult
//   - @name AddResult
//   - @description Add a result in the database
//   - @param data ScoreData
//   - @return error
func (dao *DAO) AddResult(data ScoreData) error {
	now := time.Now()
	formattedDate := now.Format("2006-01-02 15:04:05")
	query := `INSERT INTO scores (score, time, timeMins) VALUES (?, ?, ?)`

	_, err := dao.db.Exec(query, data.Score, formattedDate, formattedDate)
	if err != nil {
		return err
	}

	return nil
}

// AddAnalytics
//   - @name AddAnalytics
//   - @description Add analytics in the database. Stored statistics to improve the game.
//   - @param data AnalyticsData
//   - @return error
func (dao *DAO) AddAnalytics(data AnalyticsData) error {
	query := `INSERT INTO analytics (questionID, rightPlayID, player1ID, player2ID, score) VALUES (?, ?, ?, ?, ?)`

	_, err := dao.db.Exec(query, data.QuestionID, data.RightPlayID, data.Player1ID, data.Player2ID, data.Score)
	if err != nil {
		return err
	}

	return nil
}

// getRanNum
//   - @name getRanNum
//   - @description Get a random number between min and max, represents a Player ID.
//   - @param min int
//   - @param max int
//   - @return int
func (dao *DAO) getRanNum(min, max int) int {
	if min > max {
		min, max = max, min
	}

	return rand.Intn(max-min+1) + min
}
