package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
)

type Node struct {
	left  string
	right string
}

func part1() {
	// file, _ := os.Open("sample1.txt")
	// file, _ := os.Open("sample2.txt")
	file, _ := os.Open("input.txt")
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Scan()
	moves := []rune(scanner.Text())
	scanner.Scan() // blank line
	nodesMap := make(map[string]Node)
	re := regexp.MustCompile("[A-Z]+")

	for scanner.Scan() {
		nodes := re.FindAllString(scanner.Text(), -1)
		nodesMap[nodes[0]] = Node{left: nodes[1], right: nodes[2]}
	}

	current := "AAA"
	i := 0
	count := 0
	for current != "ZZZ" {
		if moves[i] == 'L' {
			current = nodesMap[current].left
		} else {
			current = nodesMap[current].right
		}
		count++
		i = (i + 1) % len(moves)
	}

	fmt.Println(count)
}

func lcm(a, b int) int {
	return (a * b) / gcd(a, b)
}

// Euclidean algorithm
func gcd(a, b int) int {
	if b == 0 {
		return a
	} else {
		return gcd(b, a%b)
	}
}

func part2() {
	// file, _ := os.Open("sample3.txt")
	file, _ := os.Open("input.txt")
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Scan()
	moves := []rune(scanner.Text())
	scanner.Scan() // blank line
	nodesMap := make(map[string]Node)
	starting := []string{}
	re := regexp.MustCompile("[A-Z0-9]+")

	for scanner.Scan() {
		nodes := re.FindAllString(scanner.Text(), -1)
		nodesMap[nodes[0]] = Node{left: nodes[1], right: nodes[2]}
		if nodes[0][len(nodes[0])-1] == 'A' {
			starting = append(starting, nodes[0])
		}
	}

	fmt.Println(starting)

	cycleLength := make([]int, len(starting))
	for i, s := range starting {
		cycleStart := -1
		foundCycle := false
		j := 0
		count := 1
		current := s

		for !foundCycle {
			if moves[j] == 'L' {
				current = nodesMap[current].left
			} else {
				current = nodesMap[current].right
			}
			if current[len(current)-1] == 'Z' {
				if cycleStart == -1 {
					cycleStart = count
				} else {
					foundCycle = true
					cycleLength[i] = count - cycleStart
				}
			}
			j = (j + 1) % len(moves)
			count++
		}
	}

	fmt.Println(cycleLength)

	// Using LCM doesn't work in general but happens to work in this particular input
	// This is not obvious from the problem statement. Found the hint looking online
	ans := cycleLength[0]
	for i := 1; i < len(cycleLength); i++ {
		ans = lcm(ans, cycleLength[i])
	}
	fmt.Println(ans)
}

func main() {
	// part1()
	part2()
}
