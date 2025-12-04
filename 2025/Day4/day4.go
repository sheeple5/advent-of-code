package main

import (
	"fmt"
	"os"
	"strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func getFileInput(fileName string) string {
	dat, err := os.ReadFile(fileName)
	check(err)

	return string(dat)
}

func findNextRoll(board []string, currentRoll []int) []int {
	index := -1
	for index == -1 {
		if currentRoll[1]+1 >= len(board) {
			currentRoll[0] += 1
			currentRoll[1] = -1
		}

		if currentRoll[0] >= len(board) {
			return []int{}
		}

		// fmt.Printf("%d - %d\n", currentRoll[0], currentRoll[1])
		index = strings.IndexRune(board[currentRoll[0]][currentRoll[1]+1:], '@')

		if index != -1 {
			return []int{currentRoll[0], currentRoll[1] + index + 1}
		} else {
			currentRoll[0] += 1
			currentRoll[1] = -1
		}
	}
	return []int{}
}

func calcAdjRolls(board []string, currentRoll []int) int {
	adjRolls := 0
	directions := [][]int{{0, 1}, {1, 1}, {1, 0}, {1, -1}, {0, -1}, {-1, -1}, {-1, 0}, {-1, 1}}
	for _, direction := range directions {
		checkRoll := []int{currentRoll[0] + direction[0], currentRoll[1] + direction[1]}
		if checkInBounds(board, checkRoll) {
			if board[checkRoll[0]][checkRoll[1]] == '@' {
				adjRolls += 1
			}
		}
	}
	return adjRolls
}

func checkInBounds(board []string, currentRoll []int) bool {
	if currentRoll[0] >= 0 && currentRoll[0] < len(board) && currentRoll[1] >= 0 && currentRoll[1] < len(board) {
		return true
	} else {
		return false
	}
}

func main() {
	// SETUP
	board := strings.Split(strings.TrimSpace(getFileInput("day4_input.txt")), "\n")

	// PART 1
	accessibleRolls := 0
	currentRoll := findNextRoll(board, []int{0, 0})
	for len(currentRoll) != 0 {
		if calcAdjRolls(board, currentRoll) < 4 {
			accessibleRolls += 1
			fmt.Println(currentRoll)
		}
		currentRoll = findNextRoll(board, currentRoll)
	}
	fmt.Println(accessibleRolls)
}
