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

	// TODO: This was determined via trial & error (error being "index out of range").
	// It should better be created dynamically, but that led to "out of memory" when using slices.
	// There probably is a completely different, better way to do this.
	var grid [25000][20000]uint8

	// Note: When passing slices as arguments, their underlying arrays aren't copied.
	_ = drawLine(in1, &grid, 1)
	crossings := drawLine(in2, &grid, 2)

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

func drawLine(in []wirePart, grid *[25000][20000]uint8, id int) []point {
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
			if val != 0 && val != uint8(id) {
				result = append(result, point{x: x, y: y})
			}
			grid[x][y] = uint8(id)
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
