package day2

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func Solve() (int, error) {
	file, err := os.Open("day2/input.txt")
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
	for scanner.Scan() {
		text := scanner.Text()
		plays := strings.Fields(text)
		if len(plays) != 2 {
			return 0, fmt.Errorf("invalid input: %s", text)
		}
		oponents := plays[0]
		ours := plays[1]
		score += playsScore(ours)
		if isDraw(oponents, ours) {
			score += 3
		}
		if isWin(oponents, ours) {
			score += 6
		}
	}
	return score, nil
}

func playsScore(ours string) int {
	switch ours {
	case "X":
		return 1
	case "Y":
		return 2
	case "Z":
		return 3
	default:
		return 0
	}
}

func isDraw(opponents, ours string) bool {
	switch opponents+ours {
	case "AX":
		return true
	case "BY":
		return true
	case "CZ":
		return true
	default:
		return false
	}
}

func isWin(opponents, ours string) bool {
	switch opponents {
	case "A":
		return ours == "Y"
	case "B":
		return ours == "Z"
	case "C":
		return ours == "X"
	default:
		return false
	}
}
