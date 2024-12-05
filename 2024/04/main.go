package main

import (
	"bufio"
	"fmt"
	"os"
)

func part1(W [][]rune) {
	m := len(W)
	n := len(W[0])
	xmasCount := 0

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if W[i][j] != 'X' {
				continue
			}
			if j < n-3 && W[i][j+1] == 'M' && W[i][j+2] == 'A' && W[i][j+3] == 'S' {
				// Horizontal left to right
				xmasCount++
			}
			if j >= 3 && W[i][j-1] == 'M' && W[i][j-2] == 'A' && W[i][j-3] == 'S' {
				// Horizontal right to left
				xmasCount++
			}
			if i < m-3 && W[i+1][j] == 'M' && W[i+2][j] == 'A' && W[i+3][j] == 'S' {
				// Vertical up to down
				xmasCount++
			}
			if i >= 3 && W[i-1][j] == 'M' && W[i-2][j] == 'A' && W[i-3][j] == 'S' {
				// Vertical down to up
				xmasCount++
			}
			if j < n-3 && i < m-3 && W[i+1][j+1] == 'M' && W[i+2][j+2] == 'A' && W[i+3][j+3] == 'S' {
				// Diagonal left to right and up to down
				xmasCount++
			}
			if j >= 3 && i < m-3 && W[i+1][j-1] == 'M' && W[i+2][j-2] == 'A' && W[i+3][j-3] == 'S' {
				// Diagonal right to left and up to down
				xmasCount++
			}
			if j < n-3 && i >= 3 && W[i-1][j+1] == 'M' && W[i-2][j+2] == 'A' && W[i-3][j+3] == 'S' {
				// Diagonal left to right and down to up
				xmasCount++
			}
			if j >= 3 && i >= 3 && W[i-1][j-1] == 'M' && W[i-2][j-2] == 'A' && W[i-3][j-3] == 'S' {
				// Diagonal right to left and down to up
				xmasCount++
			}
		}
	}

	fmt.Println(xmasCount)
}

func part2(W [][]rune) {
	m := len(W)
	n := len(W[0])
	xmasCount := 0

	for i := 0; i <= m-3; i++ {
		for j := 0; j <= n-3; j++ {
			if W[i+1][j+1] != 'A' {
				continue
			}
			if W[i][j] == 'M' && W[i][j+2] == 'M' && W[i+2][j] == 'S' && W[i+2][j+2] == 'S' {
				xmasCount++
			} else if W[i][j] == 'M' && W[i][j+2] == 'S' && W[i+2][j] == 'M' && W[i+2][j+2] == 'S' {
				xmasCount++
			} else if W[i][j] == 'S' && W[i][j+2] == 'S' && W[i+2][j] == 'M' && W[i+2][j+2] == 'M' {
				xmasCount++
			} else if W[i][j] == 'S' && W[i][j+2] == 'M' && W[i+2][j] == 'S' && W[i+2][j+2] == 'M' {
				xmasCount++
			}
		}
	}

	fmt.Println(xmasCount)
}

func main() {
	// file, _ := os.Open("sample.txt")
	file, _ := os.Open("input.txt")
	defer file.Close()
	scanner := bufio.NewScanner(file)

	var W [][]rune
	for scanner.Scan() {
		line := scanner.Text()
		W = append(W, []rune(line))
	}

	part1(W)
	part2(W)
}
