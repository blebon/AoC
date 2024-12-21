package main

import log "github.com/sirupsen/logrus"

func main() {
	input := "input.txt"
	log.Infof("Complexity sum 2: %v", getComplexitySum(input, 2))
	log.Infof("Complexity sum 25: %v", getComplexitySum(input, 25))
}
