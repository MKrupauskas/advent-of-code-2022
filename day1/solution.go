package day1

import (
	"bufio"
	"os"
	"strconv"
)

func Run() (int, error) {
	file, err := os.Open("day1/input.txt")
	if err != nil {
			return 0, err
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	calories, err := findMostCalories(scanner)
	if err != nil {
			return 0, err
	}
	if err := scanner.Err(); err != nil {
		return 0, err
	}
	return calories, nil
}

func maxInt(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func findMostCalories(scanner *bufio.Scanner) (int, error) {
	var max, current int
	for scanner.Scan() {
		text := scanner.Text()
		if text == "" {
			max = maxInt(max, current)
			current = 0
			continue
		}
		n, err := strconv.Atoi(text)
		if err != nil {
				return 0, err
		}
		current += n
	}
	return max, nil
}
