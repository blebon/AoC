package main

import (
	"testing"
)

const TEST_SMALL string = "test_small.txt"
const TEST_OX string = "test_ox.txt"
const TEST string = "test.txt"
const TEST_FMT string = "wrong %s: want %d, got %d"

func TestPart1Small(t *testing.T) {
	var want int = 140
	got := getPrice(TEST_SMALL, false)
	if want != got {
		t.Errorf(TEST_FMT, "cost", want, got)
	}
}

func TestPart1OX(t *testing.T) {
	var want int = 772
	got := getPrice(TEST_OX, false)
	if want != got {
		t.Errorf(TEST_FMT, "cost", want, got)
	}
}

func TestPart1(t *testing.T) {
	var want int = 1930
	got := getPrice(TEST, false)
	if want != got {
		t.Errorf(TEST_FMT, "cost", want, got)
	}
}

func TestPart2Small(t *testing.T) {
	var want int = 80
	got := getPrice(TEST_SMALL, true)
	if want != got {
		t.Errorf(TEST_FMT, "cost", want, got)
	}
}

func TestPart2OX(t *testing.T) {
	var want int = 436
	got := getPrice(TEST_OX, true)
	if want != got {
		t.Errorf(TEST_FMT, "cost", want, got)
	}
}

func TestPart2E(t *testing.T) {
	var want int = 236
	got := getPrice("test_e.txt", true)
	if want != got {
		t.Errorf(TEST_FMT, "cost", want, got)
	}
}

func TestPart2AB(t *testing.T) {
	var want int = 368
	got := getPrice("test_ab.txt", true)
	if want != got {
		t.Errorf(TEST_FMT, "cost", want, got)
	}
}

func TestPart2(t *testing.T) {
	var want int = 1206
	got := getPrice(TEST, true)
	if want != got {
		t.Errorf(TEST_FMT, "cost", want, got)
	}
}
