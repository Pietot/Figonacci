// Package algorithms provides functions to compute a Fibonacci number using different algorithms.
package algorithms

import (
	"math/big"
	"sync"
)

func FibonacciRecursiveOptimized(n int) *big.Int {
	var lookupTable = map[int]*big.Int{
		0: big.NewInt(0),
		1: big.NewInt(1),
	}
	var mu sync.Mutex

	var recursive func(int) *big.Int
	recursive = func(n int) *big.Int {
		if n <= 1 {
			return big.NewInt(int64(n))
		}
		mu.Lock()
		if val, exist := lookupTable[n]; exist {
			mu.Unlock()
			return new(big.Int).Set(val)
		}
		mu.Unlock()

		a := recursive(n - 1)
		b := recursive(n - 2)

		mu.Lock()
		lookupTable[n] = new(big.Int).Add(a, b)
		result := new(big.Int).Set(lookupTable[n])
		mu.Unlock()
		return result
	}
	return recursive(n)
}
