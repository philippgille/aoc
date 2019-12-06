package main

import (
	"fmt"
	"strconv"
	"strings"
	"time"
)

type wirePart struct {
	dir string
	len int
}

type point struct {
	x, y  int
	steps int
}

func main() {
	start := time.Now()
	smallSteps := calcSmallStepsFromWireInputs(wire1input, wire2input)
	fmt.Println(time.Now().Sub(start))
	fmt.Println(smallSteps)
}

func calcSmallStepsFromWireInputs(wire1input, wire2input string) int {
	in1 := convertInput(wire1input)
	in2 := convertInput(wire2input)

	// TODO: This was determined via trial & error (error being "index out of range").
	// It should better be created dynamically, but that led to "out of memory" when using slices.
	// There probably is a completely different, better way to do this.
	var grid [25000][20000]bool

	// Note: When passing slices as arguments, their underlying arrays aren't copied.
	_ = drawLine(in1, &grid, true)
	crossings2 := drawLine(in2, &grid, false)
	// Draw the other way around, on a new grid, only to determine the steps for line 1
	var grid2 [25000][20000]bool
	_ = drawLine(in2, &grid2, true)
	crossings1 := drawLine(in1, &grid2, false)

	result := minStepsForSameCrossing(crossings1, crossings2)

	return result
}

func convertInput(in string) []wirePart {
	inParts := strings.Split(in, ",")
	result := make([]wirePart, len(inParts))
	for i, elem := range inParts {
		len, err := strconv.Atoi(elem[1:])
		if err != nil {
			panic(err)
		}
		wp := wirePart{
			dir: string(elem[0]),
			len: len,
		}
		result[i] = wp
	}
	return result
}

func drawLine(in []wirePart, grid *[25000][20000]bool, draw bool) []point {
	result := make([]point, 0) // Start with len 0 to easily append
	steps := 0

	x := len(grid) / 2
	y := len(grid[0]) / 2

	for _, wp := range in {
		for i := 0; i < wp.len; i++ {
			// Move
			steps++
			switch wp.dir {
			case "L":
				x--
			case "R":
				x++
			case "U":
				y++
			case "D":
				y--
			}
			// Draw the first line, or check with second line?
			if draw {
				grid[x][y] = true
			} else if grid[x][y] {
				// Found existing line -> found crossing
				result = append(result, point{x: x, y: y, steps: steps})
			}
		}
	}

	return result
}

func minStepsForSameCrossing(crossings1, crossings2 []point) int {
	result := 0
	for _, crossing1 := range crossings1 {
		for _, crossing2 := range crossings2 {
			if (crossing1.x == crossing2.x && crossing1.y == crossing2.y) &&
				(result == 0 || crossing1.steps+crossing2.steps < result) {
				result = crossing1.steps + crossing2.steps
			}
		}
	}
	return result
}
