package tenniskata

var scoreMap = map[int]string{
	0: "Love",
	1: "Fifteen",
	2: "Thirty",
	3: "Forty",
}

type tennisGame1 struct {
	player1Score int
	player2Score int
	player1Name  string
	player2Name  string
}

func TennisGame1(player1Name string, player2Name string) TennisGame {
	game := &tennisGame1{
		player1Name: player1Name,
		player2Name: player2Name}

	return game
}

func (game *tennisGame1) WonPoint(playerName string) {
	if playerName == "player1" {
		game.player1Score += 1
	} else {
		game.player2Score += 1
	}
}

func (game *tennisGame1) GetScore() string {
	if arePlayersTied(game) {
		return getScoreAtTied(game)
	} else if arePlayersAtAdvantage(game) {
		return getScoreAtAdvantage(game)
	} else {
		return getScoreNormal(game)
	}
}
func arePlayersAtAdvantage(game *tennisGame1) bool {
	return game.player1Score >= 4 || game.player2Score >= 4
}

func arePlayersTied(game *tennisGame1) bool {
	return game.player1Score == game.player2Score
}

func getScoreNormal(game *tennisGame1) string {
	score := ""
	tempScore := 0
	for i := 1; i < 3; i++ {
		if i == 1 {
			tempScore = game.player1Score
		} else {
			score += "-"
			tempScore = game.player2Score
		}
		score += scoreMap[tempScore]

	}
	return score
}

func getScoreAtTied(game *tennisGame1) string {
	if game.player1Score > 2 {
		return "Deuce"
	}

	return scoreMap[game.player1Score] + "-All"
}

func getScoreAtAdvantage(game *tennisGame1) string {
	scoreDifference := game.player1Score - game.player2Score
	if scoreDifference == 1 {
		return "Advantage player1"
	} else if scoreDifference == -1 {
		return "Advantage player2"
	} else if scoreDifference >= 2 {
		return "Win for player1"
	} else {
		return "Win for player2"
	}
}
