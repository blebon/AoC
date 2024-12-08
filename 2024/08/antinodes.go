package main

import (
	"image"
	"sync"
	"sync/atomic"

	"github.com/blebon/AoC/2024/util"
	"github.com/sirupsen/logrus"
)

type antennas struct {
	Freqs                   map[rune][]image.Point
	InField                 map[image.Point]bool
	IsAntinode              map[image.Point]bool
	IsAntinodeWithHarmonics map[image.Point]bool
}

func getAntinodes(f string, withHarmonics bool) int64 {
	a := getFields(f)
	a.findAntinodes()
	return a.countUniqueAntinodes(withHarmonics)
}

func getFields(f string) antennas {
	log := logrus.New()

	lines, err := util.ReadFile(f)
	if err != nil {
		log.Fatalf("error reading input file: %v", err)
	}

	a := antennas{
		IsAntinode:              map[image.Point]bool{},
		IsAntinodeWithHarmonics: map[image.Point]bool{},
	}
	freqs := map[rune][]image.Point{}
	inField := map[image.Point]bool{}

	for j, l := range lines {
		for i, c := range l {
			p := image.Point{i, j}
			inField[p] = true
			if c != '.' {
				freqs[c] = append(freqs[c], p)
			}
		}
	}

	a.Freqs = freqs
	a.InField = inField

	return a
}

func (a *antennas) findAntinodes() {
	wg := sync.WaitGroup{}
	for _, v := range a.Freqs {
		wg.Add(1)
		a.findAntinodesForFreq(v, &wg)
	}
	wg.Wait()
}

func (a *antennas) findAntinodesForFreq(v []image.Point, wg *sync.WaitGroup) {
	defer wg.Done()

	for _, antenna1 := range v {
		for _, antenna2 := range v {
			if antenna1.Eq(antenna2) {
				continue
			}

			delta := antenna2.Sub(antenna1)
			antiNode := antenna2.Add(delta)
			if a.InField[antiNode] {
				a.IsAntinode[antiNode] = true
			}

			for ; a.InField[antenna2]; antenna2 = antenna2.Add(delta) {
				a.IsAntinodeWithHarmonics[antenna2] = true
			}
		}
	}
}

func (a *antennas) countUniqueAntinodes(withHarmonics bool) int64 {
	if withHarmonics {
		return a.countAntinodes(a.IsAntinodeWithHarmonics)
	}
	return a.countAntinodes(a.IsAntinode)
}

func (a *antennas) countAntinodes(m map[image.Point]bool) int64 {
	var r atomic.Int64
	for _, v := range m {
		if v {
			r.Add(1)
		}
	}
	return r.Load()
}
