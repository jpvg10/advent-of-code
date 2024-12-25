package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strings"
)

func main() {
	// file, _ := os.Open("sample.txt")
	file, _ := os.Open("input.txt")
	defer file.Close()
	scanner := bufio.NewScanner(file)

	graph := make(map[string][]string)

	for scanner.Scan() {
		line := scanner.Text()
		computers := strings.Split(line, "-")
		c1 := computers[0]
		c2 := computers[1]
		graph[c1] = append(graph[c1], c2)
		graph[c2] = append(graph[c2], c1)
	}

	groups := 0
	found := make(map[string]bool)

	for c1 := range graph {
		if c1[0] != 't' {
			continue
		}

		for _, c2 := range graph[c1] {
			for _, c3 := range graph[c2] {
				if slices.Contains(graph[c1], c3) {
					group := []string{c1, c2, c3}
					slices.Sort(group)
					key := group[0] + "-" + group[1] + "-" + group[2]
					if !found[key] {
						groups++
						found[key] = true
					}
				}
			}
		}
	}

	fmt.Println(groups)
}
