package main

import (
	"github.com/blebon/AoC/2024/util"
	log "github.com/sirupsen/logrus"
)

func readColumns(input string) ([][]int, error) {
	return util.ReadSpaceSeparatedFileToInt(input)
}

func isSafe(report []int) bool {
	log.Debugf("isSafe report: %v", report)

	if len(report) <= 1 {
		return true
	}

	var d int = report[1] - report[0]
	if d == 0 {
		return false
	}

	var increasing bool = d > 0
	for i := range len(report) - 1 {
		d = report[i+1] - report[i]
		if increasing {
			if d < 1 || d > 3 {
				return false
			}
		} else {
			if d < -3 || d > -1 {
				return false
			}
		}
	}

	return true
}

func isSafeWithDampener(report []int) bool {
	if isSafe(report) {
		return true
	}

	for i := range report {
		c := util.RemoveElementInSlice(report, i)
		if isSafe(c) {
			return true
		}
	}

	return false
}

type fn func([]int) bool

func countSafeReports(input string, f fn) int {
	reports, err := readColumns(input)
	if err != nil {
		log.Errorf("error reading columns: %v", err)
	}

	ans := 0
	for i, report := range reports {
		s := f(report)
		if s {
			ans += 1
		} else {
			log.Debugf("Unsafe report %d %v", i+1, report)
		}
	}

	return ans
}
