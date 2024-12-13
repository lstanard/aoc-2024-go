package day4

import (
	"aoc-2024-go/utils"
	"fmt"
)

func parseInput(sample bool) []string {
	path := ""
	if sample {
		path = "input-sample.txt"
	} else {
		path = "input.txt"
	}
	success, data := utils.ReadInputFile(path)

	if !success {
		return nil
	}

	lines := utils.SplitInputToLines(data)
	return lines
}

func validateDirection(i, j int, dir [2]int, rows, cols int) bool {
	// Can't move up or down
	if i-dir[1] < 0 || i+dir[1] >= rows {
		return false
	}
	// Can't move left or right
	if j-dir[0] < 0 || j+dir[0] >= cols {
		return false
	}
	return true
}

// Grid traversal / depth-first search
// https://www.youtube.com/watch?v=w8DYqrWjX1w
func Part1(sample bool) int {
	lines := parseInput(sample)
	fmt.Println(lines)

	// Directions for grid traversal
  // directions are [x, y]
	directions := [8][2]int{
		{0, 1},   // down
		{1, 1},   // right-down
		{1, 0},   // right
    {1, -1},  // right-up
    {0, -1},  // up
    {-1, 1},  // left-up
		{-1, 0},  // left
		{-1, -1}, // left-down
	}
	// OR word := "XMAS"? Depends
	letters := [4]string{"X", "M", "A", "S"}
	letterIndex := 0
	rows := len(lines) - 1 // this isn't great but there's an extra empty line that I should strip out
	cols := len(lines[0])
	// visited := make([][]bool, rows)

	// Loop through each row
	// Loop through each character in the row
	// Loop through each direction
	// Loop through each character in the word
	// Store cooridnates where each word is found (x, y, direction)

	// NOTE: In this case i = x and j = y
	for i := 0; i < rows; i++ {
		// fmt.Println("Row", lines[i])
		for j := 0; j < cols; j++ {
			if (lines[i]) == "" {
				continue
			}
			char := string(lines[i][j])
			if char == letters[letterIndex] {
				fmt.Println("Found", char, "at", i, j)
				for _, dir := range directions {
					validDirection := validateDirection(i, j, dir, rows, cols)
          if (validDirection) {
            fmt.Println("Valid direction", dir)
          }
				}
			}
		}
	}

	return 0
}
