package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func isPossiblePt1(design string, patternsMap map[string]bool, pos int) bool {
	designFragment := design[pos:]
	if patternsMap[designFragment] {
		return true
	}

	for i := pos + 1; i < len(design); i++ {
		prefix := design[pos:i]
		if patternsMap[prefix] && isPossiblePt1(design, patternsMap, i) {
			return true
		}
	}

	return false
}

func isPossiblePt2(design string, patternsMap map[string]bool, pos int, cacheWays map[string]int) int {
	designFragment := design[pos:]
	if count, ok := cacheWays[designFragment]; ok {
		return count
	}

	ways := 0
	if patternsMap[designFragment] {
		ways++
	}

	for i := pos + 1; i < len(design); i++ {
		prefix := design[pos:i]
		waysSuffix := isPossiblePt2(design, patternsMap, i, cacheWays)
		if patternsMap[prefix] && waysSuffix > 0 {
			ways += waysSuffix
		}
	}

	cacheWays[designFragment] = ways
	return ways
}

func main() {
	// file, _ := os.Open("sample.txt")
	file, _ := os.Open("input.txt")
	defer file.Close()
	scanner := bufio.NewScanner(file)

	scanner.Scan()
	line := scanner.Text()
	patterns := strings.Split(line, ", ")
	patternsMap := make(map[string]bool)
	for _, p := range patterns {
		patternsMap[p] = true
	}

	// Blank line after patterns
	scanner.Scan()

	possible := 0
	totalWays := 0
	cacheWays := make(map[string]int)

	for scanner.Scan() {
		design := scanner.Text()
		if isPossiblePt1(design, patternsMap, 0) {
			possible++
		}
		totalWays += isPossiblePt2(design, patternsMap, 0, cacheWays)
	}

	fmt.Println("Part 1:", possible)
	fmt.Println("Part 2:", totalWays)
}
