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

func TestRecursiveResult(test *testing.T) {
	unitTests := []struct {
		name     string
		function func(int, context.Context) *big.Int
		number   int
		expected string
	}{
		{"Recursive", algorithms.FibonacciRecursive, 0, "0"},
		{"Recursive", algorithms.FibonacciRecursive, 1, "1"},
		{"Recursive", algorithms.FibonacciRecursive, 2, "1"},
		{"Recursive", algorithms.FibonacciRecursive, 3, "2"},
		{"Recursive", algorithms.FibonacciRecursive, 4, "3"},
		{"Recursive", algorithms.FibonacciRecursive, 5, "5"},
		{"Recursive", algorithms.FibonacciRecursive, 6, "8"},
		{"Recursive", algorithms.FibonacciRecursive, 7, "13"},
		{"Recursive", algorithms.FibonacciRecursive, 8, "21"},
		{"Recursive", algorithms.FibonacciRecursive, 9, "34"},
		{"Recursive", algorithms.FibonacciRecursive, 10, "55"},
		{"Recursive", algorithms.FibonacciRecursive, 20, "6765"},
	}

	for _, unitTest := range unitTests {
		test.Run(unitTest.name, func(test *testing.T) {
			_, result := TimeNumber(unitTest.function, unitTest.number)
			if result[1] != unitTest.expected {
				test.Errorf("Expected %s, got %s", unitTest.expected, result[1])
			}
		})
	}
}
