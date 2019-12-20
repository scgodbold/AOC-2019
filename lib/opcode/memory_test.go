package opcode

import (
	"testing"
)

func memComp(a memory, b memory) bool {
	// Check Size
	if a.Size != b.Size {
		return false
	}

	// Check Initial
	if len(a.Initial) != len(b.Initial) {
		return false
	}
	for i, v := range a.Initial {
		if v != b.Initial[i] {
			return false
		}
	}

	// Check Current
	if len(a.Current) != len(b.Current) {
		return false
	}
	for i, v := range a.Current {
		if v != b.Current[i] {
			return false
		}
	}
	return true
}

func TestMemoryInitialize(t *testing.T) {
	input := []string{"1,2,3"}
	output := memory{
		[]int{1, 2, 3},
		[]int{},
		3,
	}

	var test memory
	test.Initialize(input)
	if !memComp(test, output) {
		t.Errorf("Memory not initializing properly. Expected %v, Got %v", output, test)
	}
}

func TestMemoryReset(t *testing.T) {
	input := memory{
		[]int{1, 2, 3},
		[]int{},
		3,
	}
	output := memory{
		[]int{1, 2, 3},
		[]int{1, 2, 3},
		3,
	}

	input.Reset()
	if !memComp(input, output) {
		t.Errorf("Memory not reset properly. Expected %v, Got %v", output, input)
	}
}

func TestMemoryGet(t *testing.T) {
	mem := memory{
		[]int{1, 2, 3},
		[]int{1, 2, 3},
		3,
	}

	tests := []struct {
		Input  int
		Output int
	}{
		{-1, -1},
		{0, 1},
		{2, 3},
		{3, -1},
		{4, -1},
	}

	for _, test := range tests {
		v := mem.Get(test.Input)
		if v != test.Output {
			t.Errorf("Memory not executing Get properly for value %v. Expected %v, Got %v", test.Input, test.Output, v)
		}
	}
}

func TestMemorySet(t *testing.T) {

	tests := []struct {
		Pointer  int
		Val      int
		Expected memory
	}{
		{
			-1,
			7,
			memory{
				[]int{1, 2, 3},
				[]int{1, 2, 3},
				3,
			},
		},
		{
			0,
			7,
			memory{
				[]int{1, 2, 3},
				[]int{7, 2, 3},
				3,
			},
		},
		{
			2,
			7,
			memory{
				[]int{1, 2, 3},
				[]int{1, 2, 7},
				3,
			},
		},
		{
			3,
			7,
			memory{
				[]int{1, 2, 3},
				[]int{1, 2, 3},
				3,
			},
		},
		{
			4,
			7,
			memory{
				[]int{1, 2, 3},
				[]int{1, 2, 3},
				3,
			},
		},
	}

	for _, test := range tests {
		mem := memory{
			[]int{1, 2, 3},
			[]int{1, 2, 3},
			3,
		}
		mem.Set(test.Pointer, test.Val)
		if !memComp(mem, test.Expected) {
			t.Errorf("Memory not setting properly on inputs (%v, %v). Expected %v, Got %v", test.Pointer, test.Val, test.Expected, mem)
		}
	}
}
