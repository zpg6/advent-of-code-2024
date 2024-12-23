// https://adventofcode.com/2024/day/1

package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func main() {
	fmt.Println("Advent of Code 2024 - Day 1\n---")

	// ======= PART 1 =======

	file, err := os.Open("list.txt")
	if err != nil {
		panic(err) // Fail if we can't open the file
	}
	defer file.Close()

	// Create two lists to store the left and right elements
	//
	// TODO: Use a data structure that will store the elements in order
	// so that we don't have to sort them later.
	var listLeft []int
	var listRight []int

	// Read the file "list.txt" line by line
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		var first, second int
		_, err := fmt.Sscanf(line, "%d   %d", &first, &second)
		if err != nil {
			panic(err) // Fail if we can't parse the line
		}
		listLeft = append(listLeft, first)
		listRight = append(listRight, second)
	}

	if len(listLeft) != len(listRight) {
		panic("The lists are not of the same length")
	}

	// Sort the lists
	sort.Ints(listLeft)
	sort.Ints(listRight)

	// Sum the diff between right and left element by element
	var sum int
	for i := 0; i < len(listLeft); i++ {
		diff := listRight[i] - listLeft[i]

		// add the absolute value of the diff to the sum
		if diff < 0 {
			sum += -diff
		} else {
			sum += diff
		}
	}

	fmt.Println("Sum of the differences is:", sum) // Correct answer is 2057374

	// ======= PART 2 =======

	// Now calculate similarity score
	//
	// TODO: Use a data structure to store the counts of the elements
	// so that we don't have to iterate over the listRight for each element.
	var score int
	for i := 0; i < len(listLeft); i++ {
		left_element := listLeft[i]
		count := 0
		for j := 0; j < len(listRight); j++ {
			if listRight[j] == left_element {
				count++
			}
		}
		// The req states that
		// similarity += left_element * count
		score += (left_element * count)
	}
	fmt.Println("Similarity score is:", score) // Correct answer is 23177084
}
