package main

import (
	"bufio"
	"fmt"
	"os"
)

// Day 6: Tuning Trouble
func main() {
	// open from the input file
	fileScanner := getFileScanner("day6input.txt")

	characterCount := 0

	// read line by line
	for fileScanner.Scan() {
		line := fileScanner.Text()

		lastFourCharacters := make([]string, 0)

		for _, ch := range line {
			lastFourCharacters = append(lastFourCharacters, string(ch))
			characterCount++

			if characterCount >= 14 {
				if !hasDuplicates(lastFourCharacters) {
					fmt.Println(characterCount)
					lastFourCharacters = make([]string, 0)
					characterCount = 0
					break
				}

				lastFourCharacters = lastFourCharacters[1:]
			}

		}
	}
}

func hasDuplicates(characters []string) bool {
	characterMap := make(map[string]bool)

	for _, ch := range characters {
		if characterMap[ch] {
			return true
		}
		characterMap[ch] = true
	}
	return false
}

func getFileScanner(fileName string) *bufio.Scanner {
	fileReader, err := os.Open(fileName)
	if err != nil {
		fmt.Println("unable to open file for reading: %w", err)
	}

	fileScanner := bufio.NewScanner(fileReader)
	fileScanner.Split(bufio.ScanLines)
	return fileScanner
}
