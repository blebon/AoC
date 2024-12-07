package main

import (
	"testing"
)

const TEST string = "test.txt"

func TestPart1(t *testing.T) {
	var want int64 = 3749
	got := getCalibration(TEST, false)
	if want != got {
		t.Errorf("wrong calibration result: want %d, got %d", want, got)
	}
}

func TestPart2(t *testing.T) {
	var want int64 = 11387
	got := getCalibration(TEST, true)
	if want != got {
		t.Errorf("wrong : want %d, got %d", want, got)
	}
}
