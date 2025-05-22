package my_math

import (
	"testing"
)

func TestFact(t *testing.T) {
	want := 6
	got := Fact(3)
	if got != want {
		t.Errorf("Fact() = %v, want %v", got, want)
	}
}

func TestMaxNum(t *testing.T) {
	want := 6
	nums := []int{1, 2, 3, 4, 5, 6, -5, 0}
	got := MaxNum(nums)
	if got != want {
		t.Errorf("Fact() = %v, want %v", got, want)
	}
}
