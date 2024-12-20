package main

import (
	"fmt"
	"testing"

	"github.com/blebon/AoC/2024/util"
	log "github.com/sirupsen/logrus"
)

const TEST string = "test.txt"
const TEST_FMT string = "wrong %s: want %v, got %v"
const CHEATS_FMT string = "There are %d cheats that save %d picoseconds."
const CHEAT_FMT string = "There is one cheat that saves %d picoseconds."

func getCheatsArray(f string) ([]int, []int) {
	l, err := util.ReadFile(f)
	util.FileError(err)

	cheats := make([]int, 0, len(l))
	times := make([]int, 0, len(l))

	for _, s := range l {
		var c, t int
		var err error
		_, err = fmt.Sscanf(s, CHEATS_FMT, &c, &t)
		if err != nil {
			c = 1
			_, err = fmt.Sscanf(s, CHEAT_FMT, &t)
		}
		if err == nil {
			cheats = append(cheats, c)
			times = append(times, t)
		}
	}

	for i := len(cheats) - 2; i >= 0; i-- {
		cheats[i] += cheats[i+1]
	}

	return cheats, times
}

func TestPart1(t *testing.T) {
	tr := newTrack(TEST)
	log.Debugf("\n%v", tr.String())
	tr.findDistancesFromStart()
	// log.Debugf("\n%v", tr.DistancesPlot())
	check := func(want int, n int) {
		got := tr.countCheatsSavingAtLeast(n, 2)
		if want != got {
			t.Errorf(TEST_FMT, "cheats", want, got)
		}
	}
	cheats, times := getCheatsArray("cheats_part1.txt")
	for i, t := range times {
		check(cheats[i], t)
	}
}

func TestPart2(t *testing.T) {
	tr := newTrack(TEST)
	tr.findDistancesFromStart()
	check := func(want int, n int) {
		got := tr.countCheatsSavingAtLeast(n, 20)
		if want != got {
			t.Errorf(TEST_FMT, "cheats", want, got)
		}
	}
	cheats, times := getCheatsArray("cheats_part2.txt")
	for i, t := range times {
		check(cheats[i], t)
	}
}
