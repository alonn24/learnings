package bowling

type roll struct {
	first  int
	second int
	third  int
}

type Game struct {
	rolls []roll
	index int
}

func (game Game) isNewTurn() bool {
	return len(game.rolls) == game.index
}

func (game *Game) roll(score int) *Game {
	if game.isNewTurn() {
		game.rolls = append(game.rolls, roll{})
		game.rolls[game.index].first = score
	} else {
		game.rolls[game.index].second = score
		game.index++
	}
	return game
}

func (game Game) score() int {
	score := 0
	for i := 0; i < game.index; i++ {
		roll := game.rolls[i]
		score += roll.first + roll.second
	}
	return score
}
