package main

import (
	"testing"
)

func TestLargestProduct(t *testing.T) {
	if 23514624000 != solve(13) {
		t.Errorf("Largest Product for 13 digits should be 23514624000")
	}
}
