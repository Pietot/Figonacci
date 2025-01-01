package timer

import (
	"Figonacci/algorithms"
	"context"
	"math/big"
	"testing"
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
		{"FieldExtension", algorithms.FibonacciFieldExtension, 1},
	}

	for _, unitTest := range algos {
		test.Run(unitTest.name, func(test *testing.T) {
			sentence, result := Timer(unitTest.function, unitTest.limit)
			if len(result) != 4 {
				test.Errorf("Expected 4 results, got %d", len(result))
			}
			if sentence == "" {
				test.Errorf("Expected non-empty sentence")
			}
		})
	}
}
