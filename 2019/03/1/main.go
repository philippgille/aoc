package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"
	"time"
)

type wirePart struct {
	dir string
	len int
}

type point struct {
	x, y int
}

func main() {
	start := time.Now()
	smallDist := calcSmallDistFromWireInputs(wire1input, wire2input)
	fmt.Println(time.Now().Sub(start))
	fmt.Println(smallDist)
}

func calcSmallDistFromWireInputs(wire1input, wire2input string) int {
	in1 := convertInput(wire1input)
	in2 := convertInput(wire2input)

	grid := makeGrid(in1, in2)

	// Note: When passing slices as arguments, their underlying arrays aren't copied.
	_ = drawLine(in1, grid, 1)
	crossings := drawLine(in2, grid, 2)

	start := point{x: len(grid) / 2, y: len(grid[0]) / 2}
	return calcSmallDist(start, crossings)
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

// makeGrid creates a grid with 2*x as the maximum R+L and 2*y as the maximum U+D.
// This way we can start at point x,y and be sure not to go out of range in any direction.
func makeGrid(in1, in2 []wirePart) [][]int {
	x1, y1 := getMovements(in1)
	x2, y2 := getMovements(in2)

	var x, y int
	if x1 > x2 {
		x = x1
	} else {
		x = x2
	}
	if y1 > y2 {
		y = y1
	} else {
		y = y2
	}

	// TODO: out of memory when going the safe route of creating [2*x][2*y] slices :( .
	// TODO: [2*x][2*y] works (and is required) for the test data, while [x/2][y/2] is enough for the main input.
	// result := make([][]int, 2*x)
	result := make([][]int, x/2)
	// for i := 0; i < 2*x; i++ {
	for i := 0; i < x/2; i++ {
		// result[i] = make([]int, 2*y)
		result[i] = make([]int, y/2)
	}

	return result
}

func getMovements(in []wirePart) (x, y int) {
	for _, wp := range in {
		if wp.dir == "L" || wp.dir == "R" {
			x += wp.len
		} else {
			y += wp.len
		}
	}
	return
}

func drawLine(in []wirePart, grid [][]int, id int) []point {
	result := make([]point, 0) // Start with len 0 to easily append

	x := len(grid) / 2
	y := len(grid[0]) / 2

	for _, wp := range in {
		for i := 0; i < wp.len; i++ {
			// Move
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
			// Draw
			val := grid[x][y]
			// 0 is "no line".
			// So if there is a line and it's not the current one, it's a crossing.
			if val != 0 && val != id {
				result = append(result, point{x: x, y: y})
			}
			grid[x][y] = id
		}
	}

	return result
}

func calcSmallDist(start point, crossings []point) int {
	result := 0
	for _, crossing := range crossings {
		temp := dist(start.x, crossing.x) + dist(start.y, crossing.y)
		if result == 0 || temp < result {
			result = temp
		}
	}
	return result
}

func dist(a, b int) int {
	if a == 0 {
		return int(math.Abs(float64(b)))
	}
	if b == 0 {
		return int(math.Abs(float64(a)))
	}
	if a > b {
		return a - b
	}
	return b - a
}
