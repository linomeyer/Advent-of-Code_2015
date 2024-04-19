package main

import (
	"AdventOfCode_2015/commons"
	"fmt"
)

func main() {
	lines := commons.ReadLines("./day05/input.txt")

	perfectStringCounter := part2(lines)
	fmt.Println(perfectStringCounter)
}

func part1(lines []string) int {
	perfectStringCounter := 0
	for _, line := range lines {
		vowelCounter := 0
		hasDoubleLetter := false
		hasForbiddenStrings := false
		var lastChar int32
		for _, c := range line {
			vowelCounter += isVowel(c)
			hasDoubleLetter, hasForbiddenStrings = checkWithLastChar(lastChar, c)
			lastChar = c
		}
		if hasDoubleLetter && !hasForbiddenStrings && vowelCounter >= 3 {
			perfectStringCounter++
		}
	}
	return perfectStringCounter
}

func part2(lines []string) int {
	perfectStringCounter := 0
	for _, line := range lines {
		pairs := getPairs(line)
		hasEqualPairs := hasDupes(pairs)
		hasRepeats := checkRepeats(line)
		if hasEqualPairs && hasRepeats {
			perfectStringCounter++
		}
	}
	return perfectStringCounter
}

func checkRepeats(line string) bool {
	for i := 1; i < len(line)-1; i++ {
		prevChar := line[i-1]
		curChar := line[i]
		nextChar := line[i+1]
		if prevChar == nextChar && prevChar != curChar {
			return true
		}
	}
	return false
}

func hasDupes(pairs []string) bool {
	distinct := make(map[string]bool, len(pairs))
	for _, pair := range pairs {
		if distinct[pair] {
			return true
		} else {
			distinct[pair] = true
		}
	}
	return false
}

func getPairs(line string) []string {
	var pairs []string
	for i := 1; i < len(line); i++ {
		prevChar := line[i-1]
		curChar := line[i]

		pairs = append(pairs, string(prevChar)+string(curChar))
	}
	return pairs
}

func checkWithLastChar(lastChar int32, c int32) (bool, bool) {
	hasDoubleLetter := false
	hasForbiddenStrings := false
	if lastChar == c {
		hasDoubleLetter = true
	} else {
		lastCharStr := string(lastChar) + string(c)
		switch lastCharStr {
		case "ab", "cd", "pq", "xy":
			hasForbiddenStrings = true
		}
	}
	return hasDoubleLetter, hasForbiddenStrings
}

func isVowel(c int32) int {
	switch c {
	case 'a', 'e', 'i', 'o', 'u':
		return 1
	}
	return 0
}
