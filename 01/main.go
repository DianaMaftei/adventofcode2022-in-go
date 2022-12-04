package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

// Day 1: Calorie Counting
func main() {
	// open from the input file
	fileScanner := getFileScanner("day1input.txt")
	caloriesList := make([]int, 0)

	currentTotal := 0

	// read line by line
	for fileScanner.Scan() {
		line := fileScanner.Text()
		if line != "" {
			calories, err := strconv.Atoi(line)
			if err != nil {
				fmt.Printf("unable to convert to number: %s\n", line)
				break
			}
			currentTotal += calories
		} else {
			caloriesList = append(caloriesList, currentTotal)
			currentTotal = 0
		}
	}

	caloriesList = append(caloriesList, currentTotal)

	// display the total calories of the top 3 elves
	fmt.Println(getTotalOfTopThreeElves(caloriesList))
}

func getTotalOfTopThreeElves(caloriesList []int) int {
	// sort the list descending
	sort.Slice(caloriesList, func(i, j int) bool {
		return caloriesList[i] > caloriesList[j]
	})

	return caloriesList[0] + caloriesList[1] + caloriesList[2]
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
