package main

import (
	"testing"
)

var testData = []struct {
	input          string
	expectedOutput int
}{
	{`1,0,0,0,99`, 2},
	{`2,3,0,3,99`, 2},
	{`2,4,4,5,99,0`, 2},
	{`1,1,1,4,99,5,6,0,99`,30},
}

func TestRequiredFuel(t *testing.T) {
	for i, td := range testData {
		in := convertInput(td.input)
		// no fix required for test input
		//in = fixInput(in)
		out := compute(in)
		if out != td.expectedOutput {
			t.Errorf("Test failed for input no. %v (input value \"%v\")", i, td.input)
		}
	}
}
