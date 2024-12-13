package timer

import (
	"fmt"
	"time"
)

func Timer(f func(int) int) (int, int, string) {
	computeTimeStart := time.Now()
	number := 0
	fibonacciNumber := 0
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
	return number, fibonacciNumber, computeTimeElapsed
}

func TimeNumber(f func(int) int, number int) int {
	fibonacciNumber := 0
	start := time.Now()
	fibonacciNumber = f(number)
	elapsed := time.Since(start)
	fmt.Println(elapsed)
	return fibonacciNumber
}
