package bowling

type frame struct {
	r1 int
	r2 int
}

func (frame frame) getScore() int {
	return frame.r1 + frame.r2
}

func (frame frame) isStrike() bool {
	return frame.r1 == 10
}

func (frame frame) isSpare() bool {
	return !frame.isStrike() && frame.getScore() == 10
}
