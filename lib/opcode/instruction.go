package opcode

import (
	"fmt"
	"strconv"
	"strings"
)

type instruction struct {
	Code              int
	ParamOnePointer   int
	ParamTwoPointer   int
	ParamThreePointer int

	RelativePointer int
}

func (i *instruction) SplitOpCode(input int) ([]int, int) {
	// Construct a known size opcode to parse further
	vals := []rune("00000")
	cleaned := strconv.Itoa(input)
	offset := len(vals) - len(cleaned)
	for i, v := range cleaned {
		vals[i+offset] = v
	}

	// Get Modes
	modes := []int{
		int(vals[0]) - '0',
		int(vals[1]) - '0',
		int(vals[2]) - '0',
	}

	// Build Opcode
	var sb strings.Builder
	sb.WriteRune(vals[3])
	sb.WriteRune(vals[4])
	code, _ := strconv.Atoi(sb.String())

	return modes, code
}

func (i *instruction) GetPointer(mode int, input int, mem *memory) int {
	switch mode {
	case 0:
		return mem.Get(input)
	case 1:
		return input
	case 2:
		return i.RelativePointer + mem.Get(input)
	default:
		return -1
	}
}

func (i *instruction) Initialize(pointer int, relativePointer int, mem *memory) {
	i.RelativePointer = relativePointer
	// Initialize each param into param mode hardest
	var modes []int
	modes, i.Code = i.SplitOpCode(mem.Get(pointer))

	i.ParamThreePointer = i.GetPointer(modes[0], pointer+3, mem)
	i.ParamTwoPointer = i.GetPointer(modes[1], pointer+2, mem)
	i.ParamOnePointer = i.GetPointer(modes[2], pointer+1, mem)

	log(fmt.Sprintf("[Instruction] - OpCode: %v, ParamPointer %v, ParamPointer %v, ParamPointer %v", i.Code, i.ParamOnePointer, i.ParamTwoPointer, i.ParamThreePointer))
}

func NewInstruction(pointer int, relativePointer int, mem *memory) *instruction {
	var instruct instruction
	instruct.Initialize(pointer, relativePointer, mem)
	return &instruct
}
