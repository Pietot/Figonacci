package algorithms

import (
	"context"
	"math/big"

	"github.com/Pietot/Figonacci/matrix"
)

// Use of fast exponentiation to calculate the Fibonacci number
func FibonacciMatrixOptimized(n int, ctx context.Context) *big.Int {
	result := big.NewInt(0)
	done := make(chan struct{})

	go func() {
		defer close(done)
		if n <= 1 {
			result.SetInt64(int64(n))
			return
		}
		fibonacciMatrix := &matrix.Matrix2x2{
			[2]*big.Int{big.NewInt(1), big.NewInt(1)},
			[2]*big.Int{big.NewInt(1), big.NewInt(0)},
		}
		resultMatrix := fibonacciMatrix.FastPow(n - 1)
		result.Set(resultMatrix[0][0])
	}()

	select {
	case <-ctx.Done():
		return big.NewInt(0)
	case <-done:
		return result
	}
}
