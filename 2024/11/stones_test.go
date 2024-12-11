package main

import (
	"testing"
)

const TEST string = "test.txt"

func TestPart1(t *testing.T) {
	var want int = 55312
	got := countStones(TEST, 25)
	if want != got {
		t.Errorf("wrong number of stones: want %d, got %d", want, got)
	}
}
