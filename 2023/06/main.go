package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func solve(t, d int) int {
	count := 0
	for j := 0; j <= t/2; j++ {
		if j*(t-j) > d {
			if t%2 == 0 && j == t/2 {
				count++
			} else {
				count += 2
			}
		}
	}
	return count
}

func part1() {
	// file, _ := os.Open("sample.txt")
	file, _ := os.Open("input.txt")
	defer file.Close()
	scanner := bufio.NewScanner(file)

	scanner.Scan()
	times := strings.Fields(strings.Split(scanner.Text(), ":")[1])
	scanner.Scan()
	distances := strings.Fields(strings.Split(scanner.Text(), ":")[1])
	n := len(times)
	countProduct := 1

	for i := 0; i < n; i++ {
		t, _ := strconv.Atoi(times[i])
		d, _ := strconv.Atoi(distances[i])
		count := solve(t, d)
		countProduct *= count
	}

	fmt.Println(countProduct)
}

func part2() {
	// file, _ := os.Open("sample.txt")
	file, _ := os.Open("input.txt")
	defer file.Close()
	scanner := bufio.NewScanner(file)

	scanner.Scan()
	time := strings.ReplaceAll(strings.Split(scanner.Text(), ":")[1], " ", "")
	scanner.Scan()
	distance := strings.ReplaceAll(strings.Split(scanner.Text(), ":")[1], " ", "")

	t, _ := strconv.Atoi(time)
	d, _ := strconv.Atoi(distance)
	count := solve(t, d)
	fmt.Println(count)
}

func main() {
	// part1()
	part2()
}
