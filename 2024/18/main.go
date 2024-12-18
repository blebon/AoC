package main

import log "github.com/sirupsen/logrus"

func main() {
	input := "input.txt"
	s := newSpace(input, 70)
	s.fall(1024)
	s.move()
	log.Infof("Steps: %v", s.Steps)
	s.findFirstBlockByteAfter(1024)
	log.Infof("First blocking byte: %v", s.FirstBlockingByte)
}
