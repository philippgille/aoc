package main

import (
	"strings"
	"strconv"
	"fmt"
)

func main() {
	result := 0
	values := strings.Split(input, "\n")
	for _, valueString := range values{
		value, err := strconv.Atoi(valueString)
		if err != nil{
			panic(err)
		}
		result += requiredFuel(value)
	}
	fmt.Println(result)
}

func requiredFuel(mass int) int {
	return (mass / 3) - 2
}
