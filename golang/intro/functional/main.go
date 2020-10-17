package functional

import "log"

func Map(vs []int, f func(int) int) []int {
	vsm := make([]int, len(vs))
	for i, v := range vs {
		vsm[i] = f(v)
	}
	return vsm
}

func Mapper(x int) int { return x + 1 }

func run() {
	list := []int{1, 2, 3}
	log.Printf("%v", Map(list, Mapper))
}
