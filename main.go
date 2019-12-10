package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"

	"github.com/scgodbold/AOC-2019/five"
	"github.com/scgodbold/AOC-2019/four"
	"github.com/scgodbold/AOC-2019/one"
	"github.com/scgodbold/AOC-2019/six"
	"github.com/scgodbold/AOC-2019/three"
	"github.com/scgodbold/AOC-2019/two"
)

var day int

const (
	inputDir = "./inputs"
)

func init() {
	flag.IntVar(&day, "day", 0, "Specify the day challenge to run")
	flag.Parse()
}

func readInput(path string) ([]string, error) {
	var parsed []string

	f, err := os.Open(path)
	defer f.Close()

	if err != nil {
		return parsed, err
	}

	scanner := bufio.NewScanner(f)

	for scanner.Scan() {
		parsed = append(parsed, scanner.Text())
	}
	return parsed, nil
}

func main() {
	if day == 0 {
		panic("Please provide a day to run")
	}
	fmtedDay := fmt.Sprintf("%02d", day)
	fmt.Printf("Running for day: %v\n", fmtedDay)

	inputPath := fmt.Sprintf("%v/%v.txt", inputDir, fmtedDay)
	fmt.Printf("Reading input from: %v\n", inputPath)

	input, err := readInput(inputPath)
	if err != nil {
		panic(err)
	}

	// // TODO: Dynamically call the correct day based on inputs
	switch day {
	case 1:
		one.DayOne(input)
	case 2:
		two.DayTwo(input)
	case 3:
		three.DayThree(input)
	case 4:
		four.DayFour(input)
	case 5:
		five.DayFive(input)
	case 6:
		six.DaySix(input)
	}
}
