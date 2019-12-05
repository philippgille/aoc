package main

import (
	"strconv"
	"testing"
)

var testData = []struct {
	in          int
	expectedOut bool
}{
	{112233, true},
	{123444, false},
	{111122, true},
}

func TestRequiredFuel(t *testing.T) {
	for i, td := range testData {
		s := strconv.Itoa(td.in)
		out := rule3(s) && rule4(s)
		if out != td.expectedOut {
			t.Errorf("Test failed for input no. %v. Expected %v, but was %v", i, td.expectedOut, out)
		}
	}
}
