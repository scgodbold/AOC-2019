package two

import (
	"fmt"
	"strconv"
	"strings"
)

type Program struct {
	DefaultMemory []int
	Memory        []int
	Pointer       int
}

func NewProgram(input []string) (*Program, error) {
	var prog Program
	for _, line := range input {
		for _, val := range strings.Split(line, ",") {
			cleaned, err := strconv.Atoi(val)
			if err != nil {
				return nil, err
			}
			prog.DefaultMemory = append(prog.DefaultMemory, cleaned)
		}
	}
	prog.Reset()
	return &prog, nil
}

func (p *Program) Reset() {
	p.Memory = make([]int, len(p.DefaultMemory))
	copy(p.Memory, p.DefaultMemory)
	p.Pointer = 0
}

func (p *Program) Run() {
	for {
		if len(p.Memory) < p.Pointer {
			break
		}
		switch p.Memory[p.Pointer] {
		case 1:
			p.opAdd()
		case 2:
			p.opMult()
		case 99:
			break
		default:
			break
		}
		p.Pointer += 4
	}
}

func (p *Program) opMult() {
	xVal := p.Memory[p.Memory[p.Pointer+1]]
	yVal := p.Memory[p.Memory[p.Pointer+2]]
	storePos := p.Memory[p.Pointer+3]
	p.Memory[storePos] = xVal * yVal
}

func (p *Program) opAdd() {
	xVal := p.Memory[p.Memory[p.Pointer+1]]
	yVal := p.Memory[p.Memory[p.Pointer+2]]
	storePos := p.Memory[p.Pointer+3]
	p.Memory[storePos] = xVal + yVal
}

func DayTwoPartOne(prog *Program) {
	// Alter Noun and Verb state of program
	prog.Memory[1] = 12
	prog.Memory[2] = 2

	// Execute Program
	prog.Run()
}

func DayTwoPartTwo(prog *Program) (int, int) {
	const (
		wanted    = 19690720
		nounBound = 99
		verbBound = 99
	)
	var (
		noun int
		verb int
	)

	found := false
	for i := 0; i <= nounBound; i++ {
		for j := 0; j <= verbBound; j++ {
			prog.Reset()
			prog.Memory[1] = i
			prog.Memory[2] = j
			prog.Run()

			if prog.Memory[0] == wanted {
				found = true
				noun = i
				verb = j
				break
			}

			// Since we can only make values larger if the value
			// was larger then what we want, this isnt working move along
			if prog.Memory[0] > wanted {
				break
			}
		}
		if found {
			break
		}
	}

	return noun, verb
}

func DayTwo(input []string) {
	prog, err := NewProgram(input)
	if err != nil {
		fmt.Printf("Failed to parse program for Day2 Part1: %v\n", err)
		return
	}

	// Part one solution
	DayTwoPartOne(prog)
	fmt.Printf("Return value of the program upon execution is: %v\n", prog.Memory[0])

	prog.Reset()
	// Part two solution
	noun, verb := DayTwoPartTwo(prog)
	fmt.Printf("Noun (%v) and Verb (%v) provide expected output", noun, verb)
}
