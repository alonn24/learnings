package mapreduce

import (
	"encoding/csv"
	"log"
	"os"
	"strconv"
	"sync"
)

// PlayerInfo -
type PlayerInfo struct {
	ID     string `json:"ID"`
	Name   string `json:"Name"`
	Age    int    `json:"Age"`
	Season string `json:"Season"`
}

func getLinesFromFile() [][]string {
	f, err := os.Open("./nba_players.csv")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	lines, err := csv.NewReader(f).ReadAll()
	if err != nil {
		panic(err)
	}
	return lines
}

// Map -
func Map(player []string) []PlayerInfo {
	list := []PlayerInfo{}
	age, _ := strconv.ParseFloat(player[3], 32)
	list = append(list, PlayerInfo{ID: player[1], Name: player[2], Age: int(age), Season: player[len(player)-1]})
	return list
}

// Reducer -
func Reducer(mapList chan []PlayerInfo, sendFinalValue chan []PlayerInfo) {
	final := []PlayerInfo{}
	for list := range mapList {
		for _, value := range list {
			if value.Age >= 40 {
				final = append(final, value)
			}
		}
	}
	sendFinalValue <- final
}

/*
	Do not communicate by sharing the memory;
	instead, share memory by communicating.
*/
func build() {
	lines := getLinesFromFile()
	lists := make(chan []PlayerInfo)
	finalValue := make(chan []PlayerInfo)

	var wg sync.WaitGroup
	wg.Add(len(lines))

	for _, line := range lines {
		go func(player []string) {
			defer wg.Done()
			lists <- Map(player)
		}(line)
	}

	go Reducer(lists, finalValue)

	wg.Wait()
	close(lists)

	log.Print(<-finalValue)

}
