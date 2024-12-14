package main

import (
	"testing"
)

const TEST string = "test.txt"

func TestPart1(t *testing.T) {
	tiles_x = 11
	tiles_y = 7

	var want int = 12
	got := getSafetyNumber(TEST, 100)
	if want != got {
		t.Errorf("wrong safety number: want %d, got %d", want, got)
	}
}
