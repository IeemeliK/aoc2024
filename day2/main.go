package main

import (
	"aoc24_2/read_input"
	"aoc24_2/validator"
	"fmt"
)

func validateByRemoving(list []int) bool {
	for i := 0; i < len(list); i++ {
		newList := append([]int{}, list[:i]...)
		newList = append(newList, list[i+1:]...)

		if validate(newList) {
			return true
		}
	}

	return false
}

func validate(list []int) bool {
	decr := validator.Decreasing(list)
	incr := validator.Increasing(list)
	maxDiff := validator.MaxDiff(list, 3)
	minDiff := validator.MinDiff(list, 1)

	if maxDiff && minDiff && (decr || incr) {
		return true
	}

	return false
}

func countSafeLevels(list [][]int, remove bool) int {
	safeCount := 0
	for _, outer := range list {
		if validate(outer) || (validateByRemoving(outer) && remove) {
			safeCount++
		}
	}
	return safeCount
}

func main() {
	list := readinput.ReadFile("input.txt")
	fmt.Println(countSafeLevels(list, true))
}
