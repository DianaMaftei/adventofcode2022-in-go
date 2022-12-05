package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// Day 5: Supply Stacks
func main() {
	// open from the input file
	fileScanner := getFileScanner("day5input.txt")
	stacks := make([][]string, 9)

	hasFinishedDiagram := false

	// read line by line
	for fileScanner.Scan() {
		line := fileScanner.Text()
		// empty line separating the stacks diagram from the moving instructions
		if len(line) == 0 {
			continue
		}
		split := strings.Split(line, " ")

		// get the diagram of crates in stacks
		if !hasFinishedDiagram {
			stackNo := 0
			for i := 0; i < len(split); i++ {
				if split[i] == "" {
					if i == 0 && len(split[i+1]) == 1 {
						hasFinishedDiagram = true
						break
					}
					// empty spaces, no crate, move on to next stack
					i += 3
					stackNo++
				} else {
					// add crate to stacks
					stacks[stackNo] = append(stacks[stackNo], split[i])
					stackNo++
				}
			}
		} else {
			// get move instruction
			count, _ := strconv.Atoi(split[1])
			origin, _ := strconv.Atoi(split[3])
			destination, _ := strconv.Atoi(split[5])

			// identify crates to be moved
			cratesToMove := make([]string, count)
			copy(cratesToMove, stacks[origin-1][0:count])

			// perform move
			stacks[destination-1] = append(cratesToMove, stacks[destination-1]...)
			stacks[origin-1] = stacks[origin-1][count:]
		}
	}

	topCrates := ""
	for _, stack := range stacks {
		// look at the top of each stack
		// get the 2nd character, which is the crate letter, eg [A]
		topCrates += string(stack[0][1])
	}

	// display the crate letters that end up on top of each stack
	fmt.Println(topCrates)
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
