package main

import (
	"testing"

	log "github.com/sirupsen/logrus"
)

const TEST string = "test.txt"
const TEST_FMT string = "wrong %s: want %v, got %v"

func TestNumericKeypadSequence(t *testing.T) {
	want := "<A^A>^^AvvvA"
	got := NumericKeypad.generateSequence("029A")
	if len(want) != len(got) {
		t.Errorf(TEST_FMT, "sequence", want, got)
	}
}

func TestDirectionalKeypadSequence(t *testing.T) {
	want := "v<<A>>^A<A>AvA<^AA>A<vAAA>^A"
	got := getSequenceLength("029A", 1)
	if len(want) != got {
		t.Errorf(TEST_FMT, "sequence length", want, got)
	}
}

func TestDirectionalKeypadSequence2(t *testing.T) {
	want := "<vA<AA>>^AvAA<^A>A<v<A>>^AvA^A<vA>^A<v<A>^A>AAvA^A<v<A>A>^AAAvA<^A>A"
	got := getSequenceLength("029A", 2)
	if len(want) != got {
		t.Errorf(TEST_FMT, "sequence", want, got)
	}
}

func TestComplexity029A(t *testing.T) {
	want := 29 * 68
	got := getComplexity("029A", 2)
	if want != got {
		t.Errorf(TEST_FMT, "complexity", want, got)
	}
}

func TestComplexity980A(t *testing.T) {
	want := 980 * 60
	got := getComplexity("980A", 2)
	if want != got {
		t.Errorf(TEST_FMT, "complexity", want, got)
	}
}

func TestComplexity179A(t *testing.T) {
	want := 179 * 68
	got := getComplexity("179A", 2)
	if want != got {
		t.Errorf(TEST_FMT, "complexity", want, got)
	}
}

func TestComplexity456A(t *testing.T) {
	want := 456 * 64
	got := getComplexity("456A", 2)
	if want != got {
		t.Errorf(TEST_FMT, "complexity", want, got)
	}
}

func TestComplexity379A(t *testing.T) {
	want := 379 * 64
	got := getComplexity("379A", 2)
	if want != got {
		t.Errorf(TEST_FMT, "complexity", want, got)
	}
}

func TestPart1(t *testing.T) {
	log.Debugf("\n%v", TEST)
	want := 126384
	got := getComplexitySum(TEST, 2)
	if want != got {
		t.Errorf(TEST_FMT, "complexity sum", want, got)
	}
}
