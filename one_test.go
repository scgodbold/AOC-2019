package main

import (
	"testing"
)

// These are the examples provided by AOC
func TestFuelCalculation(t *testing.T) {
	tables := []struct {
		input  int64
		output int64
	}{
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

func TestTotalFuel(t *testing.T) {
	tables := []struct {
		input  int64
		output int64
	}{
		{2, 2},
		{654, 966},
		{33583, 50346},
	}

	for _, table := range tables {
		val := TotalFuelCalculation(table.input)
		if val != table.output {
			t.Errorf("Return value for (%d) was incorrect, got: %d, wanted: %d", table.input, val, table.output)
		}
	}
}
