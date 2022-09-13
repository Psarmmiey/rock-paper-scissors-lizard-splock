package controllers

import (
	"math/rand"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"

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

func randomNum() (int, error) {
	rand.Seed(time.Now().UnixNano())
	min := 1
	max := 5
	randomNumber := rand.Intn(max-min+1) + min
	return randomNumber, nil
}

func Play(c *gin.Context) {
	var player model.Player

	if err := c.BindJSON(&player); err != nil {
		return
	}
	playwin := PlayRound(player.Player)
	c.JSON(http.StatusOK, playwin)

}
