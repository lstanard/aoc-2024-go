package day2

import "aoc-2024-go/utils"

func ParseInput(sample bool) ([]string) {
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
