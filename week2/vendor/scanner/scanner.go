package scanner

import (
	"bufio"
	"os"
	"path/filepath"
	"regexp"
	"sync"
)

var reg, _ = regexp.Compile("[^a-zA-Z0-9]+")

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
	return reg.ReplaceAllString(char, "")
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}
