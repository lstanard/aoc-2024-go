package main

import (
	"aoc-2024-go/utils"
	"fmt"
	"strconv"
)

func main() {
  success, data := utils.ReadInputFile("day1/input-sample.txt")

  if !success {
    return
  }

  listA := []int{}
  listB := []int{}
  for i, b := range data {
    fmt.Println("byte:", b)
    s := string(b)

    if s != "" {
      fmt.Println("byte to string:", s)
      num, err := strconv.Atoi(s)
      fmt.Println("number/index:", num, i)

      if err != nil {
        fmt.Println("Error converting to number:", err)
      }

      if i % 2 == 0 {
        listA = append(listA, num)
      } else {
        listB = append(listB, num)
      }
    }
  }

  fmt.Println("List A:", listA)
  fmt.Println("List B:", listB)
}
