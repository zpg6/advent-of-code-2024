// https://adventofcode.com/2024/day/9

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

type FilePiece struct {
	id int
}

type FileGap struct {
	start  int
	length int
}

type File struct {
	id     int
	start  int
	length int
}

func main() {
	fmt.Println("Advent of Code 2024 - Day 9\n---")

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

	var diskmap []FilePiece
	i := 0
	id := 0
	for _, line := range lines {
		for _, char := range line {
			asInt, err := strconv.Atoi(string(char))
			if err != nil {
				panic(err)
			}
			if i%2 == 0 {
				for j := 0; j < asInt; j++ {
					diskmap = append(diskmap, FilePiece{id: id})
				}
				id++
			} else {
				for j := 0; j < asInt; j++ {
					diskmap = append(diskmap, FilePiece{id: -1})
				}
			}
			i++
		}
	}

	// Now backfill the spaces with file pieces.
	for i := 0; i < len(diskmap); i++ {
		if diskmap[i].id == -1 {

			// replace with the last file piece
			for j := len(diskmap) - 1; j > i; j-- {
				if diskmap[j].id != -1 {
					diskmap[i].id = diskmap[j].id
					diskmap[j].id = -1
					break
				}
			}
		}
	}

	// Checksum = sum(i * id) for all i where id != -1
	checksum := 0
	for i, piece := range diskmap {
		if piece.id != -1 {
			checksum += (i * piece.id)
		}
	}

	fmt.Println("Part 1 Checksum:", checksum) // Correct answer: 6288707484810

	// ======= PART 2 =======

	diskmap = []FilePiece{}
	i = 0
	id = 0
	for _, line := range lines {
		for _, char := range line {
			asInt, err := strconv.Atoi(string(char))
			if err != nil {
				panic(err)
			}
			if i%2 == 0 {
				for j := 0; j < asInt; j++ {
					diskmap = append(diskmap, FilePiece{id: id})
				}
				id++
			} else {
				for j := 0; j < asInt; j++ {
					diskmap = append(diskmap, FilePiece{id: -1})
				}
			}
			i++
		}
	}

	// make a map of each file id to its length
	fileLengths := make(map[int]File)
	for i := 0; i < len(diskmap); i++ {
		if diskmap[i].id != -1 {
			if _, ok := fileLengths[diskmap[i].id]; !ok {
				fileLengths[diskmap[i].id] = File{id: diskmap[i].id, start: i, length: 1}
			} else {
				fileLengths[diskmap[i].id] = File{id: diskmap[i].id, start: fileLengths[diskmap[i].id].start, length: fileLengths[diskmap[i].id].length + 1}
			}
		}
	}

	// 00...111...2...333.44.5555.6666.777.888899
	// 0099.111...2...333.44.5555.6666.777.8888..
	// 0099.1117772...333.44.5555.6666.....8888..
	// 0099.111777244.333....5555.6666.....8888..
	// 00992111777.44.333....5555.6666.....8888..

	// Chart the gaps (dots)
	fileGaps := []FileGap{}
	fileGap := FileGap{start: -1, length: 0}
	for i := 0; i < len(diskmap); i++ {
		if diskmap[i].id == -1 {
			if fileGap.start == -1 {
				fileGap.start = i
			}
			fileGap.length++
		} else {
			if fileGap.start != -1 {
				fileGaps = append(fileGaps, fileGap)
				fileGap = FileGap{start: -1, length: 0}
			}
		}
	}

	// Walk from last file to first file, attempting to use it to fill a gap
	for i := 0; i < len(fileLengths); i++ {
		file := fileLengths[len(fileLengths)-1-i]

		for gapIdx, fileGap := range fileGaps {

			if file.length > fileGap.length {
				// File is too big for this gap
				continue
			}
			if file.start < fileGap.start {
				// File can only move left
				continue
			}

			// Remove the file from its original position
			for j := 0; j < file.length; j++ {
				diskmap[file.start+j].id = -1
			}

			// Fill the gap with the file
			for g := 0; g < file.length; g++ {
				diskmap[fileGap.start+g].id = file.id
			}

			// Update the gaps list by removing the gap we just filled
			fileGaps[gapIdx].start = fileGaps[gapIdx].start + file.length
			fileGaps[gapIdx].length = fileGaps[gapIdx].length - file.length

			// only use the file once
			break
		}
	}

	// Checksum = sum(i * id) for all i where id != -1
	checksum = 0
	for i, piece := range diskmap {
		if piece.id != -1 {
			checksum += (i * piece.id)
		}
	}

	fmt.Println("Part 2 Checksum:", checksum) // Correct answer: 6311837662089
}
