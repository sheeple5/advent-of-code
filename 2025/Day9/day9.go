package main

import (
	"fmt"
	"math"
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

func convertToInt(coord string) []int {
	xCoord, _ := strconv.Atoi(strings.Split(coord, ",")[0])
	yCoord, _ := strconv.Atoi(strings.Split(coord, ",")[1])

	return []int{xCoord, yCoord}
}

func calculateArea(point1 []int, point2 []int) int {
	return int((math.Abs(float64(point1[0]-point2[0])) + 1) * (math.Abs(float64(point1[1]-point2[1])) + 1))
}

func checkValid(point1 []int, point2 []int, redTiles []string) bool {
	// point3 := []int{point1[0], point2[1]}
	// point4 := []int{point2[0], point1[1]}

	// Do some kind of check if point3 and point4 equal 1 and to, i.e. the 2,3 and 7,3 case
	// To go clockwise, point order will be:
	// - point1, point4, point2, point3 if top left / bottom right rectangle (if point with smaller x has smaller y)
	// - point1, point3, point2, point4 if top right / bottom left rectangle (if point with smaller x has greater y)

	return true
}

func main() {
	// SETUP
	redTiles := strings.Split(getFileInput("day9_input.txt"), "\n")
	redTiles = redTiles[:len(redTiles)-1]

	// PART 1
	maxArea := 0
	for i := range len(redTiles) {
		for j := i + 1; j < len(redTiles); j++ {
			point1 := convertToInt(redTiles[i])
			point2 := convertToInt(redTiles[j])

			calculatedArea := calculateArea(point1, point2)
			if calculatedArea > maxArea {
				maxArea = calculatedArea
			}
		}
	}
	fmt.Printf("PART 1 LARGEST AREA: %d\n", maxArea)

	// PART 2
	// 	maxArea = 0
	// 	for i := range len(redTiles) {
	// 		for j := i + 1; j < len(redTiles); j++ {
	// 			point1 := convertToInt(redTiles[i])
	// 			point2 := convertToInt(redTiles[j])
	//
	// 			if checkValid(point1, point2, redTiles) {
	// 				calculatedArea := calculateArea(point1, point2)
	// 				if calculatedArea > maxArea {
	// 					maxArea = calculatedArea
	// 				}
	// 			}
	// 		}
	// 	}
	// 	fmt.Printf("PART 1 LARGEST AREA: %d\n", maxArea)
}
