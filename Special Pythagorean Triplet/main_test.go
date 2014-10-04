package main

import (
	"testing"
)

func TestPythagoreanTriples(t *testing.T) {
	if solve(1000) != 31875000 {
		t.Errorf("Answer should be 31875000")
	}
}

func BenchmarkPythagoreanTriples(b *testing.B) {
	for i := 0; i < b.N; i++ {
		solve(1000)
	}

}
