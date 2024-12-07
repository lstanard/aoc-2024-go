package day1

import (
	"aoc-2024-go/utils"
	"fmt"
	"sort"
	"strconv"
	"strings"
)

func ParseInput() ([]int, []int) {
  success, data := utils.ReadInputFile("day1/input.txt")

  if !success {
    return nil, nil
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

  return listA, listB
}

func Day1Part1() {
  listA, listB := ParseInput()

  totalDistance := 0
  for i := 0; i < len(listA); i++ {
    distance := 0
    if listA[i] > listB[i] {
      distance = listA[i] - listB[i]
    } else {
      distance = listB[i] - listA[i]
    }
    totalDistance += distance
  }

  fmt.Println("Total Distance:", totalDistance)
}
