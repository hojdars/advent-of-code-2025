package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

const batteryPackLength = 12

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

	return 10*tens + ones
}

func swap(selection []byte, incoming byte, position int) []byte {
	if incoming >= selection[position] {
		backup := selection[position]
		selection[position] = incoming
		if position+1 < batteryPackLength {
			return swap(selection, backup, position+1)
		} else {
			return selection
		}
	}

	// no swap = no change
	return selection
}

func SolvePartTwo(batteryPack string) int {
	if len(batteryPack) < batteryPackLength {
		fmt.Printf("error: battery pack length < %d, len=%d", batteryPackLength, len(batteryPack))
		os.Exit(1)
	}

	selection := []byte(batteryPack[len(batteryPack)-batteryPackLength:])

	for i := len(batteryPack) - 13; i >= 0; i-- {
		selection = swap([]byte(selection), batteryPack[i], 0)
	}

	intRes, err := strconv.Atoi(string(selection))
	if err != nil {
		fmt.Printf("cannot parse result to string, got=%s", string(selection))
		os.Exit(1)
	}

	return intRes
}

func Run(scanner *bufio.Scanner) (int, int) {
	part1Result := 0
	part2Result := 0
	for scanner.Scan() {
		batteryPack := scanner.Text()
		part1Result += SolvePartOne(batteryPack)
		part2Result += SolvePartTwo(batteryPack)
	}
	return part1Result, part2Result
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
