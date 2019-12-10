package six

import (
	"testing"
)

func TestOrbitCount(t *testing.T) {
	tables := []struct {
		Input  []string
		Output int
	}{
		{
			[]string{"COM)B",
				"B)C",
				"C)D",
				"D)E",
				"E)F",
				"B)G",
				"G)H",
				"D)I",
				"E)J",
				"J)K",
				"K)L",
			}, 42,
		},
	}
	for _, table := range tables {
		m := MassObjects{
			make(map[string]*Orbit),
		}
		for _, val := range table.Input {
			m.AddOrbit(val)
		}
		val := m.CountOrbits()
		if val != table.Output {
			t.Errorf("Unexpected orbit count for %v, Got: %v, Wanted: %v", table.Input, val, table.Output)
		}
	}
}
