package common

import (
	"fmt"
	"strconv"
	"strings"
)

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
	return splitToInts(input, "\n")
}

func CsvToInts(input string) []int {
	return splitToInts(input, ",")
}

func splitToInts(input string, delimiter string) []int {
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
