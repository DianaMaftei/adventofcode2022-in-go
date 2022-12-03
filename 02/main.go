package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// Day 2: Rock Paper Scissors
func main() {
	// open from the input file
	fileScanner := getFileScanner("day2input.txt")

	totalPoints := 0

	// read line by line
	for fileScanner.Scan() {
		line := fileScanner.Text()
		split := strings.Split(line, " ")

		// get the opponent's choice and the expected outcome of the game
		opponentChoice := getChoice(split[0])
		gameOutcome := getOutcome(split[1])

		// make the correct choice based on the expected outcome
		yourChoice := getYourChoice(opponentChoice, gameOutcome)

		// get the total points based on the outcome of the game and the choice that you made
		totalPoints += int(gameOutcome) + int(yourChoice)
	}

	// display the total score for the game played
	fmt.Println(totalPoints)
}

func getYourChoice(opponentChoice choice, gameOutcome outcome) choice {
	if opponentChoice == rock {
		switch gameOutcome {
		case draw:
			return rock
		case lose:
			return scissors
		case win:
			return paper
		}
	}

	if opponentChoice == paper {
		switch gameOutcome {
		case draw:
			return paper
		case lose:
			return rock
		case win:
			return scissors
		}
	}

	if opponentChoice == scissors {
		switch gameOutcome {
		case draw:
			return scissors
		case lose:
			return paper
		case win:
			return rock
		}
	}

	return scissors
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
