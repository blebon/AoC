package main

import (
	"testing"
)

const TEST_FMT string = "wrong %s: want %d, got %d"

func TestMul(t *testing.T) {
	want := 161
	got := muls("test1.txt", false)
	if want != got {
		t.Errorf(TEST_FMT, "uncorrupted multiplications sum", want, got)
	}
}

func TestDoMul(t *testing.T) {
	want := 48
	got := muls("test2.txt", true)
	if want != got {
		t.Errorf(TEST_FMT, "uncorrupted multiplications sum", want, got)
	}
}
