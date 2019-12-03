package three

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

type point struct {
	X int
	Y int
}

type position struct {
	WireA bool
	WireB bool
}

type board struct {
	Graph         [][]position
	Origin        point
	Intersections []point
	WireCount     int
}

type instruction struct {
	Direction byte
	Magnitude int
}

type wire struct {
	Steps []instruction
}

func newBoard(max int) (*board, error) {
	var b board
	edge := (max * 2) + 1

	b.Graph = make([][]position, edge)
	b.Origin = point{max, max}
	b.WireCount = 0
	for i := 0; i < edge; i++ {
		row := make([]position, edge)
		b.Graph[i] = row
	}
	return &b, nil
}

func (b *board) wireStep(step instruction) {
}

func (b *board) AddWire(wire []w) {
	for _, step := range wire.Steps {

	}
}

func newWire(input string) (*wire, error) {
	// Split wire into steps
	var w wire
	for _, val := range strings.Split(input, ",") {
		direction := val[0]
		magnitude, err := strconv.Atoi(strings.Trim(val, "DURL"))
		if err != nil {
			return nil, err
		}

		w.Steps = append(w.Steps, instruction{direction, magnitude})
	}

	return &w, nil
}

func wiresLargestDirection(wires []wire) int {
	var (
		max  float64
		curX int
		curY int
	)
	max = 0.0
	for _, w := range wires {
		curX = 0
		curY = 0
		for _, step := range w.Steps {
			switch step.Direction {
			case 'D':
				curY -= step.Magnitude
			case 'U':
				curY += step.Magnitude
			case 'L':
				curX -= step.Magnitude
			case 'R':
				curX += step.Magnitude
			}
			if math.Abs(float64(curX)) > max {
				max = math.Abs(float64(curX))
			}
			if math.Abs(float64(curY)) > max {
				max = math.Abs(float64(curY))
			}
		}
	}
	return int(max)
}

func DayThree(input []string) {
	fmt.Println("Hello World")
	b, _ := newBoard(2)
	fmt.Printf("%v\n", b.Graph)
	fmt.Printf("%v\n", b.Origin)
}
