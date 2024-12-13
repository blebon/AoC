package main

import log "github.com/sirupsen/logrus"

func main() {
	input := "input.txt"
	log.Infof("Tokens: %v", countTokens(input, false))
	log.Infof("Tokens: %v", countTokens(input, true))
}
