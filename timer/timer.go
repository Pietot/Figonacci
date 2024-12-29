package timer

import (
	"context"
	"fmt"
	"math/big"
	"time"
)

func Timer(f func(int, context.Context) *big.Int, limit ...interface{}) string {
	var duration time.Duration

	if limit_seconds, err := checkLimit(limit...); err != nil {
		return fmt.Sprintf("Error: %s", err)
	} else {
		duration = limit_seconds
	}

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
	computeTimeElapsed := time.Since(computeTimeStart).Seconds()

	sentence := fmt.Sprintf(
		"\nThe biggest Fibonacci number that has been computed in less than \033[35m%.3f\033[0m second is the \033[32m%dnth\033[0m Fibonacci number\n\n"+
			"It's value is :\n\033[32m%s\033[0m\n\n"+
			"It has \033[32m%d\033[0m digits.\n\n"+
			"It has been found in \033[32m%.3f\033[0m seconds",
		duration.Seconds(), highNumber, fibonacciNumberString, len(fibonacciNumberString), computeTimeElapsed,
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
		"\nThe \033[35m%dnth\033[0m Fibonacci number is :\n\033[32m%s\033[0m\n\n"+
			"It has \033[32m%d\033[0m digits.\n\n"+
			"It has been found in \033[32m%s\033[0m",
		number, fibonacciNumberString, len(fibonacciNumberString), computeTimeElapsed,
	)

	return sentence
}

func checkLimit(limit ...interface{}) (time.Duration, error) {
	if len(limit) == 1 {
		switch _time := limit[0].(type) {
		case float64:
			return time.Duration(_time * float64(time.Second)), nil
		default:
			return 0, fmt.Errorf("argument must be a float64")
		}
	} else if len(limit) > 1 {
		return 0, fmt.Errorf("too many arguments")
	} else {
		return time.Second, nil
	}
}
