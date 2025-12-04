package main

import (
	"fmt"
	"os"
	"slices"
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

func removeRoll(board []string, currentRoll []int) {
	currentRow := board[currentRoll[0]]
	newRow := currentRow[:currentRoll[1]] + string('.') + currentRow[currentRoll[1]+1:]
	board[currentRoll[0]] = newRow
}

func initializeRoll(board []string) []int {
	currentRoll := make([]int, 2)
	if board[0][0] == '@' {
		currentRoll[0] = 0
		currentRoll[1] = 0
	} else {
		firstRoll := findNextRoll(board, []int{0, 0})
		currentRoll[0] = firstRoll[0]
		currentRoll[1] = firstRoll[1]
	}
	return currentRoll
}

func copyBoard(board []string) []string {
	newBoard := make([]string, len(board))
	copy(newBoard, board)
	return newBoard
}

func main() {
	// SETUP
	board := strings.Split(strings.TrimSpace(getFileInput("day4_input.txt")), "\n")

	// PART 1
	accessibleRolls := 0
	currentRoll := initializeRoll(board)

	for len(currentRoll) != 0 {
		if calcAdjRolls(board, currentRoll) < 4 {
			accessibleRolls += 1
		}
		currentRoll = findNextRoll(board, currentRoll)
	}
	fmt.Printf("PART 1 ACCESSIBLE ROLLS: %d\n", accessibleRolls)

	// PART 2
	accessibleRolls = 0
	newBoard := copyBoard(board)

	for {
		currentRoll := initializeRoll(board)
		for len(currentRoll) != 0 {
			if calcAdjRolls(board, currentRoll) < 4 {
				accessibleRolls += 1
				removeRoll(newBoard, currentRoll)
			}
			currentRoll = findNextRoll(board, currentRoll)
		}

		if slices.Equal(board, newBoard) {
			break
		}

		board = copyBoard(newBoard)

	}
	fmt.Printf("PART 2 ACCESSIBLE ROLLS: %d\n", accessibleRolls)
}
