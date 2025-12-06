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

func formatCalculations(stringCalculations []string) [][]string {
	var calculations [][]string
	for _, line := range stringCalculations {
		calculations = append(calculations, strings.Fields(line))
	}

	var formattedCalculations [][]string
	var formattedCalculation []string
	numOperators := len(calculations)
	numCalculations := len(calculations[0])
	for i := range numCalculations {
		for j := range numOperators {
			formattedCalculation = append(formattedCalculation, calculations[j][i])
		}
		formattedCalculations = append(formattedCalculations, copySlice(formattedCalculation))
		formattedCalculation = formattedCalculation[:0]
	}
	return formattedCalculations
}

func constructColumnCalculations(stringCalculations []string) [][]string {
	numOperators := len(stringCalculations)
	numCalculations := len(stringCalculations[0])
	operator := string(stringCalculations[numOperators-1][0])
	var constructedCalculations [][]string
	var constructedValues []string

	for i := range numCalculations {
		if checkBlankColumn(stringCalculations, i) {
			constructedValues = append(constructedValues, operator)
			constructedCalculations = append(constructedCalculations, copySlice(constructedValues))

			operator = string(stringCalculations[numOperators-1][i+1])
			constructedValues = constructedValues[:0]
			continue
		}

		newVal := ""
		for j := range numOperators - 1 {
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

func checkBlankColumn(stringCalculations []string, index int) bool {
	for _, line := range stringCalculations {
		if string(line[index]) != " " {
			return false
		}
	}
	return true
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

	// PART 1
	formattedCalculations := formatCalculations(stringCalculations)
	totalValue := 0
	for _, calculation := range formattedCalculations {
		totalValue += calculateResult(calculation)
	}
	fmt.Printf("PART 1 TOTAL CALCULATIONS: %d\n", totalValue)

	// PART 2
	constructedCalculations := constructColumnCalculations(stringCalculations)
	totalValue = 0
	for _, calculation := range constructedCalculations {
		totalValue += calculateResult(calculation)
	}
	fmt.Printf("PART 2 TOTAL CALCULATIONS: %d\n", totalValue)
}
