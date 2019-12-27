package opcode

import (
	"testing"
)

func TestInstructionSplitOpcode(t *testing.T) {

	tests := []struct {
		Input int
		Modes []int
		Code  int
	}{
		{99, []int{0, 0, 0}, 99},
		{102, []int{0, 0, 1}, 2},
		{1002, []int{0, 1, 0}, 2},
		{10002, []int{1, 0, 0}, 2},
		{20002, []int{2, 0, 0}, 2},
		{99992, []int{9, 9, 9}, 92},
	}

	for j, test := range tests {
		i := instruction{}
		newMode, newCode := i.SplitOpCode(test.Input)

		if test.Code != newCode {
			t.Errorf("Test %v: Failed to construct proper code. Expected %v, Got %v", j, test.Code, newCode)
			break
		}

		if len(newMode) != len(test.Modes) {
			t.Errorf("Test %v: Failed to modes properly. Expected %v, Got %v", j, test.Modes, newMode)
			break
		}

		for x, v := range newMode {
			if v != test.Modes[x] {
				t.Errorf("Test %v: Failed to modes properly. Expected %v, Got %v", j, test.Modes, newMode)
				break
			}
		}

	}
}

func TestInstructionGetPointer(t *testing.T) {
	mem := &memory{
		[]int{5, -2, 2, 99, 0},
		[]int{5, -2, 2, 99, 0},
		5,
	}

	tests := []struct {
		Mode    int
		Pointer int
		Output  int
	}{
		{0, 2, 2},
		{1, 3, 3},
		{2, 3, 149},
		{2, 1, 48},
	}

	for x, test := range tests {
		i := instruction{}
		i.RelativePointer = 50
		result := i.GetPointer(test.Mode, test.Pointer, mem)
		if result != test.Output {
			t.Errorf("Test %v: Failed on (Mode; %v, Pointer %v). Expected %v, Got %v", x, test.Mode, test.Pointer, test.Output, result)
		}
	}
}
