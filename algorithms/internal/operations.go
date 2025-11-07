package matrix

import "math/big"

type Matrix2x2 [2][2]*big.Int

func (m *Matrix2x2) Multiply(b *Matrix2x2) *Matrix2x2 {
	return &Matrix2x2{
		[2]*big.Int{
			new(big.Int).Add(new(big.Int).Mul(m[0][0], b[0][0]), new(big.Int).Mul(m[0][1], b[1][0])),
			new(big.Int).Add(new(big.Int).Mul(m[0][0], b[0][1]), new(big.Int).Mul(m[0][1], b[1][1])),
		},
		[2]*big.Int{
			new(big.Int).Add(new(big.Int).Mul(m[1][0], b[0][0]), new(big.Int).Mul(m[1][1], b[1][0])),
			new(big.Int).Add(new(big.Int).Mul(m[1][0], b[0][1]), new(big.Int).Mul(m[1][1], b[1][1])),
		},
	}
}

func (m *Matrix2x2) Pow(n int) *Matrix2x2 {
	result := &Matrix2x2{
		[2]*big.Int{big.NewInt(1), big.NewInt(0)},
		[2]*big.Int{big.NewInt(0), big.NewInt(1)},
	}
	base := m
	for range n {
		result = result.Multiply(base)
	}
	return result
}

func (m *Matrix2x2) FastPow(n int) *Matrix2x2 {
	result := &Matrix2x2{
		[2]*big.Int{big.NewInt(1), big.NewInt(0)},
		[2]*big.Int{big.NewInt(0), big.NewInt(1)},
	}
	base := m
	for n > 0 {
		if n%2 == 1 {
			result = result.Multiply(base)
		}
		n >>= 1
		base = base.Multiply(base)
	}
	return result
}
