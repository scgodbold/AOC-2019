package six

import (
	"fmt"
	"strings"
)

type Orbit struct {
	Name   string
	Parent *Orbit
}

func (o *Orbit) TotalOrbits() int {
	if o.Parent != nil {
		return o.Parent.TotalOrbits() + 1
	}
	return 0
}

type Planets struct {
	Orbits map[string]*Orbit
}

func (p *Planets) AddOrbit(input string) {
	split := strings.Split(input, ")")
	parent := split[0]
	child := split[1]
	parentOrbit, ok := p.Orbits[parent]
	if !ok {
		parentOrbit = &Orbit{parent, nil}
		p.Orbits[parent] = parentOrbit
	}

	childOrbit, ok := p.Orbits[child]
	if !ok {
		childOrbit = &Orbit{child, parentOrbit}
	}
	childOrbit.Parent = parentOrbit

	p.Orbits[child] = childOrbit
}

func (p *Planets) CountOrbits() int {
	total := 0
	for _, val := range p.Orbits {
		total += val.TotalOrbits()
	}
	return total
}

func (p *Planets) ShortestDistance(source string, dest string) int {
	moves := 0
	start := p.Orbits[source].Parent
	end := p.Orbits[dest].Parent
	for {
		if start.Name == end.Name {
			break
		}
		if start.TotalOrbits() > end.TotalOrbits() {
			start = start.Parent
		} else {
			end = end.Parent
		}
		moves += 1
	}
	return moves
}

func NewPlanets(input []string) *Planets {
	p := Planets{
		make(map[string]*Orbit),
	}
	for _, val := range input {
		p.AddOrbit(val)
	}
	return &p
}

func DaySix(input []string) {
	p := NewPlanets(input)
	fmt.Printf("Total Orbit Count: %v", p.CountOrbits())
	fmt.Printf("Total Moves to reach Santa: %v", p.ShortestDistance("YOU", "SAN"))
}
