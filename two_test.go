package main

import (
	"testing"
)

func TestNewProgram(t *testing.T) {
	var input []string
	input = append(input, "1")

	prog, err := NewProgram(input)
	if err != nil {
		t.Errorf("Return for NewProgram was an error: %v", err)
	}

	if prog.Instructions[0] == 1 {
		t.Errorf("Return value for NewProgram(\"1\") returned the wrong value: wanted 1, got %v", prog)
	}
}
