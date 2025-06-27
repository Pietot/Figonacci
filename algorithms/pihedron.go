package algorithms

import (
	"math/big"
)

func Pihedron(n int) *big.Int {
	var recursive func(int) (*big.Int, *big.Int)
	recursive = func(n int) (*big.Int, *big.Int) {

		if n == 0 {
			return big.NewInt(0), big.NewInt(2)
		}

		if n&1 == 1 {
			fib, lucas := recursive(n - 1)
			result1 := new(big.Int).Add(fib, lucas)
			result1.Rsh(result1, 1)
			result2 := new(big.Int).Mul(fib, big.NewInt(5))
			result2.Add(result2, lucas)
			result2.Rsh(result2, 1)
			return result1, result2
		}

		n >>= 1
		k := n%2*2 - 1
		fib, lucas := recursive(n)
		result1 := new(big.Int).Mul(fib, lucas)
		result2 := new(big.Int).Mul(lucas, lucas)
		result2.Add(result2, big.NewInt(int64(2*k)))
		return result1, result2
	}
	result, _ := recursive(n)
	return result
}
