package algorithms

import (
	"context"
	"math/big"
)

func FibonacciRecursive(n int, ctx context.Context) *big.Int {
	result := big.NewInt(0)
	done := make(chan struct{})

	var recursive func(int, context.Context) *big.Int
	recursive = func(n int, ctx context.Context) *big.Int {
		if n <= 1 {
			return big.NewInt(int64(n))
		}
		select {
		case <-ctx.Done():
			return big.NewInt(0)
		default:
			a := recursive(n-1, ctx)
			b := recursive(n-2, ctx)
			return new(big.Int).Add(a, b)
		}
	}

	go func() {
		defer close(done)
		result.Set(recursive(n, ctx))
	}()

	select {
	case <-ctx.Done():
		return big.NewInt(0)
	case <-done:
		return result
	}
}
