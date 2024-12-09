package readinput

import (
	"bufio"
	"os"
	"strings"
)

func ReadFile(fileName string) [][]string {
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

	// Collect all lines
	lines := make([][]string, 0)
	for scanner.Scan() {
		line := scanner.Text()
		charArray := strings.Split(line, "")
		lines = append(lines, charArray)
	}

	return lines
}
