package opcode

import (
	"fmt"
)

const DEBUG = false

type Program struct {
	Memory  *memory
	Pointer int
}

func log(line string) {
	if DEBUG {
		fmt.Println(line)
	}
}

func (p *Program) Reset() {
	p.Memory.Reset()
	p.Pointer = 0
}

func (p *Program) Add(i *instruction) {
	x := p.Memory.Get(i.ParamOnePointer)
	y := p.Memory.Get(i.ParamTwoPointer)
	writePointer := i.ParamThreePointer
	log(fmt.Sprintf("[Program.Add] - Adding values %v & %v", x, y))
	p.Memory.Set(writePointer, x+y)
	p.Pointer += 4
}

func (p *Program) Multiply(i *instruction) {
	x := p.Memory.Get(i.ParamOnePointer)
	y := p.Memory.Get(i.ParamTwoPointer)
	writePointer := i.ParamThreePointer
	log(fmt.Sprintf("[Program.Multiply] - Multiplying values %v & %v", x, y))
	p.Memory.Set(writePointer, x*y)
	p.Pointer += 4
}

func (p *Program) End(i *instruction) {
	log("[Program.End] Exiting Program")
	p.Pointer = p.Memory.Size + 1
}

func (p *Program) Step() {
	log(fmt.Sprintf("[Program.Step] - Evaluating Step at %v", p.Pointer))
	i := NewInstruction(p.Pointer, p.Memory)

	switch i.Code {
	case 1:
		p.Add(i)
	case 2:
		p.Multiply(i)
	case 99:
		p.End(i)
	default:
		log(fmt.Sprintf("[Program.Error] - Reached unknown opcode %v", i.Code))
	}
}

func (p *Program) Run() {
	for {
		if p.Pointer >= p.Memory.Size {
			break
		}
		p.Step()
	}
}

func (p *Program) Initialize(input []string) {
	p.Memory = NewMemory(input)
	p.Pointer = 0
}

func NewProgram(input []string) *Program {
	prog := Program{}
	prog.Initialize(input)
	return &prog
}
