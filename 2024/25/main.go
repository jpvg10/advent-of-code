package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	// file, _ := os.Open("sample.txt")
	file, _ := os.Open("input.txt")
	defer file.Close()
	scanner := bufio.NewScanner(file)

	locks := [][]int{}
	keys := [][]int{}

	isFirstLine := true
	isLock := true
	height := 0
	currentHeight := 0

	for scanner.Scan() {
		line := scanner.Text()

		if line == "" {
			isFirstLine = true
			if height == 0 {
				height = currentHeight
			}
			continue
		}

		currentHeight++

		n := len(line)
		if isFirstLine {
			isFirstLine = false
			if line[0] == '#' {
				isLock = true
				locks = append(locks, make([]int, n))
			} else {
				isLock = false
				keys = append(keys, make([]int, n))
			}
		}

		if isLock {
			current := len(locks) - 1
			for i, c := range line {
				if c == '#' {
					locks[current][i]++
				}
			}
		} else {
			current := len(keys) - 1
			for i, c := range line {
				if c == '#' {
					keys[current][i]++
				}
			}
		}
	}

	totalFits := 0
	for _, lock := range locks {
		for _, key := range keys {
			n := len(lock)
			fits := true
			for i := 0; i < n; i++ {
				if lock[i]+key[i] > height {
					fits = false
					break
				}
			}
			if fits {
				totalFits++
			}
		}
	}

	fmt.Println(totalFits)
}
