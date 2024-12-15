package util

import log "github.com/sirupsen/logrus"

func FileError(err error) {
	if err != nil {
		log.Fatalf("error reading input file: %v", err)
	}
}
