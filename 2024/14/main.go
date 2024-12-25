package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

type robot struct {
	x  int
	y  int
	vx int
	vy int
}

func main() {
	// file, _ := os.Open("sample.txt")
	// m, n := 7, 11
	file, _ := os.Open("input.txt")
	m, n := 103, 101
	defer file.Close()
	scanner := bufio.NewScanner(file)

	re := regexp.MustCompile(`p=(-?[0-9]+),(-?[0-9]+) v=(-?[0-9]+),(-?[0-9]+)`)
	seconds := 100
	robots := []robot{}

	for scanner.Scan() {
		line := scanner.Text()
		matches := re.FindStringSubmatch(line)
		x, _ := strconv.Atoi(matches[1])
		y, _ := strconv.Atoi(matches[2])
		vx, _ := strconv.Atoi(matches[3])
		vy, _ := strconv.Atoi(matches[4])
		robots = append(robots, robot{x, y, vx, vy})
	}

	q1 := 0
	q2 := 0
	q3 := 0
	q4 := 0
	halfM := m / 2
	halfN := n / 2

	for _, r := range robots {
		finalX := ((r.x+seconds*r.vx)%n + n) % n
		finalY := ((r.y+seconds*r.vy)%m + m) % m

		if finalX < halfN && finalY < halfM {
			q1++
		} else if finalX > halfN && finalY < halfM {
			q2++
		} else if finalX < halfN && finalY > halfM {
			q3++
		} else if finalX > halfN && finalY > halfM {
			q4++
		}
	}

	fmt.Println(q1 * q2 * q3 * q4)
}
