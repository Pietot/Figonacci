package timer

import (
	"context"
	"fmt"
	"math/big"
	"time"
)

func Timer(f func(int, context.Context) *big.Int, limit float64) (string, []interface{}) {
	duration := time.Duration(limit * float64(time.Second))
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

	fibonacciNumberString := fibonacciNumber.String()
	computeTimeElapsed := time.Since(computeTimeStart)
	computeTimeFormated := formatDuration(computeTimeElapsed)

	sentence := fmt.Sprintf(
		"\nThe biggest Fibonacci number that has been computed in less than \033[35m%.3f\033[0m second is the \033[32m%dnth\033[0m Fibonacci number\n\n"+
			"It's value is :\n\033[32m%s\033[0m\n\n"+
			"It has \033[32m%d\033[0m digits.\n\n"+
			"It has been found in \033[32m%s\033[0m",
		duration.Seconds(), highNumber, fibonacciNumberString, len(fibonacciNumberString), computeTimeFormated,
	)

	return sentence, []interface{}{highNumber, fibonacciNumberString, len(fibonacciNumberString), computeTimeElapsed}
}

func TimeNumber(f func(int, context.Context) *big.Int, number int) (string, []interface{}) {
	ctx := context.Background()
	fibonacciNumber := new(big.Int)
	computeTimeStart := time.Now()

	fibonacciNumber = f(number, ctx)

	fibonacciNumberString := fibonacciNumber.String()
	computeTimeElapsed := time.Since(computeTimeStart)
	computeTimeFormated := formatDuration(computeTimeElapsed)

	sentence := fmt.Sprintf(
		"\nThe \033[35m%dnth\033[0m Fibonacci number is :\n\033[32m%s\033[0m\n\n"+
			"It has \033[32m%d\033[0m digits.\n\n"+
			"It has been found in \033[32m%s\033[0m",
		number, fibonacciNumberString, len(fibonacciNumberString), computeTimeFormated,
	)

	return sentence, []interface{}{number, fibonacciNumberString, len(fibonacciNumberString), computeTimeElapsed}
}

func formatDuration(duration time.Duration) string {
	switch {
	case duration < time.Nanosecond:
		return fmt.Sprintf("%v", duration)
	case duration < time.Microsecond:
		return fmt.Sprintf("%.3fns", float64(duration)/float64(time.Nanosecond))
	case duration < time.Millisecond:
		return fmt.Sprintf("%.3fµs", float64(duration)/float64(time.Microsecond))
	case duration < time.Second:
		return fmt.Sprintf("%.3fms", float64(duration)/float64(time.Millisecond))
	case duration < time.Minute:
		return fmt.Sprintf("%.3fs", float64(duration)/float64(time.Second))
	case duration < time.Hour:
		return fmt.Sprintf("%.3fmin", float64(duration)/float64(time.Minute))
	default:
		return fmt.Sprintf("%.3fh", float64(duration)/float64(time.Hour))
	}
}