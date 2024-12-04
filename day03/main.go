// https://adventofcode.com/2024/day/3

package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
)

func main() {
	fmt.Println("Advent of Code 2024 - Day 3\n---")

	// ======= PART 1 =======

	file, err := os.Open("input.txt")
	if err != nil {
		panic(err) // Fail if we can't open the file
	}
	defer file.Close()

	// Read the file "input.txt" into a string
	scanner := bufio.NewScanner(file)
	var file_string string
	for scanner.Scan() {
		file_string += scanner.Text()
	}

	// "mul(%d,%d)" is the pattern we're looking for
	regex := regexp.MustCompile(`mul\((\d+),(\d+)\)`)
	matches := regex.FindAllStringSubmatch(file_string, -1)

	var sum int
	for _, match := range matches {
		var a, b int
		fmt.Sscanf(match[1], "%d", &a)
		fmt.Sscanf(match[2], "%d", &b)
		sum += a * b
	}

	fmt.Println("The sum of all the products is", sum) // Correct answer: 155955228

	// ======= PART 2 =======

	var sumEnabled int = 0

	regexDo := regexp.MustCompile(`do\(()\)`)      // "do()"
	regexDont := regexp.MustCompile(`don't\(()\)`) // "don't()"

	indicesDo := regexDo.FindAllStringIndex(file_string, -1)
	indicesDont := regexDont.FindAllStringIndex(file_string, -1)

	// determine which of the original matches are enabled
	indices := regex.FindAllStringIndex(file_string, -1)
	var i int
	for _, m := range indices {
		// Check the latest do block before this match
		var doIndex int = -1
		for _, d := range indicesDo {
			if d[0] < m[0] {
				doIndex = d[0]
			}
		}

		// Check the latest don't block before this match
		var dontIndex int = -1
		for _, d := range indicesDont {
			if d[0] < m[0] {
				dontIndex = d[0]
			}
		}

		// Default to enabled if no do or don't blocks are found
		// Otherwise, "do()" enables and "don't()" disables
		if doIndex == -1 && dontIndex == -1 || doIndex > dontIndex {
			var a, b int
			fmt.Sscanf(matches[i][1], "%d", &a)
			fmt.Sscanf(matches[i][2], "%d", &b)
			sumEnabled += a * b
		}

		i++
	}

	fmt.Println("The sum of all the products in enabled blocks is", sumEnabled) // Correct answer: 100189366
}
