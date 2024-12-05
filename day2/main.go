package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func readFile(fileName string) [][]int {
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

	list := make([][]int, 0)
	for scanner.Scan() {
		temp := make([]int, 0)
		line := scanner.Text()
		arr := strings.Split(line, " ")

		if len(arr) == 0 || len(arr) == 1 {
			continue
		}

		for _, v := range arr {
			num, err := strconv.Atoi(strings.TrimSpace(v))
			if err != nil {
				panic(err)
			}
			temp = append(temp, num)
		}
		list = append(list, temp)
	}

	return list
}

func absVal(num int) int {
	if num < 0 {
		return num * -1
	}
	return num
}

func countSafeLevels(list [][]int) int {
	safeCount := 0
	for _, outer := range list {
		safeValues := true
		increasing := outer[1] > outer[0]
		decreasing := outer[1] < outer[0]
		removed := false

		for i := 1; i <= len(outer)-1; i++ {
			if increasing {
				check := outer[i] > outer[i-1]

				if !check && !removed {
					check = outer[i+1] > outer[i-1]
					removed = true
				}
				increasing = check
			}

			if decreasing {
				check := outer[i] < outer[i-1]

				if !check && !removed {
					check = outer[i+1] < outer[i-1]
					removed = true
				}
				decreasing = check
			}

			if safeValues {
				num := absVal(outer[i] - outer[i-1])
				check := num <= 3 && num >= 1

				if !check && !removed {
					num = absVal(outer[i+1] - outer[i-1])
					check = num <= 3 && num >= 1
					removed = true
				}
				safeValues = check
			}

			if (!increasing && !decreasing) || !safeValues {
				break
			}
		}

		if (increasing || decreasing) && safeValues {
			safeCount++
		}

		fmt.Printf("list: %v, decreasing: %t, increasing: %t, safeValues: %t, removed: %t\n", outer, decreasing, increasing, safeValues, removed)
	}
	return safeCount
}

func main() {
	list := readFile("input_test.txt")
	fmt.Println(countSafeLevels(list))

}
