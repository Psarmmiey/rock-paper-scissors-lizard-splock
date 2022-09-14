package data

import (
	model "rock-paper/internal/model"
)

const (
	ROCK = iota + 1
	PAPER
	SCISSORS
	LIZARD
	SPOCK
)

const (
	Win  = "win"
	Lose = "lose"
	Tie  = "tie"
)

var Choices = []model.Choice{
	{ID: ROCK, Name: "Rock"},
	{ID: PAPER, Name: "Paper"},
	{ID: SCISSORS, Name: "Scissors"},
	{ID: LIZARD, Name: "Lizard"},
	{ID: SPOCK, Name: "Spock"},
}
var Winning = map[int][2]int{
	ROCK:     {SCISSORS, LIZARD},
	PAPER:    {ROCK, SPOCK},
	SCISSORS: {PAPER, LIZARD},
	LIZARD:   {PAPER, SPOCK},
	SPOCK:    {SCISSORS, ROCK},
}
