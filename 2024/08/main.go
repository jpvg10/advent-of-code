package main

import (
	"bufio"
	"fmt"
	"os"
)

type coordinate struct {
	i int
	j int
}

func part1(grid [][]rune) {
	m := len(grid)
	n := len(grid[0])
	antennas := make(map[rune][]coordinate)
	antinodes := 0

	gridAnti := make([][]bool, m)
	for i := 0; i < m; i++ {
		gridAnti[i] = make([]bool, n)
	}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			symbol := grid[i][j]
			if symbol == '.' {
				continue
			}

			for _, c := range antennas[symbol] {
				di := i - c.i
				dj := j - c.j

				antinodeI := i + di
				antinodeJ := j + dj
				if antinodeI >= 0 && antinodeI < m && antinodeJ >= 0 && antinodeJ < n && !gridAnti[antinodeI][antinodeJ] {
					gridAnti[antinodeI][antinodeJ] = true
					antinodes++
				}

				antinodeI = c.i - di
				antinodeJ = c.j - dj
				if antinodeI >= 0 && antinodeI < m && antinodeJ >= 0 && antinodeJ < n && !gridAnti[antinodeI][antinodeJ] {
					gridAnti[antinodeI][antinodeJ] = true
					antinodes++
				}
			}

			antennas[symbol] = append(antennas[symbol], coordinate{i, j})
		}
	}

	fmt.Println("Part 1:", antinodes)
}

func part2(grid [][]rune) {
	m := len(grid)
	n := len(grid[0])
	antennas := make(map[rune][]coordinate)
	antinodes := 0

	gridAnti := make([][]bool, m)
	for i := 0; i < m; i++ {
		gridAnti[i] = make([]bool, n)
	}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			symbol := grid[i][j]
			if symbol == '.' {
				continue
			}

			for _, c := range antennas[symbol] {
				di := i - c.i
				dj := j - c.j

				antinodeI := i
				antinodeJ := j
				for {
					if antinodeI >= 0 && antinodeI < m && antinodeJ >= 0 && antinodeJ < n {
						if !gridAnti[antinodeI][antinodeJ] {
							gridAnti[antinodeI][antinodeJ] = true
							antinodes++
						}

						antinodeI += di
						antinodeJ += dj
					} else {
						break
					}
				}

				antinodeI = c.i
				antinodeJ = c.j
				for {
					if antinodeI >= 0 && antinodeI < m && antinodeJ >= 0 && antinodeJ < n {
						if !gridAnti[antinodeI][antinodeJ] {
							gridAnti[antinodeI][antinodeJ] = true
							antinodes++
						}

						antinodeI -= di
						antinodeJ -= dj
					} else {
						break
					}
				}
			}

			antennas[symbol] = append(antennas[symbol], coordinate{i, j})
		}
	}

	fmt.Println("Part 2:", antinodes)
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

	part1(grid)
	part2(grid)
}
