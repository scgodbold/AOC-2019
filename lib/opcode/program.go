package opcode

import (
	"fmt"
)

const DEBUG = true

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
	writePointer := p.Memory.Get(i.ParamThreePointer)
	log(fmt.Sprintf("[Program.Add] - Adding values %v & %v", x, y))
	p.Memory.Set(writePointer, x+y)
	p.Pointer += 4
}

func (p *Program) Multiply(i *instruction) {
	x := p.Memory.Get(i.ParamOnePointer)
	y := p.Memory.Get(i.ParamTwoPointer)
	writePointer := p.Memory.Get(i.ParamThreePointer)
	log(fmt.Sprintf("[Program.Multiply] - Multiplying values %v & %v", x, y))
	p.Memory.Set(writePointer, x*y)
	p.Pointer += 4
}

func (p *Program) End(i *instruction) {
	p.Pointer = p.Memory.Size + 1
}

func (p *Program) Step() {
	log(fmt.Sprintf("[Program.Step] - Evaluating Step at %v", p.Pointer))
	i := NewInstruction(p.Pointer, p.Memory)

	switch i.Code {
	case 1:
		p.Add(i)
	case 2:
		p.End(i)
	default:
		p.End(i)
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
