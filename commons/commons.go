package commons

import (
	"os"
	"strings"
)

func ReadFileString(filename string) string {
	file, err := os.ReadFile(filename)
	if err != nil {
		panic(err)
	}
	return string(file)
}

func ReadLines(filename string) []string {
	file, err := os.ReadFile(filename)
	if err != nil {
		panic(err)
	}
	return strings.Split(string(file), "\n")
}
