package controllers

import (
	data "rock-paper/internal/data"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

func recordScore(c *gin.Context, result string) {
	session := sessions.Default(c)
	playerScore := session.Get("playerScore")
	computerScore := session.Get("computerScore")
	totalRounds := session.Get("totalRounds")
	if playerScore == nil || computerScore == nil || totalRounds == nil {
		playerScore = 0
		computerScore = 0
		totalRounds = 0
	}

	if result == data.Win {
		playerScore = playerScore.(int) + 1
	} else if result == data.Lose {
		computerScore = computerScore.(int) + 1
	}

	session.Set("playerScore", playerScore)
	session.Set("computerScore", computerScore)
	session.Set("totalRounds", totalRounds.(int)+1)

	session.Save()
}

func ResetScoreBoard(c *gin.Context) string {
	session := sessions.Default(c)
	session.Clear()
	session.Save()
	return "Scoreboard reset successfully"
}
