package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func isReportSafe(report []int) bool {
	isSafeIncreasing := true
	for i := 1; i < len(report); i++ {
		if report[i] <= report[i-1] || report[i]-report[i-1] > 3 {
			isSafeIncreasing = false
			break
		}
	}
	if isSafeIncreasing {
		return true
	}

	isSafeDecreasing := true
	for i := 1; i < len(report); i++ {
		if report[i-1] <= report[i] || report[i-1]-report[i] > 3 {
			isSafeDecreasing = false
			break
		}
	}
	return isSafeDecreasing
}

func main() {
	// file, _ := os.Open("sample.txt")
	file, _ := os.Open("input.txt")
	defer file.Close()
	scanner := bufio.NewScanner(file)

	safeCount := 0
	safeWithDampCount := 0

	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Fields(line)
		report := []int{}
		for _, p := range parts {
			level, _ := strconv.Atoi(p)
			report = append(report, level)
		}

		if isReportSafe(report) {
			safeCount++
			continue
		}

		// Check if safe by removing one element
		for i := 0; i < len(report); i++ {
			var newReport []int
			newReport = append(newReport, report[:i]...)
			newReport = append(newReport, report[i+1:]...)
			if isReportSafe(newReport) {
				safeWithDampCount++
				break
			}
		}
	}

	fmt.Println("Part 1:", safeCount)
	fmt.Println("Part 2:", safeCount+safeWithDampCount)
}
