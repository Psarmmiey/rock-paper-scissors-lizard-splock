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

type SinglePlayerGameResult struct {
	Player        int    `json:"player"`
	Computer      int    `json:"computer"`
	Result        string `json:"results"`
	PlayerScore   int    `json:"player_score"`
	ComputerScore int    `json:"computer_score"`
	TotalRounds   int    `json:"total_rounds"`
	Message       string `json:"message"`
}

type MultiPlayerGameResult struct {
	Player1      int    `json:"player1"`
	Player2      int    `json:"player2"`
	Result       string `json:"results"`
	Player1Score int    `json:"player1_score"`
	Player2Score int    `json:"player2_score"`
	TotalRounds  int    `json:"total_rounds"`
	Message      string `json:"message"`
}

type Player struct {
	Player int `json:"player" validate:"required,min=1,max=5"`
}

type Multiplayer struct {
	Player1 int `json:"player1" validate:"required,min=1,max=5"`
	Player2 int `json:"player2" validate:"required,min=1,max=5"`
}
