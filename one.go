package main

import (
	"fmt"
	"math"
	"strconv"
)

func FuelCalculation(mass int64) int64 {
	return int64(math.Floor(float64(mass)/3.0) - 2)
}

// Perform fuel calculations for each line which represents a module
// weight. Sum the values and return total fuel required to move the
// list of modules
func DayOnePartOne(input []string) (int64, error) {
	var sum int64
	sum = 0

	for _, val := range input {
		cleaned, err := strconv.ParseInt(val, 10, 64)
		if err != nil {
			return 0, err
		}
		fuelRequired := FuelCalculation(cleaned)

		sum += fuelRequired
	}
	return sum, nil
}

// Takes base fuel required and returns how much fuel will be required for
// to move this amount of fuel
func TotalFuelCalculation(baseFuel int64) int64 {
	var (
		total          int64
		newBase        int64
		additionalFuel int64
	)
	total = baseFuel
	newBase = baseFuel

	for {
		additionalFuel = FuelCalculation(newBase)
		if additionalFuel <= 0 {
			break
		}
		total += additionalFuel
		newBase = additionalFuel
	}
	return total
}

func DayOnePartTwo(input []string) (int64, error) {
	var (
		sum       int64
		baseFuel  int64
		totalFuel int64
	)
	sum = 0

	for _, val := range input {
		cleaned, err := strconv.ParseInt(val, 10, 64)
		if err != nil {
			return 0, err
		}
		baseFuel = FuelCalculation(cleaned)
		totalFuel = TotalFuelCalculation(baseFuel)
		sum += totalFuel
	}

	return sum, nil
}

func DayOne(input []string) {
	// PartOne
	baseFuel, err := DayOnePartOne(input)
	if err != nil {
		fmt.Printf("Unable to perform calculation: Day1 Part 1 // %v\n", err)
		return
	}
	fmt.Printf("Base fuel required: %v\n", baseFuel)

	// Part 2
	totalFuel, err := DayOnePartTwo(input)
	if err != nil {
		fmt.Printf("Unable to perform calculation: Day1 Part2 // %v\n", err)
		return
	}
	fmt.Printf("Total fuel required: %v\n", totalFuel)
}
