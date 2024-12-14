package main

import "math/big"

var lookupTable = map[int]*big.Int{
	0: big.NewInt(0),
	1: big.NewInt(1),
}

func FibonacciRecursiveOptimized(n int) *big.Int {
	if val, ok := lookupTable[n]; ok {
		return val
	}
	a := FibonacciRecursiveOptimized(n-1)
	b := FibonacciRecursiveOptimized(n-2)
	lookupTable[n] = new(big.Int).Add(a, b)
	return lookupTable[n]
}
