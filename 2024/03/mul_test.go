package main

import (
	"testing"
)

func TestMul(t *testing.T) {
	want := 161
	got := muls("test1.txt", false)
	if want != got {
		t.Errorf("wrong uncorrupted multiplications sum. want %d, got %d", want, got)
	}
}

func TestDoMul(t *testing.T) {
	want := 48
	got := muls("test2.txt", true)
	if want != got {
		t.Errorf("wrong uncorrupted multiplications sum. want %d, got %d", want, got)
	}
}
