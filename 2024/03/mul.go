package main

import (
	"fmt"
	"regexp"

	"github.com/blebon/AoC/2024/util"
	log "github.com/sirupsen/logrus"
)

func findMuls(f string, dos bool) []string {

	l, err := util.ReadLine(f)
	if err != nil {
		log.Errorf("error reading input file: %v", err)
	}

	var r *regexp.Regexp
	var s []string
	if dos {
		r = regexp.MustCompile(`mul\(\d+,\d+\)|do\(\)|don't\(\)`)
	} else {
		r = regexp.MustCompile(`mul\(\d+,\d+\)`)
	}
	m := r.FindAllString(l, -1)

	enabled := true
	for _, v := range m {
		if v == "do()" {
			enabled = true
			continue
		} else if v == "don't()" {
			enabled = false
			continue
		}
		if enabled {
			s = append(s, v)
		}
	}

	return s
}

func mul(m string) int {
	var a, b int
	fmt.Sscanf(m, "mul(%d,%d)", &a, &b)
	return a * b
}

func muls(f string, dos bool) int {
	m := findMuls(f, dos)
	ans := 0
	for _, v := range m {
		d := mul(v)
		ans += d
	}
	return ans
}
