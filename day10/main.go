// https://adventofcode.com/2024/day/10

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

type Trailhead struct {
	x      int
	y      int
	score  int
	rating int
}

func main() {
	fmt.Println("Advent of Code 2024 - Day 10\n---")

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

	// ======= PART 1 =======

	// read into a 2d array
	m := make([][]string, len(lines))
	for i, line := range lines {
		m[i] = make([]string, len(line))
		for j, char := range line {
			m[i][j] = string(char)
		}
	}

	// find the trailheads (0s)
	var trailheads []Trailhead
	for i := 0; i < len(m); i++ {
		for j := 0; j < len(m[i]); j++ {
			if m[i][j] == "0" {
				trailheads = append(trailheads, Trailhead{x: i, y: j, score: 0, rating: 0})
			}
		}
	}

	// for each trailhead, walk all possible paths it can walk (up, down, left, right -- no diagonals)
	// without revisiting any location. It must be perfect incrementing order
	// (0 can only go to 1, 1 can only go to 2, etc.).
	// Add a point to the trailhead's score if it reaches the end of the trailhead.
	for i := 0; i < len(trailheads); i++ {
		visited := make([][]bool, len(m))
		for i := range visited {
			visited[i] = make([]bool, len(m[i]))
		}
		trailheads[i] = walkTrailheadForScore(m, trailheads[i], trailheads[i].x, trailheads[i].y, visited)
	}

	sumTrailheadScores := 0
	for i := 0; i < len(trailheads); i++ {
		sumTrailheadScores += trailheads[i].score
	}

	fmt.Println("Sum of trailhead scores: ", sumTrailheadScores) // Correct answer: 698

	// ======= PART 2 =======

	// Now handle trailhead ratings, which is pretty much the same as part 1
	// but you don't need to keep track of visited locations. You can just recursively
	// walk all possible paths.

	for i := 0; i < len(trailheads); i++ {
		trailheads[i] = walkTrailheadForRating(m, trailheads[i], trailheads[i].x, trailheads[i].y)
	}

	sumTrailheadRatings := 0
	for i := 0; i < len(trailheads); i++ {
		sumTrailheadRatings += trailheads[i].rating
	}

	fmt.Println("Sum of trailhead ratings: ", sumTrailheadRatings) // Correct answer: 1436
}

func walkTrailheadForScore(m [][]string, trailhead Trailhead, currentX int, currentY int, visited [][]bool) Trailhead {

	// if we're out of bounds, return
	if currentX < 0 || currentX >= len(m) || currentY < 0 || currentY >= len(m[currentX]) {
		return trailhead
	}

	// if we've already visited this spot, return
	if visited[currentX][currentY] {
		return trailhead
	}
	visited[currentX][currentY] = true

	// if we're at the end of the trailhead, increment the score
	if m[currentX][currentY] == "9" {
		trailhead.score++
		return trailhead
	}

	value := m[currentX][currentY]
	valueAsInt, err := strconv.Atoi(value)
	if err != nil {
		panic(err)
	}
	seeking := strconv.Itoa(valueAsInt + 1)

	// if we can validly move up, move up
	if currentX-1 >= 0 && m[currentX-1][currentY] == seeking {
		trailhead = walkTrailheadForScore(m, trailhead, currentX-1, currentY, visited)
	}

	// if we can validly move down, move down
	if currentX+1 < len(m) && m[currentX+1][currentY] == seeking {
		trailhead = walkTrailheadForScore(m, trailhead, currentX+1, currentY, visited)
	}

	// if we can validly move left, move left
	if currentY-1 >= 0 && m[currentX][currentY-1] == seeking {
		trailhead = walkTrailheadForScore(m, trailhead, currentX, currentY-1, visited)
	}

	// if we can validly move right, move right
	if currentY+1 < len(m[currentX]) && m[currentX][currentY+1] == seeking {
		trailhead = walkTrailheadForScore(m, trailhead, currentX, currentY+1, visited)
	}

	return trailhead
}

func walkTrailheadForRating(m [][]string, trailhead Trailhead, currentX int, currentY int) Trailhead {

	// if we're out of bounds, return
	if currentX < 0 || currentX >= len(m) || currentY < 0 || currentY >= len(m[currentX]) {
		return trailhead
	}

	// if we're at the end of the trailhead, increment the score
	if m[currentX][currentY] == "9" {
		trailhead.rating++
		return trailhead
	}

	value := m[currentX][currentY]
	valueAsInt, err := strconv.Atoi(value)
	if err != nil {
		panic(err)
	}

	seeking := strconv.Itoa(valueAsInt + 1)

	// if we can validly move up, move up
	if currentX-1 >= 0 && m[currentX-1][currentY] == seeking {
		trailhead = walkTrailheadForRating(m, trailhead, currentX-1, currentY)
	}

	// if we can validly move down, move down
	if currentX+1 < len(m) && m[currentX+1][currentY] == seeking {
		trailhead = walkTrailheadForRating(m, trailhead, currentX+1, currentY)
	}

	// if we can validly move left, move left
	if currentY-1 >= 0 && m[currentX][currentY-1] == seeking {
		trailhead = walkTrailheadForRating(m, trailhead, currentX, currentY-1)
	}

	// if we can validly move right, move right
	if currentY+1 < len(m[currentX]) && m[currentX][currentY+1] == seeking {
		trailhead = walkTrailheadForRating(m, trailhead, currentX, currentY+1)
	}

	return trailhead
}
