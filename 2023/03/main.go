package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func isSymbol(r rune) bool {
	if r != '.' && (r < '0' || r > '9') {
		return true
	} else {
		return false
	}
}

func isAdjacentToSymbol(A [][]rune, row, start, end int) bool {
	m := len(A)
	n := len(A[0])

	if start > 0 && isSymbol(A[row][start-1]) {
		// left
		return true
	}
	if start > 0 && row > 0 && isSymbol(A[row-1][start-1]) {
		// top left
		return true
	}
	if start > 0 && row < m-1 && isSymbol(A[row+1][start-1]) {
		// bottom left
		return true
	}
	if end < n-1 && isSymbol(A[row][end+1]) {
		// right
		return true
	}
	if end < n-1 && row > 0 && isSymbol(A[row-1][end+1]) {
		// top right
		return true
	}
	if end < n-1 && row < m-1 && isSymbol(A[row+1][end+1]) {
		// bottom right
		return true
	}

	for k := start; k <= end; k++ {
		if row > 0 && isSymbol(A[row-1][k]) {
			return true
		}
		if row < m-1 && isSymbol(A[row+1][k]) {
			return true
		}
	}

	return false
}

func addToGearMap(gearMap map[string][]int, key string, number int) {
	list, ok := gearMap[key]
	if ok {
		gearMap[key] = append(list, number)
	} else {
		gearMap[key] = []int{number}
	}
}

func checkAdjacentGears(A [][]rune, row, start, end int, gearMap map[string][]int) {
	number, _ := strconv.Atoi(string(A[row][start : end+1]))
	m := len(A)
	n := len(A[0])

	if start > 0 && A[row][start-1] == '*' {
		// left
		key := fmt.Sprintf("%d %d", row, start-1)
		addToGearMap(gearMap, key, number)
	}
	if start > 0 && row > 0 && A[row-1][start-1] == '*' {
		// top left
		key := fmt.Sprintf("%d %d", row-1, start-1)
		addToGearMap(gearMap, key, number)
	}
	if start > 0 && row < m-1 && A[row+1][start-1] == '*' {
		// bottom left
		key := fmt.Sprintf("%d %d", row+1, start-1)
		addToGearMap(gearMap, key, number)
	}
	if end < n-1 && A[row][end+1] == '*' {
		// right
		key := fmt.Sprintf("%d %d", row, end+1)
		addToGearMap(gearMap, key, number)
	}
	if end < n-1 && row > 0 && A[row-1][end+1] == '*' {
		// top right
		key := fmt.Sprintf("%d %d", row-1, end+1)
		addToGearMap(gearMap, key, number)
	}
	if end < n-1 && row < m-1 && A[row+1][end+1] == '*' {
		// bottom right
		key := fmt.Sprintf("%d %d", row+1, end+1)
		addToGearMap(gearMap, key, number)
	}

	for k := start; k <= end; k++ {
		if row > 0 && A[row-1][k] == '*' {
			key := fmt.Sprintf("%d %d", row-1, k)
			addToGearMap(gearMap, key, number)
		}
		if row < m-1 && A[row+1][k] == '*' {
			key := fmt.Sprintf("%d %d", row+1, k)
			addToGearMap(gearMap, key, number)
		}
	}
}

func part1() {
	// file, _ := os.Open("sample.txt")
	file, _ := os.Open("input.txt")
	defer file.Close()
	scanner := bufio.NewScanner(file)

	var lines []string
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	var A [][]rune
	for _, line := range lines {
		A = append(A, []rune(line))
	}

	m := len(A)
	n := len(A[0])
	sum := 0

	for i := 0; i < m; i++ {
		numStart := -1

		for j := 0; j < n; j++ {
			if A[i][j] >= '0' && A[i][j] <= '9' {
				if numStart == -1 {
					numStart = j
				}
			} else {
				if numStart != -1 {
					if isAdjacentToSymbol(A, i, numStart, j-1) {
						number, _ := strconv.Atoi(string(A[i][numStart:j]))
						sum += number
					}
					numStart = -1
				}
			}
		}

		if numStart != -1 {
			// number ends on the right edge
			if isAdjacentToSymbol(A, i, numStart, n-1) {
				number, _ := strconv.Atoi(string(A[i][numStart:n]))
				sum += number
			}
		}
	}

	fmt.Println(sum)
}

func part2() {
	// file, _ := os.Open("sample.txt")
	file, _ := os.Open("input.txt")
	defer file.Close()
	scanner := bufio.NewScanner(file)

	var lines []string
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	var A [][]rune
	for _, line := range lines {
		A = append(A, []rune(line))
	}

	m := len(A)
	n := len(A[0])
	gearMap := make(map[string][]int)

	for i := 0; i < m; i++ {
		numStart := -1

		for j := 0; j < n; j++ {
			if A[i][j] >= '0' && A[i][j] <= '9' {
				if numStart == -1 {
					numStart = j
				}
			} else {
				if numStart != -1 {
					checkAdjacentGears(A, i, numStart, j-1, gearMap)
					numStart = -1
				}
			}
		}

		if numStart != -1 {
			// number ends on the right edge
			checkAdjacentGears(A, i, numStart, n-1, gearMap)
		}
	}

	ratioSum := 0
	for _, val := range gearMap {
		if len(val) == 2 {
			ratio := val[0] * val[1]
			ratioSum += ratio
		}
	}

	fmt.Println(ratioSum)
}

func main() {
	// part1()
	part2()
}
