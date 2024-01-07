package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

type Mapping struct {
	dstStart int
	length   int
}

func readLinesAndCreateMap(scanner *bufio.Scanner, targetMap map[int]Mapping) {
	for {
		scanner.Scan()
		line := scanner.Text()
		if line == "" {
			break
		}
		parts := strings.Fields(line)
		dstStart, _ := strconv.Atoi(parts[0])
		srcStart, _ := strconv.Atoi(parts[1])
		length, _ := strconv.Atoi(parts[2])
		targetMap[srcStart] = Mapping{dstStart: dstStart, length: length}
	}
}

func findDestination(src int, targetMap map[int]Mapping) int {
	for srcStart, m := range targetMap {
		if src >= srcStart && src < srcStart+m.length {
			offset := src - srcStart
			return m.dstStart + offset
		}
	}
	return src
}

func part1(seeds []string, seedToSoil, soilToFertilizer, fertilizerToWater, waterToLight, lightToTemperature, temperatureToHumidity, humidityToLocation map[int]Mapping) {
	minLocation := math.MaxInt

	// follow the maps
	for _, seed := range seeds {
		s, _ := strconv.Atoi(seed)
		next := findDestination(s, seedToSoil)
		next = findDestination(next, soilToFertilizer)
		next = findDestination(next, fertilizerToWater)
		next = findDestination(next, waterToLight)
		next = findDestination(next, lightToTemperature)
		next = findDestination(next, temperatureToHumidity)
		next = findDestination(next, humidityToLocation)

		if next < minLocation {
			minLocation = next
		}
	}

	fmt.Println(minLocation)
}

// Takes 1h+
func part2(seeds []string, seedToSoil, soilToFertilizer, fertilizerToWater, waterToLight, lightToTemperature, temperatureToHumidity, humidityToLocation map[int]Mapping) {
	minLocation := math.MaxInt

	// follow the maps
	for i := 0; i < len(seeds)-1; i += 2 {
		start, _ := strconv.Atoi(seeds[i])
		lenght, _ := strconv.Atoi(seeds[i+1])
		end := start + lenght - 1

		for i := start; i <= end; i++ {
			next := findDestination(i, seedToSoil)
			next = findDestination(next, soilToFertilizer)
			next = findDestination(next, fertilizerToWater)
			next = findDestination(next, waterToLight)
			next = findDestination(next, lightToTemperature)
			next = findDestination(next, temperatureToHumidity)
			next = findDestination(next, humidityToLocation)

			if next < minLocation {
				minLocation = next
			}
		}
	}

	fmt.Println(minLocation)
}

func main() {
	// file, _ := os.Open("sample.txt")
	file, _ := os.Open("input.txt")
	defer file.Close()

	scanner := bufio.NewScanner(file)

	// seeds
	scanner.Scan()
	line := scanner.Text()
	seedsString := strings.Split(line, ":")[1]
	seeds := strings.Fields(seedsString)
	scanner.Scan() // blank line

	// seed to soil
	scanner.Scan() // header line
	seedToSoil := make(map[int]Mapping)
	readLinesAndCreateMap(scanner, seedToSoil)

	// soil to fertilizer
	scanner.Scan()
	soilToFertilizer := make(map[int]Mapping)
	readLinesAndCreateMap(scanner, soilToFertilizer)

	// fertilizer to water
	scanner.Scan()
	fertilizerToWater := make(map[int]Mapping)
	readLinesAndCreateMap(scanner, fertilizerToWater)

	// water to light
	scanner.Scan()
	waterToLight := make(map[int]Mapping)
	readLinesAndCreateMap(scanner, waterToLight)

	// light to temperature
	scanner.Scan()
	lightToTemperature := make(map[int]Mapping)
	readLinesAndCreateMap(scanner, lightToTemperature)

	// temperature to humidity
	scanner.Scan()
	temperatureToHumidity := make(map[int]Mapping)
	readLinesAndCreateMap(scanner, temperatureToHumidity)

	// humidity to location
	scanner.Scan()
	humidityToLocation := make(map[int]Mapping)
	readLinesAndCreateMap(scanner, humidityToLocation)

	// part1(seeds, seedToSoil, soilToFertilizer, fertilizerToWater, waterToLight, lightToTemperature, temperatureToHumidity, humidityToLocation)
	part2(seeds, seedToSoil, soilToFertilizer, fertilizerToWater, waterToLight, lightToTemperature, temperatureToHumidity, humidityToLocation)
}
