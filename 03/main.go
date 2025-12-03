package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func AtoiError(number string) int {
	first, err := strconv.Atoi(number)
	if err != nil {
		fmt.Printf("couldn't parse string=%s to number", number)
		os.Exit(1)
	}
	return first
}

func SolvePartOne(batteryPack string) int {
	if len(batteryPack) < 2 {
		fmt.Printf("error: battery pack too short, got=%s", batteryPack)
		os.Exit(1)
	}

	tens := int(batteryPack[0] - '0')
	ones := int(batteryPack[1] - '0')

	for i := 2; i < len(batteryPack); i++ {
		num := int(batteryPack[i] - '0')
		if tens < ones {
			tens = ones
			ones = num
		} else if ones < num {
			ones = num
		}
	}

	return int(10*tens + ones)
}

func Run(scanner *bufio.Scanner) (int, int) {
	part1Result := 0
	for scanner.Scan() {
		batteryPack := scanner.Text()
		part1Result += SolvePartOne(batteryPack)
	}
	return part1Result, 0
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
