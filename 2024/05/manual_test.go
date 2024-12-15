package main

import (
	"testing"
)

const TEST string = "test.txt"
const TEST_FMT string = "wrong %s: want %d, got %d"

func TestPart1(t *testing.T) {
	want := 143
	got := getMiddleSum(TEST, true)
	if want != got {
		t.Errorf(TEST_FMT, "correct update sum", want, got)
	}
}

func TestPart2(t *testing.T) {
	want := 123
	got := getMiddleSum(TEST, false)
	if want != got {
		t.Errorf(TEST_FMT, "fixed update sum", want, got)
	}
}
