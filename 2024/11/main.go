package main

import log "github.com/sirupsen/logrus"

func main() {
	input := "input.txt"
	log.Infof("Stones with 25 blinks: %v", countStones(input, 25))
	log.Infof("Stones with 75 blinks: %v", countStones(input, 75))
}
