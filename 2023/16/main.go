package main

import (
	"bufio"
	"fmt"
	"os"
)

type Direction int

const (
	UP    Direction = 0
	DOWN  Direction = 1
	LEFT  Direction = 2
	RIGHT Direction = 3
)

type LightPoint struct {
	row int
	col int
	dir Direction
}

func solve(A [][]rune, start LightPoint) int {
	m := len(A)
	n := len(A[0])
	visited := make([]bool, m*n)
	points := []LightPoint{}
	points = append(points, start)

	for len(points) > 0 {
		current := points[0]
		points = points[1:]
		i := current.row
		j := current.col
		dir := current.dir

		for {
			if i < 0 || i >= m || j < 0 || j >= n {
				break
			}
			if visited[i*n+j] && (A[i][j] == '-' || A[i][j] == '|') {
				break
			}

			visited[i*n+j] = true

			if A[i][j] == '.' ||
				(A[i][j] == '-' && (dir == LEFT || dir == RIGHT)) ||
				(A[i][j] == '|' && (dir == UP || dir == DOWN)) {
				switch dir {
				case UP:
					i--
				case DOWN:
					i++
				case LEFT:
					j--
				case RIGHT:
					j++
				}
				continue
			}

			if A[i][j] == '\\' {
				switch dir {
				case UP:
					dir = LEFT
					j--
				case DOWN:
					dir = RIGHT
					j++
				case LEFT:
					dir = UP
					i--
				case RIGHT:
					dir = DOWN
					i++
				}
				continue
			}

			if A[i][j] == '/' {
				switch dir {
				case UP:
					dir = RIGHT
					j++
				case DOWN:
					dir = LEFT
					j--
				case LEFT:
					dir = DOWN
					i++
				case RIGHT:
					dir = UP
					i--
				}
				continue
			}

			if A[i][j] == '-' && (dir == UP || dir == DOWN) {
				points = append(points, LightPoint{row: i, col: j + 1, dir: RIGHT})
				dir = LEFT
				j--
				continue
			}

			if A[i][j] == '|' && (dir == LEFT || dir == RIGHT) {
				points = append(points, LightPoint{row: i + 1, col: j, dir: DOWN})
				dir = UP
				i--
			}
		}
	}

	energized := 0
	for _, v := range visited {
		if v {
			energized++
		}
	}
	return energized
}

func part1(A [][]rune) {
	start := LightPoint{row: 0, col: 0, dir: RIGHT}
	energized := solve(A, start)
	fmt.Println(energized)
}

func part2(A [][]rune) {
	m := len(A)
	n := len(A[0])
	max := -1

	for i := 0; i < m; i++ {
		start := LightPoint{row: i, col: 0, dir: RIGHT}
		energized := solve(A, start)
		if energized > max {
			max = energized
		}
		start = LightPoint{row: 0, col: n - 1, dir: LEFT}
		energized = solve(A, start)
		if energized > max {
			max = energized
		}
	}

	for j := 0; j < n; j++ {
		start := LightPoint{row: 0, col: j, dir: DOWN}
		energized := solve(A, start)
		if energized > max {
			max = energized
		}
		start = LightPoint{row: m - 1, col: j, dir: UP}
		energized = solve(A, start)
		if energized > max {
			max = energized
		}
	}

	fmt.Println(max)
}

func main() {
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

	// part1(A)
	part2(A)
}
