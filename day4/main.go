package main

import (
	readinput "aoc24_4/read_input"
	"fmt"
)

func checkDiagonal(list [][]string, yindex int, xindex int, word string, directionX int, directionY int) bool {
	for i, v := range word {
		newY := yindex + i*directionY
		newX := xindex + i*directionX
		if newY < 0 || newY >= len(list) || newX < 0 || newX >= len(list[newY]) || list[newY][newX] != string(v) {
			return false
		}
	}
	return true
}

func checkVertical(list [][]string, yindex int, xindex int, word string, direction int) bool {
	for i, v := range word {
		newY := yindex + i*direction
		if newY < 0 || newY >= len(list) || xindex >= len(list[newY]) || list[newY][xindex] != string(v) {
			return false
		}
	}
	return true
}

func checkHorizontal(list []string, xindex int, word string, direction int) bool {
	for i, v := range word {
		newX := xindex + i*direction
		if newX < 0 || newX >= len(list) || list[newX] != string(v) {
			return false
		}
	}
	return true
}

func findWord(list [][]string, yindex int, xindex int, word string) int {
	count := 0
	// Checks backward
	if checkHorizontal(list[yindex], xindex, word, -1) {
		count++
	}

	// Checks forward
	if checkHorizontal(list[yindex], xindex, word, 1) {
		count++
	}

	// Check down-right
	if checkDiagonal(list, yindex, xindex, word, 1, 1) {
		count++
	}

	// Check down-left
	if checkDiagonal(list, yindex, xindex, word, -1, 1) {
		count++
	}

	// Check up-right
	if checkDiagonal(list, yindex, xindex, word, 1, -1) {
		count++
	}

	// Check up-left
	if checkDiagonal(list, yindex, xindex, word, -1, -1) {
		count++
	}

	// checks down
	if checkVertical(list, yindex, xindex, word, 1) {
		count++
	}

	// checks up
	if checkVertical(list, yindex, xindex, word, -1) {
		count++
	}
	return count
}

func main() {
	list := readinput.ReadFile("input.txt")
	fmt.Println(list)
	word := "XMAS"
	sum := 0
	for y, inner := range list {
		fmt.Println(inner)
		for x, letter := range inner {
			if letter != string(word[0]) {
				continue
			}
			sum += findWord(list, y, x, word)
		}
	}

	fmt.Println(sum)
}
