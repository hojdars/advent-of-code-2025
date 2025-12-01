package main

import (
	"bufio"
	"fmt"
	"os"
)

func Run(scanner *bufio.Scanner) (int, int) {
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
	return 0, 0
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
