package main

import (
	"context"
	"math/big"
)

func FibonacciIterative(n int, ctx context.Context) *big.Int {
	select {
	case <-ctx.Done():
		return big.NewInt(0)
	default:
		if n <= 1 {
			return big.NewInt(int64(n))
		}

		a, b := big.NewInt(0), big.NewInt(1)
		for i := 2; i <= n; i++ {
			a, b = b, a.Add(a, b)
		}

		return b
	}
}
