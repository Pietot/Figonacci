package algorithms

import (
	"context"
	"math/big"
	"sync"
)

var lookupTable = map[int]*big.Int{
	0: big.NewInt(0),
	1: big.NewInt(1),
}
var mu sync.Mutex

func FibonacciRecursiveOptimized(n int, ctx context.Context) *big.Int {
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
			mu.Lock()
			if val, exist := lookupTable[n]; exist {
				result.Set(val)
				mu.Unlock()
				return result
			}

			mu.Unlock()

			a := recursive(n-1, ctx)
			b := recursive(n-2, ctx)

			mu.Lock()
			lookupTable[n] = new(big.Int).Add(a, b)
			result.Set(lookupTable[n])
			mu.Unlock()
			return result
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
