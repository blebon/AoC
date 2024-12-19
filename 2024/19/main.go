package main

import log "github.com/sirupsen/logrus"

func main() {
	input := "input.txt"
	t := newTowels(input)
	t.check()
	log.Infof("Possible: %v", t.Possible)
	log.Infof("Combinations: %v", t.Combos)
}
