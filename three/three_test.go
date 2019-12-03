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

func TestNewWire(t *testing.T) {
	input := "U1,D1,L1,R1"
	expected := wire{
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
	input := []wire{
		wire{
			[]instruction{
				instruction{'U', 15},
				instruction{'U', 15},
			},
		},
		wire{
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
