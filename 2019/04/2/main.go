package main

import (
	"fmt"
	"runtime"
	"strconv"
	"github.com/gosuri/uiprogress"
)

func main() {
	// 6 digits, any number, would be 10^6 = 1,000,000 possibilities.
	// Between the puzzle input min and max, 245,318-765,747, would be 520,429 possibilities, without rules.
	// There's probably a clever algorithm, but let's try to go through all possibilities,
	// not counting the ones that are against the rules.
	//
	// Rule 1 and 2 are covered by just going from min to max.
	//
	// Speed up by using goroutines.
	cpus := runtime.NumCPU()
	elems := inMax - inMin
	elemsPerCPU := elems / cpus
	resChan := make(chan int, cpus)
	var remMin int
	uiprogress.Start() 

	for i := 0; i < cpus; i++ {
		cpuMin := inMin + (i * elemsPerCPU)
		cpuMax := cpuMin + elemsPerCPU
		remMin = cpuMax
		bar := uiprogress.AddBar(100).PrependElapsed().AppendCompleted()
		go possibilities(cpuMin, cpuMax, resChan, bar)
	}
	result := 0
	for i := 0; i < cpus; i++ {
		result += <-resChan
	}
	// Remainders
	remMax := inMax
	if remMax-remMin > 0 {
		result += possibilities(remMin, remMax, nil, nil)
	}
	fmt.Println(result)
}

func possibilities(min, max int, resChan chan int, bar *uiprogress.Bar) int {
	result := 0
	progress := 0
	for i := min; i <= max; i++ {
		s := strconv.Itoa(i)
		if rule3(s) && rule4(s) {
			result++
		}
		if bar != nil && i / 100 > progress{
			bar.Incr()
			progress = i/100
		}
	}
	if resChan != nil {
		resChan <- result
	}
	return result
}

// rule3 checks the pw for:
// There are two digits that are the same and there's no same digit next to them.
func rule3(pw string) bool {
	// We know the number of digits is 6.
	// Start from index 1 and go to index 4 and check if
	// *one* neighbor is the same and the *other* is different.
	// But the one next to the same one should be different again
	// (or no digit at all because it's the beginning or end of the string).
	for i := 1; i < 5; i++ {
		// Left neighbor is the same
		if pw[i-1] == pw[i] &&
		// Right neighbor is not
		pw[i] != pw[i+1] &&
		// Left neighbor of left neighbor doesn't exist or is different
		(i == 1 || pw[i-2] != pw[i-1]) {
			return true // Early return
		}
		// Right neighbor is the same
		if pw[i] == pw[i+1] &&
			// Left neighbor is not
			pw[i-1] != pw[i] &&
			// Right neighbor of right neighbor doesn't exist or is different
			(i == len(pw)-2 || pw[i] != pw[i+2])  {
			return true // Early return
		}
	}
	return false
}

// rule4 checks the pw for:
// The digits never decrease.
func rule4(pw string) bool {
	// We know the number of digits is 6
	for i := 0; i < 5; i++ {
		if pw[i+1] < pw[i] {
			return false
		}
	}
	return true
}
