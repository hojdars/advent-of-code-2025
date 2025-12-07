package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

func doOperation(data [][]string, pos int, initial int, op func(int, int) int) (result int) {
	result = initial

	for j := 0; j < len(data)-1; j++ {
		conv, err := strconv.Atoi(data[j][pos])
		if err != nil {
			fmt.Printf("cannot convert data=%s", data[j][pos])
			os.Exit(1)
		}
		result = op(result, conv)
	}

	return
}

func SolvePartOne(data [][]string) (result int) {
	result = 0
	signPos := len(data) - 1
	for i := 0; i < len(data[0]); i++ {
		switch data[signPos][i] {
		case "+":
			result += doOperation(data, i, 0, func(a, b int) int { return a + b })
		case "*":
			result += doOperation(data, i, 1, func(a, b int) int { return a * b })
		default:
			fmt.Printf("error: unexpected sign=%s\n", data[signPos][i])
		}
	}
	return
}

func Run(scanner *bufio.Scanner) (result1 int, result2 int) {
	lines := make([][]string, 0)
	for scanner.Scan() {
		line := scanner.Text()
		values := strings.Split(line, " ")
		values = slices.DeleteFunc(values, func(s string) bool {
			return s == " " || len(s) == 0
		})
		lines = append(lines, values)
	}

	result1 = SolvePartOne(lines)

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
