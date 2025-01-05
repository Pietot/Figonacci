package timer

import (
	"bufio"
	"context"
	"fmt"
	"math/big"
	"os"
	"strconv"
	"testing"

	"github.com/Pietot/Figonacci/algorithms"
)

func readFile(index int) (string, error) {
	filePath := "../fibonacci_numbers/" + strconv.Itoa(index) + ".txt"
	file, err := os.Open(filePath)
	if err != nil {
		return "", fmt.Errorf("error opening file: %w", err)
	}
	defer file.Close()

	reader := bufio.NewReader(file)

	line, err := reader.ReadString('\n')
	if err != nil && err.Error() != "EOF" {
		return "", fmt.Errorf("error reading file: %w", err)
	}

	return line, nil
}

func TestCompute(test *testing.T) {
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
			sentence, result := Compute(unitTest.function, unitTest.number)
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
		{"Recursive", algorithms.FibonacciRecursive, 30, "832040"},
	}

	for _, unitTest := range unitTests {
		test.Run(unitTest.name, func(test *testing.T) {
			_, result := Compute(unitTest.function, unitTest.number)
			if result[1] != unitTest.expected {
				test.Errorf("Expected %s, got %s", unitTest.expected, result[1])
			}
		})
	}
}

func TestMatrixResult(test *testing.T) {
	unitTests := []struct {
		name     string
		function func(int, context.Context) *big.Int
		number   int
	}{
		{"Matrix", algorithms.FibonacciMatrix, 0},
		{"Matrix", algorithms.FibonacciMatrix, 1},
		{"Matrix", algorithms.FibonacciMatrix, 10_000},
		{"Matrix", algorithms.FibonacciMatrix, 100_000},
	}

	for _, unitTest := range unitTests {
		test.Run(unitTest.name, func(test *testing.T) {
			_, result := Compute(unitTest.function, unitTest.number)
			expected, err := readFile(unitTest.number)
			if err != nil {
				test.Errorf("%v", err)
				return
			}
			if result[1] != expected {
				test.Errorf("Expected %s, got %s", expected, result[1])
			}
		})
	}
}

func TestRecursiveOptimizedResult(test *testing.T) {
	unitTests := []struct {
		name     string
		function func(int, context.Context) *big.Int
		number   int
	}{
		{"RecursiveOptimized", algorithms.FibonacciRecursiveOptimized, 0},
		{"RecursiveOptimized", algorithms.FibonacciRecursiveOptimized, 1},
		{"RecursiveOptimized", algorithms.FibonacciRecursiveOptimized, 10_000},
		{"RecursiveOptimized", algorithms.FibonacciRecursiveOptimized, 100_000},
	}

	for _, unitTest := range unitTests {
		test.Run(unitTest.name, func(test *testing.T) {
			_, result := Compute(unitTest.function, unitTest.number)
			expected, err := readFile(unitTest.number)
			if err != nil {
				test.Errorf("%v", err)
				return
			}
			if result[1] != expected {
				test.Errorf("Expected %s, got %s", expected, result[1])
			}
		})
	}
}

func TestIterativeResult(test *testing.T) {
	unitTests := []struct {
		name     string
		function func(int, context.Context) *big.Int
		number   int
	}{
		{"Iterative", algorithms.FibonacciIterative, 0},
		{"Iterative", algorithms.FibonacciIterative, 1},
		{"Iterative", algorithms.FibonacciIterative, 100_000},
		{"Iterative", algorithms.FibonacciIterative, 1_000_000},
	}

	for _, unitTest := range unitTests {
		test.Run(unitTest.name, func(test *testing.T) {
			_, result := Compute(unitTest.function, unitTest.number)
			expected, err := readFile(unitTest.number)
			if err != nil {
				test.Errorf("%v", err)
				return
			}
			if result[1] != expected {
				test.Errorf("Expected %s, got %s", expected, result[1])
			}
		})
	}
}

func TestMatrixOptimizedResult(test *testing.T) {
	unitTests := []struct {
		name     string
		function func(int, context.Context) *big.Int
		number   int
	}{
		{"MatrixOptimized", algorithms.FibonacciMatrixOptimized, 0},
		{"MatrixOptimized", algorithms.FibonacciMatrixOptimized, 1},
		{"MatrixOptimized", algorithms.FibonacciMatrixOptimized, 100_000},
		{"MatrixOptimized", algorithms.FibonacciMatrixOptimized, 1_000_000},
		{"MatrixOptimized", algorithms.FibonacciMatrixOptimized, 10_000_000},
	}

	for _, unitTest := range unitTests {
		test.Run(unitTest.name, func(test *testing.T) {
			_, result := Compute(unitTest.function, unitTest.number)
			expected, err := readFile(unitTest.number)
			if err != nil {
				test.Errorf("%v", err)
				return
			}
			if result[1] != expected {
				test.Errorf("Expected %s, got %s", expected, result[1])
			}
		})
	}
}

func TestFieldExtensionResult(test *testing.T) {
	unitTests := []struct {
		name     string
		function func(int, context.Context) *big.Int
		number   int
	}{
		{"FieldExtension", algorithms.FibonacciFieldExtension, 0},
		{"FieldExtension", algorithms.FibonacciFieldExtension, 1},
		{"FieldExtension", algorithms.FibonacciFieldExtension, 100_000},
		{"FieldExtension", algorithms.FibonacciFieldExtension, 1_000_000},
		{"FieldExtension", algorithms.FibonacciFieldExtension, 10_000_000},
	}

	for _, unitTest := range unitTests {
		test.Run(unitTest.name, func(test *testing.T) {
			_, result := Compute(unitTest.function, unitTest.number)
			expected, err := readFile(unitTest.number)
			if err != nil {
				test.Errorf("%v", err)
				return
			}
			if result[1] != expected {
				test.Errorf("Expected %s, got %s", expected, result[1])
			}
		})
	}
}
