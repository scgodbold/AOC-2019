package two

import (
	"fmt"
	"github.com/scgodbold/AOC-2019/lib/opcode"
)

func DayTwoPartOne(prog *opcode.Program) {
	// Alter Noun and Verb state of program
	prog.Memory.Set(1, 12)
	prog.Memory.Set(2, 2)

	// Execute Program
	prog.Run()
}

func DayTwoPartTwo(prog *opcode.Program) (int, int) {
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
			prog.Memory.Set(1, i)
			prog.Memory.Set(2, j)
			prog.Run()

			if prog.Memory.Get(0) == wanted {
				found = true
				noun = i
				verb = j
				break
			}

			// Since we can only make values larger if the value
			// was larger then what we want, this isnt working move along
			if prog.Memory.Get(0) > wanted {
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

	prog := opcode.NewProgram(input)

	// Part one solution
	DayTwoPartOne(prog)
	fmt.Printf("Return value of the program upon execution is: %v\n", prog.Memory.Get(0))

	prog.Reset()
	// Part two solution
	noun, verb := DayTwoPartTwo(prog)
	fmt.Printf("Noun (%v) and Verb (%v) provide expected output", noun, verb)
}
