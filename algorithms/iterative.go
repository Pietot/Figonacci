// Package algorithms provides functions to compute a Fibonacci number using different algorithms.
package algorithms

import (
	"math/big"
)

func FibonacciIterative(n int) *big.Int {
	if n <= 1 {
		return big.NewInt(int64(n))
	}
	a, b := big.NewInt(0), big.NewInt(1)
	for i := 2; i <= n; i++ {
		a, b = b, a.Add(a, b)
	}
	return b
}
