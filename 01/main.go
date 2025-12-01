package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
)

type Result struct {
	Position     int
	TotalZeros   int
	PartTwoZeros int
}

func PartTwo(rotation int, position int) (result int) {
	// distance to cross '0' depends on the direction
	distance := 0
	if rotation > 0 {
		// rotating right
		distance = 100 - position
	} else {
		// rotating left
		distance = position
	}

	hundreds := float64(rotation / 100)
	sanitized := int(math.Abs(float64(rotation - (int(hundreds) * 100))))

	// first crossing check is '(rotation > distance)'
	if distance > 0 && sanitized != distance && sanitized > distance {
		result += 1
	}
	// then add '(rotation div 100)' for every full rotation
	if int(math.Abs(hundreds)) > 0 {
		result += int(math.Abs(hundreds))
	}

	return
}

func (res *Result) ApplyRotation(rotation int) {
	// calculate Part 2 first
	res.PartTwoZeros += PartTwo(rotation, res.Position)

	// actual movement
	res.Position += rotation
	if res.Position >= 100 {
		res.Position = res.Position % 100
	} else if res.Position < 0 {
		hundreds := math.Ceil(math.Abs(float64(res.Position) / 100.0))
		res.Position += 100 * int(hundreds)
	}

	// check actual position
	if res.Position == 0 {
		res.TotalZeros += 1
	}

	// sanity check
	if res.Position < 0 {
		fmt.Printf("error: %d < 0", res.Position)
		os.Exit(1)
	}
	if res.Position > 100 {
		fmt.Printf("error: %d > 100", res.Position)
		os.Exit(1)
	}
}

func ParseLine(line string) int {
	mult := 1
	if line[0] == 'L' {
		mult = -1
	}
	number, err := strconv.Atoi(line[1:])
	if err != nil {
		fmt.Printf("error: cannot parse %s", line)
		os.Exit(1)
	}
	return mult * number
}

func Run(scanner *bufio.Scanner) (int, int) {
	result := Result{Position: 50, TotalZeros: 0}
	for scanner.Scan() {
		rotation := ParseLine(scanner.Text())
		result.ApplyRotation(rotation)
		fmt.Printf("Position=%d, TotalZeros=%d, PartTwoZeros=%d\n", result.Position, result.TotalZeros, result.PartTwoZeros)
	}

	return result.TotalZeros, result.TotalZeros + result.PartTwoZeros
}

func main() {
	if len(os.Args) > 2 {
		fmt.Printf("too many arguments, exit")
		os.Exit(1)
	}

	// run on a file
	if len(os.Args) == 2 {
		file, err := os.Open(os.Args[1])
		if err != nil {
			fmt.Printf("error opening file=%s, err=%s\n", os.Args[1], err)
			os.Exit(1)
		}
		defer file.Close()

		scanner := bufio.NewScanner(file)

		fmt.Println(Run(scanner))
		os.Exit(0)
	}

	// run from stdin
	if len(os.Args) == 1 {
		scanner := bufio.NewScanner(os.Stdin)
		fmt.Println(Run(scanner))
		os.Exit(0)
	}
}
