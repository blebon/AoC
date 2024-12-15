package main

import (
	"testing"
)

const TEST string = "test.txt"
const TEST_FMT string = "wrong %s: want %d, got %d"

func TestCountXmas(t *testing.T) {
	want := 18
	got := countXmas(TEST)
	if want != got {
		t.Errorf(TEST_FMT, "XMAS count", want, got)
	}
}

func TestCountX_Mas(t *testing.T) {
	want := 9
	got := countX_Mas(TEST)
	if want != got {
		t.Errorf(TEST_FMT, "X-MAS count", want, got)
	}
}
