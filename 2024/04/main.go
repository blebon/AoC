package main

import "github.com/sirupsen/logrus"

func main() {
	log := logrus.New()

	i := countXmas("input.txt")
	log.Infof("XMAS count: %v", i)

	j := countX_Mas("input.txt")
	log.Infof("X-MAS count: %v", j)
}
