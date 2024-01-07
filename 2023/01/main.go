package main

import (
	"bufio"
	"fmt"
	"os"
)

func part1() {
	// file, _ := os.Open("sample1.txt")
	file, _ := os.Open("input.txt")
	defer file.Close()
	scanner := bufio.NewScanner(file)

	totalSum := 0

	for scanner.Scan() {
		line := scanner.Text()
		first := -1
		last := -1

		for _, c := range line {
			if c >= '0' && c <= '9' {
				value := int(c - '0')
				if first == -1 {
					first = value
				}
				last = value
			}
		}

		totalSum += 10*first + last
	}

	fmt.Println(totalSum)
}

func part2() {
	// file, _ := os.Open("sample2.txt")
	file, _ := os.Open("input.txt")
	defer file.Close()
	scanner := bufio.NewScanner(file)

	totalSum := 0
	numbers := [10]string{"zero", "one", "two", "three", "four", "five", "six", "seven", "eight", "nine"}

	for scanner.Scan() {
		line := scanner.Text()
		first := -1
		last := -1
		i := 0

		for i < len(line) {
			c := line[i]
			value := -1

			if c >= '0' && c <= '9' {
				value = int(c - '0')
			} else if c >= 'a' && c <= 'z' {
				for j, n := range numbers {
					l := len(n)
					if len(line) >= i+l && line[i:i+l] == n {
						value = j
						break
					}
				}
			}

			if value != -1 {
				if first == -1 {
					first = value
				}
				last = value
			}

			i++
		}

		totalSum += 10*first + last
	}

	fmt.Println(totalSum)
}

func main() {
	// part1()
	part2()
}
