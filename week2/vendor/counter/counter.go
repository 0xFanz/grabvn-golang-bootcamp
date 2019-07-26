package counter

import (
	"sync"
)

var mutex sync.Mutex

// Count Increase the frequency of the word in the map
func Count(wg *sync.WaitGroup, wordBroker chan string, table map[string]int) {
	defer wg.Done()

	for {
		word, ok := <-wordBroker

		if !ok {
			return
		}

		mutex.Lock()
		table[word]++
		mutex.Unlock()
	}
}
