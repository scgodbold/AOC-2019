package nine

import (
	"fmt"
	"github.com/scgodbold/AOC-2019/lib/opcode"
)

func Run(input []string) {
	prog := opcode.NewProgram(input)
	prog.AddInput(1)
	prog.Run()
	fmt.Println(prog.Outputs)
}
