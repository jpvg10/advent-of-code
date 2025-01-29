package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func part1() {
	// file, _ := os.Open("sample1.txt")
	file, _ := os.Open("input.txt")
	defer file.Close()
	scanner := bufio.NewScanner(file)

	re := regexp.MustCompile(`mul\(([0-9]{1,3}),([0-9]{1,3})\)`)
	productSum := 0

	for scanner.Scan() {
		line := scanner.Text()
		matches := re.FindAllStringSubmatch(line, -1)

		for _, match := range matches {
			a, _ := strconv.Atoi(match[1])
			b, _ := strconv.Atoi(match[2])
			productSum += (a * b)
		}
	}

	fmt.Println("Part 1:", productSum)
}

func part2() {
	// file, _ := os.Open("sample2.txt")
	file, _ := os.Open("input.txt")
	defer file.Close()
	scanner := bufio.NewScanner(file)

	re := regexp.MustCompile(`do\(\)|don't\(\)|mul\(([0-9]{1,3}),([0-9]{1,3})\)`)
	productSum := 0
	isEnabled := true

	for scanner.Scan() {
		line := scanner.Text()
		matches := re.FindAllStringSubmatch(line, -1)

		for _, match := range matches {
			if match[0] == "do()" {
				isEnabled = true
			} else if match[0] == "don't()" {
				isEnabled = false
			} else if isEnabled && strings.HasPrefix(match[0], "mul") {
				a, _ := strconv.Atoi(match[1])
				b, _ := strconv.Atoi(match[2])
				productSum += (a * b)
			}
		}
	}

	fmt.Println("Part 2:", productSum)
}

func main() {
	part1()
	part2()
}
