package main

import (
	"testing"
)

const TEST string = "test.txt"
const TEST_FMT string = "wrong %s: want %d, got %d"

func TestSafeReports(t *testing.T) {
	want := 2
	got := countSafeReports(TEST, isSafe)
	if want != got {
		t.Errorf(TEST_FMT, "number of safe reports", want, got)
	}
}

func TestSafeReportsWithDampener(t *testing.T) {
	want := 4
	got := countSafeReports(TEST, isSafeWithDampener)
	if want != got {
		t.Errorf(TEST_FMT, "number of safe reports with dampener", want, got)
	}
}
