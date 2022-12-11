package main

import (
	"fmt"
	"os"

	"mkrup.com/advent-of-code-2022/day3"
)

func run() error {
	n, err := day3.Solve()
	if err != nil {
			return err
	}
	fmt.Println(n)
	return nil
}

func main() {
	if err := run(); err != nil {
		fmt.Fprintf(os.Stderr, "run: %s\n", err.Error())
		os.Exit(1)
	}
}
