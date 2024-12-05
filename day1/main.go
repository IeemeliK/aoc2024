package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

func readFile(fileName string) ([]int, []int) {
	file, err := os.Open(fileName)
	if err != nil {
		panic(err)
	}

	defer func() {
		if err := file.Close(); err != nil {
			panic(err)
		}
	}()

	scanner := bufio.NewScanner(file)

	list1 := make([]int, 0)
	list2 := make([]int, 0)
	for scanner.Scan() {
		line := scanner.Text()
		arr := strings.Split(line, "  ")

		if len(arr) == 0 || len(arr) == 1 {
			continue
		}

		num1, _ := strconv.Atoi(strings.TrimSpace(arr[0]))
		num2, _ := strconv.Atoi(strings.TrimSpace(arr[1]))
		list1 = append(list1, num1)
		list2 = append(list2, num2)
	}

	return list1, list2
}

func absVal(num int) int {
	if num < 0 {
		return num * -1
	}
	return num
}

func calculateDistance(list1 []int, list2 []int) int {
	if len(list1) != len(list2) {
		log.Fatal("Lists are not of equal length")
	}
	sort.Ints(list1)
	sort.Ints(list2)

	var sum int
	for i := 0; i <= len(list1)-1; i++ {
		sum += absVal(list1[i] - list2[i])
	}

	return sum
}

func similarityMap(list1 []int, list2 []int) int {
	var sum int
	cachedValues := map[int]int{}
	for _, v := range list1 {
		if cache, ok := cachedValues[v]; ok {
			sum += cache
		} else {
			count := 0
			for _, y := range list2 {
				if y == v {
					count++
				}
			}

			similarityScore := v * count
			cachedValues[v] = similarityScore
			sum += similarityScore
		}
	}

	return sum
}

func main() {
	list1, list2 := readFile("input.txt")

	fmt.Println(calculateDistance(list1, list2))
	fmt.Println(similarityMap(list1, list2))
}
