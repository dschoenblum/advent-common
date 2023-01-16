package main

import "testing"

type TestCases map[string]string

func RunTests(t *testing.T, name string, solver SolverFunc, tests TestCases) {
	for input, want := range tests {
		t.Run(input, func(t *testing.T) {
			got, err := solver(input)
			if err != nil {
				t.Errorf("%s() error = %v", name, err)
			} else if got != want {
				t.Errorf("%s() = %v, want %v", name, got, want)
			}
		})
	}
}
