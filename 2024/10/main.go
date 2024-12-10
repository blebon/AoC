package main

import "github.com/sirupsen/logrus"

func main() {
	log := logrus.New()
	input := "input.txt"
	log.Infof("Trailhead: %v", getTrailCount(input, false))
	log.Infof("Trailhead with : %v", getTrailCount(input, true))
}
