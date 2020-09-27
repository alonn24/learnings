package bowling

// Game - bowling class game
type Game struct {
	frames []frame
}

// Roll
func (game *Game) roll(score int) *Game {
	if len(game.frames) == 0 || game.frames[len(game.frames)-1].isDone() {
		game.frames = append(game.frames, frame{})
	}
	frame := &game.frames[len(game.frames)-1]
	frame.roll(score)
	return game
}

func getRollsForFrame(game Game, i int) []int {
	if len(game.frames) > i {
		return game.frames[i].rolls
	}
	return []int{}
}

func getNextTwoRolls(game Game, i int) (int, int) {
	nextRolls := append(getRollsForFrame(game, i+1), getRollsForFrame(game, i+2)...)
	r1 := 0
	if len(nextRolls) > 0 {
		r1 = nextRolls[0]
	}

	r2 := 0
	if len(nextRolls) > 1 {
		r2 = nextRolls[1]
	}
	return r1, r2
}

// Score
func (game Game) score() int {
	total := 0
	for i := 0; i < len(game.frames); i++ {
		frame := game.frames[i]
		nr1, nr2 := getNextTwoRolls(game, i)
		total += frame.getScore()
		if frame.isStrike() {
			total += nr1 + nr2
		} else if frame.isSpare() {
			total += nr1
		}
	}
	return total
}
