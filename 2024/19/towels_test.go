package main

import (
	"testing"

	log "github.com/sirupsen/logrus"
)

const TEST string = "test.txt"
const TEST_FMT string = "wrong %s: want %v, got %v"

func TestPart1(t *testing.T) {
	towels := newTowels(TEST)
	towels.check()
	log.Info(towels)
	var want int = 6
	got := towels.Possible
	if want != got {
		t.Errorf(TEST_FMT, "possible count", want, got)
	}
}

func TestPart2(t *testing.T) {
	towels := newTowels(TEST)
	towels.check()
	var want int = 16
	got := towels.Combos
	if want != got {
		t.Errorf(TEST_FMT, "combos", want, got)
	}
}
