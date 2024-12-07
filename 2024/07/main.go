package main

import "github.com/sirupsen/logrus"

func main() {
	log := logrus.New()
	input := "input.txt"
	log.Infof("Calibrarion result: %v", getCalibration(input, false))
	log.Infof("Calibrarion result with concatenation: %v", getCalibration(input, true))
}
