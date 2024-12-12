package main

import log "github.com/sirupsen/logrus"

func main() {
	input := "input.txt"
	log.Infof("Cost with perimeter: %v", getPrice(input, false))
	log.Infof("Cost with sides: %v", getPrice(input, true))
}
