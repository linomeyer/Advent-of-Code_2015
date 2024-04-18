package main

import (
	"fmt"
	"os"
)

func main() {
	file, err := os.ReadFile("./day01/input.txt")
	fileStr := string(file)
	floor := 0
	basementEntry := -1
	if err == nil {
		floor = calculateFloor(fileStr)
		basementEntry = calculateBasementEntry(fileStr)
	}
	fmt.Println(floor)
	fmt.Println(basementEntry)
}

func calculateBasementEntry(file string) int {
	floor := 0
	for i, c := range file {
		if c == '(' {
			floor++
		}
		if c == ')' {
			floor--
		}
		if floor < 0 {
			return i + 1
		}
	}
	return -1
}

func calculateFloor(file string) int {
	floor := 0
	for _, c := range file {
		if string(c) == "(" {
			floor++
		}
		if string(c) == ")" {
			floor--
		}
	}
	return floor
}
