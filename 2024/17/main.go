package main

import log "github.com/sirupsen/logrus"

func main() {
	input := "input.txt"
	d := newDebugger(input)
	d.run()
	log.Infof("Part 1 Output: %v", d.PrintOutput())
	log.Infof("Part 2 Uncorrupted A: %v", d.findUncorruptedA())
}
