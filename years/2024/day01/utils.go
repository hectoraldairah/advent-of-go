package day01

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func LoadFile() ([]int, []int, error) {
	file, err := os.Open("./years/2024/day01/input.txt")
	if err != nil {
		fmt.Printf("%v", err)
	}

	defer file.Close()

	var firstColumn []int
	var secondColumn []int

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()

		fields := strings.Fields(line)

		first, err := strconv.Atoi(fields[0])
		if err != nil {
			return nil, nil, fmt.Errorf("invalid number in first column: %s", fields[0])
		}

		second, err := strconv.Atoi(fields[1])
		if err != nil {
			return nil, nil, fmt.Errorf("invalid number in first column: %s", fields[1])
		}

		firstColumn = append(firstColumn, first)
		secondColumn = append(secondColumn, second)

	}

	return firstColumn, secondColumn, nil
}
