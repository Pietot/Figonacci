package main

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

	go func() {
		defer close(done)
		mu.Lock()
		if val, exist := lookupTable[n]; exist {
			result.Set(val)
			mu.Unlock()
			return
		}
		
		mu.Unlock()

		a := FibonacciRecursiveOptimized(n-1, ctx)
		b := FibonacciRecursiveOptimized(n-2, ctx)

		mu.Lock()
		lookupTable[n] = new(big.Int).Add(a, b)
		result.Set(lookupTable[n])
		mu.Unlock()
	}()

	select {
	case <-ctx.Done():
		return big.NewInt(0)
	case <-done:
		return result
	}
}
