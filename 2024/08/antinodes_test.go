package main

import (
	"testing"
)

const TEST string = "test.txt"
const TEST_FMT string = "wrong %s: want %d, got %d"

func TestPart1(t *testing.T) {
	var want int64 = 14
	got := getAntinodes(TEST, false)
	if want != got {
		t.Errorf(TEST_FMT, "antinode number", want, got)
	}
}

func TestPart2(t *testing.T) {
	var want int64 = 34
	got := getAntinodes(TEST, true)
	if want != got {
		t.Errorf(TEST_FMT, "antinode number with resonant harmonics", want, got)
	}
}
