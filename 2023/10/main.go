package main

import (
	"bufio"
	"fmt"
	"os"
)

func connectUp(i, j int, G [][]int, A [][]rune) {
	if i <= 0 {
		return
	}
	n := len(A[0])
	k := i*n + j
	up := (i-1)*n + j
	if A[i-1][j] == '|' || A[i-1][j] == '7' || A[i-1][j] == 'F' || A[i-1][j] == 'S' {
		G[up] = append(G[up], k)
		G[k] = append(G[k], up)
	}
}

func connectDown(i, j int, G [][]int, A [][]rune) {
	m := len(A)
	if i >= m-1 {
		return
	}
	n := len(A[0])
	k := i*n + j
	down := (i+1)*n + j
	if A[i+1][j] == '|' || A[i+1][j] == 'J' || A[i+1][j] == 'L' || A[i+1][j] == 'S' {
		G[down] = append(G[down], k)
		G[k] = append(G[k], down)
	}
}

func connectLeft(i, j int, G [][]int, A [][]rune) {
	if j <= 0 {
		return
	}
	n := len(A[0])
	k := i*n + j
	left := i*n + (j - 1)
	if A[i][j-1] == '-' || A[i][j-1] == 'L' || A[i][j-1] == 'F' || A[i][j-1] == 'S' {
		G[left] = append(G[left], k)
		G[k] = append(G[k], left)
	}
}

func connectRight(i, j int, G [][]int, A [][]rune) {
	n := len(A[0])
	if j >= n-1 {
		return
	}
	k := i*n + j
	right := i*n + (j + 1)
	if A[i][j+1] == '-' || A[i][j+1] == '7' || A[i][j+1] == 'J' || A[i][j+1] == 'S' {
		G[right] = append(G[right], k)
		G[k] = append(G[k], right)
	}
}

func dfs(u int, G [][]int, seen []bool, prev []int) {
	seen[u] = true
	for _, v := range G[u] {
		if !seen[v] {
			prev[v] = u
			dfs(v, G, seen, prev)
		}
	}
}

func part1() {
	// file, _ := os.Open("sample1.txt")
	// file, _ := os.Open("sample2.txt")
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
	G := make([][]int, m*n)
	s := -1

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			switch A[i][j] {
			case '|':
				connectUp(i, j, G, A)
				connectDown(i, j, G, A)
			case '-':
				connectLeft(i, j, G, A)
				connectRight(i, j, G, A)
			case 'L':
				connectUp(i, j, G, A)
				connectRight(i, j, G, A)
			case 'J':
				connectUp(i, j, G, A)
				connectLeft(i, j, G, A)
			case '7':
				connectDown(i, j, G, A)
				connectLeft(i, j, G, A)
			case 'F':
				connectDown(i, j, G, A)
				connectRight(i, j, G, A)
			case 'S':
				s = i*n + j
			}
		}
	}

	seen := make([]bool, m*n)
	prev := make([]int, m*n)
	dfs(s, G, seen, prev)

	loopLength := 1
	current := G[s][1]
	for current != s {
		current = prev[current]
		loopLength++
	}

	maxDist := (loopLength + 1) / 2
	fmt.Println(maxDist)
}

func main() {
	part1()
}
