// Package timer provides functions to compute the Fibonacci number of a given index and to determine the largest Fibonacci index that can be computed within a given time limit.
package timer

import (
	"testing"

	"github.com/Pietot/Figonacci/v2/algorithms"
	"github.com/Pietot/Figonacci/v2/timer"
)

func TestTimer(test *testing.T) {
	algos := []struct {
		name     string
		function any
		limit    float64
	}{
		{"Recursive", algorithms.FibonacciRecursive, 0.1},
		{"RecursiveOptimized", algorithms.FibonacciRecursiveOptimized, 0.1},
		{"Iterative", algorithms.FibonacciIterative, 0.1},
		{"Matrix", algorithms.FibonacciMatrix, 0.1},
		{"MatrixOptimized", algorithms.FibonacciMatrixOptimized, 0.1},
		{"FieldExtension", algorithms.FieldExtension, 0.1},
		{"Pihedron", algorithms.Pihedron, 0.1},
	}

	for _, unitTest := range algos {
		test.Run(unitTest.name, func(test *testing.T) {
			sentence, result := timer.Timer(unitTest.function, unitTest.limit)
			if len(result) != 4 {
				test.Errorf("Expected 4 results, got %d", len(result))
			}
			if sentence == "" {
				test.Errorf("Expected non-empty sentence")
			}
		})
	}
}

func TestCompute(test *testing.T) {
	unitTests := []struct {
		name     string
		function any
		number   int
	}{
		{"Recursive", algorithms.FibonacciRecursive, 10},
		{"RecursiveOptimized", algorithms.FibonacciRecursiveOptimized, 10},
		{"Iterative", algorithms.FibonacciIterative, 10},
		{"Matrix", algorithms.FibonacciMatrix, 10},
		{"MatrixOptimized", algorithms.FibonacciMatrixOptimized, 10},
		{"FieldExtension", algorithms.FieldExtension, 10},
	}

	for _, unitTest := range unitTests {
		test.Run(unitTest.name, func(test *testing.T) {
			sentence, result := timer.Compute(unitTest.function, unitTest.number)
			if len(result) != 4 {
				test.Errorf("Expected 4 results, got %d", len(result))
			}
			if sentence == "" {
				test.Errorf("Expected non-empty sentence")
			}
		})
	}
}
