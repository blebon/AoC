package main

import log "github.com/sirupsen/logrus"

func main() {
	input := "input.txt"
	m := getMaze(input)
	m.searchShortestPath()
	log.Infof("Minimum maze score: %v", m.getScore())
	log.Infof("Tiles: %v", m.getTiles())
}
