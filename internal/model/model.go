package interfaces

type Round struct {
	Message        string `json:"message"`
	ComputerChoice string `json:"computer_choice"`
	RoundResult    string `json:"round_result"`
}

type Choice struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type GameResult struct {
	Player   int    `json:"player"`
	Computer int    `json:"computer"`
	Result   string `json:"results"`
}

type Player struct {
	Player int `json:"player"`
}
