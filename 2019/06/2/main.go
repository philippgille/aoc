package main

import (
	"fmt"
	"strings"
)

func main() {
	orbitMap := convertInput(input)
	result := steps(orbitMap)
	fmt.Println(result)
}

func convertInput(in string) map[string]string {
	objStrings := strings.Split(in, "\n")
	result := make(map[string]string, len(objStrings))
	for _, objString := range objStrings {
		parts := strings.Split(objString, ")")
		center := parts[0]
		object := parts[1]
		result[object] = center
	}
	return result
}

func steps(orbitMap map[string]string) int {
	// For "YOU", record the number of steps for each center
	youObj := "YOU"
	youSteps := -1
	youMap := make(map[string]int)
	// Follows the "dependencies" until `center` is "COM", which will lead to ok == false.
	for center, ok := orbitMap[youObj]; ok; center, ok = orbitMap[youObj] {
		youSteps++
		youMap[youObj] = youSteps
		youObj = center
	}
	// For "SAN", go through the orbit until one matches a key in the youMap.
	sanObj := "SAN"
	sanSteps := -1
	for center, ok := orbitMap[sanObj]; ok; center, ok = orbitMap[sanObj] {
		sanSteps++
		if youSteps, ok := youMap[center]; ok{
			return youSteps + sanSteps - 1
		}
		sanObj = center
	}
	panic("should have returned earlier")
}
