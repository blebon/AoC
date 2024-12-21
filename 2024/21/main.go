package main

import log "github.com/sirupsen/logrus"

func main() {
	input := "input.txt"
	for _, i := range []int{2, 25} {
		log.Infof("Complexity sum with %d robots: %v", i, getComplexitySum(input, i))
	}
}
