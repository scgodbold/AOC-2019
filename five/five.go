package five

import (
	"fmt"

	"github.com/scgodbold/AOC-2019/lib/opcode"
)

func Run(input []string) {
	// Part 1, provide input of 1 and run
	prog := opcode.NewProgram(input)
	prog.AddInput(1)
	prog.Run()
	fmt.Println(prog.Outputs)

	// Part 2, provide input of 5 and run
	prog.Reset()
	prog.AddInput(5)
	prog.Run()
	fmt.Println(prog.Outputs)
}
