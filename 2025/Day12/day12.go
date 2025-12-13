package main

import (
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

type Present struct {
	ID        int
	Tile      [][]int
	Rotations [][][]int
}

type Tree struct {
	Region       []int
	PresentTypes []int
}

func rotateTile(tile [][]int) [][]int {
	rotatedTile := [][]int{{0, 0, 0}, {0, 0, 0}, {0, 0, 0}}

	rotatedTile[0][2] = tile[0][0]
	rotatedTile[2][2] = tile[0][2]
	rotatedTile[2][0] = tile[2][2]
	rotatedTile[0][0] = tile[2][0]

	rotatedTile[1][2] = tile[0][1]
	rotatedTile[2][1] = tile[1][2]
	rotatedTile[1][0] = tile[2][1]
	rotatedTile[0][1] = tile[1][0]

	rotatedTile[1][1] = tile[1][1]
	return rotatedTile
}

func stringToIntSlice(vals []string) []int {
	var intVals []int
	for _, val := range vals {
		intVal, _ := strconv.Atoi(val)
		intVals = append(intVals, intVal)
	}
	return intVals
}

func convertToBits(line string) []int {
	var bitSlice []int
	for _, char := range line {
		if char == '#' {
			bitSlice = append(bitSlice, 1)
		} else {
			bitSlice = append(bitSlice, 0)
		}
	}
	return bitSlice
}

func copySlice(origSlice []int) []int {
	newSlice := make([]int, len(origSlice))
	copy(newSlice, origSlice)
	return newSlice
}

func copyPresentBuild(presentBuild [][]int) [][]int {
	var newPresentBuild [][]int
	for _, bitSlice := range presentBuild {
		copyBitSlice := copySlice(bitSlice)
		newPresentBuild = append(newPresentBuild, copyBitSlice)
	}
	return newPresentBuild
}

func initializeRegion(tree Tree) [][]int {
	var regionBuild [][]int
	for range tree.Region[1] {
		newRow := make([]int, tree.Region[0])
		regionBuild = append(regionBuild, newRow)
	}
	return regionBuild
}

func presentsLeft(presents []int) bool {
	for _, present := range presents {
		if present > 0 {
			return true
		}
	}
	return false
}

func canFitPresents(presents []int, region [][]int, presentMap map[int]*Present) bool {
	if !presentsLeft(presents) {
		return true
	}
}

func main() {
	// SETUP
	presentData := strings.Split(getFileInput("day12_input.txt"), "\n")
	presentData = presentData[:len(presentData)-1]

	presentMap := make(map[int]*Present)
	var trees []*Tree
	var presentBuild [][]int
	var currentID int

	for _, line := range presentData {
		if line == "" {
			newPresent := Present{ID: currentID, Tile: copyPresentBuild(presentBuild)}
			presentMap[currentID] = &newPresent
			presentBuild = presentBuild[:0]
		} else if line[len(line)-1] == ':' {
			currentID, _ = strconv.Atoi(line[:len(line)-1])
		} else if strings.Contains(line, ":") {
			newTree := Tree{Region: stringToIntSlice(strings.Split(strings.Split(line, ": ")[0], "x")), PresentTypes: stringToIntSlice(strings.Split(strings.Split(line, ": ")[1], " "))}
			trees = append(trees, &newTree)
		} else {
			presentBuild = append(presentBuild, convertToBits(line))
		}
	}

	for _, present := range presentMap {
		present.Rotations = append(present.Rotations, present.Tile)
		present.Rotations = append(present.Rotations, rotateTile(present.Tile))
		present.Rotations = append(present.Rotations, rotateTile(present.Rotations[1]))
		present.Rotations = append(present.Rotations, rotateTile(present.Rotations[2]))
	}

	// PART 1
	totalValid := 0
	for _, tree := range trees {
		initialRegion := initializeRegion(*tree)
		if canFitPresents(tree.PresentTypes, initialRegion, presentMap) {
			totalValid += 1
		}
	}
}
