package main

import (
	"testing"
)

const TEST string = "test.txt"

func TestPart1(t *testing.T) {
	want := 143
	got := getMiddleSum(TEST, true)
	if want != got {
		t.Errorf("wrong correct update sum. want %d, got %d", want, got)
	}
}

func TestPart2(t *testing.T) {
	want := 123
	got := getMiddleSum(TEST, false)
	if want != got {
		t.Errorf("wrong fixed update sum. want %d, got %d", want, got)
	}
}
