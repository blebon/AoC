package main

import log "github.com/sirupsen/logrus"

func main() {
	input := "input.txt"
	log.Infof("Number of antinodes: %v", getAntinodes(input, false))
	log.Infof("Number of antinodes with harmonics: %v", getAntinodes(input, true))
}
