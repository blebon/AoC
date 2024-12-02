package main

import (
	"testing"
)

const TEST string = "test.txt"

func TestFindDistance(t *testing.T) {
	want := 11
	got := findDistance(TEST)
	if want != got {
		t.Errorf("wrong distance. want %d, got %d", want, got)
	}
}

func TestFindSimilarity(t *testing.T) {
	want := 31
	got := findSimilarity(TEST)
	if want != got {
		t.Errorf("wrong similarity. want %d, got %d", want, got)
	}
}
