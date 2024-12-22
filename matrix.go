package main

import (
	"context"
	"math/big"
)

type Matrix2x2 [2][2]*big.Int

func (m *Matrix2x2) Multiply(a, b *Matrix2x2) *Matrix2x2 {
	return &Matrix2x2{
		[2]*big.Int{
			new(big.Int).Add(new(big.Int).Mul(a[0][0], b[0][0]), new(big.Int).Mul(a[0][1], b[1][0])),
			new(big.Int).Add(new(big.Int).Mul(a[0][0], b[0][1]), new(big.Int).Mul(a[0][1], b[1][1])),
		},
		[2]*big.Int{
			new(big.Int).Add(new(big.Int).Mul(a[1][0], b[0][0]), new(big.Int).Mul(a[1][1], b[1][0])),
			new(big.Int).Add(new(big.Int).Mul(a[1][0], b[0][1]), new(big.Int).Mul(a[1][1], b[1][1])),
		},
	}
}

func (m *Matrix2x2) Pow(n int) *Matrix2x2 {
	result := &Matrix2x2{
		[2]*big.Int{big.NewInt(1), big.NewInt(0)},
		[2]*big.Int{big.NewInt(0), big.NewInt(1)},
	}
	base := m
	for i := 0; i < n; i++ {
		result = result.Multiply(result, base)
	}
	return result
}

func FibonacciMatrix(n int, ctx context.Context) *big.Int {
	result := big.NewInt(0)
	done := make(chan struct{})

	go func() {
		defer close(done)
		if n == 0 {
			result.SetInt64(0)
			return
		}
		fibonacciMatrix := &Matrix2x2{
			[2]*big.Int{big.NewInt(1), big.NewInt(1)},
			[2]*big.Int{big.NewInt(1), big.NewInt(0)},
		}
		resultMatrix := fibonacciMatrix.Pow(n - 1)
		result.Set(resultMatrix[0][0])
	}()
	
	select {
	case <-ctx.Done():
		return big.NewInt(0)
	case <-done:
		return result
	}
}
