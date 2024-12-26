package timer

import (
	"context"
	"fmt"
	"math/big"
	"time"
)

func Timer(f func(int, context.Context) *big.Int) (string) {
	fibonacciNumber := new(big.Int)
	lowNumber, highNumber := 0, 1
	computeTimeStart := time.Now()

	// Initial test to find the range of numbers that takes less than 1 second to compute
	for {
		ctx, cancel := context.WithTimeout(context.Background(), time.Second)
		start := time.Now()
		f(highNumber, ctx)
		elapsed := time.Since(start)
		cancel()
		if elapsed > time.Second {
			break
		}
		lowNumber = highNumber
		highNumber *= 2
	}

	// Binary search to find the biggest number that takes less than 1 second to compute
	for lowNumber <= highNumber {
		mid := (lowNumber + highNumber) / 2
		ctx, cancel := context.WithTimeout(context.Background(), time.Second)
		start := time.Now()
		fibonacciNumber = f(mid, ctx)
		elapsed := time.Since(start)
		cancel()
		if elapsed < time.Second {
			lowNumber = mid + 1
		} else {
			highNumber = mid - 1
		}
	}

	fibonacciNumberString := fibonacciNumber.String()
	computeTimeElapsed := time.Since(computeTimeStart).Seconds()

	sentence := fmt.Sprintf(
		"\nThe biggest Fibonacci number that has been computed in less than 1 second is the \033[32m%dnth\033[0m Fibonacci number\n\n"+
			"It's value is :\n\033[32m%s\033[0m\n\n"+
			"It has \033[32m%d\033[0m digits.\n\n"+
			"It has been found in \033[32m%.3f\033[0m seconds",
		highNumber, fibonacciNumberString, len(fibonacciNumberString), computeTimeElapsed,
	)

	return sentence
}

func TimeNumber(f func(int, context.Context) *big.Int, number int) string {
	ctx := context.Background()
	fibonacciNumber := new(big.Int)
	start := time.Now()
	fibonacciNumber = f(number, ctx)
	elapsed := time.Since(start)

	fibonacciNumberString := fibonacciNumber.String()
	computeTimeElapsed := fmt.Sprintf("%.2fs", elapsed.Seconds())
	
	sentence := fmt.Sprintf(
		"\nThe \033[33m%d\033[0mnth Fibonacci number is :\n\033[33m%s\033[0m\n\n"+
			"It has \033[33m%d\033[0m digits.\n\n"+
			"It has been found in \033[33m%s\033[0m",
		number, fibonacciNumberString, len(fibonacciNumberString), computeTimeElapsed,
	)
	
	return sentence
}
