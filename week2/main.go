package main

import (
	"counter"
	"fmt"
	"scanner"
	"sync"
)

const (
	numbersGoroutines = 10
	dataDir           = "data"
)

var (
	mapperWaitGroup  sync.WaitGroup
	scannerWaitGroup sync.WaitGroup
	table            = make(map[string]int)
)

func main() {
	wordBroker := make(chan string)
	files := scanner.GetPaths(dataDir)

	scannerWaitGroup.Add(len(files))
	for _, file := range files {
		go scanner.SendWords(&scannerWaitGroup, file, wordBroker)
	}

	mapperWaitGroup.Add(numbersGoroutines)
	for gr := 1; gr <= numbersGoroutines; gr++ {
		go counter.Count(&mapperWaitGroup, wordBroker, table)
	}

	scannerWaitGroup.Wait()
	close(wordBroker)
	mapperWaitGroup.Wait()

	fmt.Println(table)
}
