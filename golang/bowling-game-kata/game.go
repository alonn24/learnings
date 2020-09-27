package bowling

type Game struct {
	frames []frame
	index  int
}

func isNewTurn(game *Game) bool {
	return len(game.frames) == game.index
}

func playFirstRoll(game *Game, score int) {
	game.frames = append(game.frames, frame{})
	game.frames[game.index].r1 = score
}

func playSecondRoll(game *Game, score int) {
	game.frames[game.index].r2 = score
	game.index++
}

func (game *Game) roll(score int) *Game {
	if isNewTurn(game) {
		playFirstRoll(game, score)
	} else {
		playSecondRoll(game, score)
	}
	return game
}

func getRolls(game Game, i int) (int, int) {
	if len(game.frames) > i+1 {
		return game.frames[i+1].r1, game.frames[i+1].r2
	}
	return 0, 0
}

func (game Game) score() int {
	total := 0
	for i := 0; i < game.index; i++ {
		roll := game.frames[i]
		nr1, nr2 := getRolls(game, i)
		score, isStrike, isSpare := roll.getScore()
		total += score
		if isStrike {
			total += nr1 + nr2
		} else if isSpare {
			total += nr1
		}
	}
	return total
}
