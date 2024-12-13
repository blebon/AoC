package main

import (
	"testing"
)

const TEST string = "test.txt"

func TestPart1(t *testing.T) {
	var want int = 480
	got := countTokens(TEST, false)
	if want != got {
		t.Errorf("wrong number of tokens: want %d, got %d", want, got)
	}
}
