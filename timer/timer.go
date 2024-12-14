package timer

import (
	"context"
	"fmt"
	"math/big"
	"time"
)

func Timer(f func(int, context.Context) *big.Int) (int, string, string, string) {
	computeTimeStart := time.Now()
	lowNumber, highNumber := 0, 1

	// Initial test to find the range of numbers that takes less than 1 second to compute
	for {
		ctx, cancel := context.WithTimeout(context.Background(), time.Second)
		start := time.Now()
		f(highNumber, ctx)
		elapsed := time.Since(start)
		cancel() // Annule le contexte après utilisation
		if elapsed > time.Second {
			break
		}
		lowNumber = highNumber
		highNumber *= 2
	}

	// Binary search to find the biggest number that takes less than 1 second to compute
	var fibonacciNumber *big.Int
	for lowNumber <= highNumber {
		mid := (lowNumber + highNumber) / 2
		ctx, cancel := context.WithTimeout(context.Background(), time.Second)
		start := time.Now()
		fibonacciNumber = f(mid, ctx)
		elapsed := time.Since(start)
		cancel() // Annule le contexte après utilisation
		if elapsed < time.Second {
			lowNumber = mid + 1
		} else {
			highNumber = mid - 1
		}
	}

	fibonacciNumberString := fibonacciNumber.String()
	fibonacciNumberStringLen := fmt.Sprintf("%d digits", len(fibonacciNumberString))
	computeTimeElapsed := fmt.Sprintf("%.2fs", time.Since(computeTimeStart).Seconds())
	return highNumber, fibonacciNumberString, fibonacciNumberStringLen, computeTimeElapsed
}

func TimeNumber(f func(int) *big.Int, number int) (string, string, string) {
	fibonacciNumber := new(big.Int)
	start := time.Now()
	fibonacciNumber = f(number)
	elapsed := time.Since(start)
	fibonacciNumberString := fibonacciNumber.String()
	fibonacciNumberStringLen := fmt.Sprintf("%d digits", len(fibonacciNumberString))
	computeTimeElapsed := fmt.Sprintf("%.2fs", elapsed.Seconds())
	return fibonacciNumberString, fibonacciNumberStringLen, computeTimeElapsed
}
