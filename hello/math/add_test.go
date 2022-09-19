package math_test

import (
	"hello/math"
	"testing"
)

func TestAdd(t *testing.T) {
	got := math.Add(1, 1)
	expected := 2

	if got != expected {
		t.Fail()
	}
}

func FuzzTestAdd(f *testing.F) {
	f.Fuzz(func(t *testing.T, a, b int) {
		math.Add(a, b)
	})
}
