package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"slices"
	"strconv"
	"strings"
)

type coordinate struct {
	y int
	x int
}

func bfsShortestHop(blocks [][]bool) (int, bool) {
	n := len(blocks)
	start := coordinate{0, 0}
	end := coordinate{n - 1, n - 1}
	q := []coordinate{start}

	dist := make([][]int, n)
	for i := 0; i < n; i++ {
		dist[i] = slices.Repeat([]int{math.MaxInt}, n)
	}
	dist[start.y][start.x] = 0

	for len(q) > 0 {
		u := q[0]
		q = q[1:]

		if u.x > 0 && !blocks[u.y][u.x-1] && dist[u.y][u.x-1] == math.MaxInt {
			q = append(q, coordinate{u.y, u.x - 1})
			dist[u.y][u.x-1] = dist[u.y][u.x] + 1
		}
		if u.x < n-1 && !blocks[u.y][u.x+1] && dist[u.y][u.x+1] == math.MaxInt {
			q = append(q, coordinate{u.y, u.x + 1})
			dist[u.y][u.x+1] = dist[u.y][u.x] + 1
		}
		if u.y > 0 && !blocks[u.y-1][u.x] && dist[u.y-1][u.x] == math.MaxInt {
			q = append(q, coordinate{u.y - 1, u.x})
			dist[u.y-1][u.x] = dist[u.y][u.x] + 1
		}
		if u.y < n-1 && !blocks[u.y+1][u.x] && dist[u.y+1][u.x] == math.MaxInt {
			q = append(q, coordinate{u.y + 1, u.x})
			dist[u.y+1][u.x] = dist[u.y][u.x] + 1
		}
	}

	possible := (dist[end.y][end.x] != math.MaxInt)
	return dist[end.y][end.x], possible
}

func main() {
	// file, _ := os.Open("sample.txt")
	// n, part1Fall := 7, 12
	file, _ := os.Open("input.txt")
	n, part1Fall := 71, 1024
	defer file.Close()
	scanner := bufio.NewScanner(file)

	blocks := make([][]bool, n)
	for i := 0; i < n; i++ {
		blocks[i] = make([]bool, n)
	}
	fell := 0

	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Split(line, ",")
		x, _ := strconv.Atoi(parts[0])
		y, _ := strconv.Atoi(parts[1])
		blocks[y][x] = true
		fell++
		if fell == part1Fall {
			ans, _ := bfsShortestHop(blocks)
			fmt.Println("Part 1:", ans)
		} else if fell > part1Fall {
			_, possible := bfsShortestHop(blocks)
			if !possible {
				fmt.Printf("Part 2: %d,%d\n", x, y)
				break
			}
		}
	}
}
