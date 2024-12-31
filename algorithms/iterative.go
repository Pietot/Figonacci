package algorithms

import (
	"context"
	"math/big"
)

func FibonacciIterative(n int, ctx context.Context) *big.Int {
	result := big.NewInt(0)
	done := make(chan struct{})

	go func() {
		defer close(done)
		if n <= 1 {
			result.SetInt64(int64(n))
			return
		}
		a, b := big.NewInt(0), big.NewInt(1)
		for i := 2; i <= n; i++ {
			a, b = b, a.Add(a, b)
		}
		result.Set(b)
	}()

	select {
	case <-ctx.Done():
		return big.NewInt(0)
	case <-done:
		return result
	}
}