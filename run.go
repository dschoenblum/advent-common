package common

import (
	"fmt"
	"strconv"
	"time"
)

type SolverFunc func(input string) (string, error)

type Day struct {
	SolverA SolverFunc
	SolverB SolverFunc
	Input   string
}

func IntResult(result int) (string, error) {
	return strconv.Itoa(result), nil
}

func StringResult(result string) (string, error) {
	return result, nil
}

func ErrorResult(result error) (string, error) {
	return "", result
}

func Run(days []Day) {
	for i, day := range days {
		num := i + 1
		if num == len(days) {
			handleSolver(num, "A", day.SolverA, day.Input)
			handleSolver(num, "B", day.SolverB, day.Input)
		}
	}
}

func (s SolverFunc) runFunc(input string) (string, time.Duration, error) {
	start := time.Now()
	answer, err := s(input)
	duration := time.Since(start)
	return answer, duration, err
}

func handleSolver(dayNum int, id string, solver SolverFunc, input string) {
	answer, duration, err := solver.runFunc(input)
	if err != nil {
		fmt.Printf("*%d%s*: Failed: %s [%s]\n", dayNum, id, err, duration)
	} else {
		fmt.Printf("*%d%s*: %s [%s]\n", dayNum, id, answer, duration)
	}
}
