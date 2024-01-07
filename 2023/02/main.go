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

	possibleIdSum := 0
	redNum := 12
	greenNum := 13
	blueNum := 14
	i := 1

	for scanner.Scan() {
		game := strings.Split(scanner.Text(), ":")[1]
		game = strings.TrimSpace(game)
		game = strings.ReplaceAll(game, ";", ",")
		colors := strings.Split(game, ",")
		possible := true

		for _, color := range colors {
			color = strings.TrimSpace(color)
			parts := strings.Split(color, " ")
			num, _ := strconv.Atoi(parts[0])
			if (parts[1] == "red" && num > redNum) ||
				(parts[1] == "green" && num > greenNum) ||
				(parts[1] == "blue" && num > blueNum) {
				possible = false
				break
			}
		}

		if possible {
			possibleIdSum += i
		}

		i++
	}

	fmt.Println(possibleIdSum)
}

func part2() {
	// file, _ := os.Open("sample.txt")
	file, _ := os.Open("input.txt")
	defer file.Close()
	scanner := bufio.NewScanner(file)

	powerSum := 0

	for scanner.Scan() {
		game := strings.Split(scanner.Text(), ":")[1]
		game = strings.TrimSpace(game)
		game = strings.ReplaceAll(game, ";", ",")
		colors := strings.Split(game, ",")
		redMax := 0
		greenMax := 0
		blueMax := 0

		for _, color := range colors {
			color = strings.TrimSpace(color)
			parts := strings.Split(color, " ")
			num, _ := strconv.Atoi(parts[0])
			if parts[1] == "red" && num > redMax {
				redMax = num
			} else if parts[1] == "green" && num > greenMax {
				greenMax = num
			} else if parts[1] == "blue" && num > blueMax {
				blueMax = num
			}
		}

		power := redMax * greenMax * blueMax
		powerSum += power
	}

	fmt.Println(powerSum)
}

func main() {
	// part1()
	part2()
}
