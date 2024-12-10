package day2

import (
	"aoc-2024-go/utils"
	"strconv"
	"strings"
)

func ParseInput(sample bool) []string {
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

func isLineSafe(line []string, direction string) bool {
  // down to the last number, return true if we made it this far
  if len(line) < 2 {
    return true
  }

  num1, _ := strconv.Atoi(line[0])
  num2, _ := strconv.Atoi(line[1])

  if num1 > num2 {
    // decreasing
    if direction == "up" {
      return false
    }
    if num1 - num2 > 3 {
      return false
    }
    return isLineSafe(line[1:], "down")
  } else if num1 < num2 {
    // increasing
    if direction == "down" {
      return false
    }
    if num2 - num1 > 3 {
      return false
    }
    return isLineSafe(line[1:], "up")
  } else {
    // the two numbers are equal
    return false
  }
}

func getSafeCount(lines []string) int {
  safeCount := 0
  for _, line := range lines {
    if len(line) == 0 {
      continue
    }

    sNums := strings.Fields(line)
    safe := isLineSafe(sNums, "")
    if safe {
      safeCount++
    }
  }

  return safeCount
}

func Part1(sample bool) int {
  lines := ParseInput(sample)
  safeCount := getSafeCount(lines)

  return safeCount
}
