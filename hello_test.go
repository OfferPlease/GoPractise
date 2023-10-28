package main

import (
	"testing"
)

func Sum(number [5]int) (sum int) {
	sum = 0
	for i := 0; i < 5; i++ {
		sum += number[i]
	}
	return
}
func TestSum(t *testing.T) {
	numbers := [5]int{1, 2, 3, 4, 5}
	got := Sum(numbers)
	want := 15
	if want != got {
		t.Errorf("got %d want %d given %v", got, want, numbers)
	}
}
