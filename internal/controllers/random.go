package controllers

import (
	"math/rand"
	"time"
)

func randomNum() (int, error) {
	rand.Seed(time.Now().UnixNano())
	min := 1
	max := 5
	randomNumber := rand.Intn(max-min+1) + min
	return randomNumber, nil
}
