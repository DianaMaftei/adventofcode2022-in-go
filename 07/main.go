package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

// Day 7: No Space Left On Device
func main() {
	// open from the input file
	fileScanner := getFileScanner("day7input.txt")

	diskSpaceAvailable := 70000000
	diskSpaceNeeded := 30000000

	separator := "/"
	currentDir := ""
	directories := make(map[string]int)

	// read line by line
	for fileScanner.Scan() {
		line := fileScanner.Text()

		// parse command list & populate a map of directories and their sizes, based on the files within
		if strings.HasPrefix(line, "$ cd") {
			destination := line[5:]

			if destination == ".." {
				lastIndexOfSeparator := strings.LastIndex(currentDir[:len(currentDir)-1], separator)
				currentDir = currentDir[:lastIndexOfSeparator+1]
			} else if destination == "/" {
				currentDir = destination
			} else {
				currentDir += destination + separator
			}
		} else if strings.HasPrefix(line, "$ ls") {
			continue
		} else if strings.HasPrefix(line, "dir") {
			continue
		} else {
			detail := strings.Split(line, " ")
			size, _ := strconv.Atoi(detail[0])
			updateSizes(directories, currentDir, size)
		}
	}

	totalUsedSpace := 0

	// put all the directory sizes in a list
	allDirSizes := make([]int, 0)
	for dir, size := range directories {
		allDirSizes = append(allDirSizes, size)
		if dir == "/" {
			totalUsedSpace = size
		}
	}

	// sort the list of directory sizes ascendingly
	sort.Slice(allDirSizes, func(i, j int) bool {
		return allDirSizes[i] < allDirSizes[j]
	})

	// calculate the amount of space we'd need to free up in order to meet the requirements
	spaceToFreeUp := diskSpaceNeeded - (diskSpaceAvailable - totalUsedSpace)

	// get the largest directory that we can delete to meet the space freeing requirements
	sizeOfDirectoryToDelete := 0
	for _, dirSize := range allDirSizes {
		diskSizeIsMoreThanRequired := dirSize > spaceToFreeUp
		if diskSizeIsMoreThanRequired {
			sizeOfDirectoryToDelete = dirSize
			break
		}
	}

	fmt.Println(sizeOfDirectoryToDelete)
}

func updateSizes(directories map[string]int, currentDir string, size int) {
	// update current directory size
	directories[currentDir] = directories[currentDir] + size

	// recursively update size up tree
	for currentDir != "/" {
		lastIndexOfSeparator := strings.LastIndex(currentDir[:len(currentDir)-1], "/")
		currentDir = currentDir[:lastIndexOfSeparator+1]
		directories[currentDir] = directories[currentDir] + size
	}
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
