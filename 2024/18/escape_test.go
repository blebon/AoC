package main

import (
	"testing"

	log "github.com/sirupsen/logrus"
)

const TEST string = "test.txt"
const TEST_FMT string = "wrong %s: want %v, got %v"

func TestPart1(t *testing.T) {
	s := newSpace(TEST, 6)
	s.fall(12)
	s.move()
	log.Debugf("\n%v", s.String())
	var want int = 22
	got := s.Steps
	if want != got {
		t.Errorf(TEST_FMT, "steps", want, got)
	}
}

func TestPart2(t *testing.T) {
	s := newSpace(TEST, 6)
	s.findFirstBlockByteAfter(12)
	var want string = "6,1"
	got := s.FirstBlockingByte
	if want != got {
		t.Errorf(TEST_FMT, "first byte", want, got)
	}
}
