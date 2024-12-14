package timer

import (
	"fmt"
	"math/big"
	"time"
)

func Timer(f func(int) *big.Int) (int, string, string, string) {
    computeTimeStart := time.Now()
    low_number, high_number := 0, 1

    for {
        start := time.Now()
        f(high_number)
        elapsed := time.Since(start)
        if elapsed > time.Second {
            break
        }
        low_number = high_number
        high_number *= 2
    }

    var fibonacciNumber *big.Int
    for low_number <= high_number {
        mid := (low_number + high_number) / 2
        start := time.Now()
        fibonacciNumber = f(mid)
        elapsed := time.Since(start)
        if elapsed < time.Second {
            low_number = mid + 1
        } else {
            high_number = mid - 1
        }
    }

	fibonacciNumberString := fibonacciNumber.String()
    fibonacciNumberStringLen := fmt.Sprintf("%d digits", len(fibonacciNumberString))
	computeTimeElapsed := fmt.Sprintf("%.2fs", time.Since(computeTimeStart).Seconds())
	return high_number, fibonacciNumberString, fibonacciNumberStringLen, computeTimeElapsed
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
