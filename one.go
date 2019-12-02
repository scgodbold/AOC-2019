package main

import (
	"bufio"
	"io"
	"math"
	"strconv"
	"strings"
)

func FuelCalculation(mass int64) int {
	return int(math.Floor(float64(mass)/3.0) - 2)
}

func CalculateDayOne(input io.Reader) (int, error) {
	inputReader := bufio.NewReader(input)
	var sum int
	sum = 0

	for {
		val, err := inputReader.ReadString('\n')
		if err != nil {
			break
		}
		cleaned, err := strconv.ParseInt(strings.TrimSuffix(val, "\n"), 10, 64)
		if err != nil {
			return sum, err
		}
		fuelRequired := FuelCalculation(cleaned)

		sum += fuelRequired
	}
	return sum, nil
}
