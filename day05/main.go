// https://adventofcode.com/2024/day/5

package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	fmt.Println("Advent of Code 2024 - Day 5\n---")

	file, err := os.Open("input.txt")
	if err != nil {
		panic(err) // Fail if we can't open the file
	}
	defer file.Close()

	// Read the file "input.txt" into a string
	scanner := bufio.NewScanner(file)
	var first_section_lines []string
	var past_separator bool
	var second_section_lines []string
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			past_separator = true
			continue
		}
		if !past_separator {
			first_section_lines = append(first_section_lines, line)
		} else {
			second_section_lines = append(second_section_lines, line)
		}
	}

	// ======= PART 1 =======

	// Each of the first lines looks like 'x|y' where x and y are integers.
	// It means X must come before Y in the ordering. We'll track with a map.

	// Create a map to store the order of the elements
	order := make(map[int][]int)
	for _, line := range first_section_lines {
		var x, y int
		_, err := fmt.Sscanf(line, "%d|%d", &x, &y)
		if err != nil {
			panic(err) // Fail if we can't parse the line
		}
		order[x] = append(order[x], y)
	}

	var incorrectly_ordered_lines []string
	var sum int = 0

	// Each of the second lines looks like "a,b,c,d,e" or "f,g,h"
	// and we need to see if they follow the ordering rules from the first section
	for _, line := range second_section_lines {
		// Split the line by commas
		elements := strings.Split(line, ",")
		element_ints := make([]int, len(elements))
		for i, element := range elements {
			fmt.Sscanf(element, "%d", &element_ints[i])
		}

		// Check if the elements are in the correct order according to the first section
		// a,b,c,d,e should check that e doesn't have any ordering rules that it come before
		// a,b,c,d. If order[e] contains any of a,b,c,d, then the ordering is incorrect.

		// We'll iterate through the elements in reverse order
		// and check if any of the elements that come after it are in the ordering rules
		// for the current element.
		var correct bool = true
		for i := len(element_ints) - 1; i >= 0; i-- {

			// Check if the element has any ordering rules
			if _, ok := order[element_ints[i]]; ok {
				// any elements < index i should not be in the ordering rules for element_ints[i]
				for j := i - 1; j >= 0; j-- {
					if Contains(order[element_ints[i]], element_ints[j]) {
						correct = false
						break
					}
				}
			}
			if !correct {
				break
			}

		}
		if correct {
			// Puzzle wants the sum of the middle elements for the correctly ordered lines
			midpoint := len(element_ints) / 2
			sum += element_ints[midpoint]
		} else {
			incorrectly_ordered_lines = append(incorrectly_ordered_lines, line)
		}
	}

	fmt.Println("Sum of the middle elements for the correctly ordered lines:", sum) // Correct answer is 6384

	// ======= PART 2 =======

	// For the incorrectly ordered lines, we need to order them correctly
	// and then sum the middle elements.
	var sum_incorrect int = 0

	// We will start from the last element and then move backwards, so when a rule
	// is broken we will place the element at the end of the list and restart the process.
	// We restart in case the element's new placement breaks another rule.
	for _, line := range incorrectly_ordered_lines {
		// Split the line by commas
		elements := strings.Split(line, ",")
		element_ints := make([]int, len(elements))
		for i, element := range elements {
			fmt.Sscanf(element, "%d", &element_ints[i])
		}

		complete := false
		for !complete {
			var correct bool = true
			var offending_element int
			for i := len(element_ints) - 1; i >= 0; i-- {

				// Check if the element has any ordering rules
				if _, ok := order[element_ints[i]]; ok {

					// any elements < index i should not be in the ordering rules for element_ints[i]
					for j := i - 1; j >= 0; j-- {
						if Contains(order[element_ints[i]], element_ints[j]) {
							offending_element = element_ints[j]

							// Move (not copy) the offending element to the end of the list
							element_ints = append(element_ints[:j], element_ints[j+1:]...)
							element_ints = append(element_ints, offending_element)

							correct = false
							break
						}
					}
				}
			}
			if correct {
				// Puzzle wants the sum of the middle elements for the correctly ordered lines
				midpoint := len(element_ints) / 2
				sum_incorrect += element_ints[midpoint]
				complete = true
			}
		}
	}

	fmt.Println("Sum of the middle elements for the incorrectly ordered lines:", sum_incorrect) // Correct answer is 5353
}

func Contains(slice []int, element int) bool {
	for _, e := range slice {
		if e == element {
			return true
		}
	}
	return false
}
