package main

import (
	"fmt"
	"time"

	"github.com/hectoraldairah/advent2024/years/2024/day01"
)

func main() {
	start := time.Now()

	day01.Day02()

	elpased := time.Since(start)

	fmt.Printf("Execution time %s\n", elpased)
}
