// https://adventofcode.com/2024/day/11

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("Advent of Code 2024 - Day 11\n---")

	file, err := os.Open("input.txt")
	if err != nil {
		panic(err) // Fail if we can't open the file
	}
	defer file.Close()

	// Read the file "input.txt" into a string
	scanner := bufio.NewScanner(file)
	var lines []string
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	// read line 0 into a list of stones (they are stored as space-separated integers)
	var stonesPart1 []string
	var stonesPart2 []string
	for _, val := range strings.Split(lines[0], " ") {
		stonesPart1 = append(stonesPart1, val)
		stonesPart2 = append(stonesPart2, val)
	}

	// ======= PART 1 =======

	blinks := 25
	stoneMap := make(map[string]int)
	for _, val := range stonesPart1 {
		stoneMap[val]++
	}
	sum := processStones(stoneMap, blinks)
	fmt.Println("There are a total of", sum, "stones after", blinks, "blinks") // Correct answer: 193899 stones after 25 blinks

	// ======= PART 2 =======

	blinks = 75
	stoneMap = make(map[string]int)
	for _, val := range stonesPart2 {
		stoneMap[val]++
	}
	sum = processStones(stoneMap, blinks)
	fmt.Println("There are a total of", sum, "stones after", blinks, "blinks") // Correct answer: 229682160383225 stones after 75 blinks
}

func processStones(stones map[string]int, blinks int) int {
	for blink := 0; blink < blinks; blink++ {
		newStones := make(map[string]int)

		for stone, count := range stones {
			left, right := processStone(stone)
			newStones[left] += count
			if right != "" {
				newStones[right] += count
			}
		}

		stones = newStones
	}

	return aggregateCount(stones)
}

func aggregateCount(stones map[string]int) int {
	total := 0
	for _, count := range stones {
		total += count
	}
	return total
}

func processStone(stone string) (string, string) {
	if stone == "0" {
		return "1", ""
	} else if len(stone)%2 == 1 {
		stoneInt, err := strconv.Atoi(stone)
		if err != nil {
			panic(err)
		}
		// multiply the stone value by 2024
		stoneInt *= 2024
		return strconv.Itoa(stoneInt), ""
	}

	left := stone[:len(stone)/2]
	right := stone[len(stone)/2:]

	// remove leading zeros from right (wouldn't happen on left)
	right = strings.TrimLeft(right, "0")
	if right == "" {
		right = "0"
	}

	return left, right
}
