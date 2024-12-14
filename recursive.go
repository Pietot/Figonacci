package main

import (
	"context"
	"math/big"
)

func FibonacciRecursive(n int, ctx context.Context) *big.Int {
	select {
	case <-ctx.Done():
		return big.NewInt(0)
	default:
		if n <= 1 {
			return big.NewInt(int64(n))
		}
		a := FibonacciRecursive(n-1, ctx)
		b := FibonacciRecursive(n-2, ctx)
		return new(big.Int).Add(a, b)
	}
}
