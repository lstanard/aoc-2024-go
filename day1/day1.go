package day1

import (
	"aoc-2024-go/utils"
	"fmt"
	"sort"
	"strconv"
	"strings"
)

func ParseInput(sample bool) ([]int, []int) {
  path := ""
  if sample {
    path = "input-sample.txt"
  } else {
    path = "input.txt"
  }
  success, data := utils.ReadInputFile(path)

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

func Part1(sample bool) int {
  listA, listB := ParseInput(sample)

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
  return totalDistance
}

func Part2(sample bool) int {
  A, B := ParseInput(sample)
  countMap := make(map[int]int)
  for _, b := range B {
    countMap[b]++
  }

  result := 0
  for _, a := range A {
    if count, ok := countMap[a]; ok {
      result += a * count
    }
  }

  fmt.Println("Result:", result)
  return result
}
