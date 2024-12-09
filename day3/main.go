package main

import (
	readinput "aoc24_3/read_input"
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

func parseNumbers(text string) []int {
	stringNums := regexp.MustCompile(`\d+`).FindAllString(text, -1)

	intSlice := make([]int, 0, len(stringNums))

	for _, v := range stringNums {
		i, err := strconv.Atoi(v)
		if err != nil {
			panic(err)
		}
		intSlice = append(intSlice, i)
	}

	return intSlice
}

func findExpressions(text string) []string {

	// Finds all strings that look exactly like "mul(digits,digits)"
	list := regexp.MustCompile(`(mul\(\d+,\d+\))`).FindAllString(text, -1)

	return list
}

func removeDisabled(inputText string) string {
	// Finds everything between don'() and do()
	replacer := regexp.MustCompile(`don't\(\).*?do\(\)`)
	text := replacer.ReplaceAllString(inputText, "")

	// Cuts everything after last don't()
	if before, _, found := strings.Cut(text, "don't()"); found {
		text = before
	}

	return text
}

func main() {
	list := readinput.ReadFile("input.txt")
	list = removeDisabled(list)
	sum := 0
	for _, string := range findExpressions(list) {
		nums := parseNumbers(string)
		fmt.Println(nums)
		sum += nums[0] * nums[1]
	}

	fmt.Println(sum)
}
