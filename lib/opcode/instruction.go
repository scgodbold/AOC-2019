package opcode

import (
	"fmt"
)

type instruction struct {
	Code              int
	ParamOnePointer   int
	ParamTwoPointer   int
	ParamThreePointer int
}

func (i *instruction) Initialize(pointer int, mem *memory) {
	// Initialize each param into param mode hardest
	i.Code = mem.Get(pointer)
	i.ParamOnePointer = mem.Get(pointer + 1)
	i.ParamTwoPointer = mem.Get(pointer + 2)
	i.ParamThreePointer = mem.Get(pointer + 3)

	if i.Code > 10000 {
		i.Code -= 10000
		i.ParamThreePointer = pointer + 3
	}
	if i.Code > 1000 {
		i.Code -= 1000
		i.ParamTwoPointer = pointer + 2
	}
	if i.Code > 100 {
		i.Code -= 100
		i.ParamOnePointer = pointer + 1
	}

	log(fmt.Sprintf("[Instruction] - OpCode: %v, ParamPointer %v, ParamPointer %v, ParamPointer %v", i.Code, i.ParamOnePointer, i.ParamTwoPointer, i.ParamThreePointer))
}

func NewInstruction(pointer int, mem *memory) *instruction {
	var instruct instruction
	instruct.Initialize(pointer, mem)
	return &instruct
}
