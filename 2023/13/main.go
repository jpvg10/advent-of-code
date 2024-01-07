package main

import (
	"bufio"
	"fmt"
	"os"
)

func part1(lines []string) int {
	return solve(lines, 0)
}

func part2(lines []string) int {
	return solve(lines, 1)
}

// Assumes that there is only one dividing line
func solve(lines []string, smudges int) int {
	m := len(lines)
	n := len(lines[0])
	A := make([][]rune, 0)

	for _, l := range lines {
		A = append(A, []rune(l))
	}

	for j := 0; j < (n - 1); j++ {
		l := j
		r := j + 1
		diffs := 0
		for l >= 0 && r < n {
			for i := 0; i < m; i++ {
				if A[i][l] != A[i][r] {
					diffs++
				}
			}
			l--
			r++
		}
		if diffs == smudges {
			return (j + 1)
		}
	}

	for i := 0; i < (m - 1); i++ {
		a := i
		b := i + 1
		diffs := 0
		for a >= 0 && b < m {
			for j := 0; j < n; j++ {
				if A[a][j] != A[b][j] {
					diffs++
				}
			}
			a--
			b++
		}
		if diffs == smudges {
			return 100 * (i + 1)
		}
	}

	return 0
}

func main() {
	// file, _ := os.Open("sample.txt")
	file, _ := os.Open("input.txt")
	defer file.Close()

	scanner := bufio.NewScanner(file)
	lines := []string{}
	totalSum := 0

	for scanner.Scan() {
		line := scanner.Text()
		if line != "" {
			lines = append(lines, line)
		} else {
			// totalSum += part1(lines)
			totalSum += part2(lines)
			lines = []string{}
		}
	}
	// totalSum += part1(lines)
	totalSum += part2(lines)

	fmt.Println(totalSum)
}
