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
	ID      int
	X       int
	Y       int
	Z       int
	Edges   map[int]float64
	Circuit int
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
		box.Circuit = id
	}
}

func main() {
	// SETUP
	boxCoords := strings.Split(getFileInput("day8_input.txt"), "\n")
	boxCoords = boxCoords[:len(boxCoords)-1]

	var boxes []*Box
	boxMap := make(map[int]*Box)
	for i, coord := range boxCoords {
		splitCoords := strings.Split(coord, ",")
		xInt, _ := strconv.Atoi(splitCoords[0])
		yInt, _ := strconv.Atoi(splitCoords[1])
		zInt, _ := strconv.Atoi(splitCoords[2])

		newBox := Box{ID: i, X: xInt, Y: yInt, Z: zInt, Edges: make(map[int]float64)}

		boxes = append(boxes, &newBox)
		boxMap[i] = &newBox
	}

	// PART 1
	distanceMap := make(map[float64]string)
	var distances []float64
	for i := range boxes {
		for j := i + 1; j < len(boxes); j++ {
			distance := boxDistance(boxMap[i], boxMap[j])
			boxMap[i].Edges[j] = distance
			boxMap[j].Edges[i] = distance

			distanceMap[distance] = fmt.Sprintf("%d-%d", i, j)
			distances = append(distances, distance)
		}
	}

	sort.Float64s(distances)

	circuitMap := make(map[int]*Circuit)
	circuitID := 1
	for _, distance := range distances[:1000] {
		boxIDs := distanceMap[distance]
		boxID1, _ := strconv.Atoi(strings.Split(boxIDs, "-")[0])
		boxID2, _ := strconv.Atoi(strings.Split(boxIDs, "-")[1])

		box1 := boxMap[boxID1]
		box2 := boxMap[boxID2]

		if box1.Circuit == 0 && box2.Circuit == 0 {
			newCircuit := Circuit{ID: circuitID, Boxes: []*Box{box1, box2}}
			box1.Circuit = circuitID
			box2.Circuit = circuitID
			circuitMap[circuitID] = &newCircuit
			circuitID += 1
		} else if box1.Circuit != 0 && box2.Circuit == 0 {
			circuit := circuitMap[box1.Circuit]
			circuit.Boxes = append(circuit.Boxes, box2)
			box2.Circuit = box1.Circuit
		} else if box1.Circuit == 0 && box2.Circuit != 0 {
			circuit := circuitMap[box2.Circuit]
			circuit.Boxes = append(circuit.Boxes, box1)
			box1.Circuit = box2.Circuit
		} else if box1.Circuit != box2.Circuit {
			circuit1 := circuitMap[box1.Circuit]
			circuit2 := circuitMap[box2.Circuit]
			circuit1.Boxes = append(circuit1.Boxes, circuit2.Boxes...)

			updateIDs(circuit2.Boxes, circuit1.ID)
			delete(circuitMap, circuit2.ID)
		}
	}

	var circutLengths []int
	for _, circuit := range circuitMap {
		circutLengths = append(circutLengths, len(circuit.Boxes))
	}
	sort.Ints(circutLengths)

	finalValue := circutLengths[len(circutLengths)-1] * circutLengths[len(circutLengths)-2] * circutLengths[len(circutLengths)-3]
	fmt.Printf("PART 1 CIRCUITS VALUE: %d\n", finalValue)
}
