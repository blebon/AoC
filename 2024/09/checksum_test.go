package main

import (
	"testing"
)

const TEST string = "test.txt"

func TestPart1(t *testing.T) {
	var want int64 = 1928
	got := getChecksum(TEST, false)
	if want != got {
		t.Errorf("wrong checksum: want %d, got %d", want, got)
	}
}

func TestPart2(t *testing.T) {
	var want int64 = 2858
	got := getChecksum(TEST, true)
	if want != got {
		t.Errorf("wrong checksum with less fragmentation: want %d, got %d", want, got)
	}
}
