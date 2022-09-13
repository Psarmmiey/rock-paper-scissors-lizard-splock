package controllers

import (
	data "rock-paper/internal/data"
	model "rock-paper/internal/model"
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

func PlayRound(player int) model.GameResult {
	computerChoice := GetComputerChoice()
	return model.GameResult{Player: player, Computer: computerChoice.ID, Result: PlayGame(player, computerChoice.ID)}
}
