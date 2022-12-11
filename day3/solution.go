package day3

import (
	"bufio"
	"fmt"
	"os"
	"unicode"
)

type entry struct {
	left, right int
}

func Solve() (int, error) {
	file, err := os.Open("day3/input.txt")
	if err != nil {
			return 0, err
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	score, err := calculateScore(scanner)
	if err != nil {
			return 0, err
	}
	if err := scanner.Err(); err != nil {
		return 0, err
	}
	return score, nil
}

func calculateScore(scanner *bufio.Scanner) (int, error) {
	var score int
	fmt.Println(int('a'))
	fmt.Println(int('A'))
	for scanner.Scan() {
		line := scanner.Text()
		length := len(line)
		if length % 2 != 0 {
			return 0, fmt.Errorf("line length not even: %d", length)
		}
		char, err := findCommon([]rune(line))
		if err != nil {
				return 0, err
		}
		score += charScore(char)
	}
	return score, nil
}

func findCommon(line []rune) (rune, error) {
	half := len(line) / 2
	common := make(map[rune]entry)
	for i := range line {
		char := line[i]
		e := common[char]
		if i < half {
			e.left++
			common[char] = e
			continue
		}
		e.right++
		common[char] = e
	}
	for char, count := range common {
		if count.left > 0 && count.right > 0 {
			return char, nil
		}
	}
	return 0, nil
}

func charScore(char rune) int {
	if unicode.IsUpper(char) {
		return int(char) - 38 
	}
	return int(char) - 96
} 
