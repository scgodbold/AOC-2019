package two

import (
	"testing"
)

func progCompare(x []int, y []int) bool {
	if len(x) != len(y) {
		return false
	}

	for i, v := range x {
		if y[i] != v {
			return false
		}
	}

	return true
}

func TestProgRun(t *testing.T) {
	tables := []struct {
		input  []int
		output []int
	}{
		{[]int{1, 0, 0, 0, 99}, []int{2, 0, 0, 0, 99}},
		{[]int{2, 3, 0, 3, 99}, []int{2, 3, 0, 6, 99}},
		{[]int{2, 4, 4, 5, 99, 0}, []int{2, 4, 4, 5, 99, 9801}},
		{[]int{1, 1, 1, 4, 99, 5, 6, 0, 99}, []int{30, 1, 1, 4, 2, 5, 6, 0, 99}},
	}
	for _, val := range tables {
		input_prog := make([]int, len(val.input))
		copy(input_prog, val.input)
		prog := Program{val.input, input_prog, 0}
		prog.Run()
		if !progCompare(prog.Memory, val.output) {
			t.Errorf("Program (%v) did not execute as expected: got: %v, wanted: %v", val.input, prog.Memory, val.output)
		}
	}
}

func TestProgReset(t *testing.T) {
	input := []int{1, 0, 0, 0, 99}

	// Use distinct copies so input maynot be modified
	copyOne := make([]int, len(input))
	copyTwo := make([]int, len(input))
	copy(copyOne, input)
	copy(copyTwo, input)
	prog := Program{copyOne, copyTwo, 0}

	// Execute + Reset
	prog.Run()
	prog.Reset()

	// Validate that Memory is back to starting position
	if !progCompare(prog.Memory, input) {
		t.Errorf("Program (%v) did not reset correctly: got: %v, wanted: %v", input, prog.Memory, input)
	}
}
