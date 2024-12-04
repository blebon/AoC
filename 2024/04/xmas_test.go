package main

import (
	"testing"
)

func TestCountXmas(t *testing.T) {
	want := 18
	got := countXmas("test.txt")
	if want != got {
		t.Errorf("wrong XMAS count. want %d, got %d", want, got)
	}
}

func TestCountX_Mas(t *testing.T) {
	want := 9
	got := countX_Mas("test.txt")
	if want != got {
		t.Errorf("wrong X-MAS count. want %d, got %d", want, got)
	}
}
