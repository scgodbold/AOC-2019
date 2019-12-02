package main

import (
	"fmt"
	"strconv"
	"strings"
)

type Program struct {
	Instructions []int
	Pointer      int
}

func NewProgram(input []string) (*Program, error) {
	var prog Program
	for _, line := range input {

		for _, val := range strings.Split(line, ",") {
			cleaned, err := strconv.Atoi(val)
			if err != nil {
				return nil, err
			}
			prog.Instructions = append(prog.Instructions, cleaned)
		}
	}
	prog.Pointer = 0
	return &prog, nil
}

func (p *Program) opCodeAdd() {
}

func DayTwo(input []string) {
	prog, err := NewProgram(input)
	if err != nil {
		fmt.Printf("Failed to process program")
		return
	}
	fmt.Println(prog.Instructions)
}
