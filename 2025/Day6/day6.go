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

func copySlice(origSlice []string) []string {
	newSlice := make([]string, len(origSlice))
	copy(newSlice, origSlice)
	return newSlice
}

func totalCalculations(calculations [][]string) int {
	numCalculations := len(calculations[0])
	numOperators := len(calculations)

	summedCalculations := 0
	for i := range numCalculations {
		operator := calculations[numOperators-1][i]
		var calculationResult int

		switch operator {
		case "+":
			calculationResult = 0
		case "*":
			calculationResult = 1
		}

		for j := range numOperators - 1 {
			intVal, _ := strconv.Atoi(calculations[j][i])
			switch operator {
			case "+":
				calculationResult += intVal
			case "*":
				calculationResult *= intVal
			}
		}
		summedCalculations += calculationResult
	}
	return summedCalculations
}

func checkBlankColumn(stringCalculations []string, index int) bool {
	for _, line := range stringCalculations {
		if string(line[index]) != " " {
			return false
		}
	}
	return true
}

func constructColumnCalculations(stringCalculations []string) [][]string {
	numOperators := len(stringCalculations)
	var constructedCalculations [][]string
	var constructedValues []string
	operator := string(stringCalculations[numOperators-1][0])

	for i := 0; i < len(stringCalculations[0]); i++ {
		if checkBlankColumn(stringCalculations, i) {
			constructedValues = append(constructedValues, operator)
			constructedCalculations = append(constructedCalculations, copySlice(constructedValues))

			operator = string(stringCalculations[numOperators-1][i+1])
			constructedValues = constructedValues[:0]
			continue
		}

		newVal := ""
		for j := 0; j < numOperators-1; j++ {
			char := string(stringCalculations[j][i])
			if char != " " {
				newVal += char
			}
		}
		constructedValues = append(constructedValues, newVal)
	}

	constructedValues = append(constructedValues, operator)
	constructedCalculations = append(constructedCalculations, copySlice(constructedValues))
	return constructedCalculations
}

func calculateResult(calculation []string) int {
	operator := calculation[len(calculation)-1]
	var totalValue int

	switch operator {
	case "+":
		totalValue = 0
	case "*":
		totalValue = 1
	}

	for _, value := range calculation[:len(calculation)-1] {
		intVal, _ := strconv.Atoi(value)
		switch operator {
		case "+":
			totalValue += intVal

		case "*":
			totalValue *= intVal
		}
	}
	return totalValue
}

func main() {
	// SETUP
	stringCalculations := strings.Split(getFileInput("day6_input.txt"), "\n")
	stringCalculations = stringCalculations[:len(stringCalculations)-1]
	var calculations [][]string

	for _, line := range stringCalculations {
		calculations = append(calculations, strings.Fields(line))
	}

	// PART 1
	fmt.Printf("PART 1 TOTAL CALCULATIONS: %d\n", totalCalculations(calculations))

	// PART 2
	constructedCalculations := constructColumnCalculations(stringCalculations)
	totalValue := 0
	for _, calculation := range constructedCalculations {
		totalValue += calculateResult(calculation)
	}
	fmt.Printf("PART 2 TOTAL CALCULATIONS: %d\n", totalValue)
}
