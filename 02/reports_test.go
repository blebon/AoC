package main

import (
	"testing"
)

const TEST string = "test.txt"

func TestSafeReports(t *testing.T) {
	want := 2
	got := countSafeReports(TEST, isSafe)
	if want != got {
		t.Errorf("wrong number of safe reports. want %d, got %d", want, got)
	}
}

func TestSafeReportsWithDampener(t *testing.T) {
	want := 4
	got := countSafeReports(TEST, isSafeWithDampener)
	if want != got {
		t.Errorf("wrong number of safe reports with dampener. want %d, got %d", want, got)
	}
}
