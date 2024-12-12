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

// Grid traversal / depth-first search
// https://www.youtube.com/watch?v=w8DYqrWjX1w
func Part1(sample bool) int {
  lines := parseInput(sample)
  fmt.Println(lines)

  // Directions for grid traversal
  // directions := [8][2]int{
  //   {0,1},    // up
  //   {1,1},    // up-right
  //   {1,0},    // right
  //   {-1,0},   // left
  //   {-1,-1},  // down-left
  //   {0,-1},   // down
  //   {1,-1},   // down-right
  //   {-1,1},   // up-left
  // }
  // OR word := "XMAS"? Depends
  letters := [4]string{"X", "M", "A", "S"}
  letterIndex := 0
  rows := len(lines)
  // cols := len(lines[0])
  // visited := make([][]bool, rows)

  // Loop through each row
  for i := 0; i < rows; i++ {
    // Loop through each character in the row
    row := lines[i]
    for j := 0; j < len(row); j++ {
      // OR char := lines[i][j]
      char := string(row[j])
      if char == letters[letterIndex] {
        // for _, dir := range directions {
        // }
      }
    }
  }

  return 0
}
