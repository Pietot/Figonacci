package main

import (
	"context"
	"math/big"
)

var lookupTable = map[int]*big.Int{
	0: big.NewInt(0),
	1: big.NewInt(1),
}

func FibonacciRecursiveOptimized(n int, ctx context.Context) *big.Int {
	result := big.NewInt(0)
	done := make(chan struct{})

	go func() {
		defer close(done)
		if val, ok := lookupTable[n]; ok {
			result.Set(val)
			return
		}
		a := FibonacciRecursiveOptimized(n-1, ctx)
		b := FibonacciRecursiveOptimized(n-2, ctx)
		lookupTable[n] = new(big.Int).Add(a, b)
		result.Set(lookupTable[n])
	}()
	
	select {
	case <-ctx.Done():
		return big.NewInt(0)
	case <-done:
		return result
	}
}
