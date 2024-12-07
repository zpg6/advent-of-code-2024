// https://adventofcode.com/2024/day/6

package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	fmt.Println("Advent of Code 2024 - Day 6\n---")

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

	// Create a 2D matrix from the input
	var m [][]string
	for _, line := range lines {
		var row []string
		items := strings.Split(line, "")
		row = append(row, items...)
		m = append(m, row)
	}

	// ======= PART 1 =======

	rows := len(m)
	cols := len(m[0])

	// Find the current position of the player (^ character)
	var playerX, playerY int
	for i, row := range m {
		for j, char := range row {
			if char == "^" {
				playerX = j
				playerY = i
			}
		}
	}

	// Move the player in the direction of the arrow (^=>N, v=>S, <=W, =>E)
	// Mark each visited cell with an X. Stop when the player reaches the leaves the grid
	var direction string = "^"
	for {
		m[playerY][playerX] = "X"

		switch direction {
		case "^":
			playerY--
		case "v":
			playerY++
		case "<":
			playerX--
		case ">":
			playerX++
		}

		if playerY < 0 || playerY >= rows || playerX < 0 || playerX >= cols {
			break
		}

		// Check if the player is now at a wall (#)
		if m[playerY][playerX] == "#" {
			// undo the move and turn clockwise
			switch direction {
			case "^":
				playerY++
				direction = ">"
			case "v":
				playerY--
				direction = "<"
			case "<":
				playerX++
				direction = "^"
			case ">":
				playerX--
				direction = "v"
			}
		}

	}

	// Count the unique cells visited by the player
	steps := 0
	for _, row := range m {
		for _, cell := range row {
			if cell == "X" {
				steps++
			}
		}
	}
	fmt.Printf("Part 1: The player visited %d cells\n", steps) // Correct answer is: 4778

	// ======= PART 2 =======

	// Now find which cells putting a "O" (acts like a wall "#") will
	// cause the player to loop forever.
	var loop_cells int

	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			// Skip cells that are already walls
			if m[i][j] == "#" || (i == playerY && j == playerX) {
				continue
			}

			// Create a copy of the grid
			var mCopy [][]string
			for _, line := range lines {
				var row []string
				items := strings.Split(line, "")
				row = append(row, items...)
				mCopy = append(mCopy, row)
			}

			// Place an O in the cell
			mCopy[i][j] = "O"

			// Find the current position of the player (^ character)
			var playerX, playerY int
			for i, row := range mCopy {
				for j, char := range row {
					if char == "^" {
						playerX = j
						playerY = i
					}
				}
			}

			var direction string = "^"
			var steps int
			for {
				mCopy[playerY][playerX] = "X"
				steps++

				switch direction {
				case "^":
					playerY--
				case "v":
					playerY++
				case "<":
					playerX--
				case ">":
					playerX++
				}

				// Check if the player has left the grid
				if playerY < 0 || playerY >= rows || playerX < 0 || playerX >= cols {
					break
				}

				// Check if the player is now at a wall (#)
				if mCopy[playerY][playerX] == "#" || mCopy[playerY][playerX] == "O" {
					// undo the move and turn clockwise
					switch direction {
					case "^":
						playerY++
						direction = ">"
					case "v":
						playerY--
						direction = "<"
					case "<":
						playerX++
						direction = "^"
					case ">":
						playerX--
						direction = "v"
					}
				}

				// Check if the player has taken more than 10K steps
				// (this is our greedy solution to detect loops)
				if steps > 10_000 {
					loop_cells++
					break
				}
			}
		}
	}

	fmt.Printf("Part 2: There are %d cells that will cause the player to loop forever\n", loop_cells) // Correct answer is: 1618
}
