package main

import (
	"testing"
)

var testData = []struct {
	input          int
	expectedOutput int
}{
	{12, 2},
	{14, 2},
	{1969, 654},
	{100756, 33583},
}

func TestRequiredFuel(t *testing.T) {
	for i, td := range testData {
		if requiredFuel(td.input) != td.expectedOutput {
			t.Errorf("Test failed for input no. %v (input value \"%v\")", i, td.input)
		}
	}
}
