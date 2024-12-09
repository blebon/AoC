package main

import (
	"strings"

	"github.com/blebon/AoC/2024/util"
	log "github.com/sirupsen/logrus"
)

func getChecksum(f string, noFrag bool) int64 {
	blocks := getBlocks(f)
	return calculateChecksum(blocks, noFrag)
}

func getBlocks(f string) []rune {
	m := getDiskmap(f)
	s := 1
	r := make([]rune, 0, len(m))
	for i, c := range m {
		b := rune(i/2) + '0'
		if s < 0 {
			b = '.'
		}
		for range int(c - '0') {
			r = append(r, b)
		}
		s *= -1
	}
	return r
}

func getDiskmap(f string) string {
	l, err := util.ReadLine(f)
	if err != nil {
		log.Fatalf("error reading file: %v", err)
	}
	return strings.TrimSpace(l)
}

func calculateChecksum(blocks []rune, noFrag bool) int64 {
	a := append(make([]rune, 0, len(blocks)), blocks...)
	if noFrag {
		compressNoFrag(a)
	} else {
		compress(a)
	}
	r := int64(0)
	for i, d := range a {
		if d == '.' {
			continue
		}
		r += int64(i * int(d-'0'))
	}
	return r
}

func compress(blocks []rune) {
	i := 0
	j := len(blocks) - 1

	for i < j {
		if blocks[i] == '.' {
			for j > i && blocks[j] == '.' {
				j--
			}
			if i < j {
				blocks[i], blocks[j] = blocks[j], blocks[i]
			}
		}
		if blocks[i] != 46 {
			i++
		}
	}
}

func compressNoFrag(a []rune) {
	for i := len(a) - 1; i >= 0; i-- {
		k, size := findLastBlockSize(a, i)
		i = max(k, 0)
		j := findFirstSpaceBlock(a, size, i)
		if j >= 0 {
			swapBlock(a, i, j, size)
			log.Debugf("Updated slice by moving [%2d:%2d] to [%2d:%2d]: %v", i, i+size, j, j+size, print(a))
		}

	}
}

func findLastBlockSize(a []rune, i int) (int, int) {
	size := 0
out:
	for ; i >= 0; i-- {
		if a[i] != '.' {
			c := a[i]
			for j := i; j >= 0; j-- {
				if a[j] != c {
					break out
				}
				size++
			}
		}
	}
	return i - size + 1, size
}

func findFirstSpaceBlock(a []rune, size, i int) int {
	if size <= 0 {
		return -1
	}
	spaceSize := 0
	for j := 0; j < i; j++ {
		if a[j] != '.' {
			spaceSize = 0
		} else {
			if spaceSize == size-1 {
				return j - spaceSize
			}
			spaceSize += 1
		}
	}
	return -1
}

func swapBlock(a []rune, i, j, size int) {
	c := a[i]
	for k := j; k < j+size; k++ {
		a[k] = c
	}
	for k := i; k < i+size; k++ {
		a[k] = '.'
	}
}

func print(a []rune) string {
	var s string
	for _, c := range a {
		s += string(c)
	}
	return s
}
