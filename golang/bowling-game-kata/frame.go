package bowling

type frame struct {
	rolls []int
}

func (frame *frame) roll(score int) {
	frame.rolls = append(frame.rolls, score)
}

func (frame frame) isStrike() bool {
	return (len(frame.rolls) == 1 && frame.rolls[0] == 10)
}

func (frame frame) getScore() int {
	sum := 0
	for _, s := range frame.rolls {
		sum += s
	}
	return sum
}

func (frame frame) isSpare() bool {
	return !frame.isStrike() && frame.getScore() == 10
}

func (frame frame) isDone() bool {
	return len(frame.rolls) == 2 || frame.isStrike()
}
