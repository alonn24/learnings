package points

type entry struct {
	line  Line
	score int
}

func maxPoints(points [][]int) int {
	if len(points) == 1 {
		return 1
	}
	scores := make(map[string]entry)
	for i := 0; i < len(points); i++ {
		p1 := points[i]
		for j := i + 1; j < len(points); j++ {
			p2 := points[j]
			line := LineByTwoPoints(p1, p2)
			scores[line.toString()] = entry{line, 0}
		}
	}

	for i := range scores {
		for _, point := range points {
			if scores[i].line.isOnLine(point) {
				scores[i] = entry{scores[i].line, scores[i].score + 1}
			}
		}
	}
	// log.Printf("%v", scores)
	max := 0
	for _, entry := range scores {
		if entry.score > max {
			max = entry.score
		}
	}
	return max
}
