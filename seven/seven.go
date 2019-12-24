package seven

import (
	"fmt"
	"github.com/scgodbold/AOC-2019/lib/opcode"
)

type signalproc func([]int, []string) int

func permutations(elems []int) [][]int {
	var helper func([]int, int)
	res := [][]int{}

	helper = func(arr []int, n int) {
		if n == 1 {
			tmp := make([]int, len(arr))
			copy(tmp, arr)
			res = append(res, tmp)
		} else {
			for i := 0; i < n; i++ {
				helper(arr, n-1)
				if n%2 == 1 {
					tmp := arr[i]
					arr[i] = arr[n-1]
					arr[n-1] = tmp
				} else {
					tmp := arr[0]
					arr[0] = arr[n-1]
					arr[n-1] = tmp
				}
			}
		}
	}

	helper(elems, len(elems))
	return res
}

func getSignal(setting []int, input []string) int {
	program := opcode.NewProgram(input)
	signal := 0
	for _, v := range setting {
		program.Reset()
		program.AddInput(v)
		program.AddInput(signal)
		program.Run()

		signal = program.Outputs[len(program.Outputs)-1]
	}
	return signal
}

func getFeedbackSignal(setting []int, input []string) int {
	// Setup Programs
	progs := make([]*opcode.Program, len(setting))
	for i, v := range setting {
		progs[i] = opcode.NewProgram(input)
		progs[i].AddInput(v)
	}
	signal := 0
	pointer := 0
	var prog *opcode.Program
	for {
		prog = progs[pointer]
		prog.AddInput(signal)
		prog.Run()
		signal = prog.Outputs[len(prog.Outputs)-1]
		if prog.State == 2 && pointer == len(progs)-1 {
			break
		}
		if pointer == len(progs)-1 {
			pointer = 0
		} else {

			pointer += 1
		}
	}

	return signal
}

func findOptimalPhase(settings [][]int, input []string, f signalproc) int {
	signal := 0
	for _, v := range settings {
		newSignal := f(v, input)
		if newSignal > signal {
			signal = newSignal
		}
	}
	return signal
}

func Run(input []string) {
	// Part one
	phaseSettings := permutations([]int{0, 1, 2, 3, 4})
	maxSignal := findOptimalPhase(phaseSettings, input, getSignal)
	fmt.Println(maxSignal)

	// Part two
	phaseSettings = permutations([]int{5, 6, 7, 8, 9})
	maxSignal = findOptimalPhase(phaseSettings, input, getFeedbackSignal)
	fmt.Println(maxSignal)
}
