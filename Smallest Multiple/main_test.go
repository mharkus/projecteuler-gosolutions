package main

import (
	"testing"
)

func TestLCM(t *testing.T) {
	prev := 1
	for x := 1; x <= 10; x++ {
		prev = lcm(x, prev)
	}

	if prev != 2520 {
		t.Errorf("LCM of 10 should be 2520")
	}
}
