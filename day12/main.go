// https://adventofcode.com/2024/day/12

package main

import (
	"bufio"
	"fmt"
	"os"
)

type Coord struct {
	x, y int
}

type FloatCoord struct {
	x, y float64
}

type Cluster struct {
	letter string
	coords []Coord
}

func main() {
	fmt.Println("Advent of Code 2024 - Day 12\n---")

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

	// Read into 2d map of strings
	m := make(map[int]map[int]string)
	for y, line := range lines {
		m[y] = make(map[int]string)
		for x, val := range line {
			m[y][x] = string(val)
		}
	}

	// Find all clusters
	clusters := []Cluster{}
	visited := make(map[Coord]bool)
	for y := 0; y < len(m); y++ {
		for x := 0; x < len(m[y]); x++ {
			if visited[Coord{x, y}] {
				continue
			}

			// explore this letter in all directions
			clusterLetter := m[y][x]
			clusterCoords := []Coord{{x, y}}
			visited[Coord{x, y}] = true

			// explore all directions
			if x < len(m[y])-1 && m[y][x+1] == clusterLetter {
				clusterCoords = explore(m, x+1, y, visited, clusterLetter, clusterCoords)
			}
			if x > 0 && m[y][x-1] == clusterLetter {
				clusterCoords = explore(m, x-1, y, visited, clusterLetter, clusterCoords)
			}
			if y < len(m)-1 && m[y+1][x] == clusterLetter {
				clusterCoords = explore(m, x, y+1, visited, clusterLetter, clusterCoords)
			}
			if y > 0 && m[y-1][x] == clusterLetter {
				clusterCoords = explore(m, x, y-1, visited, clusterLetter, clusterCoords)
			}

			// add the cluster
			clusters = append(clusters, Cluster{clusterLetter, clusterCoords})
		}
	}

	// Calculate the fence length
	fence := 0
	for _, cluster := range clusters {

		clusterArea := len(cluster.coords)
		clusterPerimeter := 0

		// find the perimeter (edges that are not shared with another coord in the cluster)
		for _, coord := range cluster.coords {
			x, y := coord.x, coord.y
			if x == 0 || m[y][x-1] != cluster.letter {
				clusterPerimeter++
			}
			if x == len(m[y])-1 || m[y][x+1] != cluster.letter {
				clusterPerimeter++
			}
			if y == 0 || m[y-1][x] != cluster.letter {
				clusterPerimeter++
			}
			if y == len(m)-1 || m[y+1][x] != cluster.letter {
				clusterPerimeter++
			}
		}
		fence += clusterArea * clusterPerimeter
	}
	fmt.Println("Fence length with perimeter:", fence) // Correct answer: 1485656

	// ======= PART 2 =======

	// Calculate the fence length with sides
	fence = 0

	// Like the perimeter, but sides (where two consecutive perimeter edges are connected)
	for _, cluster := range clusters {

		// Instead of the squares, get the corners between the squares.
		// Use a map as a set to avoid duplicates
		corner_coords := map[FloatCoord]struct{}{}
		for x := 0; x < len(m[0]); x++ {
			for y := 0; y < len(m); y++ {

				// check if any of the corner coords are in the cluster
				corner_coords[FloatCoord{float64(x) - 0.5, float64(y) - 0.5}] = struct{}{}
				corner_coords[FloatCoord{float64(x) - 0.5, float64(y) + 0.5}] = struct{}{}
				corner_coords[FloatCoord{float64(x) + 0.5, float64(y) - 0.5}] = struct{}{}
				corner_coords[FloatCoord{float64(x) + 0.5, float64(y) + 0.5}] = struct{}{}
			}
		}

		// For each corner coordinate pair, see if it is a corner to the cluster.
		// 4-corners are in cluster -> interior to the cluster
		// 3-corners are in cluster -> corner
		// 2-corners are in cluster -> if they are opposite, 2 corners, otherwise not a corner
		// 1-corner is in cluster -> corner
		corners := 0
		for corner_coord := range corner_coords {
			x, y := corner_coord.x, corner_coord.y
			tl, tr, bl, br := Coord{int(x - 0.5), int(y - 0.5)}, Coord{int(x + 0.5), int(y - 0.5)}, Coord{int(x - 0.5), int(y + 0.5)}, Coord{int(x + 0.5), int(y + 0.5)}
			inCluster := []bool{containsCoord(cluster.coords, tl), containsCoord(cluster.coords, tr), containsCoord(cluster.coords, bl), containsCoord(cluster.coords, br)}
			numInCluster := 0
			for _, in := range inCluster {
				if in {
					numInCluster++
				}
			}
			if numInCluster == 0 {
				continue
			} else if numInCluster == 4 {
				continue // interior to the cluster
			} else if numInCluster == 2 && ((inCluster[0] && inCluster[3]) || (inCluster[1] && inCluster[2])) {
				corners += 2
			} else if numInCluster == 1 {
				corners++
			} else if numInCluster == 3 {
				corners++
			}
		}

		fence += len(cluster.coords) * corners
	}

	fmt.Println("Fence length with sides:", fence) // Correct answer: 899196
}

func containsCoord(coords []Coord, c Coord) bool {
	for _, coord := range coords {
		if coord == c {
			return true
		}
	}
	return false
}

func explore(m map[int]map[int]string, x, y int, visited map[Coord]bool, clusterLetter string, clusterCoords []Coord) []Coord {
	if visited[Coord{x, y}] {
		return clusterCoords
	}

	visited[Coord{x, y}] = true
	clusterCoords = append(clusterCoords, Coord{x, y})

	// explore all directions
	if m[y][x+1] == clusterLetter {
		clusterCoords = explore(m, x+1, y, visited, clusterLetter, clusterCoords)
	}
	if m[y][x-1] == clusterLetter {
		clusterCoords = explore(m, x-1, y, visited, clusterLetter, clusterCoords)
	}
	if m[y+1][x] == clusterLetter {
		clusterCoords = explore(m, x, y+1, visited, clusterLetter, clusterCoords)
	}
	if m[y-1][x] == clusterLetter {
		clusterCoords = explore(m, x, y-1, visited, clusterLetter, clusterCoords)
	}

	return clusterCoords
}
