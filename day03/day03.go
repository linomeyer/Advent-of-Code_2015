package main

import (
	"AdventOfCode_2015/commons"
	"fmt"
	"slices"
	"sort"
	"strconv"
)

func main() {
	fileString := commons.ReadFileString("./day03/input.txt")

	grid := createGrid(fileString)

	sort.Strings(grid)
	grid = slices.Compact(grid)

	fmt.Println(len(grid))
}

func createGrid(fileString string) []string {
	xS, yS, xR, yR := 0, 0, 0, 0
	grid := []string{strconv.Itoa(xS) + "," + strconv.Itoa(yS)}
	for i, c := range fileString {
		if i%2 == 0 {
			move(c, &xS, &yS)
			grid = append(grid, strconv.Itoa(xS)+","+strconv.Itoa(yS))
		}
		if i%2 == 1 {
			move(c, &xR, &yR)
			grid = append(grid, strconv.Itoa(xR)+","+strconv.Itoa(yR))
		}
	}
	return grid
}

func move(c int32, x *int, y *int) {
	if c == '^' {
		*y++
	}
	if c == 'v' {
		*y--
	}
	if c == '>' {
		*x++
	}
	if c == '<' {
		*x--
	}
}
