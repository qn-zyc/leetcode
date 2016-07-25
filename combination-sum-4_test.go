package leetcode

import "testing"

func TestCombinationSum4(t *testing.T) {
	s := CombinationSum4([]int{1, 2, 3}, 32)
	if s != 181997601 {
		t.Errorf("want 181997601, got %d", s)
	}
}

func BenchmarkCombinationSum4(b *testing.B) {
	for i := 0; i < b.N; i++ {
		CombinationSum4([]int{1, 2, 3}, 32)
	}
}

func TestCombinationSum4_2(t *testing.T) {
	if s := CombinationSum4_2([]int{1, 2, 3}, 32); s != 181997601 {
		t.Errorf("want 181997601, got %d", s)
	}
}

func BenchmarkCombinationSum4_2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		CombinationSum4_2([]int{1, 2, 3}, 32)
	}
}
