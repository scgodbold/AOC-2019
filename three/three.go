package three

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

type point struct {
	X     int
	Y     int
	WireA int
	WireB int
}

type position struct {
	WireA int
	WireB int
}

type board struct {
	Graph         [][]position
	Origin        point
	Intersections []point
	wireCount     int
}

type instruction struct {
	Direction byte
	Magnitude int
}

type wire struct {
	CurrentPosition point
	Steps           []instruction
	Distance        int
}

func newBoard(max int) (*board, error) {
	var b board
	edge := (max * 2) + 1

	b.Graph = make([][]position, edge)
	b.Origin = point{max, max, 0, 0}
	b.wireCount = 0
	for i := 0; i < edge; i++ {
		row := make([]position, edge)
		b.Graph[i] = row
	}
	b.Graph[b.Origin.X][b.Origin.Y].WireA = 0
	b.Graph[b.Origin.X][b.Origin.Y].WireB = 0
	return &b, nil
}

func (b *board) WirePresent(p point, distance int) {
	if b.wireCount == 0 {
		b.Graph[p.X][p.Y].WireA = distance
	} else {
		b.Graph[p.X][p.Y].WireB = distance
	}

	if b.Graph[p.X][p.Y].WireA > 0 && b.Graph[p.X][p.Y].WireB > 0 {
		b.Intersections = append(b.Intersections, point{p.X, p.Y, b.Graph[p.X][p.Y].WireA, b.Graph[p.X][p.Y].WireB})
	}
}

func (b *board) wireStep(step instruction, w *wire) {
	for i := 0; i < step.Magnitude; i++ {
		switch step.Direction {
		case 'U':
			w.CurrentPosition.Y += 1
		case 'D':
			w.CurrentPosition.Y -= 1
		case 'R':
			w.CurrentPosition.X += 1
		case 'L':
			w.CurrentPosition.X -= 1
		}
		w.Distance += 1
		b.WirePresent(w.CurrentPosition, w.Distance)
	}

}

func (b *board) AddWires(wires []*wire) {
	for _, w := range wires {
		w.CurrentPosition = point{b.Origin.X, b.Origin.Y, 0, 0}
		for _, step := range w.Steps {
			b.wireStep(step, w)
		}
		b.wireCount += 1
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
	w.Distance = 0

	return &w, nil
}

func wiresLargestDirection(wires []*wire) int {
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

func Run(input []string) {
	// Input -> wires
	wires := make([]*wire, len(input))
	for i, val := range input {
		w, _ := newWire(val)
		wires[i] = w
	}

	// Build our board, laydown the wires
	maxMagnitude := wiresLargestDirection(wires)
	b, _ := newBoard(maxMagnitude)
	b.AddWires(wires)

	// Calculate the minimum
	lowestVal := -1
	lowestSteps := -1
	for _, p := range b.Intersections {
		xMagnitude := math.Abs(float64(p.X - b.Origin.X))
		yMagnitude := math.Abs(float64(p.Y - b.Origin.Y))
		totalSteps := p.WireA + p.WireB
		if lowestVal == -1 || lowestVal > int(xMagnitude+yMagnitude) {
			lowestVal = int(xMagnitude + yMagnitude)
		}
		if lowestSteps == -1 || lowestSteps > totalSteps {
			lowestSteps = totalSteps
		}
	}
	fmt.Printf("Closest intersection to the Origin: %v\n", lowestVal)
	fmt.Printf("Intersection with the fewest steps: %v\n", lowestSteps)
}
