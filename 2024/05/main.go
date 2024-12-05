package main

import "github.com/sirupsen/logrus"

func main() {
	log := logrus.New()
	input := "input.txt"
	log.Infof("Middle sum with correct updates: %v", getMiddleSum(input, true))
	log.Infof("Middle sum with fixed updates: %v", getMiddleSum(input, false))
}
