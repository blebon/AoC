package main

import log "github.com/sirupsen/logrus"

func main() {
	input := "input.txt"
	log.Infof("Part 1: %v", getSafetyNumber(input, 100))
	log.Infof("Part2: %v", getXmasTree(input, true))
}
