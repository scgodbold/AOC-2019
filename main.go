package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"

	"github.com/scgodbold/AOC-2019/eight"
	"github.com/scgodbold/AOC-2019/five"
	"github.com/scgodbold/AOC-2019/four"
	"github.com/scgodbold/AOC-2019/nine"
	"github.com/scgodbold/AOC-2019/one"
	"github.com/scgodbold/AOC-2019/seven"
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
		one.Run(input)
	case 2:
		two.Run(input)
	case 3:
		three.Run(input)
	case 4:
		four.Run(input)
	case 5:
		five.Run(input)
	case 6:
		six.Run(input)
	case 7:
		seven.Run(input)
	case 8:
		eight.Run(input)
	case 9:
		nine.Run(input)
	}
}
