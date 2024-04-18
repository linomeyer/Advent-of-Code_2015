package main

import (
	"AdventOfCode_2015/commons"
	"fmt"
	"sort"
	"strconv"
	"strings"
)

func main() {
	lines := commons.ReadLines("./day02/input.txt")
	sumWrapping := calculateWrappingNeeded(lines)
	sumRibbon := calculateRibbonNeeded(lines)

	fmt.Println(sumWrapping)
	fmt.Println(sumRibbon)
}

func calculateRibbonNeeded(lines []string) int {
	sum := 0
	for _, line := range lines {
		lwh := toIntArray(line)
		sort.Ints(lwh)
		sum += (2*lwh[0] + 2*lwh[1]) + (lwh[0] * lwh[1] * lwh[2])
	}
	return sum
}

func calculateWrappingNeeded(lines []string) int {
	sum := 0

	for _, line := range lines {
		lwh := toIntArray(line)

		sides := getSides(lwh)

		sum += 2*sides[0] + 2*sides[1] + 2*sides[2]

		sort.Ints(sides)
		sum += sides[0]
	}
	return sum
}

func getSides(lwh []int) []int {
	lw := lwh[0] * lwh[1]
	wh := lwh[1] * lwh[2]
	hl := lwh[2] * lwh[0]
	return []int{lw, wh, hl}
}

func toIntArray(line string) []int {
	var lwh []int
	for _, element := range strings.Split(line, "x") {
		number, _ := strconv.Atoi(element)
		lwh = append(lwh, number)
	}
	return lwh
}
