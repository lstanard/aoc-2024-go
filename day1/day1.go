package main

import (
	"aoc-2024-go/utils"
	"fmt"
	"sort"
	"strconv"
	"strings"
)

func main() {
  success, data := utils.ReadInputFile("day1/input-sample.txt")

  if !success {
    return
  }

  lines := utils.SplitInputToLines(data)

  listA := []int{}
  listB := []int{}
  for _, line := range lines {
    nums := strings.Fields(line)

    if len(nums) <= 0 {
      continue
    }
    valA, _ := strconv.Atoi(nums[0])
    valB, _ := strconv.Atoi(nums[1])

    listA = append(listA, valA)
    listB = append(listB, valB)
  }

  sort.Ints(listA)
  sort.Ints(listB)
  fmt.Println("List A:", listA)
  fmt.Println("List B:", listB)

  totalDistance := 0

  return totalDistance
}
