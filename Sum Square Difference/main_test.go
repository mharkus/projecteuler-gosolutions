package main

import (
	//"fmt"
	"testing"
)

func TestSolveSumOfSquares(t *testing.T) {
	pipe := make(chan int)
	defer close(pipe)

	go solveSumOfSquares(pipe, 10, 1, 0)

	x := <-pipe

	if x != 385 {
		t.Errorf("Sum of Squares of 10 should be 385")
	}
}

func TestSolveSquareOfSums(t *testing.T) {
	pipe := make(chan int)
	defer close(pipe)

	go solveSquareOfSums(pipe, 10, 1, 0)

	x := <-pipe

	if x != 3025 {
		t.Errorf("Sum of Squares of 10 should be 3025")
	}
}
