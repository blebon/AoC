package main

import (
	"testing"
)

const TEST string = "test.txt"
const TEST_FMT string = "wrong %s: want %d, got %d"

func TestPart1(t *testing.T) {
	var want int64 = 1928
	got := getChecksum(TEST, false)
	if want != got {
		t.Errorf(TEST_FMT, "checksum", want, got)
	}
}

func TestPart2(t *testing.T) {
	var want int64 = 2858
	got := getChecksum(TEST, true)
	if want != got {
		t.Errorf(TEST_FMT, "checksum without fragmentation", want, got)
	}
}
