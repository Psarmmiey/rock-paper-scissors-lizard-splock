package controllers

import (
	"net/http"
	data "rock-paper/internal/data"
	logger "rock-paper/internal/log"
	model "rock-paper/internal/model"
	"strconv"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

func GetComputerChoice() *model.Choice {
	computer, err := randomNum()
	if err != nil {
		return nil
	}
	return &data.Choices[computer-1]
}

func PlayGame(playerID, computerID int) string {
	if playerID == computerID {
		return data.Tie
	}

	defeats := data.Winning[playerID]
	if computerID == defeats[0] || computerID == defeats[1] {
		return data.Win
	}

	return data.Lose
}

func PlayRound(c *gin.Context, mode int, player1 int, player2 int) interface{} {
	session := sessions.Default(c)

	var gameOutcome interface{}

	result := PlayGame(player1, player2)
	recordScore(c, result)

	playerScore := session.Get("playerScore")
	computerScore := session.Get("computerScore")
	totalRounds := session.Get("totalRounds")
	logger.LogInfo("----Game Round: " + strconv.Itoa(totalRounds.(int)) + "----")
	message := "Game in progress"

	if playerScore == nil || computerScore == nil || totalRounds == nil {
		playerScore = 0
		computerScore = 0
		totalRounds = 0
	}
	if totalRounds.(int) >= 10 {
		message = ResetScoreBoard(c)
		logger.LogInfo("----Game Round: " + strconv.Itoa(totalRounds.(int)) + "----" + message)
	} else if totalRounds.(int) == 9 {
		message = "Final round"
		logger.LogInfo("----Game Round: " + strconv.Itoa(totalRounds.(int)) + "----" + message)
	}
	if mode == 1 {
		gameOutcome = singlePlayerGameResult(c,
			player1,
			player2,
			result,
			playerScore.(int),
			computerScore.(int),
			totalRounds.(int),
			message,
		)
	} else if mode == 2 {
		if result == data.Win {
			result = "Player 1 wins"
		} else if result == data.Lose {
			result = "Player 2 wins"
		}
		gameOutcome = multiplayerPlayerGameResult(c,
			player1,
			player2,
			result,
			playerScore.(int),
			computerScore.(int),
			totalRounds.(int),
			message,
		)
	}
	return gameOutcome
}

func ResetGameManually(c *gin.Context) {
	ResetScoreBoard(c)
	logger.LogInfo("Game reset manually")
	c.JSON(http.StatusOK, gin.H{
		"message": "Game reset successfully",
	})
}
