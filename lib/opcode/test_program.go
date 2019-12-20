package opcode

import (
	"testing"
)

func compProgram(a Program, b Program) bool {
	if a.Pointer != b.Pointer {
		return false
	}

	if a.Memory.Size != b.Memory.Size {
		return false
	}

	// Compare Initial Content
	for i, v := range a.Memory.Current {
		if v != b.Memory.Current[i] {
			return false
		}
	}

	return true
}

func TestProgramAdd(t *testing.T) {
	test := Program{
		&memory{
			[]int{1, 2, 2, 0},
			[]int{1, 2, 2, 0},
			4,
		},
		0,
	}

	input := &instruction{
		1,
		2,
		2,
		0,
	}

	expected := Program{
		&memory{
			[]int{1, 2, 2, 0},
			[]int{4, 2, 2, 0},
			4,
		},
		4,
	}

	test.Add(input)

	if compProgram(test, expected) {
		t.Errorf("Program failed add operation. Expected %v, Got %v", expected, test)
	}
}

func TestProgramMultiply(t *testing.T) {
	test := Program{
		&memory{
			[]int{2, 0, 4, 0, 3},
			[]int{2, 0, 4, 0, 3},
			5,
		},
		0,
	}

	input := &instruction{
		1,
		2,
		2,
		0,
	}

	expected := Program{
		&memory{
			[]int{2, 0, 4, 0, 3},
			[]int{6, 0, 4, 0, 3},
			5,
		},
		4,
	}

	test.Multiply(input)

	if compProgram(test, expected) {
		t.Errorf("Program failed add operation. Expected %v, Got %v", expected, test)
	}
}
