package main

import (
	"testing"
)

const TEST string = "test.txt"
const TEST2 string = "test2.txt"
const TEST_FMT string = "wrong %s: want %d, got %d"

func TestPart1(t *testing.T) {
	DEBUG = true
	var want int = 7036
	m := getMaze(TEST)
	m.searchShortestPath()
	got := m.getScore()
	if want != got {
		t.Errorf(TEST_FMT, "maze score", want, got)
	}
}

func TestPart12(t *testing.T) {
	DEBUG = true
	var want int = 11048
	m := getMaze(TEST2)
	m.searchShortestPath()
	got := m.getScore()
	if want != got {
		t.Errorf(TEST_FMT, "minimum maze score", want, got)
	}
}

func TestPart2(t *testing.T) {
	var want int = 45
	m := getMaze(TEST)
	m.searchShortestPath()
	got := m.getTiles()
	if want != got {
		t.Errorf(TEST_FMT, "tiles", want, got)
	}
}

func TestPart22(t *testing.T) {
	var want int = 64
	m := getMaze(TEST2)
	m.searchShortestPath()
	got := m.getTiles()
	if want != got {
		t.Errorf(TEST_FMT, "minimum maze score", want, got)
	}
}
