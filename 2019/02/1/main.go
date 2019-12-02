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
	pos := 0
	for {
		// Check for index out of range
		if pos > len(in)-1 {
			return in[0]
		}
		// Re-read op
		op := in[pos]
		// Check for exit
		if op == 99 {
			return in[0]
		}
		// Compute
		x := in[in[pos+1]]
		y := in[in[pos+2]]
		res := 0
		if op == 1 {
			res = x + y
		} else if op == 2 {
			res = x * y
		} else {
			panic("bad op")
		}
		in[in[pos+3]] = res
		// Advance pos
		pos = pos + 4
	}
}
