package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

func part1(left []int, right []int) {
	slices.Sort(left)
	slices.Sort(right)

	totalDiff := 0

	for i := 0; i < len(left); i++ {
		if left[i] < right[i] {
			totalDiff += (right[i] - left[i])
		} else {
			totalDiff += (left[i] - right[i])
		}
	}

	fmt.Println("Part 1:", totalDiff)
}

func part2(left []int, right []int) {
	freqInRight := make(map[int]int)
	for _, r := range right {
		if freq, ok := freqInRight[r]; ok {
			freqInRight[r] = freq + 1
		} else {
			freqInRight[r] = 1
		}
	}

	score := 0
	for _, l := range left {
		if freq, ok := freqInRight[l]; ok {
			score += l * freq
		}
	}

	fmt.Println("Part 2:", score)
}

func main() {
	// file, _ := os.Open("sample.txt")
	file, _ := os.Open("input.txt")
	defer file.Close()
	scanner := bufio.NewScanner(file)

	left := []int{}
	right := []int{}

	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Fields(line)
		l, _ := strconv.Atoi(parts[0])
		r, _ := strconv.Atoi(parts[1])
		left = append(left, l)
		right = append(right, r)
	}

	part1(left, right)
	part2(left, right)
}
