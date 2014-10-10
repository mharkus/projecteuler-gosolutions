package main

import (
	"testing"
)

func TestSolve(t *testing.T) {

	sum := solve(2000000)

	if sum != 142913828922 {
		t.Errorf("sum of primes below 2000000 should be 142913828922")
	}
}
