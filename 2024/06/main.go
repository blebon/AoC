package main

import log "github.com/sirupsen/logrus"

func main() {
	input := "input.txt"
	log.Infof("Number of guard positions: %v", getGuardPositions(input))
	log.Infof("Number of obstacle positions: %v", getObstaclePositions(input))
}
