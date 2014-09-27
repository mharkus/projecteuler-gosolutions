package main

import (
	"testing"
)

func TestPrime(t *testing.T) {
	pipe := make(chan int)
	defer close(pipe)

	go solve(pipe, 10001)

	if <-pipe != 104743 {
		t.Errorf("10001st prime should be 104743")
	}
}
