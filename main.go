package main

import (
	"Figonacci/timer"
	"Figonacci/algorithms"
	"context"
	"fmt"
	"math/big"
	"os"
	"strconv"
)

func main() {
	if len(os.Args) < 2 {
		printUsage()
		return
	}

	command := os.Args[1]

	switch command {
	case "timer":
		err := handleTimer(os.Args[2:])
		if err != nil {
			fmt.Println("\033[31mError:", err, "\033[0m")
			printUsage()
		}

	case "compute":
		err := handleCompute(os.Args[2:])
		if err != nil {
			fmt.Println("\033[31mError:", err, "\033[0m")
			printUsage()
		}

	default:
		fmt.Println("\033[31mUnknown command:", command, "\033[0m")
		printUsage()
	}
}

func printUsage() {
	fmt.Println("")
	fmt.Println("    Usage:")
	fmt.Println("")
	fmt.Println("        timer   --algorithm [--limit] (optional, in second, default to 1 second)")
	fmt.Println("        \033[35mTo find the largest Fibonacci index calculable in less than a second\033[0m")
	fmt.Println("")
	fmt.Println("        compute --algorithm --value")
	fmt.Println("        \033[35mTo compute the Fibonacci number at the specified index\033[0m")
	fmt.Println("")
	fmt.Println("    Available algorithms (from \033[31mslowest\033[0m to \033[34mfastest\033[0m):")
	fmt.Println("")
	fmt.Println("        1 - \033[31mrecursive (r)\033[0m")
	fmt.Println("        2 - \033[33mrecursive_optimized (ro)\033[0m")
	fmt.Println("        3 - \033[33miterative (i)\033[0m")
	fmt.Println("        4 - \033[32mmatrix (m)\033[0m")
	fmt.Println("        5 - \033[36mmatrix_optimized (mo)\033[0m")
	fmt.Println("        6 - \033[34mfield_extension (fe)\033[0m")
	fmt.Println("")
}

func handleTimer(args []string) error {
	if len(args) > 2 || args[0][:2] != "--" || (len(args) == 2 && args[1][:2] != "--") {
		return fmt.Errorf("invalid syntax for timer. Expected: --algorithm [--limit] (optional, in second, default to 1 second)")
	}

	algorithm := args[0][2:]

	algoFunc, algoName := isValidAlgorithm(algorithm)
	if algoFunc == nil {
		return fmt.Errorf("invalid algorithm: %s", algorithm)
	}

	var limit string

	if len(args) == 2 {
		limit = args[1][2:]
	} else {
		limit = "1"
	}

	limitFloat, err := strconv.ParseFloat(limit, 64)
	if err != nil {
		return fmt.Errorf("invalid limit: %s. Expected an integer or float", limit)
	}

	fmt.Printf("Executing timer with algorithm \033[35m%s\033[0m...\n", algoName)

	sentence, _ := timer.Timer(algoFunc, limitFloat)
	fmt.Println(sentence)
	return nil
}

func handleCompute(args []string) error {
	if len(args) != 2 || len(args[0]) < 3 || args[0][:2] != "--" || len(args[1]) < 3 || args[1][:2] != "--" {
		return fmt.Errorf("invalid syntax for compute. Expected: --algorithm --value")
	}

	algorithm := args[0][2:]
	valueStr := args[1][2:]

	algoFunc, algoName := isValidAlgorithm(algorithm)
	if algoFunc == nil {
		return fmt.Errorf("invalid algorithm: %s", algorithm)
	}

	value, err := strconv.Atoi(valueStr)
	if err != nil {
		return fmt.Errorf("invalid value: %s. Expected an integer", valueStr)
	} else if value < 0 {
		return fmt.Errorf("invalid value: %d. Expected a positive integer greater than 0", value)
	}

	fmt.Printf("Executing compute with algorithm \033[35m%s\033[0m...\n", algoName)

	sentence, _ := timer.TimeNumber(algoFunc, value)
	fmt.Println(sentence)
	return nil
}

func isValidAlgorithm(algorithm string) (func(int, context.Context) *big.Int, string) {
	validAlgorithms := map[string]struct {
		function func(int, context.Context) *big.Int
		name     string
	}{
		"recursive":           {algorithms.FibonacciRecursive, "recursive"},
		"r":                   {algorithms.FibonacciRecursive, "recursive"},
		"recursive_optimized": {algorithms.FibonacciRecursiveOptimized, "recursive optimized"},
		"ro":                  {algorithms.FibonacciRecursiveOptimized, "recursive optimized"},
		"iterative":           {algorithms.FibonacciIterative, "iterative"},
		"i":                   {algorithms.FibonacciIterative, "iterative"},
		"matrix":              {algorithms.FibonacciMatrix, "matrix"},
		"m":                   {algorithms.FibonacciMatrix, "matrix"},
		"matrix_optimized":    {algorithms.FibonacciMatrixOptimized, "matrix optimized"},
		"mo":                  {algorithms.FibonacciMatrixOptimized, "matrix optimized"},
		"field_extension":     {algorithms.FibonacciFieldExtension, "field extension"},
		"fe":                  {algorithms.FibonacciFieldExtension, "field extension"},
	}

	if entry, exists := validAlgorithms[algorithm]; exists {
		return entry.function, entry.name
	}
	return nil, ""
}
