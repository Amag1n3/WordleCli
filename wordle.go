package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strings"
)

func getWordList(filepath string) ([]string, error) {
	file, err := os.Open(filepath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var words []string
	reader := bufio.NewReader(file)
	for {
		line, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("Error reading the file")
		}
		line = strings.TrimSpace(line)
		if len(line) == 5 {
			words = append(words, line)
		}
	}
	return words, nil
}

func wordleGame(wordList []string) {
	target := wordList[rand.Intn(len(wordList))]
	maxAttempts := 6
	fmt.Println("Welcome to Wordle CLI game")

}
