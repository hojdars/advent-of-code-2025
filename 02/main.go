package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func SolvePartOne(intervalStart int, intervalEnd int) int {
	resultPart1 := 0 // total INVALID intervals

	value := intervalStart
	for value <= intervalEnd {
		totalLength := int(math.Floor(math.Log10(float64(value))) + 1)

		// numbers with odd length can be skipped
		if totalLength%2 != 0 {
			value = int(math.Pow10(int(totalLength)))
			continue
		}

		strVal := strconv.Itoa(value)
		if strVal[:totalLength/2] == strVal[totalLength/2:] {
			resultPart1 += value
		}

		value += 1
	}

	return resultPart1
}

func SolvePartTwo(intervalStart int, intervalEnd int) int {
	result := 0

	value := intervalStart
	for value <= intervalEnd {
		strVal := strconv.Itoa(value)
		totalLength := int(math.Floor(math.Log10(float64(value))) + 1)
		for length := 1; length <= totalLength/2; length++ {
			// skip non-divisible patterns
			if totalLength%length != 0 {
				continue
			}

			// verify all repeats
			isSame := true
			for i := 1; i < totalLength/length; i++ {
				if strVal[0:length] != strVal[i*length:(i+1)*length] {
					isSame = false
					break
				}
			}

			if isSame {
				result += value
				break
			}
		}

		value += 1
	}

	return result
}

func Run(scanner *bufio.Scanner) (int, int) {
	scanner.Scan()
	line := scanner.Text()
	intervals := strings.SplitSeq(line, ",")

	total1 := 0
	total2 := 0

	for interval := range intervals {
		intervalParts := strings.Split(interval, "-")

		intervalStart, err := strconv.Atoi(intervalParts[0])
		if err != nil {
			fmt.Printf("error parsing interval start, val=%s, err=%s", intervalParts[0], err)
			os.Exit(1)
		}

		intervalEnd, err := strconv.Atoi(intervalParts[1])
		if err != nil {
			fmt.Printf("error parsing interval end, val=%s, err=%s", intervalParts[1], err)
			os.Exit(1)
		}

		part1 := SolvePartOne(intervalStart, intervalEnd)
		total1 += part1

		part2 := SolvePartTwo(intervalStart, intervalEnd)
		total2 += part2
	}

	return total1, total2
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
