package main

import (
	"fmt"
	"log"
	"sync"
	"time"
)

func worker(c chan<- string, id int, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 0; i < 5; i++ {
		c <- fmt.Sprintf("ping %v: %v/10", id, i)
		time.Sleep(time.Second)
	}
}

func printer(c <-chan string, finished chan<- bool) {
	defer func() { finished <- true }()

	timec := time.After(time.Second * 5)
	for {
		select {
		case msg := <-c:
			log.Println(msg)
		case <-timec:
			log.Println("timeout")
			return
		default:
			log.Println("nothing ready")
			time.Sleep(time.Second * 1)
		}
	}
}

func concurrency() {
	var wg sync.WaitGroup
	finished := make(chan bool)
	var c chan string = make(chan string, 500)

	for i := 0; i < 5; i++ {
		wg.Add(1)
		go worker(c, i, &wg)
	}
	go printer(c, finished)
	wg.Wait()
	close(c)
	<-finished
	for elem := range c {
		fmt.Printf("items left: %v\n", elem)
	}
	fmt.Print("press to continue")
	var input string
	fmt.Scanln(&input)
}
