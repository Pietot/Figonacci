package main

import (
	"context"
	"math/big"
)

type Matrix2x2 [2][2]*big.Int

func (m *Matrix2x2) Mul(a, b *Matrix2x2, ctx context.Context) *Matrix2x2 {
	select {
	case <-ctx.Done():
		return &Matrix2x2{
			[2]*big.Int{big.NewInt(0), big.NewInt(0)},
			[2]*big.Int{big.NewInt(0), big.NewInt(0)},
		}
	default:
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
}

func (m *Matrix2x2) Pow(n int, ctx context.Context) *Matrix2x2 {
	result := &Matrix2x2{
		[2]*big.Int{big.NewInt(1), big.NewInt(0)},
		[2]*big.Int{big.NewInt(0), big.NewInt(1)},
	}
	base := m
	for i := 0; i < n; i++ {
		select {
		case <-ctx.Done():
			return &Matrix2x2{
				[2]*big.Int{big.NewInt(0), big.NewInt(0)},
				[2]*big.Int{big.NewInt(0), big.NewInt(0)},
			}
		default:
			result = result.Mul(result, base, ctx)
		}
	}
	return result
}

func FibonacciMatrix(n int, ctx context.Context) *big.Int {
	if n == 0 {
		return big.NewInt(0)
	}
	fibonacciMatrix := &Matrix2x2{
		[2]*big.Int{big.NewInt(1), big.NewInt(1)},
		[2]*big.Int{big.NewInt(1), big.NewInt(0)},
	}
	resultMatrix := fibonacciMatrix.Pow(n-1, ctx)
	return resultMatrix[0][0]
}
