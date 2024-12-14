package timer

import (
	"fmt"
	"math/big"
	"time"
)

func Timer(f func(int) *big.Int) (int, string, string, string) {
	computeTimeStart := time.Now()
	number := 0
	fibonacciNumber := new(big.Int)
	for {
		start := time.Now()
		fibonacciNumber = f(number)
		elapsed := time.Since(start)
		if elapsed > time.Second {
			break
		}
		number++
	}
	computeTimeElapsed := fmt.Sprintf("%.2fs", time.Since(computeTimeStart).Seconds())
    fibonacciNumberString := fibonacciNumber.String()
    fibonacciNumberStringLen := fmt.Sprintf("%d digits", len(fibonacciNumberString))
	return number, fibonacciNumberString, fibonacciNumberStringLen, computeTimeElapsed
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
