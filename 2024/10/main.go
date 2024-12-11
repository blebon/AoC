package main

import "github.com/sirupsen/logrus"

func main() {
	log := logrus.New()
	input := "input.txt"
	log.Infof("Trailheads: %v", getTrailCount(input, false))
	log.Infof("Ratings: %v", getTrailCount(input, true))
}
