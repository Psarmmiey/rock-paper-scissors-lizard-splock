# rock-paper-scissors-lizard-spock-go

This is a simple implementation of the game Rock Paper Scissors Lizard Spock in Go.

Run locally
   1. Clone the files.
   2. Open the console, stand on the root folder of this project and type: go run main.go
   3. Paste on your browser: http://localhost:8080/
   4. Play.

You can also run the Docker image or the Makefile.

## Game rules

  - Scissors cuts Paper
  - Paper covers Rock
  - Rock crushes Lizard
  - Lizard poisons Spock
  - Spock smashes Scissors
  - Scissors decapitates Lizard
  - Lizard eats Paper
  - Paper disproves Spock
  - Spock vaporizes Rock
  - Rock crushes Scissors

  *Each game play is reset automatically after 10 rounds.*



## Game Modes

   1. Player vs Computer
   2. Multiplayer

## Endpoints

 
1. /choices [GET] - Get all choices
   
```json
[
    {
        "id": 1,
        "name": "Rock"
    },
    ...
    {
        "id": 5,
        "name": "Spock"
    }
]
```


2. /choice  [GET] - Get a random choice
   
```json
{
    "id": 1,
    "name": "Rock"
}
```

3. /play [POST] - Play a game
```json
{
    "player": 2
}
```
```json
{
    "player": 2,
    "computer": 4,
    "results": "lose",
    "player_score": 2,
    "computer_score": 6,
    "total_rounds": 9,
    "message": "Final round"
}
```

4. /multiplayer [POST] - Play a multiplayer game

```json
{
    "player1":2,
    "player2":5
}
```json
{
     "player1": 2,
    "player2": 3,
    "results": "Player 2 wins",
    "player1_score": 0,
    "player2_score": 1,
    "total_rounds": 1,
    "message": "Game in progress"
}
```

5. /reset [GET] - Reset the game

```json
{
    "message": "Game reset successfully"
}
```


## Credits
Sam Kass & Karen Bryla http://www.samkass.com/theories/RPSSL.html