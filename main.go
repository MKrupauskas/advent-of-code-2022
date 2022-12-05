package main

import (
	"fmt"
	"os"

	"mkrup.com/advent-of-code-2022/day2"
)

func run() error {
	n, err := day2.Solve()
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
