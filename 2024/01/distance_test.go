package main

import (
	"testing"
)

const TEST string = "test.txt"
const TEST_FMT string = "wrong %s: want %d, got %d"

func TestFindDistance(t *testing.T) {
	want := 11
	got := findDistance(TEST)
	if want != got {
		t.Errorf(TEST_FMT, "distance", want, got)
	}
}

func TestFindSimilarity(t *testing.T) {
	want := 31
	got := findSimilarity(TEST)
	if want != got {
		t.Errorf(TEST_FMT, "similarity", want, got)
	}
}
