package main

import (
	"flag"
	"fmt"
	"os"
)

var day int

const (
	inputDir = "./inputs"
)

func init() {
	flag.IntVar(&day, "day", 0, "Specify the day challenge to run")
	flag.Parse()
}

func main() {
	if day == 0 {
		panic("Please provide a day to run")
	}
	fmtedDay := fmt.Sprintf("%02d", day)
	fmt.Printf("Running for day: %v\n", fmtedDay)

	inputPath := fmt.Sprintf("%v/%v.txt", inputDir, fmtedDay)
	fmt.Printf("Reading input from: %v\n", inputPath)

	input, err := os.Open(inputPath)
	defer input.Close()
	if err != nil {
		panic(err)
	}

	// TODO: Dynamically call the correct day based on inputs
	val, err := CalculateDayOne(input)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Answer for day %d: %v", day, val)
}
