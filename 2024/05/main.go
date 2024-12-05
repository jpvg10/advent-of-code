package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

func main() {
	// file, _ := os.Open("sample.txt")
	file, _ := os.Open("input.txt")
	defer file.Close()
	scanner := bufio.NewScanner(file)

	mapPages := make(map[string][]string)
	rulesSection := true
	middleSumOk := 0
	middleSumNotOk := 0

	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			rulesSection = false
			continue
		}

		if rulesSection {
			parts := strings.Split(line, "|")
			u, v := parts[0], parts[1]
			mapPages[u] = append(mapPages[u], v)
		} else {
			update := strings.Split(line, ",")
			isOk := true

			for i := 1; i < len(update); i++ {
				u, v := update[i-1], update[i]
				if !slices.Contains(mapPages[u], v) {
					isOk = false
					break
				}
			}

			middlePos := len(update) / 2
			if isOk {
				middlePage, _ := strconv.Atoi(update[middlePos])
				middleSumOk += middlePage
			} else {
				slices.SortFunc(update, func(u, v string) int {
					if slices.Contains(mapPages[u], v) {
						return -1
					} else if slices.Contains(mapPages[v], u) {
						return 1
					}
					return 0
				})
				middlePage, _ := strconv.Atoi(update[middlePos])
				middleSumNotOk += middlePage
			}
		}
	}

	fmt.Println("Part 1:", middleSumOk)
	fmt.Println("Part 2:", middleSumNotOk)
}
