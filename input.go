package common

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func FromFile(filename string) string {
	input, err := os.ReadFile(filename)
	if err != nil {
		panic("Unable to read file")
	}
	return string(input)
}

func FromFileTrimmed(filename string) string {
	input, err := os.ReadFile(filename)
	if err != nil {
		panic("Unable to read file")
	}
	return strings.TrimSpace(string(input))
}

func ToLines(input string) []string {
	return strings.Split(input, "\n")
}

func ToTrimmedLines(input string) []string {
	lines := strings.Split(input, "\n")
	trimmed := make([]string, len(lines))
	for i, line := range lines {
		trimmed[i] = strings.TrimSpace(line)
	}
	return trimmed
}

func ToInts(input string) []int {
	return SplitToInts(input, "\n")
}

func CsvToInts(input string) []int {
	return SplitToInts(input, ",")
}

func SplitToInts(input string, delimiter string) []int {
	lines := strings.Split(input, delimiter)
	ints := make([]int, len(lines))
	for i, line := range lines {
		ints[i], _ = strconv.Atoi(strings.TrimSpace(line))
	}
	return ints
}

func Atoi(input string) int {
	num, err := strconv.Atoi(input)
	if err != nil {
		panic(fmt.Sprintf("Failed to convert '%s' to int", input))
	}
	return num
}
