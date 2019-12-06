package five

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// {{{1 operations
type operation struct {
	Code       int
	ParamMode0 bool
	ParamMode1 bool
	ParamMode2 bool
}

func newOperation(input int) *operation {
	op := operation{99, false, false, false}

	if input >= 10000 {
		op.ParamMode2 = true
		input -= 10000
	}
	if input >= 1000 {
		op.ParamMode1 = true
		input -= 1000
	}
	if input >= 100 {
		op.ParamMode0 = true
		input -= 100
	}
	op.Code = input
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

func (p *Program) opAdd(op *operation) {
	var (
		x int
		y int
	)
	if op.ParamMode0 {
		x = p.Memory[p.Pointer+1]
	} else {
		x = p.Memory[p.Memory[p.Pointer+1]]
	}
	if op.ParamMode0 {
		y = p.Memory[p.Pointer+2]
	} else {
		y = p.Memory[p.Memory[p.Pointer+2]]
	}

	fmt.Println(x)
	fmt.Println(y)
	fmt.Println(p.Pointer + 3)
	fmt.Println(p.Memory[p.Pointer+3])

	p.Memory[p.Pointer+3] = (x + y)
}

func (p *Program) opMult(op *operation) {
	var (
		x int
		y int
	)
	if op.ParamMode0 {
		x = p.Memory[p.Pointer+1]
	} else {
		x = p.Memory[p.Memory[p.Pointer+1]]
	}
	if op.ParamMode0 {
		y = p.Memory[p.Pointer+2]
	} else {
		y = p.Memory[p.Memory[p.Pointer+2]]
	}

	p.Memory[p.Pointer+3] = x * y
}

func (p *Program) opPut(op *operation) {
	var out int
	if op.ParamMode0 {
		out = p.Memory[p.Pointer+1]
	} else {
		out = p.Memory[p.Memory[p.Pointer+1]]
	}
	fmt.Printf("%v\n", out)
}

func (p *Program) opGet(op *operation) {
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
		fmt.Println("Try again, must be an integer")
	}
	p.Memory[p.Memory[p.Pointer+1]] = cleaned
}

func (p *Program) Run() {
	for {
		if len(p.Memory) < p.Pointer {
			break
		}
		op := newOperation(p.Memory[p.Pointer])
		switch op.Code {
		case 1:
			p.opAdd(op)
		case 2:
			p.opMult(op)
		case 3:
			p.opGet(op)
		case 4:
			p.opPut(op)
		case 99:
			break
		default:
			break
		}
		p.Pointer += 4
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
