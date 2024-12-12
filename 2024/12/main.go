package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
)

func visitRegion(graph [][]int, visited []bool, perimeterContrib []int, start int) (int, int) {
	queue := []int{start}
	area := 0
	perimeter := 0
	visited[start] = true

	for len(queue) > 0 {
		u := queue[0]
		queue = queue[1:]

		area++
		perimeter += perimeterContrib[u]

		for _, v := range graph[u] {
			if !visited[v] {
				visited[v] = true
				queue = append(queue, v)
			}
		}
	}

	return area, perimeter
}

func main() {
	// file, _ := os.Open("sample1.txt")
	// file, _ := os.Open("sample2.txt")
	file, _ := os.Open("input.txt")
	defer file.Close()
	scanner := bufio.NewScanner(file)

	grid := [][]rune{}
	for scanner.Scan() {
		line := scanner.Text()
		grid = append(grid, []rune(line))
	}

	m := len(grid)
	n := len(grid[0])
	graph := make([][]int, m*n)
	perimeterContrib := slices.Repeat([]int{4}, m*n)

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			u := i*n + j
			if i < m-1 && grid[i][j] == grid[i+1][j] {
				v := (i+1)*n + j
				graph[u] = append(graph[u], v)
				graph[v] = append(graph[v], u)
				perimeterContrib[u]--
				perimeterContrib[v]--
			}
			if j < n-1 && grid[i][j] == grid[i][j+1] {
				v := i*n + (j + 1)
				graph[u] = append(graph[u], v)
				graph[v] = append(graph[v], u)
				perimeterContrib[u]--
				perimeterContrib[v]--
			}
		}
	}

	result := 0
	visited := make([]bool, m*n)

	for u := 0; u < m*n; u++ {
		if !visited[u] {
			area, perimeter := visitRegion(graph, visited, perimeterContrib, u)
			result += (area * perimeter)
		}
	}

	fmt.Println(result)
}
