// https://adventofcode.com/2024/day/13

package main

import (
	"bufio"
	"fmt"
	"os"
)

type Game struct {
	ax, ay, bx, by, px, py int
}

func main() {
	fmt.Println("Advent of Code 2024 - Day 13\n---")

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
	var games []Game
	for i := 0; i < len(lines); i += 4 {
		var game Game
		fmt.Sscanf(lines[i], "Button A: X+%d, Y+%d", &game.ax, &game.ay)
		fmt.Sscanf(lines[i+1], "Button B: X+%d, Y+%d", &game.bx, &game.by)
		fmt.Sscanf(lines[i+2], "Prize: X=%d, Y=%d", &game.px, &game.py)
		games = append(games, game)
	}

	// ======= PART 1 =======

	// Pressing A costs 3 tokens, B costs 1 token.
	// Find the cheapest way to get to the prize.
	// The prize is at (px, py), and we start at (0, 0).
	// Max 100 moves, so calculate 100 As,0 Bs, 99 As,1 B, etc.
	totalCost := 0
	for _, game := range games {
		loser := 3*100 + 100 // Max cost
		minCost := loser

		for i := 0; i <= 100; i++ {
			for j := 0; j <= 100; j++ {
				cost := 3*i + j

				px := game.ax*i + game.bx*j
				py := game.ay*i + game.by*j

				if px == game.px && py == game.py {
					if cost < minCost {
						minCost = cost
					}
				}
			}
		}

		if minCost < loser {
			totalCost += minCost
		}
	}

	fmt.Println("Total cost:", totalCost) // Correct answer: 34393

	// ======= PART 2 =======

	// Add 10000000000000 to every prize's X and Y
	for gameIdx, game := range games {
		game.px += 10000000000000
		game.py += 10000000000000
		games[gameIdx] = game
	}

	totalCost = 0
	for _, game := range games {
		if game.px < game.ax && game.px < game.bx {
			// If the prize is to the left of both buttons, we can't reach it
			continue
		}
		if game.py < game.ay && game.py < game.by {
			// If the prize is below both buttons, we can't reach it
			continue
		}

		// Solving the system of equations:
		// AMoves = (Px*By - Py*Bx) / (Ax*By - Ay*Bx)
		// BMoves = (Px - Ax*Ca) / Bx
		a := float64(game.px*game.by-game.py*game.bx) / float64(game.ax*game.by-game.ay*game.bx)
		b := (float64(game.px) - float64(game.ax)*a) / float64(game.bx)

		if isWholeNumber(a) && isWholeNumber(b) {
			totalCost += int(a)*3 + int(b)
		}
	}

	fmt.Println("Total cost:", totalCost) // Correct answer: 83551068361379
}

func isWholeNumber(num float64) bool {
	return num == float64(int(num)) // Equivalent check
}
