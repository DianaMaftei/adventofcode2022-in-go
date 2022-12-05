package main

import (
	"bufio"
	"fmt"
	"os"
)

var priorities = map[string]byte{
	"a": 1,
	"b": 2,
	"c": 3,
	"d": 4,
	"e": 5,
	"f": 6,
	"g": 7,
	"h": 8,
	"i": 9,
	"j": 10,
	"k": 11,
	"l": 12,
	"m": 13,
	"n": 14,
	"o": 15,
	"p": 16,
	"q": 17,
	"r": 18,
	"s": 19,
	"t": 20,
	"u": 21,
	"v": 22,
	"w": 23,
	"x": 24,
	"y": 25,
	"z": 26,
	"A": 27,
	"B": 28,
	"C": 29,
	"D": 30,
	"E": 31,
	"F": 32,
	"G": 33,
	"H": 34,
	"I": 35,
	"J": 36,
	"K": 37,
	"L": 38,
	"M": 39,
	"N": 40,
	"O": 41,
	"P": 42,
	"Q": 43,
	"R": 44,
	"S": 45,
	"T": 46,
	"U": 47,
	"V": 48,
	"W": 49,
	"X": 50,
	"Y": 51,
	"Z": 52,
}

// Day 3: Rucksack Reorganization
func main() {
	// open from the input file
	fileScanner := getFileScanner("day3input.txt")

	sumOfPriorities := 0
	elfNumberInGroup := 1
	commonItems := make(map[string]int, 0)
	itemsInFirstBackpack := make(map[string]bool, 0)

	// read line by line
	for fileScanner.Scan() {
		line := fileScanner.Text()

		// get common items for all 3 elves
		commonItemsInThisBackpack := make(map[string]bool, 0)
		for _, item := range line {
			// tack all the items that elf 1 had
			if elfNumberInGroup == 1 {
				itemsInFirstBackpack[string(item)] = true
				continue
			}

			itemWasInFirstBackpack := itemsInFirstBackpack[string(item)]
			itemAlreadySeenInCurrentBackpack := commonItemsInThisBackpack[string(item)]

			// check if current item was present in the 1st backpack
			// and that we haven't already counted it
			if itemWasInFirstBackpack && !itemAlreadySeenInCurrentBackpack {
				commonItems[string(item)] = commonItems[string(item)] + 1
				commonItemsInThisBackpack[string(item)] = true
			}
		}

		badge := ""
		if elfNumberInGroup == 3 {
			for item, count := range commonItems {
				// the badge is the only item that was common to the other 2 elves as well
				if count == 2 {
					badge = item
					break
				}
			}
		}

		// get priority of item type
		priority := priorities[badge]

		// add to sum
		sumOfPriorities += int(priority)

		if elfNumberInGroup == 3 {
			elfNumberInGroup = 0
			commonItems = make(map[string]int, 0)
			itemsInFirstBackpack = make(map[string]bool, 0)
		}
		elfNumberInGroup++
	}

	// display the sum of the priorities of the item types
	fmt.Println(sumOfPriorities)
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
