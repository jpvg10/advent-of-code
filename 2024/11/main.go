package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func convert(n int) []int {
	if n == 0 {
		return []int{1}
	}
	s := strconv.Itoa(n)
	if len(s)%2 == 0 {
		mid := len(s) / 2
		l := s[:mid]
		r := s[mid:]
		ln, _ := strconv.Atoi(l)
		rn, _ := strconv.Atoi(r)
		return []int{ln, rn}
	} else {
		return []int{n * 2024}
	}
}

func main() {
	// file, _ := os.Open("sample.txt")
	file, _ := os.Open("input.txt")
	defer file.Close()
	scanner := bufio.NewScanner(file)
	scanner.Scan()
	line := scanner.Text()

	fields := strings.Fields(line)
	nums := make(map[int]int)
	for _, s := range fields {
		n, _ := strconv.Atoi(s)
		nums[n]++
	}

	// for b := 1; b <= 25; b++ {
	for b := 1; b <= 75; b++ {
		current := make(map[int]int)
		for n, count := range nums {
			new := convert(n)
			for _, m := range new {
				current[m] += count
			}
		}
		nums = current
	}

	total := 0
	for _, count := range nums {
		total += count
	}
	fmt.Println(total)
}
