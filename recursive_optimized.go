package main

var lookupTable = map[int]int{
	0: 0,
	1: 1,
}

func FibonacciRecursiveOptimized(n int) int {
	if val, ok := lookupTable[n]; ok {
		return val
	}
	lookupTable[n] = FibonacciRecursiveOptimized(n-1) + FibonacciRecursiveOptimized(n-2)
	return lookupTable[n]
}
