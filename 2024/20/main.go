package main

import log "github.com/sirupsen/logrus"

func main() {
	input := "input.txt"
	tr := newTrack(input)
	tr.findDistancesFromStart()
	log.Infof("Cheats 2: %v", tr.countCheatsSavingAtLeast(100, 2))
	log.Infof("Cheats 20: %v", tr.countCheatsSavingAtLeast(100, 20))
}
