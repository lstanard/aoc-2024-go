package day3_test

import (
	"aoc-2024-go/day3"
	"testing"
)

func TestPart1_SampleData(t *testing.T) {
	result := day3.Part1(true)
	if result != 161 {
		t.Errorf("Expected 161, got %d", result)
	}
}

func TestPart2_SampleData(t *testing.T) {
	result := day3.Part1(false)
	if result != 173785482 {
		t.Errorf("Expected 161, got %d", result)
	}
}

