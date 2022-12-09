package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

// Day 8: Treetop Tree House
func main() {
	// open from the input file
	fileScanner := getFileScanner("day8input.txt")

	matrixSize := 99
	treeArrangement := getMatrix(matrixSize)
	mostVisibility := 0

	i := 0
	// read line by line
	for fileScanner.Scan() {
		line := fileScanner.Text()

		// build the treeArrangement matrix
		for index, ch := range line {
			treeHeight, _ := strconv.Atoi(string(ch))
			treeArrangement[i][index] = treeHeight
		}
		i++
	}

	// check tree to see which offers the most visibility
	for i := 0; i < matrixSize; i++ {
		for j := 0; j < matrixSize; j++ {
			currentTreeHeight := treeArrangement[i][j]

			// check top
			counterTreesTop := 0
			for k := i - 1; k >= 0; k-- {
				if treeArrangement[k][j] < currentTreeHeight {
					counterTreesTop++
				} else {
					counterTreesTop++
					break
				}
			}

			// check bottom
			counterTreesBottom := 0
			for k := i + 1; k < matrixSize; k++ {
				if treeArrangement[k][j] < currentTreeHeight {
					counterTreesBottom++
				} else {
					counterTreesBottom++
					break
				}
			}

			// check left
			counterTreesLeft := 0
			for k := j - 1; k >= 0; k-- {
				if treeArrangement[i][k] < currentTreeHeight {
					counterTreesLeft++
				} else {
					counterTreesLeft++
					break
				}
			}

			// check right
			counterTreesRight := 0
			for k := j + 1; k < matrixSize; k++ {
				if treeArrangement[i][k] < currentTreeHeight {
					counterTreesRight++
				} else {
					counterTreesRight++
					break
				}
			}

			scenicScore := counterTreesTop * counterTreesBottom * counterTreesLeft * counterTreesRight
			if scenicScore > mostVisibility {
				mostVisibility = scenicScore
			}

			// reset tree counters
			counterTreesBottom = 0
			counterTreesRight = 0
			counterTreesTop = 0
			counterTreesLeft = 0
		}
	}

	fmt.Println(mostVisibility)
}

func getMatrix(size int) [][]int {
	var matrix = make([][]int, size)
	for i := range matrix {
		matrix[i] = make([]int, size)
	}
	return matrix
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
