package main

import (
	"testing"
)

const TEST string = "test.txt"

func TestPart1(t *testing.T) {
	var want int64 = 14
	got := getAntinodes(TEST, false)
	if want != got {
		t.Errorf("wrong antinode number: want %d, got %d", want, got)
	}
}

func TestPart2(t *testing.T) {
	var want int64 = 34
	got := getAntinodes(TEST, true)
	if want != got {
		t.Errorf("wrong antinode number with resonant harmonics : want %d, got %d", want, got)
	}
}
