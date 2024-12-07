package day1_test

import (
  "aoc-2024-go/day1"
  "testing"
)

func TestPart1(t *testing.T) {
  result := day1.Part1()
  if result != 1646452 {
    t.Errorf("Expected 1646452, got %d", result)
  }
}

func TestPart2(t *testing.T) {
  result := day1.Part2()
  if result != 23609874 {
    t.Errorf("Expected 23609874, got %d", result)
  }
}
