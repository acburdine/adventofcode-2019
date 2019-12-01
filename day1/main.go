package main

import (
	"fmt"
	"log"

	"github.com/acburdine/adventofcode-2019/input"
)

func try(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	inputData, err := input.GetInputFileAsInt("./day1/input.txt")
	try(err)

	fmt.Printf("Part 1: %d\n", part1(inputData))
	fmt.Printf("Part 2: %d\n", part2(inputData))
}

func part1(data []int) int {
	sum := 0

	for _, n := range data {
		sum += (n / 3) - 2
	}

	return sum
}

func part2(data []int) int {
	sum := 0

	for _, n := range data {
		sum += part2recurse(n)
	}

	return sum
}

func part2recurse(n int) int {
	fuel := (n / 3) - 2

	if fuel <= 0 {
		return 0
	}

	return fuel + part2recurse(fuel)
}
