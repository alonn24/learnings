package bowling

// Game - bowling class game
type Game struct {
	frames []frame
	index  int
}

// Roll
func (game *Game) roll(score int) *Game {
	isNewTurn := len(game.frames) == game.index
	if isNewTurn {
		game.frames = append(game.frames, frame{})
		game.frames[game.index].r1 = score
		if score == 10 {
			game.index++
		}
	} else {
		game.frames[game.index].r2 = score
		game.index++
	}
	return game
}

func getFrameAt(game Game, i int) *frame {
	if len(game.frames) > i {
		return &game.frames[i]
	}
	return nil
}

func getNextRolls(game Game, i int) (int, int) {
	nextFrame := getFrameAt(game, i+1)
	next2Frame := getFrameAt(game, i+2)
	if nextFrame == nil {
		return 0, 0
	}

	r1 := nextFrame.r1
	r2 := nextFrame.r2
	if nextFrame.isStrike() && next2Frame != nil {
		r2 = next2Frame.r1
	}
	return r1, r2
}

// Score
func (game Game) score() int {
	total := 0
	for i := 0; i < len(game.frames); i++ {
		frame := game.frames[i]
		nr1, nr2 := getNextRolls(game, i)
		total += frame.getScore()
		if frame.isStrike() {
			total += nr1 + nr2
		} else if frame.isSpare() {
			total += nr1
		}
	}
	return total
}
