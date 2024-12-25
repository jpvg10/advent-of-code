package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	// file, _ := os.Open("sample.txt")
	file, _ := os.Open("input.txt")
	defer file.Close()
	scanner := bufio.NewScanner(file)

	totalSum := 0

	for scanner.Scan() {
		line := scanner.Text()
		num, _ := strconv.Atoi(line)

		for i := 0; i < 2000; i++ {
			temp := num * 64
			num = (num ^ temp) % 16777216
			temp = num / 32
			num = (num ^ temp) % 16777216
			temp = num * 2048
			num = (num ^ temp) % 16777216
		}
		// fmt.Println(num)

		totalSum += num
	}

	fmt.Println(totalSum)
}
