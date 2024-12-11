package main

import log "github.com/sirupsen/logrus"

func main() {
	input := "input.txt"
	log.Infof("Middle sum with correct updates: %v", getMiddleSum(input, true))
	log.Infof("Middle sum with fixed updates: %v", getMiddleSum(input, false))
}
