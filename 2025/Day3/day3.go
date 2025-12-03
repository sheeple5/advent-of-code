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

func getLargestJoltage(bank string) int {
	largestMap := findLargestVal(bank[:len(bank)-1])
	secondLargestMap := findLargestVal(bank[largestMap["index"]+1:])
	stringFirst := strconv.Itoa(largestMap["value"])
	stringSecond := strconv.Itoa(secondLargestMap["value"])
	mergedVal, _ := strconv.Atoi(stringFirst + stringSecond)
	return mergedVal
}

func main() {
	banks := strings.Split(getFileInput(("day3_input.txt")), "\n")
	totalJoltage := 0
	for _, bank := range banks[:len(banks)-1] {
		totalJoltage += getLargestJoltage(bank)
	}
	fmt.Printf("PART 1 TOTAL JOLTAGE: %d\n", totalJoltage)
}
