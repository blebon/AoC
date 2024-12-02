package main

import (
	"github.com/sirupsen/logrus"
)

func main() {
	log := logrus.New()

	// input := "test.txt"
	input := "input.txt"

	i := countSafeReports(input, isSafe)
	j := countSafeReports(input, isSafeWithDampener)

	log.Infof("Safe reports: %d", i)
	log.Infof("Safe reports with dampener: %d", j)
}
