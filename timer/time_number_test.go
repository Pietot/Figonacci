package timer

import (
	"Figonacci/algorithms"
	"bufio"
	"context"
	"fmt"
	"math/big"
	"os"
	"strconv"
	"testing"
)

func readFile(index int) (string, error) {
	filePath := "../fibonacci_numbers/" + strconv.Itoa(index) + ".txt"
	file, err := os.Open(filePath)
	if err != nil {
		return "", fmt.Errorf("error opening file: %w", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	if scanner.Scan() {
		return scanner.Text(), nil
	} else {
		return "", fmt.Errorf("error reading file: %w", scanner.Err())
	}

}

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

func TestRecursiveOptimizedResult(test *testing.T) {
	unitTests := []struct {
		name     string
		function func(int, context.Context) *big.Int
		number   int
	}{
		{"RecursiveOptimized", algorithms.FibonacciRecursiveOptimized, 0},
		{"RecursiveOptimized", algorithms.FibonacciRecursiveOptimized, 1},
		{"RecursiveOptimized", algorithms.FibonacciRecursiveOptimized, 100000},
		{"RecursiveOptimized", algorithms.FibonacciRecursiveOptimized, 200000},
		{"RecursiveOptimized", algorithms.FibonacciRecursiveOptimized, 300000},
		{"RecursiveOptimized", algorithms.FibonacciRecursiveOptimized, 400000},
		{"RecursiveOptimized", algorithms.FibonacciRecursiveOptimized, 500000},
		{"RecursiveOptimized", algorithms.FibonacciRecursiveOptimized, 600000},
		{"RecursiveOptimized", algorithms.FibonacciRecursiveOptimized, 700000},
		{"RecursiveOptimized", algorithms.FibonacciRecursiveOptimized, 800000},
		{"RecursiveOptimized", algorithms.FibonacciRecursiveOptimized, 900000},
		{"RecursiveOptimized", algorithms.FibonacciRecursiveOptimized, 1000000},
	}

	for _, unitTest := range unitTests {
		test.Run(unitTest.name, func(test *testing.T) {
			_, result := TimeNumber(unitTest.function, unitTest.number)
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
		{"Iterative", algorithms.FibonacciIterative, 100000},
		{"Iterative", algorithms.FibonacciIterative, 200000},
		{"Iterative", algorithms.FibonacciIterative, 300000},
		{"Iterative", algorithms.FibonacciIterative, 400000},
		{"Iterative", algorithms.FibonacciIterative, 500000},
		{"Iterative", algorithms.FibonacciIterative, 600000},
		{"Iterative", algorithms.FibonacciIterative, 700000},
		{"Iterative", algorithms.FibonacciIterative, 800000},
		{"Iterative", algorithms.FibonacciIterative, 900000},
		{"Iterative", algorithms.FibonacciIterative, 1000000},
	}

	for _, unitTest := range unitTests {
		test.Run(unitTest.name, func(test *testing.T) {
			_, result := TimeNumber(unitTest.function, unitTest.number)
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