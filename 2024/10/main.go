package main

import (
	"bufio"
	"fmt"
	"os"
)

func dfs(graph [][]int, visited []int, start int) {
	visited[start] += 1
	for _, v := range graph[start] {
		dfs(graph, visited, v)
	}
}

func main() {
	// file, _ := os.Open("sample.txt")
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
	zeros := []int{}
	nines := []int{}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			numA := int(grid[i][j] - '0')
			u := i*n + j

			if numA == 0 {
				zeros = append(zeros, u)
			} else if numA == 9 {
				nines = append(nines, u)
			}

			if i < m-1 {
				numB := int(grid[i+1][j] - '0')
				v := (i+1)*n + j
				if numB == numA+1 {
					graph[u] = append(graph[u], v)
				}
				if numA == numB+1 {
					graph[v] = append(graph[v], u)
				}
			}

			if j < n-1 {
				numB := int(grid[i][j+1] - '0')
				v := i*n + j + 1
				if numB == numA+1 {
					graph[u] = append(graph[u], v)
				}
				if numA == numB+1 {
					graph[v] = append(graph[v], u)
				}
			}
		}
	}

	totalScore := 0
	totalRating := 0
	for _, zero := range zeros {
		visited := make([]int, m*n)
		dfs(graph, visited, zero)
		score := 0
		rating := 0
		for _, nine := range nines {
			if visited[nine] > 0 {
				score++
				rating += visited[nine]
			}
		}
		totalScore += score
		totalRating += rating
	}

	fmt.Println("Part 1:", totalScore)
	fmt.Println("Part 2:", totalRating)
}
