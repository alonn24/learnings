package bowling

// Game - bowling class game
type Game struct {
	frames []frame
}

func isNewFrame(game *Game) bool {
	return len(game.frames) == 0 || game.frames[len(game.frames)-1].isDone()
}

func addFrame(game *Game) *Game {
	if len(game.frames) == 9 {
		game.frames = append(game.frames, &lastFrame{})
	} else {
		game.frames = append(game.frames, &singleFrame{})
	}
	return game
}

// Roll
func (game *Game) roll(score int) *Game {
	if isNewFrame(game) {
		addFrame(game)
	}
	frame := game.frames[len(game.frames)-1]
	frame.roll(score)
	return game
}

func getRollsForFrame(game Game, i int) []int {
	if len(game.frames) > i {
		return game.frames[i].getRolls()
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
