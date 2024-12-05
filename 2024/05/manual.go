package main

import (
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"

	"github.com/sirupsen/logrus"
)

func readUpdates(f string) ([]string, error) {
	b, err := os.ReadFile(f)
	if err != nil {
		return nil, err
	}

	manual := strings.Split(string(b), "\n\n")
	if len(manual) != 2 {
		return manual, fmt.Errorf("wrong manual size: expected 2, got %d", len(manual))
	}

	return manual, nil
}

func getMiddleSum(f string, ordered bool) int {
	log := logrus.New()

	manual, err := readUpdates(f)
	if err != nil {
		log.Fatalf("error reading file: %v", err)
		return 0
	}

	order := [][]string{}
	for _, v := range strings.Split(manual[0], "\n") {
		order = append(order, strings.Split(v, "|"))
	}
	log.Debugf("ordering pairs: %v", order)

	cmp := func(a, b string) int {
		for _, s := range order {
			if s[0] == a && s[1] == b {
				return -1
			}
			if s[0] == b && s[1] == a {
				return 1
			}
		}
		return 0
	}

	var ans int = 0
	for _, l := range strings.Split(manual[1], "\n") {
		s := strings.Split(l, ",")
		if slices.IsSortedFunc(s, cmp) == ordered {
			if !ordered {
				slices.SortFunc(s, cmp)
			}
			n, _ := strconv.Atoi(s[len(s)/2])
			ans += n
		}
	}
	return ans
}
