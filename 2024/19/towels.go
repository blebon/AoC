package main

import (
	"strings"

	"github.com/blebon/AoC/2024/util"
	log "github.com/sirupsen/logrus"
)

type Towels struct {
	Available []string
	Desired   []string
	Possible  int
	Combos    int
	Memo      map[string]int
}

func newTowels(f string) Towels {
	l, err := util.ReadFile(f)
	util.FileError(err)

	memo := map[string]int{}
	memo[""] = 1

	return Towels{
		Available: strings.Split(l[0], ", "),
		Desired:   l[2:],
		Possible:  0,
		Combos:    0,
		Memo:      memo,
	}
}

func (t *Towels) check() {
	t.Possible = 0
	t.Combos = 0
	for _, d := range t.Desired {
		c := t.checkDesign(d)
		log.Debugf("%v %v", d, c)
		if c > 0 {
			t.Possible++
			t.Combos += c
		}
	}
}

func (t *Towels) checkDesign(d string) int {
	n, ok := t.Memo[d]
	if ok {
		return n
	}

	for _, a := range t.Available {
		if strings.HasPrefix(d, a) {
			n += t.checkDesign(d[len(a):])
		}
	}
	t.Memo[d] = n
	return n
}
