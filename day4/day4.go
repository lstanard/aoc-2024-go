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

  // Turn each line into an array of strings
  // Go through each line and stop whenever an "X" is found
  // Begin checking for occurences of "M"
  // Horizontally:
  // Is "M" the next letter?
  // Is "M" the previous letter?
  // Vertically:
  // Does "M" occur directly above "X"?
  // Does "M" occur directly below "X"?
  // Diagonally:
  // Keep in mind if there are fewer rows above/below we can exit early. i.e. no need to check up if this is the first 3 rows
  // Does "M" occur on the line above, but i - 1?
  // Does "M" occur on the line above, but i + 1?
  // Does "M" occur on the line below, but i - 1?
  // Does "M" occur on the line below, but i + 1?
  // If no occurence of "M" is found we can exit early
  // If "M" is found in one of those positions...
  // Repeat these instructions for each letter in "XMAS" until the end

  // Half-baked idea:
  // Could some kind of [x,y] coordinates be helpful here?
  // Map each letter to an [x,y] coordinate, then apply rules to every occurence of "X"?
  // Example: "X" at [5,1] would check for "M" at [4,1] [6,1] [5,2] [4,2] [6,2]
  // x coordinate is less than 4 so no need to check up
  // As we continue this pattern, would need to know which "direction" we're going in
  // For example, we could find an "X" at [5,1], then a "M" at [6,2], but the only valid "A" would be at [7,3]
  // Also need to keep in mind that a single "X" could have multiple instances of "XMAS" branching off of it in different directions

  return 0
}
