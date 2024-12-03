package day01

import (
	"fmt"
	"math"
	"sort"
)

func Day01() {
	left, right, err := LoadFile()
	if err != nil {
		panic("an error ocurred")
	}

	sort.Ints(left)
	sort.Ints(right)

	res := 0
	for k := range left {
		diff := int(math.Abs(float64(left[k]) - float64(right[k])))
		res += diff
	}

	fmt.Printf("%d\n", res)
}
