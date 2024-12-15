package main

import (
	"testing"
)

const TEST_SMALL string = "test_small.txt"
const TEST string = "test.txt"

func TestPart1Small(t *testing.T) {
	var want int = 2028
	got := getGPSSum(TEST_SMALL, false)
	if want != got {
		t.Errorf("wrong GPS sum: want %d, got %d", want, got)
	}
}

func TestPart1(t *testing.T) {
	var want int = 10092
	got := getGPSSum(TEST, false)
	if want != got {
		t.Errorf("wrong GPS sum: want %d, got %d", want, got)
	}
}

func TestPart2Small(t *testing.T) {
	var want int = 618
	got := getGPSSum("test_small2.txt", true)
	if want != got {
		t.Errorf("wrong GPS sum: want %d, got %d", want, got)
	}
}

func TestPart2(t *testing.T) {
	var want int = 9021
	got := getGPSSum(TEST, true)
	if want != got {
		t.Errorf("wrong GPS sum: want %d, got %d", want, got)
	}
}
