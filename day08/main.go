// https://adventofcode.com/2024/day/8

package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	fmt.Println("Advent of Code 2024 - Day 8\n---")

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

	// Read into a 2d character array
	var m [][]string
	for _, line := range lines {
		var row []string
		for _, c := range line {
			row = append(row, string(c))
		}
		m = append(m, row)
	}

	// Find the alphanumeric characters
	var alphanum []string
	for _, c := range "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789" {
		alphanum = append(alphanum, string(c))
	}

	// ======= PART 1 =======

	// Create a 2d array to store the antinodes
	var antinodeMap [][]string
	for i := 0; i < len(m); i++ {
		var row []string
		for j := 0; j < len(m[i]); j++ {
			row = append(row, "")
		}
		antinodeMap = append(antinodeMap, row)
	}

	// Find the pairs of same alphanumeric characters
	var pairsCount int = 0
	for _, an := range alphanum {
		var locationsX []int
		var locationsY []int
		for i, row := range m {
			for j, col := range row {
				if col == an {
					locationsX = append(locationsX, i)
					locationsY = append(locationsY, j)
				}
			}
		}

		// Form the unique pairs of locations
		var pairs [][]int
		for i := 0; i < len(locationsX); i++ {
			for j := i + 1; j < len(locationsX); j++ {
				if locationsX[i] != locationsX[j] && locationsY[i] != locationsY[j] {
					pairs = append(pairs, []int{locationsX[i], locationsY[i], locationsX[j], locationsY[j]})
				}
			}
		}
		pairsCount += len(pairs)

		// For each pair, mark the antinodes
		for _, pair := range pairs {
			// Antinode is like the opposite of a midpoint.
			// antinode1 = (x1-(x2-x1), y1-(y2-y1))
			// antinode2 = (x2-(x1-x2), y2-(y1-y2))

			antinode1X := pair[0] - (pair[2] - pair[0])
			antinode1Y := pair[1] - (pair[3] - pair[1])
			antinode2X := pair[2] - (pair[0] - pair[2])
			antinode2Y := pair[3] - (pair[1] - pair[3])

			if antinode1X >= 0 && antinode1X < len(m) && antinode1Y >= 0 && antinode1Y < len(m[0]) {
				antinodeMap[antinode1X][antinode1Y] = "#"
			}
			if antinode2X >= 0 && antinode2X < len(m) && antinode2Y >= 0 && antinode2Y < len(m[0]) {
				antinodeMap[antinode2X][antinode2Y] = "#"
			}

		}
	}

	// Count the number of antinodes
	var antinodeCount int = 0
	for _, row := range antinodeMap {
		for _, col := range row {
			if col == "#" {
				antinodeCount++
			}
		}
	}

	fmt.Println("Count of Antinodes:", antinodeCount) // Correct answer: 311

	// ======= PART 2 =======

	// Find the antinodes based on the updated definition
	var antinodeMapPart2 [][]string
	for i := 0; i < len(m); i++ {
		var row []string
		for j := 0; j < len(m[i]); j++ {
			row = append(row, "")
		}
		antinodeMapPart2 = append(antinodeMapPart2, row)
	}

	// Find the pairs of same alphanumeric characters
	var pairsCountPart2 int = 0
	for _, an := range alphanum {
		var locationsX []int
		var locationsY []int
		for i, row := range m {
			for j, col := range row {
				if col == an {
					locationsX = append(locationsX, i)
					locationsY = append(locationsY, j)
				}
			}
		}

		// Form the unique pairs of locations
		var pairs [][]int
		for i := 0; i < len(locationsX); i++ {
			for j := i + 1; j < len(locationsX); j++ {
				if locationsX[i] != locationsX[j] && locationsY[i] != locationsY[j] {
					pairs = append(pairs, []int{locationsX[i], locationsY[i], locationsX[j], locationsY[j]})
				}
			}
		}
		pairsCountPart2 += len(pairs)

		for _, pair := range pairs {
			// Antinodes now keep walking in the same direction until they go out of bounds
			antinode1X := pair[0]
			antinode1Y := pair[1]
			antinode2X := pair[2]
			antinode2Y := pair[3]

			// Any in-line nodes are also antinodes now
			antinodeMapPart2[antinode1X][antinode1Y] = "#"
			antinodeMapPart2[antinode2X][antinode2Y] = "#"

			for {
				antinode1X -= (pair[2] - pair[0])
				antinode1Y -= (pair[3] - pair[1])
				if antinode1X >= 0 && antinode1X < len(m) && antinode1Y >= 0 && antinode1Y < len(m[0]) {
					antinodeMapPart2[antinode1X][antinode1Y] = "#"
				} else {
					break
				}
			}

			for {
				antinode2X -= (pair[0] - pair[2])
				antinode2Y -= (pair[1] - pair[3])
				if antinode2X >= 0 && antinode2X < len(m) && antinode2Y >= 0 && antinode2Y < len(m[0]) {
					antinodeMapPart2[antinode2X][antinode2Y] = "#"
				} else {
					break
				}
			}
		}
	}

	// Count the number of antinodes
	var antinodeCountPart2 int = 0
	for _, row := range antinodeMapPart2 {
		for _, col := range row {
			if col == "#" {
				antinodeCountPart2++
			}
		}
	}

	fmt.Println("Count of Antinodes Part 2:", antinodeCountPart2) // Correct answer: 1115
}
