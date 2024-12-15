package main

import (
	"testing"
)

const TEST string = "test.txt"
const TEST_FMT string = "wrong %s: want %d, got %d"

func TestPart1(t *testing.T) {
	var want int64 = 3749
	got := getCalibration(TEST, false)
	if want != got {
		t.Errorf(TEST_FMT, "calibration result", want, got)
	}
}

func TestPart2(t *testing.T) {
	var want int64 = 11387
	got := getCalibration(TEST, true)
	if want != got {
		t.Errorf(TEST_FMT, "calibration result with concatenation", want, got)
	}
}
