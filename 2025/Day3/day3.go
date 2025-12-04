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

func findLargestVal(bank string) map[string]int {
	largestMap := make(map[string]int)
	for i := 9; i >= 0; i-- {
		index := strings.IndexRune(bank, rune('0'+i))

		if index != -1 {
			largestMap["index"] = index
			largestMap["value"] = i
			break
		}
	}
	return largestMap
}

func getLargestJoltage(bank string, digits int) int {
	var constructedVal string
	digitRange := digits - 1
	previousIndex := 0

	for range digits {
		largestMap := findLargestVal(bank[previousIndex : len(bank)-digitRange])
		previousIndex += largestMap["index"] + 1
		digitRange -= 1
		constructedVal += strconv.Itoa(largestMap["value"])
	}

	largestJoltage, _ := strconv.Atoi(constructedVal)
	return largestJoltage
}

func main() {
	// SETUP
	banks := strings.Split(getFileInput(("day3_input.txt")), "\n")

	// PART 1
	totalJoltage := 0
	for _, bank := range banks[:len(banks)-1] {
		totalJoltage += getLargestJoltage(bank, 2)
	}
	fmt.Printf("PART 1 TOTAL JOLTAGE: %d\n", totalJoltage)

	// PART 2
	totalJoltage = 0
	for _, bank := range banks[:len(banks)-1] {
		totalJoltage += getLargestJoltage(bank, 12)
	}
	fmt.Printf("PART 2 TOTAL JOLTAGE: %d", totalJoltage)
}
