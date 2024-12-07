// https://adventofcode.com/2024/day/7

package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strings"
)

type Line struct {
	key    int
	values []int
}

func main() {
	fmt.Println("Advent of Code 2024 - Day 7\n---")

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

	// Each line looks like this:
	// int: int int int ...
	var m []Line

	// Parse the lines
	for _, line := range lines {
		var key int
		var values []int

		_, err := fmt.Sscanf(line, "%d:", &key)
		if err != nil {
			panic(err) // Fail if we can't parse the key
		}
		values = parseValues(line[strings.Index(line, ":")+2:]) // skip ": "

		m = append(m, Line{key, values})
	}

	// ======= PART 1 =======

	var operations = []string{"+", "*"}

	// For each key in the map, see if any permutation of operations results in the target value
	sumOfCalculated := 0
	for _, line := range m {
		perms := getPermutations(operations, len(line.values)-1)
		for _, p := range perms {
			var result int
			for i, v := range line.values {
				if i == 0 {
					result = v
					continue
				}
				if p[i-1] == "+" {
					result += v
				} else {
					result *= v
				}
			}
			if result == line.key {
				sumOfCalculated += result
				break
			}
		}
	}

	fmt.Println("Sum of calculated values (+,*) is:", sumOfCalculated) // Correct answer is 2654749936343

	// ======= PART 2 =======

	var operationsPart2 = []string{"+", "*", "||"}

	// For each key in the map, see if any permutation of operations results in the target value
	sumOfCalculatedPart2 := 0
	for _, line := range m {
		perms := getPermutations(operationsPart2, len(line.values)-1)
		for _, p := range perms {
			var result int
			for i, v := range line.values {
				if i == 0 {
					result = v
					continue
				}
				if p[i-1] == "+" {
					result += v
				} else if p[i-1] == "*" {
					result *= v
				} else if p[i-1] == "||" {
					// concatenate the integers
					digits := int(math.Log10(float64(v))) + 1
					result *= int(math.Pow10(digits))
					result += v
				} else {
					panic("Invalid operation")
				}
			}
			if result == line.key {
				sumOfCalculatedPart2 += result
				break
			}
		}
	}

	fmt.Println("Sum of calculated values (+,*,||) is:", sumOfCalculatedPart2) // Correct answer is 124060392153684
}

func parseValues(line string) []int {
	var values []int
	for _, v := range strings.Split(line, " ") {
		var value int
		_, err := fmt.Sscanf(v, "%d", &value)
		if err != nil {
			panic(err) // Fail if we can't parse the value
		}
		values = append(values, value)
	}
	return values
}

func getPermutations(operations []string, n int) [][]string {
	if n <= 0 {
		return [][]string{{}}
	}
	if n == 1 {
		var perms [][]string
		for _, op := range operations {
			perms = append(perms, []string{op})
		}
		return perms
	}

	perms := getPermutations(operations, n-1)
	var newPerms [][]string
	for _, perm := range perms {
		for _, op := range operations {
			newPerm := make([]string, len(perm))
			copy(newPerm, perm)
			newPerm = append(newPerm, op)
			newPerms = append(newPerms, newPerm)
		}
	}
	return newPerms
}
