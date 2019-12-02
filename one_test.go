package main

import (
	"testing"
)

// These are the examples provided by AOC
func TestFuelCalculation(t *testing.T) {
	tables := []struct {
		input  int
		output int
	}{
		{12, 2},
		{14, 2},
		{1969, 654},
		{100756, 33583},
	}

	for _, table := range tables {
		val := FuelCalculation(table.input)
		if val != table.output {
			t.Errorf("Return value for (%d) was incorrect, got: %d, wanted: %d", table.input, val, table.output)
		}
	}
}
