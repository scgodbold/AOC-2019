package five

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

const (
	DEBUG = false
)

func log(line string) {
	if DEBUG {
		fmt.Println(line)
	}
}

// {{{1 operations
type operation struct {
	//Operations XYZ0C
	Code int
	PosX bool
	PosY bool
	PosZ bool
}

func (p *Program) NewOperation(input int) *operation {
	op := operation{99, true, true, true}
	copy_input := input

	if copy_input >= 10000 {
		log(fmt.Sprintf("Setting Z to immediate for input %v", input))
		op.PosZ = false
		copy_input -= 10000
	}
	if copy_input >= 1000 {
		log(fmt.Sprintf("Setting Y to immediate for input %v", input))
		op.PosY = false
		copy_input -= 1000
	}
	if copy_input >= 100 {
		log(fmt.Sprintf("Setting X to immediate for input %v", input))
		op.PosX = false
		copy_input -= 100
	}
	op.Code = copy_input
	return &op
}

// }}}

// {{{1 Program
type Program struct {
	InitialState []int
	Memory       []int
	Pointer      int
}

func (p *Program) Reset() {
	p.Memory = make([]int, len(p.InitialState))
	copy(p.Memory, p.InitialState)
	p.Pointer = 0
}

func (p *Program) Write(pos int, val int) {
	pointer := p.Memory[pos]
	log(fmt.Sprintf("Writing %v to position %v", val, pointer))
	p.Memory[pointer] = val
}

func (p *Program) fetchVal(pointer int, pos bool) int {
	if pos {
		log(fmt.Sprintf("Fetching value from [Positional] %v", p.Memory[pointer]))
		return p.Memory[p.Memory[pointer]]
	}
	log(fmt.Sprintf("Fetching value from [Immediate] %v", pointer))
	return p.Memory[pointer]
}

func (p *Program) Add(op *operation) {
	x := p.fetchVal(p.Pointer+1, op.PosX)
	y := p.fetchVal(p.Pointer+2, op.PosY)
	log(fmt.Sprintf("Adding values %v and %v", x, y))

	p.Write(p.Pointer+3, x+y)
}

func (p *Program) Multiply(op *operation) {
	x := p.fetchVal(p.Pointer+1, op.PosX)
	y := p.fetchVal(p.Pointer+2, op.PosY)
	log(fmt.Sprintf("Multiplying values %v and %v", x, y))

	p.Write(p.Pointer+3, x*y)
}

func (p *Program) Print(op *operation) {
	val := p.fetchVal(p.Pointer+1, op.PosX)
	fmt.Printf("Writing val: %v\n", val)
}

func (p *Program) Get(op *operation) {
	var (
		cleaned int
		err     error
	)
	for {
		reader := bufio.NewReader(os.Stdin)
		raw, _ := reader.ReadString('\n')
		cleaned, err = strconv.Atoi(strings.TrimSuffix(raw, "\n"))
		if err == nil {
			break
		}
		fmt.Println("Error: Unable to convert '%v'to integer", raw)
	}
	p.Write(p.Pointer+1, cleaned)
}

func (p *Program) JumpIfTrue(op *operation) {
	val := p.fetchVal(p.Pointer+1, op.PosX)
	log(fmt.Sprintf("Jumping if true on value: %v", val))
	if val != 0 {
		new_pointer := p.fetchVal(p.Pointer+2, op.PosY)
		log(fmt.Sprintf("Setting program pointer to: %v", new_pointer))
		p.Pointer = new_pointer
	} else {
		p.Pointer += 3
	}
}

func (p *Program) JumpIfFalse(op *operation) {
	val := p.fetchVal(p.Pointer+1, op.PosX)
	log(fmt.Sprintf("Jumping if false on value: %v", val))
	if val == 0 {
		new_pointer := p.fetchVal(p.Pointer+2, op.PosY)
		log(fmt.Sprintf("Setting program pointer to: %v", new_pointer))
		p.Pointer = new_pointer
	} else {
		p.Pointer += 3
	}
}

func (p *Program) CheckLessThan(op *operation) {
	first := p.fetchVal(p.Pointer+1, op.PosX)
	second := p.fetchVal(p.Pointer+2, op.PosY)
	if first < second {
		p.Write(p.Pointer+3, 1)
	} else {
		p.Write(p.Pointer+3, 0)
	}
}

func (p *Program) CheckEqual(op *operation) {
	first := p.fetchVal(p.Pointer+1, op.PosX)
	second := p.fetchVal(p.Pointer+2, op.PosY)
	if first == second {
		p.Write(p.Pointer+3, 1)
	} else {
		p.Write(p.Pointer+3, 0)
	}
}

func (p *Program) Run() {
	for {
		if len(p.Memory) <= p.Pointer {
			break
		}
		op := p.NewOperation(p.Memory[p.Pointer])
		log(fmt.Sprintf("Executing Operation: %v", op))
		switch op.Code {
		case 1:
			p.Add(op)
			p.Pointer += 4
		case 2:
			p.Multiply(op)
			p.Pointer += 4
		case 3:
			p.Get(op)
			p.Pointer += 2
		case 4:
			p.Print(op)
			p.Pointer += 2
		case 5:
			p.JumpIfTrue(op)
		case 6:
			p.JumpIfFalse(op)
		case 7:
			p.CheckLessThan(op)
			p.Pointer += 4
		case 8:
			p.CheckEqual(op)
			p.Pointer += 4
		case 99:
			p.Pointer = len(p.Memory)
			break
		default:
			log(fmt.Sprintf("Got invalid operation: %v", op))
			p.Pointer = len(p.Memory)
			break
		}
	}
}

func NewProgram(input []string) (*Program, error) {
	var prog Program
	for _, line := range input {
		for _, val := range strings.Split(line, ",") {
			cleaned, err := strconv.Atoi(val)
			if err != nil {
				return nil, err
			}
			prog.InitialState = append(prog.InitialState, cleaned)
		}
	}
	prog.Reset()
	return &prog, nil
}

//}}}

func DayFive(input []string) {
	prog, err := NewProgram(input)
	if err != nil {
		fmt.Println("Unable to parse input for day 5")
		return
	}
	prog.Run()
}
