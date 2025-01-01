package timer

import (
	"Figonacci/algorithms"
	"context"
	"math/big"
	"testing"
)

func TestTimeNumber(test *testing.T) {
	unitTests := []struct {
		name     string
		function func(int, context.Context) *big.Int
		number   int
	}{
		{"Recursive", algorithms.FibonacciRecursive, 10},
		{"RecursiveOptimized", algorithms.FibonacciRecursiveOptimized, 10},
		{"Iterative", algorithms.FibonacciIterative, 10},
		{"Matrix", algorithms.FibonacciMatrix, 10},
		{"MatrixOptimized", algorithms.FibonacciMatrixOptimized, 10},
		{"FieldExtension", algorithms.FibonacciFieldExtension, 10},
	}

	for _, unitTest := range unitTests {
		test.Run(unitTest.name, func(test *testing.T) {
			sentence, result := TimeNumber(unitTest.function, unitTest.number)
			if len(result) != 4 {
				test.Errorf("Expected 4 results, got %d", len(result))
			}
			if sentence == "" {
				test.Errorf("Expected non-empty sentence")
			}
		})
	}
}
