// Package algorithms provides functions to compute a Fibonacci number using different algorithms.
package algorithms

import (
	"math/big"

	matrix "github.com/Pietot/Figonacci/v2/algorithms/internal"
)

func FibonacciMatrix(n int) *big.Int {
	if n <= 1 {
		return big.NewInt(int64(n))
	}

	fibonacciMatrix := &matrix.Matrix2x2{
		[2]*big.Int{big.NewInt(1), big.NewInt(1)},
		[2]*big.Int{big.NewInt(1), big.NewInt(0)},
	}

	resultMatrix := fibonacciMatrix.Pow(n - 1)
	return resultMatrix[0][0]
}
