package main

import (
	"testing"
)

const TEST string = "test.txt"
const TEST_FMT string = "wrong %s: want %d, got %d"

func TestPart1(t *testing.T) {
	var want int = 55312
	got := countStones(TEST, 25)
	if want != got {
		t.Errorf(TEST_FMT, "number of stones", want, got)
	}
}
