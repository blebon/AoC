package main

import log "github.com/sirupsen/logrus"

func main() {
	input := "input.txt"
	log.Infof("Part 1: %v", getGPSSum(input, false))
	log.Infof("Part 2: %v", getGPSSum(input, true))
}
