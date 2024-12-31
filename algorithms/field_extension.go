package main

import (
	"context"
	"math/big"
)

type fieldExtension struct {
	// Represebtation of a + b * sqrt(5)
	A *big.Int
	B *big.Int
}

// Mutiplication in Z[sqrt(5)] = {a + b * sqrt(5) | a, b in Z}
func (f *fieldExtension) multiply(firstField *fieldExtension, secondField *fieldExtension) *fieldExtension {
	// Compute (a1 * a2 + 5 * b1 * b2) for A part
	firstTerm := new(big.Int).Mul(firstField.A, secondField.A)
	secondTerm := new(big.Int).Mul(firstField.B, secondField.B)
	secondTerm.Mul(secondTerm, big.NewInt(5))
	newA := new(big.Int).Add(firstTerm, secondTerm)

	// Compute (a1 * b2 + a2 * b1) for B part
	newB := new(big.Int).Add(
		new(big.Int).Mul(firstField.A, secondField.B),
		new(big.Int).Mul(firstField.B, secondField.A),
	)

	return &fieldExtension{A: newA, B: newB}
}

// Right shift to divide by 2
func (fieldExtension *fieldExtension) rightShift() {
	fieldExtension.A.Rsh(fieldExtension.A, 1)
	fieldExtension.B.Rsh(fieldExtension.B, 1)
}

func FibonacciFieldExtension(n int, ctx context.Context) *big.Int {
	result := big.NewInt(0)
	done := make(chan struct{})

	go func() {
		defer close(done)
		if n <= 1 {
			result.SetInt64(int64(n))
			return
		}

		// Initialisation : step = 1 + sqrt(5), fib = 1 + sqrt(5)
		step := &fieldExtension{
			A: big.NewInt(1),
			B: big.NewInt(1),
		}
		fib := &fieldExtension{
			A: big.NewInt(1),
			B: big.NewInt(1),
		}

		n--
		for n > 0 {
			if n%2 == 1 {
				fib = fib.multiply(fib, step)
				fib.rightShift()
			}
			step = step.multiply(step, step)
			step.rightShift()
			n >>= 1
		}
		result.Set(fib.B)
	}()

	select {
	case <-ctx.Done():
		return big.NewInt(0)
	case <-done:
		return result
	}
}
