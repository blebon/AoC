package main

import (
	"testing"
)

const TEST string = "test.txt"
const TEST_FMT string = "wrong %s: want %d, got %d"

func TestPart1(t *testing.T) {
	want := 41
	got := getGuardPositions(TEST)
	if want != got {
		t.Errorf(TEST_FMT, "number of guard positions", want, got)
	}
}

func TestPart2(t *testing.T) {
	want := 6
	got := getObstaclePositions(TEST)
	if want != got {
		t.Errorf(TEST_FMT, "number of obstacle positions", want, got)
	}
}
