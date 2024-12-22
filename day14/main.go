// https://adventofcode.com/2024/day/14

package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Robot struct {
	x, y, dx, dy int
}

func main() {
	fmt.Println("Advent of Code 2024 - Day 14\n---")

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

	// Read the games
	var robots []Robot
	for i := 0; i < len(lines); i++ {
		var robot Robot
		// (may be negative)
		patchedLine := strings.ReplaceAll(lines[i], "p=", "p= ")
		patchedLine = strings.ReplaceAll(patchedLine, "v=", "v= ")
		patchedLine = strings.ReplaceAll(patchedLine, ",", " , ")
		fmt.Sscanf(patchedLine, "p= %d , %d v= %d , %d", &robot.x, &robot.y, &robot.dx, &robot.dy)
		robots = append(robots, robot)
	}

	BOARD_WIDTH, BOARD_HEIGHT := 101, 103

	// ======= PART 1 =======

	// For 100 cycles, move the robots
	// if they run off the edge, wrap around
	for cycle := 0; cycle < 100; cycle++ {

		// Move the robots
		for i := range robots {
			robots[i].x += robots[i].dx
			robots[i].y += robots[i].dy

			if robots[i].x < 0 {
				robots[i].x += BOARD_WIDTH
			}
			if robots[i].y < 0 {
				robots[i].y += BOARD_HEIGHT
			}
			if robots[i].x >= BOARD_WIDTH {
				robots[i].x -= BOARD_WIDTH
			}
			if robots[i].y >= BOARD_HEIGHT {
				robots[i].y -= BOARD_HEIGHT
			}
		}
	}

	q1Sum, q2Sum, q3Sum, q4Sum := quadrantSums(robots, BOARD_WIDTH, BOARD_HEIGHT)

	if q1Sum == 0 {
		fmt.Println("No robots in Q1")
		q1Sum = 1
	}
	if q2Sum == 0 {
		fmt.Println("No robots in Q2")
		q2Sum = 1
	}
	if q3Sum == 0 {
		fmt.Println("No robots in Q3")
		q3Sum = 1
	}
	if q4Sum == 0 {
		fmt.Println("No robots in Q4")
		q4Sum = 1
	}
	safetyFactor := q1Sum * q2Sum * q3Sum * q4Sum

	fmt.Println("Safety Factor:", safetyFactor) // Correct answer: 208437768

	// ======= PART 2 =======

	// Keep running until one quadrant has more than half the robots.
	// This problem is purposefully a mystery, so this was just my guess.
	// ... it just tells you to find a Christmas tree ....
	// I assumed that the robots would eventually form a pattern that would
	// repeat, and that the majority of robots would be in one quadrant.

	// Reset the robots
	robots = []Robot{}
	for i := 0; i < len(lines); i++ {
		var robot Robot
		// (may be negative)
		patchedLine := strings.ReplaceAll(lines[i], "p=", "p= ")
		patchedLine = strings.ReplaceAll(patchedLine, "v=", "v= ")
		patchedLine = strings.ReplaceAll(patchedLine, ",", " , ")
		fmt.Sscanf(patchedLine, "p= %d , %d v= %d , %d", &robot.x, &robot.y, &robot.dx, &robot.dy)
		robots = append(robots, robot)
	}

	cycle := 0
	MAJORITY_FACTOR := 2
	for {
		cycle++
		// Move the robots
		for i := range robots {
			robots[i].x += robots[i].dx
			robots[i].y += robots[i].dy

			if robots[i].x < 0 {
				robots[i].x += BOARD_WIDTH
			}
			if robots[i].y < 0 {
				robots[i].y += BOARD_HEIGHT
			}
			if robots[i].x >= BOARD_WIDTH {
				robots[i].x -= BOARD_WIDTH
			}
			if robots[i].y >= BOARD_HEIGHT {
				robots[i].y -= BOARD_HEIGHT
			}
		}

		q1Sum, q2Sum, q3Sum, q4Sum = quadrantSums(robots, BOARD_WIDTH, BOARD_HEIGHT)

		if q1Sum > len(robots)/MAJORITY_FACTOR || q2Sum > len(robots)/MAJORITY_FACTOR || q3Sum > len(robots)/MAJORITY_FACTOR || q4Sum > len(robots)/MAJORITY_FACTOR {

			// Print the robots
			for x := 0; x < BOARD_WIDTH; x++ {
				for y := 0; y < BOARD_HEIGHT; y++ {

					if x == BOARD_WIDTH/2 || y == BOARD_HEIGHT/2 {
						fmt.Print(" ")
						continue
					}

					count := 0
					for _, robot := range robots {
						if robot.x == x && robot.y == y {
							count++
						}
					}
					if count > 0 {
						fmt.Print(count)
					} else {
						fmt.Print(" ")
					}
				}
				fmt.Println()
			}

			// Wait for a key press from me while I look at the output
			fmt.Println("CYCLE:", cycle, " -- Press Enter to continue") // Correct answer: 7492
			bufio.NewReader(os.Stdin).ReadBytes('\n')
		}
	}
}

func quadrantSums(robots []Robot, BOARD_WIDTH int, BOARD_HEIGHT int) (int, int, int, int) {
	q1Sum := 0
	q2Sum := 0
	q3Sum := 0
	q4Sum := 0

	for x := 0; x < BOARD_WIDTH; x++ {
		for y := 0; y < BOARD_HEIGHT; y++ {

			if x < BOARD_WIDTH/2 && y < BOARD_HEIGHT/2 {
				for _, robot := range robots {
					if robot.x == x && robot.y == y {
						q1Sum++
					}
				}
			}
			if x > BOARD_WIDTH/2 && y < BOARD_HEIGHT/2 {
				for _, robot := range robots {
					if robot.x == x && robot.y == y {
						q2Sum++
					}
				}
			}
			if x < BOARD_WIDTH/2 && y > BOARD_HEIGHT/2 {
				for _, robot := range robots {
					if robot.x == x && robot.y == y {
						q3Sum++
					}
				}
			}
			if x > BOARD_WIDTH/2 && y > BOARD_HEIGHT/2 {
				for _, robot := range robots {
					if robot.x == x && robot.y == y {
						q4Sum++
					}
				}
			}
		}
	}

	return q1Sum, q2Sum, q3Sum, q4Sum
}
