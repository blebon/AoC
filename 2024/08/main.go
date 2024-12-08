package main

import "github.com/sirupsen/logrus"

func main() {
	log := logrus.New()
	input := "input.txt"
	log.Infof("Number of antinodes: %v", getAntinodes(input, false))
	log.Infof("Number of antinodes with harmonics: %v", getAntinodes(input, true))
}
