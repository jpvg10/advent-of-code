package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
)

func part1(M [][]rune, m int, n int, gI int, gJ int) {
	visited := make(map[int]bool)
	gDirection := 0

	for {
		current := gI*n + gJ
		visited[current] = true
		nextGI, nextGJ := gI, gJ

		if gDirection == 0 {
			nextGI--
		} else if gDirection == 1 {
			nextGJ++
		} else if gDirection == 2 {
			nextGI++
		} else {
			nextGJ--
		}

		if nextGI < 0 || nextGI >= m || nextGJ < 0 || nextGJ >= n {
			break
		}

		if M[nextGI][nextGJ] == '#' {
			gDirection = (gDirection + 1) % 4
			continue
		}

		gI, gJ = nextGI, nextGJ
	}

	fmt.Println("Part 1:", len(visited))
}

func part2(M [][]rune, m int, n int, gI int, gJ int) {
	visited := make(map[int]bool)
	gDirection := 0
	loops := 0

	for {
		current := gI*n + gJ
		visited[current] = true
		nextGI, nextGJ := gI, gJ

		if gDirection == 0 { // Going up
			nextGI--
		} else if gDirection == 1 { // Going right
			nextGJ++
		} else if gDirection == 2 { // Going down
			nextGI++
		} else { // Going left
			nextGJ--
		}

		if nextGI < 0 || nextGI >= m || nextGJ < 0 || nextGJ >= n {
			break
		}

		if M[nextGI][nextGJ] == '#' {
			gDirection = (gDirection + 1) % 4
			continue
		}

		next := nextGI*m + nextGJ
		// If haven't visited the next position, check if putting an obstacle there will create a loop
		if _, ok := visited[next]; !ok {
			doesLoop := testLoop(M, m, n, gI, gJ, nextGI, nextGJ, gDirection)
			if doesLoop {
				loops++
			}
		}

		gI, gJ = nextGI, nextGJ
	}

	fmt.Println("Part 2:", loops)
}

func testLoop(M [][]rune, m int, n int, gI int, gJ int, obstacleI int, obstacleJ int, gDirection int) bool {
	visited := make(map[int][]int)

	for {
		current := gI*n + gJ

		// If have already visited this position while going in the same direction, there is a loop
		if dirs, ok := visited[current]; ok && slices.Contains(dirs, gDirection) {
			return true
		}

		visited[current] = append(visited[current], gDirection)
		nextGI, nextGJ := gI, gJ

		if gDirection == 0 {
			nextGI--
		} else if gDirection == 1 {
			nextGJ++
		} else if gDirection == 2 {
			nextGI++
		} else {
			nextGJ--
		}

		if nextGI < 0 || nextGI >= m || nextGJ < 0 || nextGJ >= n {
			break
		}

		if M[nextGI][nextGJ] == '#' || (nextGI == obstacleI && nextGJ == obstacleJ) {
			gDirection = (gDirection + 1) % 4
			continue
		}

		gI, gJ = nextGI, nextGJ
	}

	return false
}

func main() {
	// file, _ := os.Open("sample.txt")
	file, _ := os.Open("input.txt")
	defer file.Close()
	scanner := bufio.NewScanner(file)

	var M [][]rune
	for scanner.Scan() {
		line := scanner.Text()
		M = append(M, []rune(line))
	}

	m := len(M)
	n := len(M[0])

	gI, gJ := 0, 0
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if M[i][j] == '^' {
				gI, gJ = i, j
			}
		}
	}

	part1(M, m, n, gI, gJ)
	part2(M, m, n, gI, gJ)
}
