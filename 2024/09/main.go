package main

import (
	"bufio"
	"fmt"
	"os"
)

type blockChunk struct {
	id       int
	startPos int
	length   int
}

func getChecksum(blockChunks []blockChunk) int {
	checksum := 0

	for _, b := range blockChunks {
		if b.length == 0 {
			continue
		}

		endPos := b.startPos + b.length - 1
		sum := (b.length * (b.startPos + endPos)) / 2
		checksum += b.id * sum
	}

	return checksum
}

func getCopy(chunks []blockChunk) []blockChunk {
	newChunks := make([]blockChunk, len(chunks))
	copy(newChunks, chunks)
	return newChunks
}

func part1(blockChunks []blockChunk, spaceChunks []blockChunk) {
	i := 0
	j := len(blockChunks) - 1
	newBlockChunks := []blockChunk{}

	for spaceChunks[i].startPos < blockChunks[j].startPos {
		new := blockChunk{
			id:       blockChunks[j].id,
			startPos: spaceChunks[i].startPos,
		}

		spaces := spaceChunks[i].length
		blocks := blockChunks[j].length

		if spaces <= blocks {
			// No empty space left. Some file blocks might be
			new.length = spaces
			blockChunks[j].length -= spaces
			if blocks == spaces {
				j--
			}
			i++
		} else {
			// Some empty space left. No file blocks left
			new.length = blocks
			blockChunks[j].length = 0
			spaceChunks[i].length -= blocks
			spaceChunks[i].startPos += blocks
			j--
		}

		newBlockChunks = append(newBlockChunks, new)
	}

	checksum := getChecksum(append(blockChunks, newBlockChunks...))
	fmt.Println("Part 1:", checksum)
}

func part2(blockChunks []blockChunk, spaceChunks []blockChunk) {
	newBlockChunks := []blockChunk{}

	for j := len(blockChunks) - 1; j >= 0; j-- {
		for i := 0; i < len(spaceChunks); i++ {
			if spaceChunks[i].startPos >= blockChunks[j].startPos {
				break
			}

			spaces := spaceChunks[i].length
			blocks := blockChunks[j].length

			if spaces >= blocks {
				new := blockChunk{
					id:       blockChunks[j].id,
					startPos: spaceChunks[i].startPos,
					length:   blocks,
				}
				newBlockChunks = append(newBlockChunks, new)
				blockChunks[j].length = 0
				spaceChunks[i].length -= blocks
				spaceChunks[i].startPos += blocks
				break
			}
		}
	}

	checksum := getChecksum(append(blockChunks, newBlockChunks...))
	fmt.Println("Part 2:", checksum)
}

func main() {
	// file, _ := os.Open("sample.txt")
	file, _ := os.Open("input.txt")
	defer file.Close()
	scanner := bufio.NewScanner(file)
	scanner.Scan()
	line := scanner.Text()

	nums := []int{}
	for _, c := range line {
		n := int(c - '0')
		nums = append(nums, n)
	}

	blockChunks := []blockChunk{}
	spaceChunks := []blockChunk{}
	startPos := 0

	for i := 0; i < len(nums); i++ {
		length := nums[i]
		if i%2 == 0 {
			id := i / 2
			blockChunks = append(blockChunks, blockChunk{id, startPos, length})
		} else {
			spaceChunks = append(spaceChunks, blockChunk{id: -1, startPos: startPos, length: length})
		}
		startPos += length
	}

	part1(getCopy(blockChunks), getCopy(spaceChunks))
	part2(getCopy(blockChunks), getCopy(spaceChunks))
}
