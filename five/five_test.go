package five

import (
	"testing"
)

func TestNewOperation(t *testing.T) {
	tables := []struct {
		Input  int
		Output *operation
	}{
		{3, &operation{3, false, false, false}},
		{103, &operation{3, true, false, false}},
		{1003, &operation{3, false, true, false}},
		{10003, &operation{3, false, false, true}},
	}

	for _, val := range tables {
		op := newOperation(val.Input)
		if op != val.Output {
			t.Errorf("Bad operation for %v, Wanted: %v, Got: %v", val.Input, val.Output, op)
		}
	}
}
