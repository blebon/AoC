package main

import (
	"strconv"
	"strings"

	"github.com/blebon/AoC/2024/util"
	log "github.com/sirupsen/logrus"
)

func countStones(f string, b int) int {
	s := getStones(f)

	var r int = 0
	for _, v := range s {
		r += blink(v, b)
	}

	return r
}

func getStones(f string) []string {
	l, err := util.ReadLine(f)
	if err != nil {
		log.Fatalf("error reading file: %v", err)
	}
	return strings.Fields(l)
}

type key struct {
	Value string
	Blink int
}

var memo = map[key]int{}

func blink(v string, b int) int {
	if b == 0 {
		memo[key{Value: v, Blink: b}] = 1
		return 1
	}

	r, ok := memo[key{Value: v, Blink: b}]
	if ok {
		return r
	}

	if v == "0" {
		r = blink("1", b-1)
	} else if len(v)%2 == 0 {
		r = getSplit(v, b)
	} else {
		i, _ := strconv.Atoi(v)
		r = blink(strconv.Itoa(2024*i), b-1)
	}

	memo[key{Value: v, Blink: b}] = r
	return r
}

func getSplit(v string, b int) int {
	var r int = 0

	f, _ := strconv.Atoi(v[:len(v)/2])
	r += blink(strconv.Itoa(f), b-1)

	s, _ := strconv.Atoi(v[len(v)/2:])
	r += blink(strconv.Itoa(s), b-1)

	return r
}
