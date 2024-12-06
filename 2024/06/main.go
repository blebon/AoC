package main

import "github.com/sirupsen/logrus"

func main() {
	log := logrus.New()
	input := "input.txt"
	log.Infof("Number of guard positions: %v", getGuardPositions(input))
	log.Infof("Number of obstacle positions: %v", getObstaclePositions(input))
}
