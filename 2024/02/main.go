package main

import (
	log "github.com/sirupsen/logrus"
)

func main() {
	input := "input.txt"

	i := countSafeReports(input, isSafe)
	j := countSafeReports(input, isSafeWithDampener)

	log.Infof("Safe reports: %d", i)
	log.Infof("Safe reports with dampener: %d", j)
}
