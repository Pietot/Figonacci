// Package algorithms provides functions to compute a Fibonacci number using different algorithms.
package algorithms

import (
	"math/big"
)

func FibonacciRecursiveOptimized(n int) *big.Int {
	var lookupTable = map[int]*big.Int{
		0: big.NewInt(0),
		1: big.NewInt(1),
	}

	var recursive func(int) *big.Int
	recursive = func(n int) *big.Int {
		if n <= 1 {
			return big.NewInt(int64(n))
		}
		if val, exist := lookupTable[n]; exist {
			return new(big.Int).Set(val)
		}

		a := recursive(n - 1)
		b := recursive(n - 2)

		lookupTable[n] = new(big.Int).Add(a, b)
		result := new(big.Int).Set(lookupTable[n])
		return result
	}
	return recursive(n)
}
