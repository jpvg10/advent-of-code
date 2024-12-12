package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func checkTruePt1(result int, values []int, currentPos int, accum int) bool {
	current := values[currentPos]
	accumSum := accum + current
	var accumProd int
	if currentPos == 0 {
		accumProd = current
	} else {
		accumProd = accum * current
	}

	if currentPos == len(values)-1 {
		if accumSum == result || accumProd == result {
			return true
		} else {
			return false
		}
	}

	return checkTruePt1(result, values, currentPos+1, accumSum) ||
		checkTruePt1(result, values, currentPos+1, accumProd)
}

func checkTruePt2(result int, values []int, currentPos int, accum int) bool {
	current := values[currentPos]
	accumSum := accum + current
	accumConcat, _ := strconv.Atoi(strconv.Itoa(accum) + strconv.Itoa(current))
	var accumProd int
	if currentPos == 0 {
		accumProd = current
	} else {
		accumProd = accum * current
	}

	if currentPos == len(values)-1 {
		if accumSum == result || accumProd == result || accumConcat == result {
			return true
		} else {
			return false
		}
	}

	return checkTruePt2(result, values, currentPos+1, accumSum) ||
		checkTruePt2(result, values, currentPos+1, accumProd) ||
		checkTruePt2(result, values, currentPos+1, accumConcat)
}

func main() {
	// file, _ := os.Open("sample.txt")
	file, _ := os.Open("input.txt")
	defer file.Close()
	scanner := bufio.NewScanner(file)

	totalPt1 := 0
	totalPt2 := 0

	for scanner.Scan() {
		line := scanner.Text()
		line = strings.Replace(line, ":", "", -1)
		nums := strings.Split(line, " ")

		result, _ := strconv.Atoi(nums[0])
		values := []int{}
		for i := 1; i < len(nums); i++ {
			n, _ := strconv.Atoi(nums[i])
			values = append(values, n)
		}

		if checkTruePt1(result, values, 0, 0) {
			totalPt1 += result
		}

		if checkTruePt2(result, values, 0, 0) {
			totalPt2 += result
		}
	}

	fmt.Println("Part 1:", totalPt1)
	fmt.Println("Part 2:", totalPt2)
}
