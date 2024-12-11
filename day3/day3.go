package day3

import (
	"aoc-2024-go/utils"
	"regexp"
	"strconv"
)

func parseInput(sample bool, part int) string {
	path := ""
	if sample {
		if part == 2 {
			path = "input-sample-part2.txt"
		} else {
			path = "input-sample.txt"
		}
	} else {
		path = "input.txt"
	}
	success, data := utils.ReadInputFile(path)

	if !success {
		return "error!"
	}

	return string(data[:])
}

func multiply(mul string) int {
	nums := regexp.MustCompile(`\d+`).FindAllString(mul, -1)
	num1 := nums[0]
	num2 := nums[1]
	num1Int, _ := strconv.Atoi(num1)
	num2Int, _ := strconv.Atoi(num2)
	product := num1Int * num2Int
	return product
}

func Part1(sample bool) int {
	input := parseInput(sample, 1)

	pattern := `mul\(\d+,\d+\)`
	re := regexp.MustCompile(pattern)
	allMatches := re.FindAllString(input, -1)

	total := 0
	for _, match := range allMatches {
		product := multiply(match)
		total += product
	}

	return total
}

func Part2(sample bool) int {
	input := parseInput(sample, 2)

	pattern := `mul\(\d+,\d+\)|don't|do`
	re := regexp.MustCompile(pattern)
	allMatches := re.FindAllString(input, -1)

	total := 0
	shouldMultiply := true
	for _, match := range allMatches {
		if match == "don't" {
			shouldMultiply = false
		} else if match == "do" {
			shouldMultiply = true
		} else if shouldMultiply {
			product := multiply(match)
			total += product
		}
	}

	return total
}
