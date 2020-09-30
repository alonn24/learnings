package main

import (
	"log"
)

func main() {
	// log.SetPrefix("main: ")
	log.SetFlags(0)
	hello()
	structures()
	concurrency()
}
