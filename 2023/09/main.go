package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func part1() {
	// file, _ := os.Open("sample.txt")
	file, _ := os.Open("input.txt")
	defer file.Close()

	scanner := bufio.NewScanner(file)
	totalSum := 0

	for scanner.Scan() {
		stringSeq := strings.Fields(scanner.Text())
		n := len(stringSeq)
		seq := make([][]int, n+1)
		seq[0] = make([]int, n+1)
		for j, s := range stringSeq {
			seq[0][j], _ = strconv.Atoi(s)
		}

		last := 0
		for i := 1; i < n; i++ {
			seq[i] = make([]int, n+1)
			for j := 0; j < n-(i-1); j++ {
				seq[i][j] = seq[i-1][j-1] - seq[i-1][j]
			}
			allZero := true
			for j := 0; j < n-i; j++ {
				if seq[i][j] != 0 {
					allZero = false
					break
				}
			}
			if allZero {
				last = i - 1
				break
			}
		}

		for i := last; i >= 0; i-- {
			seq[i][n-i] = seq[i][n-i-1] + seq[i+1][n-i-1]
		}

		totalSum += seq[0][n]
	}

	fmt.Println(totalSum)
}

func part2() {
	// file, _ := os.Open("sample.txt")
	file, _ := os.Open("input.txt")
	defer file.Close()

	scanner := bufio.NewScanner(file)
	totalSum := 0

	for scanner.Scan() {
		stringSeq := strings.Fields(scanner.Text())
		n := len(stringSeq)
		seq := make([][]int, n+1)
		seq[0] = make([]int, n+1)
		for j, s := range stringSeq {
			seq[0][j+1], _ = strconv.Atoi(s)
		}

		last := 0
		for i := 1; i < n; i++ {
			seq[i] = make([]int, n+1)
			for j := n; j > i; j-- {
				seq[i][j] = seq[i-1][j] - seq[i-1][j-1]
			}
			allZero := true
			for j := n; j > i; j-- {
				if seq[i][j] != 0 {
					allZero = false
					break
				}
			}
			if allZero {
				last = i - 1
				break
			}
		}

		for i := last; i >= 0; i-- {
			seq[i][i] = seq[i][i+1] - seq[i+1][i+1]
		}

		totalSum += seq[0][0]
	}

	fmt.Println(totalSum)
}

func main() {
	// part1()
	part2()
}
