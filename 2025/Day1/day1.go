package main

import (
	"fmt"
	"os"
	"strconv"
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

func abs(val int) int {
	if val < 0 {
		return val * -1
	}
	return val
}

func main() {
	// SETUP
	moves := strings.Split(getFileInput("day1_input.txt"), "\n")

	// PART 1
	currentVal := 50
	zeroCount := 0

	for _, move := range moves {
		if move == "" {
			continue
		}

		direction := move[0]
		distance, err := strconv.Atoi(move[1:])
		check(err)

		if direction == 'L' {
			distance = distance * -1
		}

		currentVal = (currentVal + distance) % 100
		if currentVal < 0 {
			currentVal = 100 + currentVal
		}

		if currentVal == 0 {
			zeroCount += 1
		}
	}
	fmt.Printf("PART 1 ZERO COUNT: %d\n", zeroCount)

	// PART 2
	currentVal = 50
	zeroCount = 0

	for _, move := range moves {
		if move == "" {
			continue
		}

		direction := move[0]
		distance, err := strconv.Atoi(move[1:])
		check(err)

		if direction == 'L' {
			if currentVal != 0 {
				zeroCount += (100 - currentVal + distance) / 100
			} else {
				zeroCount += distance / 100
			}
			distance = distance * -1
		} else {
			zeroCount += (currentVal + distance) / 100
		}

		currentVal = (currentVal + distance) % 100
		if currentVal < 0 {
			currentVal = 100 + currentVal
		}

	}
	fmt.Printf("PART 2 ZERO COUNT: %d\n", zeroCount)
}
