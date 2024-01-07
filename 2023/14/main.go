package main

import (
	"bufio"
	"fmt"
	"os"
)

func part1() {
	// file, _ := os.Open("sample.txt")
	file, _ := os.Open("input.txt")
	defer file.Close()

	scanner := bufio.NewScanner(file)

	lines := []string{}
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	A := [][]rune{}
	for _, l := range lines {
		A = append(A, []rune(l))
	}

	m := len(A)
	n := len(A[0])
	totalSum := 0

	for j := 0; j < n; j++ {
		startRow := -1
		count := 0
		for i := 0; i < m; i++ {
			if (A[i][j] == 'O' || A[i][j] == '.') && startRow == -1 {
				startRow = i
			}
			if A[i][j] == 'O' {
				weight := m - startRow - count
				totalSum += weight
				count++
			}
			if A[i][j] == '#' {
				count = 0
				startRow = -1
			}
		}
	}

	fmt.Println(totalSum)
}

func main() {
	part1()
}
