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

type Machine struct {
	TargetLights string
	Buttons      []*Button
}

type Button struct {
	Indexes []int
	ButtonFuncs
}

type ButtonFuncs interface {
	pushButton()
}

func (button *Button) pushButton(currentLights string) string {
	for _, i := range button.Indexes {
		if string(currentLights[i]) == "." {
			currentLights = currentLights[:i] + "#" + currentLights[i+1:]
		} else {
			currentLights = currentLights[:i] + "." + currentLights[i+1:]
		}
	}
	return currentLights
}

func pushButtons(machine *Machine, currentLights string, cutoff int) bool {
	savedLights := currentLights
	for _, button := range machine.Buttons {
		currentLights = button.pushButton(currentLights)

		if machine.TargetLights == currentLights {
			return true
		} else if cutoff != 0 {
			if pushButtons(machine, currentLights, cutoff-1) {
				return true
			}
		}

		currentLights = savedLights
	}
	return false
}

func findMinPushes(machine *Machine) int {
	pushes := 1

	for {
		if pushButtons(machine, strings.Repeat(".", len(machine.TargetLights)), pushes-1) {
			return pushes
		} else {
			pushes += 1
		}
	}
}

func main() {
	// SETUP
	machineList := strings.Split(getFileInput("day10_input.txt"), "\n")
	machineList = machineList[:len(machineList)-1]

	var machines []*Machine
	for _, machineString := range machineList {
		machineChunks := strings.Split(machineString, " ")
		machineLights := machineChunks[0][1 : len(machineChunks[0])-1]

		newMachine := Machine{TargetLights: machineLights}
		for _, chunk := range machineChunks[1 : len(machineChunks)-1] {
			stringPositions := strings.Split(chunk[1:len(chunk)-1], ",")
			var positions []int

			for _, stringPosition := range stringPositions {
				intPosition, _ := strconv.Atoi(stringPosition)
				positions = append(positions, intPosition)
			}

			newButton := Button{Indexes: positions}
			newMachine.Buttons = append(newMachine.Buttons, &newButton)
		}

		machines = append(machines, &newMachine)
	}

	// PART 1
	totalPushes := 0
	for _, machine := range machines {
		totalPushes += findMinPushes(machine)
	}
	fmt.Printf("PART 1 TOTAL PUSHES: %d\n", totalPushes)
}
