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
	scanner.Scan()
	steps := strings.Split(scanner.Text(), ",")
	totalSum := 0

	for _, s := range steps {
		hash := 0
		for _, c := range s {
			hash += int(c)
			hash *= 17
			hash %= 256
		}
		totalSum += hash
	}

	fmt.Println(totalSum)
}

func main() {
	part1()
}
