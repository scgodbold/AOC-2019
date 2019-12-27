package opcode

import (
	"testing"
)

func compProgram(a Program, b Program) bool {
	if a.Pointer != b.Pointer {
		return false
	}

	if a.State != b.State {
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

	// Compare Outputs
	for i, v := range a.Outputs {
		if v != b.Outputs[i] {
			return false
		}
	}

	// Compare Inputs
	for i, v := range a.Inputs {
		if v != b.Inputs[i] {
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
		[]int{},
		[]int{},
		0,
		0,
		0,
	}

	input := &instruction{
		1,
		2,
		2,
		0,
		0,
	}

	expected := Program{
		&memory{
			[]int{1, 2, 2, 0},
			[]int{4, 2, 2, 0},
			4,
		},
		4,
		[]int{},
		[]int{},
		0,
		0,
		0,
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
		[]int{},
		[]int{},
		0,
		0,
		0,
	}

	input := &instruction{
		1,
		2,
		2,
		0,
		0,
	}

	expected := Program{
		&memory{
			[]int{2, 0, 4, 0, 3},
			[]int{16, 0, 4, 0, 3},
			5,
		},
		4,
		[]int{},
		[]int{},
		0,
		0,
		0,
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
				[]int{},
				[]int{},
				0,
				0,
				0,
			},
			&instruction{
				1,
				2,
				4,
				-1,
				0,
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
				[]int{},
				[]int{},
				0,
				0,
				0,
			},
			&instruction{
				1,
				2,
				4,
				-1,
				0,
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
				[]int{},
				[]int{},
				0,
				0,
				0,
			},
			&instruction{
				1,
				2,
				4,
				-1,
				0,
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
				[]int{},
				[]int{},
				0,
				0,
				0,
			},
			&instruction{
				1,
				2,
				4,
				-1,
				0,
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
				[]int{},
				[]int{},
				0,
				0,
				0,
			},
			&instruction{
				1,
				1,
				2,
				3,
				0,
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
				[]int{},
				[]int{},
				0,
				0,
				0,
			},
			&instruction{
				1,
				1,
				2,
				3,
				0,
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
				[]int{},
				[]int{},
				0,
				0,
				0,
			},
			&instruction{
				1,
				0,
				0,
				3,
				0,
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
				[]int{},
				[]int{},
				0,
				0,
				0,
			},
			&instruction{
				1,
				1,
				2,
				3,
				0,
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

func TestProgramOutput(t *testing.T) {
	tests := []struct {
		Program  Program
		Input    *instruction
		Expected Program
	}{
		{
			Program{
				&memory{
					[]int{5, 2, 2, 99, 0},
					[]int{5, 2, 2, 99, 0},
					5,
				},
				0,
				[]int{},
				[]int{},
				0,
				0,
				0,
			},
			&instruction{
				1,
				0,
				0,
				3,
				0,
			},
			Program{
				&memory{
					[]int{5, 2, 2, 99, 0},
					[]int{5, 2, 2, 99, 0},
					5,
				},
				2,
				[]int{5},
				[]int{},
				0,
				0,
				0,
			},
		},
	}

	for _, test := range tests {
		test.Program.Output(test.Input)

		if !compProgram(test.Program, test.Expected) {
			t.Errorf("Program failed Output operation. Expected (%v, %v), Got (%v, %v)", test.Program.Outputs, test.Program.Pointer, test.Expected.Outputs, test.Expected.Pointer)
		}
	}
}

func TestProgramAddInput(t *testing.T) {
	tests := []struct {
		Program  Program
		Input    int
		Expected Program
	}{
		{
			Program{
				&memory{
					[]int{5, 2, 2, 99, 0},
					[]int{5, 2, 2, 99, 0},
					5,
				},
				0,
				[]int{},
				[]int{},
				0,
				0,
				0,
			},
			5,
			Program{
				&memory{
					[]int{5, 2, 2, 99, 0},
					[]int{5, 2, 2, 99, 0},
					5,
				},
				0,
				[]int{},
				[]int{5},
				0,
				0,
				0,
			},
		},
	}

	for _, test := range tests {
		test.Program.AddInput(test.Input)

		if !compProgram(test.Program, test.Expected) {
			t.Errorf("Program failed AddInput operation. Expected %v, Got %v", test.Program.Inputs, test.Expected.Inputs)
		}
	}
}

func TestProgramInput(t *testing.T) {
	tests := []struct {
		Program  Program
		Input    *instruction
		Expected Program
	}{
		{
			Program{
				&memory{
					[]int{5, 2, 2, 99, 0},
					[]int{5, 2, 2, 99, 0},
					5,
				},
				0,
				[]int{},
				[]int{10},
				0,
				0,
				0,
			},
			&instruction{
				1,
				0,
				0,
				3,
				0,
			},
			Program{
				&memory{
					[]int{5, 2, 2, 99, 0},
					[]int{10, 2, 2, 99, 0},
					5,
				},
				2,
				[]int{},
				[]int{10},
				0,
				0,
				0,
			},
		},
		{
			Program{
				&memory{
					[]int{5, 2, 2, 99, 0},
					[]int{5, 2, 2, 99, 0},
					5,
				},
				0,
				[]int{},
				[]int{},
				0,
				1,
				0,
			},
			&instruction{
				1,
				0,
				0,
				3,
				0,
			},
			Program{
				&memory{
					[]int{5, 2, 2, 99, 0},
					[]int{5, 2, 2, 99, 0},
					5,
				},
				0,
				[]int{},
				[]int{},
				0,
				1,
				0,
			},
		},
		{
			Program{
				&memory{
					[]int{5, 2, 2, 99, 0},
					[]int{5, 2, 2, 99, 0},
					5,
				},
				0,
				[]int{},
				[]int{10},
				1,
				0,
				0,
			},
			&instruction{
				1,
				0,
				0,
				3,
				0,
			},
			Program{
				&memory{
					[]int{5, 2, 2, 99, 0},
					[]int{5, 2, 2, 99, 0},
					5,
				},
				0,
				[]int{},
				[]int{10},
				1,
				1,
				0,
			},
		},
	}

	for _, test := range tests {
		test.Program.Input(test.Input)
		if !compProgram(test.Program, test.Expected) {
			t.Errorf("Program failed to fetch Input. Expected (%v, %v), Got (%v, %v)", test.Expected.Memory.Current, test.Expected.Pointer, test.Program.Memory.Current, test.Program.Pointer)
		}
	}
}

func TestMoveRelative(t *testing.T) {
	p := Program{
		&memory{
			[]int{5, -2, 2, 99, 0},
			[]int{5, -2, 2, 99, 0},
			5,
		},
		0,
		[]int{},
		[]int{10},
		1,
		1,
		5,
	}
	tests := []struct {
		Input  *instruction
		Output int
	}{
		{&instruction{1, 0, 0, 2, 0}, 10},
		{&instruction{1, 1, 0, 2, 0}, 3},
	}

	for i, test := range tests {
		p.RelativePointer = 5
		p.MoveRelative(test.Input)

		if p.RelativePointer != test.Output {
			t.Errorf("Test %v: Failed moving relative pointer. Got %v, Expected %v", i, p.RelativePointer, test.Output)
		}
	}
}

// Provided programs w/ expected outputs
func TestsDayTwo(t *testing.T) {
	tests := []struct {
		Input  string
		Output []int
	}{
		// Day 2 Provided Tests
		{
			"1,0,0,0,99",
			[]int{2, 0, 0, 0, 99},
		},
		{
			"2,3,0,3,99",
			[]int{2, 3, 0, 6, 99},
		},
		{
			"2,4,4,5,99,0",
			[]int{2, 4, 4, 5, 99, 9801},
		},
		{
			"1,1,1,4,99,5,6,0,99",
			[]int{30, 1, 1, 4, 2, 5, 6, 0, 99},
		},
	}

	for _, test := range tests {
		p := NewProgram([]string{test.Input})
		p.Run()
		for i, v := range test.Output {
			if v != p.Memory.Current[i] {
				t.Errorf("Program Failed on input %v. Expected %v, got %v", test.Input, test.Output, p.Memory.Current)
				break
			}
		}
	}
}

func TestDayNine(t *testing.T) {
	tests := []struct {
		Program  string
		Input    []int
		Expected []int
	}{
		{
			"109,1,204,-1,1001,100,1,100,1008,100,16,101,1006,101,0,99",
			[]int{},
			[]int{109, 1, 204, -1, 1001, 100, 1, 100, 1008, 100, 16, 101, 1006, 101, 0, 99},
		},
		{
			"1102,34915192,34915192,7,4,7,99,0",
			[]int{},
			[]int{1219070632396864},
		},
		{
			"104,1125899906842624,99",
			[]int{},
			[]int{1125899906842624},
		},
		// Reddit Tests
		{
			"109,-1,4,1,99",
			[]int{},
			[]int{-1},
		},
		{
			"109,-1,104,1,99",
			[]int{},
			[]int{1},
		},
		{
			"109,-1,204,1,99",
			[]int{},
			[]int{109},
		},
		{
			"109,1,9,2,204,-6,99",
			[]int{},
			[]int{204},
		},
		{
			"109,1,109,9,204,-6,99",
			[]int{},
			[]int{204},
		},
		{
			"109,1,209,-1,204,-106,99",
			[]int{},
			[]int{204},
		},
		{
			"109,1,3,3,204,2,99",
			[]int{1502},
			[]int{1502},
		},
		{
			"109,1,203,2,204,2,99",
			[]int{1502},
			[]int{1502},
		},
	}
	for _, test := range tests {
		p := NewProgram([]string{test.Program})
		for _, v := range test.Input {
			p.AddInput(v)
		}
		p.Run()

		for i, v := range test.Expected {
			if v != p.Outputs[i] {
				t.Errorf("Failed test %v: Expected %v, Got %v", test.Program, test.Expected, p.Outputs)
			}
		}
	}
}
