package readinput

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

func ReadFile(fileName string) [][]int {
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
