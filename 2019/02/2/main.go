package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	for noun := 0; noun <= 99; noun++{
		for verb := 0; verb <= 99; verb++{
			in := convertInput(input)
			in = fixInput(in, noun, verb)
			out := compute(in)
			if out == 19690720{
				fmt.Println(100 * noun + verb)
				return
			}
		}
	}
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

func fixInput(in []int, noun, verb int) []int {
	in[1] = noun
	in[2] = verb
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
