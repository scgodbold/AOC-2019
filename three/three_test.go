package three

import (
	"testing"
)

func compareWires(x *wire, y *wire) bool {
	if len(x.Steps) != len(y.Steps) {
		return false
	}
	for i, v := range x.Steps {
		if v.Direction != y.Steps[i].Direction || v.Magnitude != y.Steps[i].Magnitude {
			return false
		}
	}
	return true
}

func compareIntersections(x []point, y []point) bool {
	if len(x) != len(y) {
		return false
	}
	for i, v := range x {
		if y[i].Y != v.Y || y[i].X != v.X {
			return false
		}
	}
	return true
}

func TestNewWire(t *testing.T) {
	input := "U1,D1,L1,R1"
	expected := wire{
		point{0, 0},
		[]instruction{
			instruction{'U', 1},
			instruction{'D', 1},
			instruction{'L', 1},
			instruction{'R', 1},
		},
	}

	out, err := newWire(input)
	if err != nil {
		t.Errorf("Failed to create new wire: %v\n", err)
	}

	if !compareWires(&expected, out) {
		t.Errorf("Unexpected result for wire creation: Wanted: %v, Got %v", expected, out)
	}

}

func TestWireLargestDirection(t *testing.T) {
	input := []*wire{
		&wire{
			point{0, 0},
			[]instruction{
				instruction{'U', 15},
				instruction{'U', 15},
			},
		},
		&wire{
			point{0, 0},
			[]instruction{
				instruction{'R', 15},
			},
		},
	}
	expected := 30

	out := wiresLargestDirection(input)
	if out != expected {
		t.Errorf("Unexpected Largest Direction: Wanted: %v, Got: %v", expected, out)
	}
}

func TestAddWires(t *testing.T) {
	b, _ := newBoard(8)
	input := []*wire{
		&wire{
			point{0, 0},
			[]instruction{
				instruction{'R', 8},
				instruction{'U', 5},
				instruction{'L', 5},
				instruction{'D', 3},
			},
		},
		&wire{
			point{0, 0},
			[]instruction{
				instruction{'U', 7},
				instruction{'R', 6},
				instruction{'D', 4},
				instruction{'L', 4},
			},
		},
	}
	expected := []point{
		point{14, 13},
		point{11, 11},
	}
	b.AddWires(input)

	if !compareIntersections(b.Intersections, expected) {
		t.Errorf("Unexpected intesections of the wires, Got: %v, Wanted %v", b.Intersections, expected)
	}

}
