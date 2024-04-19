package main

import (
	"AdventOfCode_2015/commons"
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

type InstructionSet int

const (
	Toggle InstructionSet = iota
	TurnOn
	TurnOff
)

func main() {
	lines := commons.ReadLines("./day06/input.txt")
	var grid [1000][1000]int

	for _, line := range lines {
		instructions := getInstructions(line)
		if strings.HasPrefix(line, "turn on") {
			toggles(&grid, instructions, TurnOn)
		}
		if strings.HasPrefix(line, "turn off") {
			toggles(&grid, instructions, TurnOff)
		}
		if strings.HasPrefix(line, "toggle") {
			toggles(&grid, instructions, Toggle)
		}
	}
	fmt.Println(eval(grid))
}

func eval(grid [1000][1000]int) int {
	sum := 0
	for _, row := range grid {
		for _, cell := range row {
			sum += cell
		}
	}
	return sum
}

// 0 0 - 1 0
// 0 1 - 1 1
func getInstructions(line string) [][]int {
	re := regexp.MustCompile("[0-9]+")
	flat := re.FindAllString(line, -1)

	num1, _ := strconv.Atoi(flat[0])
	num2, _ := strconv.Atoi(flat[1])
	num3, _ := strconv.Atoi(flat[2])
	num4, _ := strconv.Atoi(flat[3])
	return [][]int{{num1, num2}, {num3, num4}}
}

func toggles(grid *[1000][1000]int, instructions [][]int, instructType InstructionSet) {
	for i := instructions[0][0]; i <= instructions[1][0]; i++ {
		for j := instructions[0][1]; j <= instructions[1][1]; j++ {
			switch instructType {
			case Toggle:
				grid[i][j] += 2
			case TurnOff:
				if grid[i][j] > 0 {
					grid[i][j]--
				}
			case TurnOn:
				grid[i][j]++
			}
		}
	}

}
