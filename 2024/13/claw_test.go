package main

import (
	"testing"
)

const TEST string = "test.txt"
const TEST_FMT string = "wrong %s: want %d, got %d"

func TestPart1(t *testing.T) {
	var want int = 480
	got := countTokens(TEST, false)
	if want != got {
		t.Errorf(TEST_FMT, "number of tokens", want, got)
	}
}
