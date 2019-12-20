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

func (p *Program) JumpIfTrue(i *instruction) {
	val := p.Memory.Get(i.ParamOnePointer)
	log(fmt.Sprintf("[Program.JumpIfTrue] - Jumping on value %v", val))

	if val != 0 {
		newPointer := p.Memory.Get(i.ParamTwoPointer)
		log(fmt.Sprintf("[Program.JumpIfTrue] - Setting pointer to %v", newPointer))
		p.Pointer = newPointer
	} else {
		p.Pointer += 3
	}
}

func (p *Program) JumpIfFalse(i *instruction) {
	val := p.Memory.Get(i.ParamOnePointer)
	log(fmt.Sprintf("[Program.JumpIfFalse] - Jumping on value %v", val))

	if val == 0 {
		newPointer := p.Memory.Get(i.ParamTwoPointer)
		log(fmt.Sprintf("[Program.JumpIfFalse] - Setting pointer to %v", newPointer))
		p.Pointer = newPointer
	} else {
		p.Pointer += 3
	}
}

func (p *Program) LessThan(i *instruction) {
	x := p.Memory.Get(i.ParamOnePointer)
	y := p.Memory.Get(i.ParamTwoPointer)
	log(fmt.Sprintf("[Program.LessThan] - Checking if %v < %v", x, y))
	if x < y {
		p.Memory.Set(i.ParamThreePointer, 1)
	} else {
		p.Memory.Set(i.ParamThreePointer, 0)
	}
	p.Pointer += 4
}

func (p *Program) Equal(i *instruction) {
	x := p.Memory.Get(i.ParamOnePointer)
	y := p.Memory.Get(i.ParamTwoPointer)
	log(fmt.Sprintf("[Program.Equal] - Checking if %v == %v", x, y))
	if x == y {
		p.Memory.Set(i.ParamThreePointer, 1)
	} else {
		p.Memory.Set(i.ParamThreePointer, 0)
	}
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
	case 3:
		break
	case 4:
		p.Put(i)
	case 5:
		p.JumpIfTrue(i)
	case 6:
		p.JumpIfFalse(i)
	case 7:
		p.LessThan(i)
	case 8:
		p.Equal(i)
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
