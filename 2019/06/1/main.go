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
	result := 0
	// For all objects, find their distance to the COM (center of mass)
	for obj := range orbitMap {
		tempObj := obj
		// Follows the "dependencies" until `center` is "COM", which will lead to ok == false.
		for center, ok := orbitMap[tempObj]; ok; center, ok = orbitMap[tempObj] {
			result++
			tempObj = center
		}
	}
	return result
}
