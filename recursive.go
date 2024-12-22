package main

import (
	"context"
	"math/big"
)

func FibonacciRecursive(n int, ctx context.Context) *big.Int {
	result := big.NewInt(0)
	done := make(chan struct{})

	go func() {
		defer close(done)
		if n <= 1 {
			result.SetInt64(int64(n))
			return
		}
		result.Add(FibonacciRecursive(n-1, ctx), FibonacciRecursive(n-2, ctx))
	}()

	select {
	case <-ctx.Done():
		return big.NewInt(0)
	case <-done:
		return result
	}
}
