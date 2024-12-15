package main

import (
	"testing"
)

const TEST string = "test.txt"
const TEST_FMT string = "wrong %s: want %d, got %d"

func TestPart1(t *testing.T) {
	tiles_x = 11
	tiles_y = 7

	var want int = 12
	got := getSafetyNumber(TEST, 100)
	if want != got {
		t.Errorf(TEST_FMT, "safety number", want, got)
	}
}
