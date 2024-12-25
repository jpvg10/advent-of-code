package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

func solve(xA, xB, yA, yB, xPrize, yPrize int) int {
	d := xA*yB - xB*yA
	dA := xPrize*yB - xB*yPrize
	dB := xA*yPrize - xPrize*yA

	if d != 0 && dA%d == 0 && dB%d == 0 {
		a := dA / d
		b := dB / d
		return (a*3 + b)
	} else {
		return 0
	}
}

func main() {
	// file, _ := os.Open("sample.txt")
	file, _ := os.Open("input.txt")
	defer file.Close()
	scanner := bufio.NewScanner(file)

	reButtons := regexp.MustCompile(`Button [A,B]: X\+([0-9]+), Y\+([0-9]+)`)
	rePrize := regexp.MustCompile(`Prize: X=([0-9]+), Y=([0-9]+)`)
	totalTokensPt1 := 0
	totalTokensPt2 := 0

	for scanner.Scan() {
		line := scanner.Text()
		matches := reButtons.FindStringSubmatch(line)
		xA, _ := strconv.Atoi(matches[1])
		yA, _ := strconv.Atoi(matches[2])

		scanner.Scan()
		line = scanner.Text()
		matches = reButtons.FindStringSubmatch(line)
		xB, _ := strconv.Atoi(matches[1])
		yB, _ := strconv.Atoi(matches[2])

		scanner.Scan()
		line = scanner.Text()
		matches = rePrize.FindStringSubmatch(line)
		xPrize, _ := strconv.Atoi(matches[1])
		yPrize, _ := strconv.Atoi(matches[2])

		// Blank line after group
		scanner.Scan()

		totalTokensPt1 += solve(xA, xB, yA, yB, xPrize, yPrize)
		totalTokensPt2 += solve(xA, xB, yA, yB, xPrize+10000000000000, yPrize+10000000000000)
	}

	fmt.Println("Part 1:", totalTokensPt1)
	fmt.Println("Part 2:", totalTokensPt2)
}
