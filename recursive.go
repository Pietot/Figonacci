package main

import "math/big"

func FibonacciRecursive(n int) *big.Int {
	if n <= 1 {
		return big.NewInt(int64(n))
	}
	a := FibonacciRecursive(n - 1)
	b := FibonacciRecursive(n - 2)
	return new(big.Int).Add(a, b)
}
