// https://adventofcode.com/2024/day/4

package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	fmt.Println("Advent of Code 2024 - Day 4\n---")

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

	rows := len(lines)
	cols := len(lines[0])

	// Create a 2D matrix from the input
	var m [][]string
	for i := 0; i < rows; i++ {
		var row []string
		for j := 0; j < cols; j++ {
			row = append(row, string(lines[i][j]))
		}
		m = append(m, row)
	}

	// ======= PART 1 =======

	// Word search for XMAS is carried out by checking the 8 directions from every X
	var found int
	for i, row := range m {
		for j, char := range row {
			if char == "X" {
				// Check the 8 directions
				// N
				if i-1 >= 0 && m[i-1][j] == "M" {
					if i-2 >= 0 && m[i-2][j] == "A" {
						if i-3 >= 0 && m[i-3][j] == "S" {
							found++
						}
					}
				}
				// NE
				if i-1 >= 0 && j+1 < cols && m[i-1][j+1] == "M" {
					if i-2 >= 0 && j+2 < cols && m[i-2][j+2] == "A" {
						if i-3 >= 0 && j+3 < cols && m[i-3][j+3] == "S" {
							found++
						}
					}
				}
				// E
				if j+1 < cols && m[i][j+1] == "M" {
					if j+2 < cols && m[i][j+2] == "A" {
						if j+3 < cols && m[i][j+3] == "S" {
							found++
						}
					}
				}
				// SE
				if i+1 < rows && j+1 < cols && m[i+1][j+1] == "M" {
					if i+2 < rows && j+2 < cols && m[i+2][j+2] == "A" {
						if i+3 < rows && j+3 < cols && m[i+3][j+3] == "S" {
							found++
						}
					}
				}
				// S
				if i+1 < rows && m[i+1][j] == "M" {
					if i+2 < rows && m[i+2][j] == "A" {
						if i+3 < rows && m[i+3][j] == "S" {
							found++
						}
					}
				}
				// SW
				if i+1 < rows && j-1 >= 0 && m[i+1][j-1] == "M" {
					if i+2 < rows && j-2 >= 0 && m[i+2][j-2] == "A" {
						if i+3 < rows && j-3 >= 0 && m[i+3][j-3] == "S" {
							found++
						}
					}
				}
				// W
				if j-1 >= 0 && m[i][j-1] == "M" {
					if j-2 >= 0 && m[i][j-2] == "A" {
						if j-3 >= 0 && m[i][j-3] == "S" {
							found++
						}
					}
				}
				// NW
				if i-1 >= 0 && j-1 >= 0 && m[i-1][j-1] == "M" {
					if i-2 >= 0 && j-2 >= 0 && m[i-2][j-2] == "A" {
						if i-3 >= 0 && j-3 >= 0 && m[i-3][j-3] == "S" {
							found++
						}
					}
				}
			}
		}
	}

	fmt.Println("Word search XMAS count: ", found) // Correct answer is 2401

	// ======= PART 2 =======

	// Word search for MAS x's.
	found2 := 0
	for i, row := range m {
		for j, char := range row {

			if char == "A" {

				// Option 1:
				// NE=M, NW=S, SE=M, SW=S
				if i-1 >= 0 && j+1 < cols && m[i-1][j+1] == "M" {
					if i-1 >= 0 && j-1 >= 0 && m[i-1][j-1] == "S" {
						if i+1 < rows && j+1 < cols && m[i+1][j+1] == "M" {
							if i+1 < rows && j-1 >= 0 && m[i+1][j-1] == "S" {
								found2++
							}
						}
					}
				}

				// Option 2:
				// NE=S, NW=M, SE=S, SW=M
				if i-1 >= 0 && j+1 < cols && m[i-1][j+1] == "S" {
					if i-1 >= 0 && j-1 >= 0 && m[i-1][j-1] == "M" {
						if i+1 < rows && j+1 < cols && m[i+1][j+1] == "S" {
							if i+1 < rows && j-1 >= 0 && m[i+1][j-1] == "M" {
								found2++
							}
						}
					}
				}

				// Option 3:
				// NE=M, NW=M, SE=S, SW=S
				if i-1 >= 0 && j+1 < cols && m[i-1][j+1] == "M" {
					if i-1 >= 0 && j-1 >= 0 && m[i-1][j-1] == "M" {
						if i+1 < rows && j+1 < cols && m[i+1][j+1] == "S" {
							if i+1 < rows && j-1 >= 0 && m[i+1][j-1] == "S" {
								found2++
							}
						}
					}
				}

				// Option 4:
				// NE=S, NW=S, SE=M, SW=M
				if i-1 >= 0 && j+1 < cols && m[i-1][j+1] == "S" {
					if i-1 >= 0 && j-1 >= 0 && m[i-1][j-1] == "S" {
						if i+1 < rows && j+1 < cols && m[i+1][j+1] == "M" {
							if i+1 < rows && j-1 >= 0 && m[i+1][j-1] == "M" {
								found2++
							}
						}
					}
				}
			}
		}
	}

	fmt.Println("Word search X-'MAS' x's found: ", found2) // Correct answer is 1822
}
