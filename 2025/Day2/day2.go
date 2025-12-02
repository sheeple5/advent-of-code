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

func processRange(idRange string, part string) int {
	firstID, _ := strconv.Atoi(strings.Split(idRange, "-")[0])
	secondID, _ := strconv.Atoi(strings.TrimSpace(strings.Split(idRange, "-")[1]))

	invalidSum := 0
	for id := firstID; id <= secondID; id++ {
		switch part {
		case "PART 1":
			if mirrorInvalidID(id) {
				invalidSum += id
			}
		case "PART 2":
			if repetitionInvalidID(id) {
				invalidSum += id
			}
		}
	}

	return invalidSum
}

func mirrorInvalidID(id int) bool {
	stringID := strconv.Itoa(id)

	if len(stringID)%2 != 0 {
		return false
	} else {
		idHalfLen := len(stringID) / 2
		idFirstHalf := stringID[:idHalfLen]
		idSecondHalf := stringID[idHalfLen:]

		if idFirstHalf == idSecondHalf {
			return true
		} else {
			return false
		}
	}
}

func repetitionInvalidID(id int) bool {
	stringID := strconv.Itoa(id)
	var sequence string

	for _, char := range stringID {
		sequence += string(char)
		sequenceLen := len(sequence)
		if sequenceLen > len(stringID)/2 {
			return false
		}

		invalid := true
		for i := 0; i < len(stringID); i += sequenceLen {
			if i+sequenceLen > len(stringID) {
				invalid = false
				break
			}
			if stringID[i:i+sequenceLen] != sequence {
				invalid = false
				break
			}
		}

		if invalid {
			return true
		}

	}
	return false
}

func main() {
	// SETUP
	idRanges := strings.Split(getFileInput("day2_input.txt"), ",")

	// PART 1
	totalInvalidSum := 0
	for _, idRange := range idRanges {
		totalInvalidSum += processRange(idRange, "PART 1")
	}
	fmt.Printf("PART 1 INVALID ID SUM: %d\n", totalInvalidSum)

	// PART 2
	totalInvalidSum = 0
	for _, idRange := range idRanges {
		totalInvalidSum += processRange(idRange, "PART 2")
	}
	fmt.Printf("PART 2 INVALID ID SUM: %d\n", totalInvalidSum)
}
