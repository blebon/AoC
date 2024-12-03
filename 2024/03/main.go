package main

import "github.com/sirupsen/logrus"

func main() {
	log := logrus.New()

	i := muls("input.txt", false)
	log.Infof("Uncorrupted multiplication sums: %d", i)

	j := muls("input.txt", true)
	log.Infof("Uncorrupted multiplication sums with do's: %d", j)
}
