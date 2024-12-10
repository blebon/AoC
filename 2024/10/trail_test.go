package main

import (
	"testing"
)

const TEST string = "test.txt"

func TestPart1(t *testing.T) {
	var want int = 36
	got := getTrailCount(TEST, false)
	if want != got {
		t.Errorf("wrong trailhead: want %d, got %d", want, got)
	}
}

func TestPart2(t *testing.T) {
	var want int = 81
	got := getTrailCount(TEST, true)
	if want != got {
		t.Errorf("wrong rating: want %d, got %d", want, got)
	}
}
