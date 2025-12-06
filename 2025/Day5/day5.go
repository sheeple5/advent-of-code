package main

import (
	"fmt"
	"os"
	"slices"
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

func checkInRange(idRange string, id int) bool {
	lowID, _ := strconv.Atoi(strings.Split(idRange, "-")[0])
	highID, _ := strconv.Atoi(strings.Split(idRange, "-")[1])

	if id >= lowID && id <= highID {
		return true
	} else {
		return false
	}
}

func calculateRangeCount(idRange string) int {
	lowID, _ := strconv.Atoi(strings.Split(idRange, "-")[0])
	highID, _ := strconv.Atoi(strings.Split(idRange, "-")[1])

	return highID - lowID + 1
}

func copySlice(origSlice []string) []string {
	newSlice := make([]string, len(origSlice))
	copy(newSlice, origSlice)
	return newSlice
}

func mergeAllRanges(idRanges []string) []string {
	for {
		setRanges := copySlice(idRanges)
		superBreak := false
		for i, idRange1 := range idRanges {
			for j, idRange2 := range idRanges {
				if idRange1 == idRange2 && i != j {
					setRanges = slices.Delete(setRanges, i, i+1)

					if i < j {
						setRanges = slices.Delete(setRanges, j-1, j)
					} else {
						setRanges = slices.Delete(setRanges, j, j+1)
					}
					setRanges = append(setRanges, idRange1)
					superBreak = true
					break

				}
				mergedRange := mergeRanges(idRange1, idRange2)
				if mergedRange != "" {
					setRanges = slices.Delete(setRanges, i, i+1)

					if i < j {
						setRanges = slices.Delete(setRanges, j-1, j)
					} else {
						setRanges = slices.Delete(setRanges, j, j+1)
					}
					setRanges = append(setRanges, mergedRange)
					superBreak = true
					break
				}
			}
			if superBreak {
				break
			}
		}
		if slices.Equal(idRanges, setRanges) {
			return idRanges
		} else {
			idRanges = setRanges
		}
	}
}

func mergeRanges(idRange1 string, idRange2 string) string {
	lowID1, _ := strconv.Atoi(strings.Split(idRange1, "-")[0])
	highID1, _ := strconv.Atoi(strings.Split(idRange1, "-")[1])
	lowID2, _ := strconv.Atoi(strings.Split(idRange2, "-")[0])
	highID2, _ := strconv.Atoi(strings.Split(idRange2, "-")[1])

	if highID1 < lowID2 || lowID1 > highID2 {
		return ""
	}

	if lowID1 >= lowID2 && highID1 <= highID2 {
		return ""
	} else if lowID1 < lowID2 && highID1 <= highID2 {
		return fmt.Sprintf("%d-%d", lowID1, highID2)
	} else if lowID1 >= lowID2 && highID1 > highID2 {
		return fmt.Sprintf("%d-%d", lowID2, highID1)
	} else if lowID1 < lowID2 && highID1 > highID2 {
		return fmt.Sprintf("%d-%d", lowID1, highID1)
	}
	return ""
}

func main() {
	// SETUP
	lines := strings.Split(getFileInput("day5_input.txt"), "\n")
	var idRanges []string
	var idList []int
	switchType := false
	for _, line := range lines {
		if line == "" {
			switchType = true
			continue
		}
		if !switchType {
			idRanges = append(idRanges, line)
		} else {
			intID, _ := strconv.Atoi(line)
			idList = append(idList, intID)
		}
	}

	// PART 1
	idCount := 0
	for _, id := range idList {
		for _, idRange := range idRanges {
			if checkInRange(idRange, id) {
				idCount += 1
				break
			}
		}
	}
	fmt.Printf("PART 1 FRESH IDS: %d\n", idCount)

	// PART 2
	idCount = 0
	mergedRanges := mergeAllRanges(idRanges)
	for _, idRange := range mergedRanges {
		idCount += calculateRangeCount(idRange)
	}

	fmt.Printf("PART 2 FRESH IDS: %d\n", idCount)
}
