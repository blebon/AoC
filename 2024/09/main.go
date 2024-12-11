package main

import log "github.com/sirupsen/logrus"

func main() {
	input := "input.txt"
	log.Infof("Checksum: %v", getChecksum(input, false))
	log.Infof("Checksum with less fragmentation: %v", getChecksum(input, true))
}
