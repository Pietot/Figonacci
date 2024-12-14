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
	select {
	case <-ctx.Done():
		return big.NewInt(0)
	default:
		if val, ok := lookupTable[n]; ok {
			return val
		}
		a := FibonacciRecursiveOptimized(n-1, ctx)
		b := FibonacciRecursiveOptimized(n-2, ctx)
		lookupTable[n] = new(big.Int).Add(a, b)
		return lookupTable[n]
	}
}
