package main

import (
	"testing"
)

const TEST string = "test.txt"
const TEST_FMT string = "wrong %s: want %d, got %d"

func TestPart1(t *testing.T) {
	var want int = 36
	got := getTrailCount(TEST, false)
	if want != got {
		t.Errorf(TEST_FMT, "trailhead", want, got)
	}
}

func TestPart2(t *testing.T) {
	var want int = 81
	got := getTrailCount(TEST, true)
	if want != got {
		t.Errorf(TEST_FMT, "rating", want, got)
	}
}
