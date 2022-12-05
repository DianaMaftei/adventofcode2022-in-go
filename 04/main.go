package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// Day 4: Camp Cleanup
func main() {
	// open from the input file
	fileScanner := getFileScanner("day4input.txt")

	rangeContainedCount := 0

	// read line by line
	for fileScanner.Scan() {
		line := fileScanner.Text()
		pairs := strings.Split(line, ",")
		elfOneSectionsRange := strings.Split(pairs[0], "-")
		elfTwoSectionsRange := strings.Split(pairs[1], "-")

		elfOneSectionsStart, _ := strconv.Atoi(elfOneSectionsRange[0])
		elfOneSectionsEnd, _ := strconv.Atoi(elfOneSectionsRange[1])

		elfTwoSectionsStart, _ := strconv.Atoi(elfTwoSectionsRange[0])
		elfTwoSectionsEnd, _ := strconv.Atoi(elfTwoSectionsRange[1])

		elfOneSections := make(map[int]bool, 0)

		for i := elfOneSectionsStart; i <= elfOneSectionsEnd; i++ {
			elfOneSections[i] = true
		}

		for i := elfTwoSectionsStart; i <= elfTwoSectionsEnd; i++ {
			if elfOneSections[i] {
				rangeContainedCount++
				break
			}
		}

	}

	// display in how many assignment pairs one range fully contain the other
	fmt.Println(rangeContainedCount)
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
