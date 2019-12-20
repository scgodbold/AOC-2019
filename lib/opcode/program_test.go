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

	if !compProgram(test, expected) {
		t.Errorf("Program failed add operation. Expected (%v,%v), Got (%v,%v)", expected.Memory.Current, expected.Pointer, test.Memory.Current, test.Pointer)
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
			[]int{16, 0, 4, 0, 3},
			5,
		},
		4,
	}

	test.Multiply(input)

	if !compProgram(test, expected) {
		t.Errorf("Program failed multiply operation. Expected (%v,%v), Got (%v,%v)", expected.Memory.Current, expected.Pointer, test.Memory.Current, test.Pointer)
	}
}

func TestProgramJumpIfTrue(t *testing.T) {
	tests := []struct {
		Program  Program
		Input    *instruction
		Expected int
	}{
		{
			Program{
				&memory{
					[]int{5, 2, 4, 99, 0},
					[]int{5, 2, 4, 99, 0},
					5,
				},
				0,
			},
			&instruction{
				1,
				2,
				4,
				-1,
			},
			0,
		},
		{
			Program{
				&memory{
					[]int{5, 2, 0, 99, 0},
					[]int{5, 2, 0, 99, 0},
					5,
				},
				0,
			},
			&instruction{
				1,
				2,
				4,
				-1,
			},
			3,
		},
	}

	for _, test := range tests {
		test.Program.JumpIfTrue(test.Input)

		if test.Program.Pointer != test.Expected {
			t.Errorf("Program failed JumpIfTrue operation. Expected %v, Got %v", test.Expected, test.Program.Pointer)
		}
	}
}

func TestProgramJumpIfFalse(t *testing.T) {
	tests := []struct {
		Program  Program
		Input    *instruction
		Expected int
	}{
		{
			Program{
				&memory{
					[]int{5, 2, 4, 99, 0},
					[]int{5, 2, 4, 99, 0},
					5,
				},
				0,
			},
			&instruction{
				1,
				2,
				4,
				-1,
			},
			3,
		},
		{
			Program{
				&memory{
					[]int{5, 2, 0, 99, 0},
					[]int{5, 2, 0, 99, 0},
					5,
				},
				0,
			},
			&instruction{
				1,
				2,
				4,
				-1,
			},
			0,
		},
	}

	for _, test := range tests {
		test.Program.JumpIfFalse(test.Input)

		if test.Program.Pointer != test.Expected {
			t.Errorf("Program failed JumpIfFalse operation. Expected %v, Got %v", test.Expected, test.Program.Pointer)
		}
	}
}
func TestProgramLessThan(t *testing.T) {
	tests := []struct {
		Program  Program
		Input    *instruction
		Expected int
	}{
		{
			Program{
				&memory{
					[]int{5, 2, 4, 99, 0},
					[]int{5, 2, 4, 99, 0},
					5,
				},
				0,
			},
			&instruction{
				1,
				1,
				2,
				3,
			},
			1,
		},
		{
			Program{
				&memory{
					[]int{5, 2, 0, 99, 0},
					[]int{5, 100, 0, 99, 0},
					5,
				},
				0,
			},
			&instruction{
				1,
				1,
				2,
				3,
			},
			0,
		},
	}

	for _, test := range tests {
		test.Program.LessThan(test.Input)

		if test.Program.Pointer != 4 || test.Program.Memory.Get(3) != test.Expected {
			t.Errorf("Program failed LessThan operation. Expected (%v, %v), Got (%v, %v)", test.Expected, 4, test.Program.Memory.Get(3), test.Program.Pointer)
		}
	}
}

func TestProgramEqual(t *testing.T) {
	tests := []struct {
		Program  Program
		Input    *instruction
		Expected int
	}{
		{
			Program{
				&memory{
					[]int{5, 2, 2, 99, 0},
					[]int{5, 2, 2, 99, 0},
					5,
				},
				0,
			},
			&instruction{
				1,
				0,
				0,
				3,
			},
			1,
		},
		{
			Program{
				&memory{
					[]int{5, 2, 0, 99, 0},
					[]int{5, 100, 0, 99, 0},
					5,
				},
				0,
			},
			&instruction{
				1,
				1,
				2,
				3,
			},
			0,
		},
	}

	for _, test := range tests {
		test.Program.Equal(test.Input)

		if test.Program.Pointer != 4 || test.Program.Memory.Get(3) != test.Expected {
			t.Errorf("Program failed EqualTo operation (%v). Expected (%v, %v), Got (%v, %v)", test.Input, test.Expected, 4, test.Program.Memory.Get(3), test.Program.Pointer)
		}
	}
}
