package opcode

import (
	"testing"
)

func TestInstructionInitialize(t *testing.T) {
	tests := []struct {
		Pointer  int
		Memory   *memory
		Expected instruction
	}{
		{
			0,
			&memory{
				[]int{1, 12, 13, 14},
				[]int{1, 12, 13, 14},
				4,
			},
			instruction{
				1,
				12,
				13,
				14,
			},
		},
		{
			0,
			&memory{
				[]int{11101, 12, 13, 14},
				[]int{11101, 12, 13, 14},
				4,
			},
			instruction{
				1,
				1,
				2,
				3,
			},
		},
		{
			0,
			&memory{
				[]int{10101, 12, 13, 14},
				[]int{10101, 12, 13, 14},
				4,
			},
			instruction{
				1,
				1,
				13,
				3,
			},
		},
		{
			0,
			&memory{
				[]int{1101, 12, 13, 14},
				[]int{1101, 12, 13, 14},
				4,
			},
			instruction{
				1,
				1,
				2,
				14,
			},
		},
		{
			0,
			&memory{
				[]int{11001, 12, 13, 14},
				[]int{11001, 12, 13, 14},
				4,
			},
			instruction{
				1,
				12,
				2,
				3,
			},
		},
		{
			0,
			&memory{
				[]int{99, -1, -1, -1},
				[]int{99, -1, -1, -1},
				4,
			},
			instruction{
				99,
				-1,
				-1,
				-1,
			},
		},
	}

	for _, test := range tests {
		i := instruction{}
		i.Initialize(test.Pointer, test.Memory)

		if i != test.Expected {
			t.Errorf("Instruction failed to initiliaze on inputs (%v, %v). Expected %v, Got %v", test.Pointer, test.Memory, test.Expected, i)
		}
	}
}
