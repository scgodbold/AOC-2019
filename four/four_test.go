package four

import (
	"testing"
)

func TestPasswordIncreasing(t *testing.T) {
	tables := []struct {
		Input  int
		Output bool
	}{
		{123456, true},
		{111111, true},
		{111123, true},
		{123219, false},
	}
	for _, table := range tables {
		val := passwordIncreasing(table.Input)
		if val != table.Output {
			t.Errorf("Input of (%v) did not provide expected output: Wanted %v, Got: %v", table.Input, table.Output, val)
		}
	}
}

func TestPasswordLength(t *testing.T) {
	tables := []struct {
		Input  int
		Output bool
	}{
		{1, false},
		{123456, true},
		{1234567, false},
	}

	for _, table := range tables {
		val := passwordLength(table.Input)
		if val != table.Output {
			t.Errorf("Input of (%v) did not provide expected output: Wanted %v, Got: %v", table.Input, table.Output, val)
		}
	}
}

func TestPasswordRepeatDigits(t *testing.T) {
	tables := []struct {
		Input  int
		Output bool
	}{
		{111111, true},
		{123456, false},
		{123455, true},
		{111123, true},
	}

	for _, table := range tables {
		val := passwordRepeatDigits(table.Input)
		if val != table.Output {
			t.Errorf("Input of (%v) did not provide expected output: Wanted %v, Got: %v", table.Input, table.Output, val)
		}
	}
}

func TestPasswordRepeatPairs(t *testing.T) {
	tables := []struct {
		Input  int
		Output bool
	}{
		{111111, false},
		{111456, false},
		{123455, true},
		{112233, true},
		{111122, true},
	}

	for _, table := range tables {
		val := passwordRepeatPairs(table.Input)
		if val != table.Output {
			t.Errorf("Input of (%v) did not provide expected output: Wanted %v, Got: %v", table.Input, table.Output, val)
		}
	}
}
