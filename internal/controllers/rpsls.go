package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"gopkg.in/validator.v2"

	data "rock-paper/internal/data"
	model "rock-paper/internal/model"
)

func GetChoices(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, data.Choices)
}

func GetSingleChoice(c *gin.Context) {
	id, err := randomNum()

	if err != nil {
		return
	}
	for _, choice := range data.Choices {
		if choice.ID == id {
			c.IndentedJSON(http.StatusOK, choice)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"error": "Choice not found"})
}

func Play(c *gin.Context) {
	var player model.Player
	computerChoice := GetComputerChoice()

	if err := c.BindJSON(&player); err != nil {
		return
	}

	if err := validator.Validate(player); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
	}
	playwin := PlayRound(c, 1, player.Player, computerChoice.ID)
	c.JSON(http.StatusOK, playwin)

}

func Multiplayer(c *gin.Context) {
	var player model.Multiplayer

	if err := c.BindJSON(&player); err != nil {
		return
	}

	if err := validator.Validate(player); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
	}
	playRockPaper := PlayRound(c, 2, player.Player1, player.Player2)
	c.JSON(http.StatusOK, playRockPaper)

}

func singlePlayerGameResult(c *gin.Context, player1 int, player2 int, result string, player1Score int, player2Score int, totalRounds int, message string) model.SinglePlayerGameResult {
	return model.SinglePlayerGameResult{
		Player:        player1,
		Computer:      player2,
		Result:        result,
		PlayerScore:   player1Score,
		ComputerScore: player2Score,
		TotalRounds:   totalRounds,
		Message:       message,
	}
}

func multiplayerPlayerGameResult(c *gin.Context, player1 int, player2 int, result string, player1Score int, player2Score int, totalRounds int, message string) model.MultiPlayerGameResult {
	return model.MultiPlayerGameResult{
		Player1:      player1,
		Player2:      player2,
		Result:       result,
		Player1Score: player1Score,
		Player2Score: player2Score,
		TotalRounds:  totalRounds,
		Message:      message,
	}
}
