package day01

import (
	"fmt"
)

func Day02() {
	left, right, err := LoadFile()
	if err != nil {
		panic("an error ocurred")
	}

	m := make(map[int]int)

	for _, value := range right {
		m[value]++
	}

	times := 0
	for _, value := range left {
		times += (value * m[value])
	}

	fmt.Printf("%d\n", times)
}
