package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func part1() {
	// file, _ := os.Open("sample.txt")
	file, _ := os.Open("input.txt")
	defer file.Close()

	scanner := bufio.NewScanner(file)
	workflowSection := true
	workflows := make(map[string][]string, 0)
	totalSum := 0

	for scanner.Scan() {
		line := scanner.Text()

		if line == "" {
			workflowSection = false
			continue
		}

		if workflowSection {
			pos := strings.Index(line, "{")
			name := line[0:pos]
			workflows[name] = strings.Split(line[pos+1:len(line)-1], ",")
		} else {
			re := regexp.MustCompile("[0-9]+")
			match := re.FindAllString(line, -1)
			x, _ := strconv.Atoi(match[0])
			m, _ := strconv.Atoi(match[1])
			a, _ := strconv.Atoi(match[2])
			s, _ := strconv.Atoi(match[3])

			w := "in"
			for w != "R" && w != "A" {
				rules := workflows[w]
				for _, r := range rules {
					pos := strings.Index(r, ":")
					if pos == -1 {
						w = r
						break
					}

					letter := r[0]
					operator := r[1]
					parts := strings.Split(r[2:], ":")
					value, _ := strconv.Atoi(parts[0])
					dest := parts[1]

					switch letter {
					case 'x':
						if (operator == '<' && x < value) || (operator == '>' && x > value) {
							w = dest
						}
					case 'm':
						if (operator == '<' && m < value) || (operator == '>' && m > value) {
							w = dest
						}
					case 'a':
						if (operator == '<' && a < value) || (operator == '>' && a > value) {
							w = dest
						}
					case 's':
						if (operator == '<' && s < value) || (operator == '>' && s > value) {
							w = dest
						}
					}

					if w == dest {
						break
					}
				}
			}

			if w == "A" {
				totalSum += x + m + a + s
			}
		}
	}

	fmt.Println(totalSum)
}

func main() {
	part1()
}
