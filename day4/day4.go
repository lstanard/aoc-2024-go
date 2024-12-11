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

func Part1(sample bool) int {
  lines := parseInput(sample)
  fmt.Println(lines)

  return 0
}
