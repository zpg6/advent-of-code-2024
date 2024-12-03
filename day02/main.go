// https://adventofcode.com/2024/day/2

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// A "safe" array is all-increasing or all-decreasing,
// with a delta between 1 and 3 (inclusive) for all elements
func isSafe(arr []int) bool {
	if len(arr) < 2 || arr[0] == arr[1] {
		return false
	}
	var checkingIncreasing bool = arr[0] < arr[1]
	var checkingDecreasing bool = arr[0] > arr[1]

	for i := 0; i < len(arr)-1; i++ {
		if checkingIncreasing {
			delta := arr[i+1] - arr[i]
			if delta < 1 || delta > 3 {
				return false
			}
		} else if checkingDecreasing {
			delta := arr[i] - arr[i+1]
			if delta < 1 || delta > 3 {
				return false
			}
		}
	}

	return true
}

func main() {
	fmt.Println("Advent of Code 2024 - Day 2\n---")

	file, err := os.Open("data.txt")
	if err != nil {
		panic(err) // Fail if we can't open the file
	}
	defer file.Close()

	var countSafeRecords int
	var dampenedCountSafeRecords int

	// Read the file "list.txt" line by line
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		var lineItems []int

		// each line is n space-separated integers
		split := strings.Split(line, " ")
		for _, item := range split {
			num, err := strconv.Atoi(item)
			if err != nil {
				panic(err) // Fail if we can't parse the line
			}
			lineItems = append(lineItems, num)
		}

		// Part 1: Count the number of "safe" records
		if isSafe(lineItems) {
			countSafeRecords++
			dampenedCountSafeRecords++
		} else {
			// Part 2: "Dampener" sees if it passes as "safe" with
			// any one element removed
			for i := 0; i < len(lineItems); i++ {
				var dampenedLineItems []int

				// only add non-i elements
				for j := 0; j < len(lineItems); j++ {
					if j != i {
						dampenedLineItems = append(dampenedLineItems, lineItems[j])
					}
				}

				// check if safe with arr[i] removed
				if isSafe(dampenedLineItems) {
					dampenedCountSafeRecords++
					break
				}
			}
		}
	}

	fmt.Printf("Number of safe records: %d\n", countSafeRecords)                  // Correct answer: 598
	fmt.Printf("Number of dampened safe records: %d\n", dampenedCountSafeRecords) // Correct answer: 634
}
