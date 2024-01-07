package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

type Hand struct {
	hand     string
	bid      int
	handType int
}

func getType(hand string) int {
	cardMap := make(map[rune]int)
	for _, r := range []rune(hand) {
		count, ok := cardMap[r]
		if ok {
			cardMap[r] = count + 1
		} else {
			cardMap[r] = 1
		}
	}

	if len(cardMap) == 1 {
		return 7 // five of a kind
	}
	if len(cardMap) == 4 {
		return 2 // 1 pair
	}
	if len(cardMap) == 5 {
		return 1 // high card
	}

	values := make([]int, 0, len(cardMap))
	for _, v := range cardMap {
		values = append(values, v)
	}
	if len(cardMap) == 2 {
		if values[0] == 4 || values[1] == 4 {
			return 6 // 4 of a kind
		} else {
			return 5 // full house
		}
	} else { // len(cardMap) == 3
		if values[0] == 3 || values[1] == 3 || values[2] == 3 {
			return 4 // 3 of a kind
		} else {
			return 3 // 2 pair
		}
	}
}

func getTypeWithJoker(hand string) int {
	jokers := 0
	for _, r := range []rune(hand) {
		if r == 'J' {
			jokers++
		}
	}
	if jokers == 0 {
		return getType(hand)
	}

	cardMap := make(map[rune]int)
	for _, r := range []rune(hand) {
		if r == 'J' {
			continue
		}
		count, ok := cardMap[r]
		if ok {
			cardMap[r] = count + 1
		} else {
			cardMap[r] = 1
		}
	}

	if len(cardMap) == 0 || len(cardMap) == 1 {
		return 7 // five of a kind
	}
	if len(cardMap) == 4 {
		return 2 // 1 pair
	}

	cards := 5 - jokers
	if len(cardMap) == 2 {
		if cards == 2 || cards == 3 {
			return 6 // 4 of a kind
		} else { // cards == 4
			values := make([]int, 0, len(cardMap))
			for _, v := range cardMap {
				values = append(values, v)
			}
			if values[0] == 3 || values[1] == 3 {
				return 6 // 4 of a kind
			} else {
				return 5 // full house
			}
		}
	} else { // len(cardMap) == 3
		return 4 // 3 of a kind
	}
}

func getCardNumber(a rune, withJoker bool) int {
	switch a {
	case 'T':
		return 10
	case 'J':
		if withJoker {
			return 1
		} else {
			return 11
		}
	case 'Q':
		return 12
	case 'K':
		return 13
	case 'A':
		return 14
	default:
		return int(a - '0')
	}
}

func part1() {
	solve(false)
}

func part2() {
	solve(true)
}

func solve(withJoker bool) {
	// file, _ := os.Open("sample.txt")
	file, _ := os.Open("input.txt")
	defer file.Close()
	scanner := bufio.NewScanner(file)

	hands := []Hand{}
	for scanner.Scan() {
		line := strings.Split(scanner.Text(), " ")
		hand := line[0]
		bid, _ := strconv.Atoi(line[1])
		handType := -1
		if withJoker {
			handType = getTypeWithJoker(hand)
		} else {
			handType = getType(hand)
		}
		hands = append(hands, Hand{hand: hand, bid: bid, handType: handType})
	}

	sort.Slice(hands, func(i, j int) bool {
		if hands[i].handType == hands[j].handType {
			for k := 0; k < 5; k++ {
				cardI := getCardNumber([]rune(hands[i].hand)[k], withJoker)
				cardJ := getCardNumber([]rune(hands[j].hand)[k], withJoker)
				if cardI < cardJ {
					return true
				} else if cardI > cardJ {
					return false
				}
			}
			return false
		} else {
			return hands[i].handType < hands[j].handType
		}
	})

	total := 0
	for i := 0; i < len(hands); i++ {
		total += (i + 1) * hands[i].bid
	}
	fmt.Println(total)
}

func main() {
	// part1()
	part2()
}
