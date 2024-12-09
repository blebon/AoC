package main

import "github.com/sirupsen/logrus"

func main() {
	log := logrus.New()
	input := "input.txt"
	log.Infof("Checksum: %v", getChecksum(input, false))
	log.Infof("Checksum with less fragmentation: %v", getChecksum(input, true))
}
