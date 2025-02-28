// Package timer provides functions to compute the Fibonacci number of a given index and to determine the largest Fibonacci index that can be computed within a given time limit.
package timer

import (
	"context"
	"math/big"
	"testing"

	"github.com/Pietot/Figonacci/v2/algorithms"
	"github.com/Pietot/Figonacci/v2/timer"
)

func TestTimer(test *testing.T) {
	algos := []struct {
		name     string
		function func(int, context.Context) *big.Int
		limit    float64
	}{
		{"Recursive", algorithms.FibonacciRecursive, 1},
		{"RecursiveOptimized", algorithms.FibonacciRecursiveOptimized, 1},
		{"Iterative", algorithms.FibonacciIterative, 1},
		{"Matrix", algorithms.FibonacciMatrix, 1},
		{"MatrixOptimized", algorithms.FibonacciMatrixOptimized, 1},
		{"FieldExtension", algorithms.FieldExtension, 1},
		{"Pihedron", algorithms.Pihedron, 1},
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
