package main

import log "github.com/sirupsen/logrus"

func main() {
	input := "input.txt"
	log.Infof("Trailheads: %v", getTrailCount(input, false))
	log.Infof("Ratings: %v", getTrailCount(input, true))
}
