package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	in := convertInput(input)
	in = fixInput(in)
	out := compute(in)
	fmt.Println(out)
}

func convertInput(in string) []int {
	valStrings := strings.Split(in, ",")
	result := make([]int, len(valStrings))
	for i, valString := range valStrings {
		var err error
		result[i], err = strconv.Atoi(valString)
		if err != nil {
			panic(err)
		}
	}
	return result
}

func fixInput(in []int) []int {
	in[1] = 12
	in[2] = 2
	return in
}

func compute(in []int) int {
	var op, x, y int
	for pos, val := range in {
		switch pos % 4 {
		case 0:
			// val is an opcode
			op = val
			if op == 99 {
				return in[0]
			}
		case 1:
			// val points to x
			x = in[val]
		case 2:
			// val points to y
			y = in[val]
		case 3:
			// val points to the target
			// Store the computed value, depending on the opcode
			switch op {
			case 1:
				in[val] = x + y
			case 2:
				in[val] = x * y
			default:
				panic("unknown opcode")
			}
		}
	}
	return in[0]
}
