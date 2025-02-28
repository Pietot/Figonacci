// Package timer provides functions to compute the Fibonacci number of a given index and to determine the largest Fibonacci index that can be computed within a given time limit.
package timer

import (
	"context"
	"fmt"
	"math/big"
	"time"
)

// Timer determines the largest Fibonacci index that can be computed within a given time limit.
//
// This function uses an exponential search followed by a binary search to efficiently find the largest
// Fibonacci number index for which the computation completes within the specified time limit.
//
// The function returns a string containing the result and a slice of interfaces containing the following values:
//
// - The index of the largest Fibonacci number that has been computed in less than the specified time limit
//
// - The value of the number at this index
//
// - The number of digits of the value
//
// - The time taken to search for the largest Fibonacci number
func Timer(f func(int, context.Context) *big.Int, limit_second float64) (string, []any) {
	duration := time.Duration(limit_second * float64(time.Second))
	fibonacciNumber := new(big.Int)
	fibonacciTemp := new(big.Int)
	zero := big.NewInt(0)
	lowNumber, highNumber := 0, 1
	computeTimeStart := time.Now()

	// Initial test to find the range of numbers that takes less than 1 second to compute
	for {
		ctx, cancel := context.WithTimeout(context.Background(), duration)
		if f(highNumber, ctx).Cmp(zero) == 0 {
			cancel()
			break
		}
		cancel()
		lowNumber = highNumber
		highNumber *= 2
	}

	// Binary search to find the biggest number that takes less than 1 second to compute
	for lowNumber <= highNumber {
		mid := (lowNumber + highNumber) / 2
		ctx, cancel := context.WithTimeout(context.Background(), duration)
		fibonacciTemp = f(mid, ctx)
		if fibonacciTemp.Cmp(zero) == 0 {
			cancel()
			highNumber = mid - 1
		} else {
			cancel()
			lowNumber = mid + 1
			fibonacciNumber.Set(fibonacciTemp)
		}
	}

	computeTimeElapsed := time.Since(computeTimeStart)
	computeTimeFormated := FormatDuration(computeTimeElapsed, 3)
	fibonacciNumberString := fibonacciNumber.String()

	sentence := fmt.Sprintf(
		"\nThe biggest Fibonacci index that has been computed in less than \033[35m%s\033[0m second is \033[32m%d\033[0m\n\n"+
			"It's value is :\n\033[32m%s\033[0m\n\n"+
			"It has \033[32m%d\033[0m digits.\n\n"+
			"It has been found in \033[32m%s\033[0m",
		FormatDuration(duration, 0), highNumber, fibonacciNumberString, len(fibonacciNumberString), computeTimeFormated,
	)

	return sentence, []any{highNumber, fibonacciNumberString, len(fibonacciNumberString), computeTimeElapsed}
}

// Compute the Fibonacci number of a given index
// The function returns a string containing the result and a slice of interfaces containing the following values:
//
// - The index provided in parameter
//
// - The value of the number at this index
//
// - The number of digits of the value
//
// - The time taken to compute the index
func Compute(f func(int, context.Context) *big.Int, index int) (string, []any) {
	ctx := context.Background()
	fibonacciNumber := new(big.Int)
	computeTimeStart := time.Now()

	fibonacciNumber = f(index, ctx)

	computeTimeElapsed := time.Since(computeTimeStart)
	computeTimeFormated := FormatDuration(computeTimeElapsed, 3)
	fibonacciNumberString := fibonacciNumber.String()

	sentence := fmt.Sprintf(

		"\nFibonacci of \033[35m%d\033[0m is equal to :\n"+
			"\033[32m%s\033[0m\n\n"+
			"It has \033[32m%d\033[0m digits.\n\n"+
			"It has been found in \033[32m%s\033[0m",
		index, fibonacciNumberString, len(fibonacciNumberString), computeTimeFormated,
	)

	return sentence, []any{index, fibonacciNumberString, len(fibonacciNumberString), computeTimeElapsed}
}

func FormatDuration(duration time.Duration, precision int) string {
	switch {
	case duration < time.Nanosecond:
		return fmt.Sprintf("%v", duration)
	case duration < time.Microsecond:
		return fmt.Sprintf("%.*fns", precision, float64(duration)/float64(time.Nanosecond))
	case duration < time.Millisecond:
		return fmt.Sprintf("%.*fµs", precision, float64(duration)/float64(time.Microsecond))
	case duration < time.Second:
		return fmt.Sprintf("%.*fms", precision, float64(duration)/float64(time.Millisecond))
	case duration < time.Minute:
		return fmt.Sprintf("%.*fs", precision, float64(duration)/float64(time.Second))
	case duration < time.Hour:
		return fmt.Sprintf("%.*fmin", precision, float64(duration)/float64(time.Minute))
	default:
		return fmt.Sprintf("%.*fh", precision, float64(duration)/float64(time.Hour))
	}
}
