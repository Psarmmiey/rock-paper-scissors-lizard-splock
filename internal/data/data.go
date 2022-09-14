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
	{ID: ROCK, Name: "rock"},
	{ID: PAPER, Name: "paper"},
	{ID: SCISSORS, Name: "scissors"},
	{ID: LIZARD, Name: "lizard"},
	{ID: SPOCK, Name: "spock"},
}
var Winning = map[int][2]int{
	ROCK:     {SCISSORS, LIZARD},
	PAPER:    {ROCK, SPOCK},
	SCISSORS: {PAPER, LIZARD},
	LIZARD:   {PAPER, SPOCK},
	SPOCK:    {SCISSORS, ROCK},
}
