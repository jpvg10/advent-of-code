package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func part1() {
	// file, _ := os.Open("sample.txt")
	file, _ := os.Open("input.txt")
	defer file.Close()
	scanner := bufio.NewScanner(file)

	pointSum := 0

	for scanner.Scan() {
		line := scanner.Text()
		line = strings.Split(line, ":")[1]
		parts := strings.Split(line, "|")
		win := strings.Fields(parts[0])
		have := strings.Fields(parts[1])
		numMap := make(map[string]bool)

		for _, num := range win {
			numMap[num] = true
		}
		score := 0
		for _, num := range have {
			_, ok := numMap[num]
			if ok {
				if score == 0 {
					score = 1
				} else {
					score = score * 2
				}
			}
		}
		pointSum += score
	}

	fmt.Println(pointSum)
}

func part2() {
	// file, _ := os.Open("sample.txt")
	file, _ := os.Open("input.txt")
	defer file.Close()
	scanner := bufio.NewScanner(file)

	i := 1
	copies := make(map[int]int)

	for scanner.Scan() {
		line := scanner.Text()
		line = strings.Split(line, ":")[1]
		parts := strings.Split(line, "|")
		win := strings.Fields(parts[0])
		have := strings.Fields(parts[1])
		numMap := make(map[string]bool)

		val, ok := copies[i]
		if ok {
			copies[i] = val + 1
		} else {
			copies[i] = 1
		}

		for _, num := range win {
			numMap[num] = true
		}
		matching := 0
		for _, num := range have {
			_, ok := numMap[num]
			if ok {
				matching++
			}
		}

		for j := 1; j <= matching; j++ {
			val, ok := copies[i+j]
			if ok {
				copies[i+j] = val + copies[i]
			} else {
				copies[i+j] = copies[i]
			}
		}

		i++
	}

	totalCards := 0
	for _, val := range copies {
		totalCards += val
	}
	// fmt.Println(copies)
	fmt.Println(totalCards)
}

func main() {
	// part1()
	part2()
}
