package main

import log "github.com/sirupsen/logrus"

func main() {
	input := "input.txt"
	log.Infof("XMAS count: %v", countXmas(input))
	log.Infof("X-MAS count: %v", countX_Mas(input))
}
