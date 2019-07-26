package scanner

import (
	"bufio"
	"log"
	"os"
	"path/filepath"
	"regexp"
	"sync"
)

// GetPaths scan all files in folder then return to array
func GetPaths(dir string) []string {
	var files []string

	err := filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if !info.IsDir() {
			files = append(files, path)
		}
		return nil
	})
	check(err)
	return files
}

// SendWords send single word to wordBroker
func SendWords(wg *sync.WaitGroup, path string, wordBroker chan string) {

	file, fileErr := os.Open(path)
	check(fileErr)

	defer file.Close()
	defer wg.Done()

	scanner := bufio.NewScanner(file)

	scanner.Split(bufio.ScanWords)

	for scanner.Scan() {
		wordBroker <- filterChar(scanner.Text())
	}
	scanErr := scanner.Err()
	check(scanErr)
}

func filterChar(char string) string {
	reg, err := regexp.Compile("[^a-zA-Z0-9]+")
	if err != nil {
		log.Fatal(err)
	}
	return reg.ReplaceAllString(char, "")
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}
