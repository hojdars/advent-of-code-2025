package main

import (
	"bufio"
	"cmp"
	"fmt"
	"math"
	"os"
	"slices"
	"strconv"
	"strings"
)

type Interval struct {
	Start int
	End   int
}

func mergeIntervals(inputIntervals []Interval) []Interval {
	intervals := make([]Interval, len(inputIntervals))
	copy(intervals, inputIntervals)
	for i := 0; i < len(intervals); {
		current := intervals[i]

		if i == len(intervals)-1 {
			break
		}

		// no overlap, continue
		if intervals[i+1].Start-current.End > 1 {
			i += 1
			continue
		}

		// overlapping
		newInt := Interval{Start: current.Start, End: int(math.Max(float64(current.End), float64(intervals[i+1].End)))}
		intervals[i+1] = newInt
		intervals = slices.Delete(intervals, i, i+1)
	}

	return intervals
}

func parseInterval(interval string) Interval {
	split := strings.Split(interval, "-")
	if len(split) != 2 {
		fmt.Printf("error: interval is not as expected, got=%s", interval)
		os.Exit(1)
	}

	start, err := strconv.Atoi(split[0])
	if err != nil {
		fmt.Printf("error parsing int, got=%s", split[0])
		os.Exit(1)
	}

	end, err := strconv.Atoi(split[1])
	if err != nil {
		fmt.Printf("error parsing int, got=%s", split[1])
		os.Exit(1)
	}

	return Interval{start, end}
}

func SolvePartOne(query int, intervals []Interval) bool {
	n, found := slices.BinarySearchFunc(intervals, query, func(i Interval, q int) int {
		return cmp.Compare(i.Start, q)
	})

	// found means an interval starts with 'query'
	if found {
		return true
	}
	// 'n-1 < 0' means 'query' is lower than the lowest interval
	if n-1 < 0 {
		return false
	}

	// BinarySearchFunc returns the position 'query' would sort into = the closest interval is the previous one
	closest := intervals[n-1]
	return query >= closest.Start && query <= closest.End
}

func SolverPartTwo(intervals []Interval) (result int) {
	for _, interval := range intervals {
		// +1 because the intervals are inclusive on both sides
		result += interval.End - interval.Start + 1
	}
	return
}

func Run(scanner *bufio.Scanner) (result1 int, result2 int) {
	// load the intervals
	intervals := make([]Interval, 0)
	for scanner.Scan() {
		line := scanner.Text()
		if len(line) == 0 {
			break
		}
		intervals = append(intervals, parseInterval(line))
	}

	// postprocess the intervals
	slices.SortFunc(intervals, func(a, b Interval) int {
		return cmp.Compare(a.Start, b.Start)
	})
	mergedIntervals := mergeIntervals(intervals)

	// run the queries
	for scanner.Scan() {
		line := scanner.Text()
		query, err := strconv.Atoi(line)
		if err != nil {
			fmt.Printf("cannot parse int, got=%s", line)
			os.Exit(1)
		}

		if SolvePartOne(query, mergedIntervals) {
			result1 += 1
		}
	}

	// part two is already solved
	result2 = SolverPartTwo(mergedIntervals)

	return
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
