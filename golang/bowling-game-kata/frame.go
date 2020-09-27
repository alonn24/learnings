package bowling

type frame interface {
	getRolls() []int
	roll(score int)
	isStrike() bool
	getScore() int
	isSpare() bool
	isDone() bool
}

/* singleFrame */
type singleFrame struct {
	rolls []int
}

func (singleFrame singleFrame) getRolls() []int {
	return singleFrame.rolls
}

func (singleFrame *singleFrame) roll(score int) {
	singleFrame.rolls = append(singleFrame.rolls, score)
}

func (singleFrame singleFrame) isStrike() bool {
	return (len(singleFrame.rolls) == 1 && singleFrame.rolls[0] == 10)
}

func (singleFrame singleFrame) getScore() int {
	sum := 0
	for _, s := range singleFrame.rolls {
		sum += s
	}
	return sum
}

func (singleFrame singleFrame) isSpare() bool {
	return !singleFrame.isStrike() && singleFrame.getScore() == 10
}

func (singleFrame singleFrame) isDone() bool {
	return len(singleFrame.rolls) == 2 || singleFrame.isStrike()
}
