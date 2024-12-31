package edu

import "testing"

func Test_fib(t *testing.T) {
	// верный результат
	nums := 7
	want := 13
	got := fib(nums)
	if got != want {
		t.Errorf("получено %d, ожидалось %d\n", got, want)
	}
}
