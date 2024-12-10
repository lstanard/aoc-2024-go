package day1_test

import (
  "aoc-2024-go/day1"
  "testing"
)

func TestPart1_SampleData(t *testing.T) {
  result := day1.Part1(true)
  if result != 11 {
    t.Errorf("Expected 1646452, got %d", result)
  }
}

func TestPart2_SampleData(t *testing.T) {
  result := day1.Part2(true)
  if result != 31 {
    t.Errorf("Expected 23609874, got %d", result)
  }
}

func TestPart1(t *testing.T) {
  result := day1.Part1(false)
  if result != 1646452 {
    t.Errorf("Expected 1646452, got %d", result)
  }
}

func TestPart2(t *testing.T) {
  result := day1.Part2(false)
  if result != 23609874 {
    t.Errorf("Expected 23609874, got %d", result)
  }
}
