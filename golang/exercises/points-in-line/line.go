package points

import (
	"fmt"
)

// Line of type y=m*x+n or x=n
type Line interface {
	toString() string
	isOnLine(p []int) bool
}

// LineByTwoPoints - two points as [x,y]
func LineByTwoPoints(p1 []int, p2 []int) Line {
	var res Line
	if p1[0] == p2[0] && p1[1] == p2[1] {
		res = pointLine{p1[0], p1[1]}
	} else if p1[0] == p2[0] {
		res = verticalLine{p1[0]}
	} else {
		x1 := float64(p1[0])
		y1 := float64(p1[1])
		x2 := float64(p2[0])
		y2 := float64(p2[1])
		m := (y1 - y2) / (x1 - x2)
		n := (x1*y2 - x2*y1) / (x1 - x2)
		res = regLine{m, n, p1}
	}
	// log.Printf("calculate for lines - %v, %v = %v", p1, p2, res.toString())
	return res
}

type regLine struct {
	m    float64
	n    float64
	base []int
}

func (line regLine) toString() string {
	m := line.m
	if m == 0 {
		m = 0 // remove -0
	}
	n := line.n
	if n == 0 {
		n = 0 // remove -0
	}
	return fmt.Sprintf("y=%vx+%v", m, n)
}

func (line regLine) isOnLine(p []int) bool {
	if p[0] == line.base[0] && p[1] == line.base[1] {
		return true
	}
	newLine := LineByTwoPoints(p, line.base)
	return line.toString() == newLine.toString()
	// return float64(p[1]) == line.m*float64(p[0])+line.n
}

type verticalLine struct {
	x int
}

func (line verticalLine) toString() string {
	return fmt.Sprintf("x=%v", line.x)
}

func (line verticalLine) isOnLine(p []int) bool {
	return p[0] == line.x
}

type pointLine struct {
	x int
	y int
}

func (line pointLine) toString() string {
	return fmt.Sprintf("x=%v, y=%v", line.x, line.y)
}

func (line pointLine) isOnLine(p []int) bool {
	return p[0] == line.x && p[1] == line.y
}
