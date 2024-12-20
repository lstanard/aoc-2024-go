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

func TestPart1(t *testing.T) {
	result := day3.Part1(false)
	if result != 173785482 {
		t.Errorf("Expected 173785482, got %d", result)
	}
}

func TestPart2_SampleData(t *testing.T) {
	result := day3.Part2(true)
	if result != 48 {
		t.Errorf("Expected 48, got %d", result)
	}
}

func TestPart2(t *testing.T) {
	result := day3.Part2(false)
	if result != 83158140 {
		t.Errorf("Expected 83158140, got %d", result)
	}
}
