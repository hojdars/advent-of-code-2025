package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

func SolvePartOne(data [][]string) (result int) {
	result = 0
	signPos := len(data) - 1
	for i := 0; i < len(data[0]); i++ {
		if data[signPos][i] == "+" {
			sumRes := 0
			for j := 0; j < len(data)-1; j++ {
				conv, err := strconv.Atoi(data[j][i])
				if err != nil {
					fmt.Printf("cannot convert data=%s", data[j][i])
					os.Exit(1)
				}
				sumRes += conv
			}
			result += sumRes
		} else if data[signPos][i] == "*" {
			multRes := 1
			for j := 0; j < len(data)-1; j++ {
				conv, err := strconv.Atoi(data[j][i])
				if err != nil {
					fmt.Printf("cannot convert data=%s", data[j][i])
					os.Exit(1)
				}
				multRes *= conv
			}
			result += multRes
		} else {
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
