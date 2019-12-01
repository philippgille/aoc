package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	result := 0
	values := strings.Split(input, "\n")
	for _, valueString := range values {
		value, err := strconv.Atoi(valueString)
		if err != nil {
			panic(err)
		}
		result += requiredFuel(value)
	}
	fmt.Println(result)
}

func requiredFuel(mass int) int {
	result := (mass / 3) - 2
	if result <= 0 {
		return 0
	}
	return result + requiredFuel(result)
}
