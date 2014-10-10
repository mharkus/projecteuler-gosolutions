package main

// The algorithm for this solution was taken from this blog: http://www.mathblog.dk/sum-of-all-primes-below-2000000-problem-10/. All credits goes to the author/s.

import (
	"fmt"
	"math"
)

func solve(n int) int {
	sieveBound := int((n - 1) / 2)
	upperSqrt := int((math.Sqrt(float64(n)) - 1) / 2)

	var boolArray = make([]bool, sieveBound+1)
	for i := 1; i <= upperSqrt; i++ {
		if !boolArray[i] {
			for j := i * 2 * (i + 1); j <= sieveBound; j += 2*i + 1 {
				boolArray[j] = true
			}
		}
	}

	sum := 2
	for i := 1; i <= sieveBound; i++ {
		if !boolArray[i] {
			val := 2*i + 1
			sum += val
		}
	}

	return sum
}

func main() {
	fmt.Println(solve(2000000))
}
