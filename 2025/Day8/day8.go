package main

import (
	"fmt"
	"math"
	"os"
	"sort"
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

type Box struct {
	ID        int
	X         int
	Y         int
	Z         int
	CircuitID int
}

type Circuit struct {
	ID    int
	Boxes []*Box
}

func boxDistance(box1 *Box, box2 *Box) float64 {
	return math.Sqrt(math.Pow(float64(box1.X-box2.X), 2) + math.Pow(float64(box1.Y-box2.Y), 2) + math.Pow(float64(box1.Z-box2.Z), 2))
}

func updateIDs(boxes []*Box, id int) {
	for _, box := range boxes {
		box.CircuitID = id
	}
}

func updateCircuits(box1 *Box, box2 *Box, circuitMap map[int]*Circuit, circuitID *int) {
	if box1.CircuitID == 0 && box2.CircuitID == 0 {
		newCircuit := Circuit{ID: *circuitID, Boxes: []*Box{box1, box2}}
		box1.CircuitID = *circuitID
		box2.CircuitID = *circuitID
		circuitMap[*circuitID] = &newCircuit
		*circuitID += 1
	} else if box1.CircuitID != 0 && box2.CircuitID == 0 {
		circuit := circuitMap[box1.CircuitID]
		circuit.Boxes = append(circuit.Boxes, box2)
		box2.CircuitID = box1.CircuitID
	} else if box1.CircuitID == 0 && box2.CircuitID != 0 {
		circuit := circuitMap[box2.CircuitID]
		circuit.Boxes = append(circuit.Boxes, box1)
		box1.CircuitID = box2.CircuitID
	} else if box1.CircuitID != box2.CircuitID {
		circuit1 := circuitMap[box1.CircuitID]
		circuit2 := circuitMap[box2.CircuitID]
		circuit1.Boxes = append(circuit1.Boxes, circuit2.Boxes...)

		updateIDs(circuit2.Boxes, circuit1.ID)
		delete(circuitMap, circuit2.ID)
	}
}

func main() {
	// SETUP
	boxCoords := strings.Split(getFileInput("day8_input.txt"), "\n")
	boxCoords = boxCoords[:len(boxCoords)-1]

	boxMap := make(map[int]*Box)
	for i, coord := range boxCoords {
		splitCoords := strings.Split(coord, ",")
		xInt, _ := strconv.Atoi(splitCoords[0])
		yInt, _ := strconv.Atoi(splitCoords[1])
		zInt, _ := strconv.Atoi(splitCoords[2])

		newBox := Box{ID: i, X: xInt, Y: yInt, Z: zInt}

		boxMap[i] = &newBox
	}

	// PART 1
	distanceMap := make(map[float64][]*Box)
	var distances []float64
	for i := range len(boxMap) {
		for j := i + 1; j < len(boxMap); j++ {
			distance := boxDistance(boxMap[i], boxMap[j])

			distanceMap[distance] = []*Box{boxMap[i], boxMap[j]}
			distances = append(distances, distance)
		}
	}

	sort.Float64s(distances)

	circuitMap := make(map[int]*Circuit)
	circuitID := 1
	for _, distance := range distances[:1000] {
		boxes := distanceMap[distance]
		box1 := boxes[0]
		box2 := boxes[1]

		updateCircuits(box1, box2, circuitMap, &circuitID)
	}

	var circutLengths []int
	for _, circuit := range circuitMap {
		circutLengths = append(circutLengths, len(circuit.Boxes))
	}
	sort.Ints(circutLengths)

	finalValue := circutLengths[len(circutLengths)-1] * circutLengths[len(circutLengths)-2] * circutLengths[len(circutLengths)-3]
	fmt.Printf("PART 1 CIRCUITS VALUE: %d\n", finalValue)

	// PART 2
	circuitMap = make(map[int]*Circuit)
	circuitID = 1
	for _, box := range boxMap {
		box.CircuitID = 0
	}

	for _, distance := range distances {
		boxes := distanceMap[distance]
		box1 := boxes[0]
		box2 := boxes[1]

		updateCircuits(box1, box2, circuitMap, &circuitID)

		breakLoop := false
		if len(circuitMap) == 1 {
			for _, circuit := range circuitMap {
				if len(circuit.Boxes) == len(boxMap) {
					fmt.Printf("PART 2 CIRCUITS VALUE: %d\n", box1.X*box2.X)
					breakLoop = true
				}
			}
		}

		if breakLoop {
			break
		}
	}
}
