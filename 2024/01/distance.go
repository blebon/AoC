package main

import (
	"sort"

	util "github.com/blebon/AoC/2024/util"
	log "github.com/sirupsen/logrus"
)

func readColumns(input string) ([]int, []int, error) {
	fields, err := util.ReadSpaceSeparatedFileToInt(input)
	if err != nil {
		log.Errorf("error reading input file: %v", err)
		return nil, nil, err
	}

	a := make([]int, 0, len(fields))
	b := make([]int, 0, len(fields))

	for _, l := range fields {
		a = append(a, l[0])
		b = append(b, l[1])
	}

	log.Debugf("A: %v, B: %v", a, b)
	return a, b, nil
}

func findDistance(input string) int {
	a, b, _ := readColumns(input)
	sort.Ints(a)
	sort.Ints(b)

	ans := 0
	for i, v := range a {
		d := v - b[i]
		if d > 0 {
			ans += d
		} else {
			ans -= d
		}
	}

	log.Infof("Distance: %v", ans)

	return ans
}

func findSimilarity(input string) int {
	a, b, _ := readColumns(input)

	bmap := make(map[int]int)
	for _, v := range b {
		_, ok := bmap[v]
		if ok {
			bmap[v] += 1
		} else {
			bmap[v] = 1
		}
	}

	counts := make(map[int]int)
	ans := 0
	for _, v := range a {
		val, ok := counts[v]
		if ok {
			ans += val
		} else {
			bv, ok := bmap[v]
			if ok {
				counts[v] = v * bv
				ans += counts[v]
			} else {
				counts[v] = 0
			}
		}
	}

	log.Infof("Similarity: %v", ans)

	return ans
}
