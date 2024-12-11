package day4_test

import (
	"aoc-2024-go/day4"
	"testing"
)

func TestPart1_SampleData(t *testing.T) {
	result := day4.Part1(true)
	if result != 18 {
		t.Errorf("Expected 18, got %d", result)
	}
}
