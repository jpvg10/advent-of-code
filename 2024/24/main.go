package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

type instruction struct {
	left           string
	operand        string
	right          string
	resultLocation string
}

func main() {
	// file, _ := os.Open("sample.txt")
	file, _ := os.Open("input.txt")
	defer file.Close()
	scanner := bufio.NewScanner(file)

	wires := make(map[string]int)
	instructions := []instruction{}
	wiresSection := true
	zWiresCount := 0

	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			wiresSection = false
			continue
		}

		if wiresSection {
			parts := strings.Split(line, ": ")
			value, _ := strconv.Atoi(parts[1])
			wires[parts[0]] = value
		} else {
			parts := strings.Split(line, " ")
			instructions = append(instructions, instruction{parts[0], parts[1], parts[2], parts[4]})
			if parts[4][0] == 'z' {
				zWiresCount++
			}
		}
	}

	zWires := make([]string, zWiresCount)

	for len(instructions) > 0 {
		current := instructions[0]
		instructions = instructions[1:]

		leftValue, hasLeft := wires[current.left]
		rightValue, hasRight := wires[current.right]

		if hasLeft && hasRight {
			result := -1
			if current.operand == "AND" {
				result = leftValue & rightValue
			} else if current.operand == "OR" {
				result = leftValue | rightValue
			} else {
				result = leftValue ^ rightValue
			}
			wires[current.resultLocation] = result

			if current.resultLocation[0] == 'z' {
				number := current.resultLocation[1:]
				index, _ := strconv.Atoi(number)
				resultString := strconv.Itoa(result)
				zWires[index] = resultString
			}
		} else {
			instructions = append(instructions, current)
		}
	}

	slices.Reverse(zWires)
	binaryString := strings.Join(zWires, "")
	resultNumber, _ := strconv.ParseInt(binaryString, 2, 0)
	fmt.Println(resultNumber)
}
