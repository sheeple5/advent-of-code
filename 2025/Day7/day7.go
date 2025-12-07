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

type Beam struct {
	Location []int
	Weight   int
	BeamFuncs
}

type BeamFuncs interface {
	moveBeam()
}

func (beam *Beam) moveBeam(board []string, splitCount *int) any {
	if beam.Location[0] == len(board)-1 {
		return nil
	}

	if string(board[beam.Location[0]+1][beam.Location[1]]) == "." {
		beam.Location = []int{beam.Location[0] + 1, beam.Location[1]}
		return nil
	} else if string(board[beam.Location[0]+1][beam.Location[1]]) == "^" {
		newBeam := Beam{Weight: beam.Weight, Location: []int{beam.Location[0] + 1, beam.Location[1] + 1}}
		beam.Location = []int{beam.Location[0] + 1, beam.Location[1] - 1}
		*splitCount += 1
		return newBeam
	}
	return nil
}

func findStart(board []string) []int {
	topLine := board[0]
	startingX := strings.IndexRune(topLine, 'S')

	return []int{0, startingX}
}

func assertBeam(beam any) Beam {
	if b, ok := beam.(Beam); ok {
		return b
	} else {
		panic("Failed to convert any to beam")
	}
}

func checkDupes(beams []*Beam) []*Beam {
	var newBeams []*Beam
	var seenLocations []string
	for _, beam := range beams {
		locationString := fmt.Sprintf("%d-%d", beam.Location[0], beam.Location[1])
		if !slices.Contains(seenLocations, locationString) {
			newBeams = append(newBeams, beam)
			seenLocations = append(seenLocations, locationString)
		} else {
			index := slices.Index(seenLocations, locationString)
			newBeams[index].Weight += beam.Weight
		}
	}
	return newBeams
}

func main() {
	// SETUP
	board := strings.Split(strings.TrimSpace(getFileInput("day7_input.txt")), "\n")

	// PART 1
	splitCount := 0
	startLocation := findStart(board)
	beams := []*Beam{{Weight: 1, Location: []int{startLocation[0], startLocation[1]}}}
	var newBeams []*Beam

	for {
		for _, beam := range beams {
			newBeam := beam.moveBeam(board, &splitCount)
			if newBeam != nil {
				assertedBeam := assertBeam(newBeam)
				newBeams = append(newBeams, &assertedBeam)
			}
		}

		beams = checkDupes(append(beams, newBeams...))
		newBeams = newBeams[:0]

		if beams[0].Location[0] == len(board)-1 {
			break
		}
	}
	fmt.Printf("PART 1 TOTAL SPLITS: %d\n", splitCount)

	// PART 2
	totalRealities := 0
	for _, beam := range beams {
		totalRealities += beam.Weight
	}
	fmt.Printf("PART 2 TOTAL REALITIES: %d\n", totalRealities)
}
