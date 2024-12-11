package day3

import (
	"aoc-2024-go/utils"
	"regexp"
	"strconv"
)

func parseInput(sample bool) string {
	path := ""
	if sample {
		path = "input-sample.txt"
	} else {
		path = "input.txt"
	}
	success, data := utils.ReadInputFile(path)

	if !success {
		return "error!"
	}

	return string(data[:])
}

func Part1(sample bool) int {
	input := parseInput(sample)

	pattern := `mul\(\d+,\d+\)`
	re := regexp.MustCompile(pattern)
	allMatches := re.FindAllString(input, -1)

	total := 0
	for _, match := range allMatches {
		nums := regexp.MustCompile(`\d+`).FindAllString(match, -1)
		num1 := nums[0]
		num2 := nums[1]
		num1Int, _ := strconv.Atoi(num1)
		num2Int, _ := strconv.Atoi(num2)
		product := num1Int * num2Int
		total += product
	}

	return total
}
