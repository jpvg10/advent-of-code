package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

type Galaxy struct {
	row int
	col int
}

func part1() {
	solve(2)
}

func part2() {
	solve(1e6)
}

func solve(replace int) {
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

	rows := []int{}
	for i := 0; i < m; i++ {
		allDot := true
		for j := 0; j < n; j++ {
			if A[i][j] != '.' {
				allDot = false
				break
			}
		}
		if !allDot {
			rows = append(rows, i)
		}
	}

	cols := []int{}
	for j := 0; j < n; j++ {
		allDot := true
		for i := 0; i < m; i++ {
			if A[i][j] != '.' {
				allDot = false
				break
			}
		}
		if !allDot {
			cols = append(cols, j)
		}
	}

	galaxies := []Galaxy{}
	for _, row := range rows {
		for _, col := range cols {
			if A[row][col] == '#' {
				galaxies = append(galaxies, Galaxy{row: row, col: col})
			}
		}
	}

	totalSum := 0

	for _, g1 := range galaxies {
		for _, g2 := range galaxies {
			if g1 == g2 {
				continue
			}
			rowIndex1 := -1
			rowIndex2 := -1
			colIndex1 := -1
			colIndex2 := -1
			for i, row := range rows {
				if row == g1.row {
					rowIndex1 = i
				}
				if row == g2.row {
					rowIndex2 = i
				}
			}
			for i, col := range cols {
				if col == g1.col {
					colIndex1 = i
				}
				if col == g2.col {
					colIndex2 = i
				}
			}

			rowDif := int(math.Abs(float64(g1.row) - float64(g2.row)))
			rowIndexDif := int(math.Abs(float64(rowIndex1) - float64(rowIndex2)))
			emptyRowsInBetween := rowDif - rowIndexDif

			colDif := int(math.Abs(float64(g1.col) - float64(g2.col)))
			colIndexDif := int(math.Abs(float64(colIndex1) - float64(colIndex2)))
			emptyColsInBetween := colDif - colIndexDif

			dist := rowIndexDif + replace*emptyRowsInBetween + colIndexDif + replace*emptyColsInBetween
			totalSum += dist
		}
	}

	fmt.Println(totalSum / 2)
}

func main() {
	// part1()
	part2()
}
