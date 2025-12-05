package main

import (
	"bufio"
	"fmt"
	"os"
)

type coord struct {
	x int
	y int
}

var neighbours = []coord{
	{-1, -1}, {0, -1}, {1, -1},
	{-1, 0} /*, 0 */, {1, 0},
	{-1, 1}, {0, 1}, {1, 1},
}

func isPositionValid(pos coord, width int, height int) bool {
	if pos.x < 0 || pos.x >= height {
		return false
	}
	if pos.y < 0 || pos.y >= width {
		return false
	}
	return true
}

func checkNeighbours(field []string, pos coord, width int, height int) bool {
	total := 0
	for _, neiCoord := range neighbours {
		neighbour := coord{pos.x + neiCoord.x, pos.y + neiCoord.y}

		if !isPositionValid(neighbour, width, height) {
			continue
		}

		if field[neighbour.x][neighbour.y] == '@' {
			total += 1
		}
	}

	return total < 4
}

func SolvePartOne(field []string, width int, height int) (result int) {
	for x, row := range field {
		for y, col := range row {
			if col != '@' {
				continue
			}

			if checkNeighbours(field, coord{x, y}, width, height) {
				result += 1
			}
		}
	}
	return
}

func Run(scanner *bufio.Scanner) (partOne int, partTwo int) {
	field := []string{}

	width := 0
	height := 0
	for scanner.Scan() {
		row := scanner.Text()
		if width == 0 {
			width = len(row)
		} else if width != len(row) {
			fmt.Printf("error: inconsistent widths, expected=%d, got=%d\n", width, len(row))
			os.Exit(1)
		}

		field = append(field, row)
		height += 1
	}

	partOne += SolvePartOne(field, width, height)

	return partOne, partTwo
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
