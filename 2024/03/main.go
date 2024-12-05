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

	fmt.Println(productSum)
}

func part2() {
	// file, _ := os.Open("sample2.txt")
	file, _ := os.Open("input.txt")
	defer file.Close()
	scanner := bufio.NewScanner(file)

	mulRE := regexp.MustCompile(`mul\(([0-9]{1,3}),([0-9]{1,3})\)`)
	doRE := regexp.MustCompile(`do\(\)`)
	dontRE := regexp.MustCompile(`don't\(\)`)
	productSum := 0
	enabled := true

	for scanner.Scan() {
		line := scanner.Text()
		doIndices := doRE.FindAllStringIndex(line, -1)
		dontIndices := dontRE.FindAllStringIndex(line, -1)
		mulSubIndices := mulRE.FindAllStringSubmatchIndex(line, -1)
		// fmt.Println(doIndices, dontIndices, mulSubIndices)

		mapIndices := make(map[int]string)
		for _, do := range doIndices {
			mapIndices[do[0]] = "do"
		}
		for _, dont := range dontIndices {
			mapIndices[dont[0]] = "dont"
		}
		for _, mul := range mulSubIndices {
			mapIndices[mul[0]] = "mul"
		}

		nextMatch := 0

		for i := 0; i < len(line); i++ {
			val, ok := mapIndices[i]
			if ok && val == "do" {
				enabled = true
			} else if ok && val == "dont" {
				enabled = false
			} else if ok && val == "mul" {
				if enabled {
					match := mulSubIndices[nextMatch]
					a, _ := strconv.Atoi(string(line[match[2]:match[3]]))
					b, _ := strconv.Atoi(string(line[match[4]:match[5]]))
					productSum += (a * b)
				}
				nextMatch++
			}
		}
	}

	fmt.Println(productSum)
}

func part2B() {
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

	fmt.Println(productSum)
}

func main() {
	// part1()
	part2B()
}
