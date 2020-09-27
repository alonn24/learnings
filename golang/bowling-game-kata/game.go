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

	r1 := 0
	r2 := 0
	if len(nextFrame.rolls) >= 1 {
		r1 = nextFrame.rolls[0]
	}
	if nextFrame.isStrike() && next2Frame != nil && len(next2Frame.rolls) >= 1 {
		r2 = next2Frame.rolls[0]
	} else if len(nextFrame.rolls) >= 2 {
		r2 = nextFrame.rolls[1]
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
