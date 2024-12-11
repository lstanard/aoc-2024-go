package day2_test

import (
	"aoc-2024-go/day2"
	"testing"
)

func TestPart1_SampleData(t *testing.T) {
	result := day2.Part1(true)
	if result != 2 {
		t.Errorf("Expected 2, got %d", result)
	}
}

func TestPart1(t *testing.T) {
	result := day2.Part1(false)
	if result != 213 {
		t.Errorf("Expected 213, got %d", result)
	}
}

func TestPart2_SampleData(t *testing.T) {
	result := day2.Part2(true)
	if result != 4 {
		t.Errorf("Expected 4, got %d", result)
	}
}

func TestPart2(t *testing.T) {
	result := day2.Part2(false)
	if result != 285 { // it's between 273 and 295
		t.Errorf("Expected 4, got %d", result)
	}
}
