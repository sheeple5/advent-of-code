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

func copySlice(origSlice []string) []string {
	newSlice := make([]string, len(origSlice))
	copy(newSlice, origSlice)
	return newSlice
}

type Device struct {
	ID    string
	Edges []string
}

func findNumPathes(currentDevice string, visitedDevices []string, deviceMap map[string]*Device, start string) int {
	if currentDevice == "out" {
		switch start {
		case "you":
			return 1
		case "svr":
			if slices.Contains(visitedDevices, "fft") && slices.Contains(visitedDevices, "dac") {
				return 1
			} else {
				return 0
			}
		}
	}

	paths := 0
	for _, edge := range deviceMap[currentDevice].Edges {
		if !slices.Contains(visitedDevices, edge) {
			newVisited := copySlice(visitedDevices)
			newVisited = append(newVisited, edge)
			paths += findNumPathes(edge, newVisited, deviceMap, start)
		}
	}
	return paths
}

func main() {
	// SETUP
	deviceList := strings.Split(getFileInput("day11_input.txt"), "\n")
	deviceList = deviceList[:len(deviceList)-1]

	deviceMap := make(map[string]*Device)
	for _, device := range deviceList {
		deviceID := strings.Split(device, ": ")[0]
		deviceEdges := strings.Split(strings.Split(device, ": ")[1], " ")
		deviceMap[deviceID] = &Device{ID: deviceID, Edges: deviceEdges}
	}

	// PART 1
	fmt.Printf("PART 1 NUM PATHS: %d\n", findNumPathes("you", []string{"you"}, deviceMap, "you"))

	// PART 2 - Currently doesn't work due to optimizations. Need to use a cache that can combine paths to see if total path has both fft and dac
	// So maybe have a path cache for three parts: number of paths that contain just dac, number of paths that contain fft, and both
	// Then combine that knowledge with current path (i.e. if cached result has a dac path and current step has fft, can use the dac path cache count)
	fmt.Printf("PART 2 NUM PATHS: %d\n", findNumPathes("svr", []string{"svr"}, deviceMap, "svr"))
}
