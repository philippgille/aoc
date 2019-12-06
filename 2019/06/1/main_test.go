package main

import (
	"testing"
)

var testData = []struct {
	in          string
	expectedOut int
}{
	{`COM)B
B)C
C)D
D)E
E)F
B)G
G)H
D)I
E)J
J)K
K)L`, 42},
}

func TestRequiredFuel(t *testing.T) {
	for i, td := range testData {
		out := steps(convertInput(td.in))
		if out != td.expectedOut {
			t.Errorf("Test failed for input no. %v. Expected %v, but was %v", i, td.expectedOut, out)
		}
	}
}
