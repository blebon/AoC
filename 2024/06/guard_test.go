package main

import (
	"testing"
)

const TEST string = "test.txt"

func TestPart1(t *testing.T) {
	want := 41
	got := getGuardPositions(TEST)
	if want != got {
		t.Errorf("wrong number of guard positions: want %d, got %d", want, got)
	}
}

func TestPart2(t *testing.T) {
	want := 6
	got := getObstaclePositions(TEST)
	if want != got {
		t.Errorf("wrong number of obstacle positions: want %d, got %d", want, got)
	}
}
