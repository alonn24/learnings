package bowling

type frame struct {
	r1 int
	r2 int
}

func (frame frame) getScore() (score int, isStrike bool, isSpare bool) {
	score = frame.r1 + frame.r2
	isStrike = frame.r1 == 10
	isSpare = !isStrike && score == 10
	return
}
