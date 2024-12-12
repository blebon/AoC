package main

import (
	"slices"
	"strings"
	"sync"

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

type runes []rune

func calculateChecksum(blocks []rune, noFrag bool) int64 {
	a := append(make(runes, 0, len(blocks)), blocks...)
	if noFrag {
		a.compressNoFrag()
	} else {
		a.compress()
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

func (blocks runes) compress() {
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

type space struct {
	i    int
	size int
}

type spaces []*space

func (a runes) compressNoFrag() {
	s := a.findSpaces()

	var wg sync.WaitGroup
	for i, size := len(a)-1, 0; i >= 0 && len(s) > 0 && s[0].i < i; i-- {
		i, size = a.findLastBlockSize(i)
		if size > 0 {
			j := s.findFirstSpaceBlock(size, i)
			if j >= 0 {
				wg.Add(1)
				log.Debugf("Sending rune #%v %v with size %d j %d and i %d", int(a[i]-'0'), string(a[i]), size, j, i)
				go a.swapBlock(i, j, size, &wg)
			}
		}
	}
	log.Debugf("Waiting for all swaps")
	wg.Wait()
	log.Debugf("Swaps complete")
}

func (a runes) findSpaces() spaces {
	r := make(spaces, 0)
	spaceIndex := -1
	spaceSize := 0
	for j := 0; j < len(a); j++ {
		if a[j] != '.' {
			if spaceIndex >= 0 {
				r = append(r, &space{i: spaceIndex, size: spaceSize})
			}
			spaceIndex = -1
			spaceSize = 0
		} else {
			if spaceIndex == -1 {
				spaceIndex = j
			}
			spaceSize += 1
		}
	}
	return r
}

func (a runes) findLastBlockSize(i int) (int, int) {
	size := 0
out:
	for ; i > 0; i-- {
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

func (s *spaces) findFirstSpaceBlock(size, i int) int {
	for j, space := range *s {
		if space.i > i {
			return -1
		}
		if space.size >= size {
			space.size -= size
			ans := space.i
			space.i += size
			if space.size == 0 {
				*s = slices.Delete(*s, j, j+1)
			}
			return ans
		}
	}
	return -1
}

func (a runes) swapBlock(i, j, size int, wg *sync.WaitGroup) {
	jm := j + size
	for ; j < jm; i, j = i+1, j+1 {
		a[j], a[i] = a[i], a[j]
	}
	// log.Debugf("Updated slice by moving [%2d:%2d] to [%2d:%2d]: %v", i, i+size, j, j+size, a.print())
	wg.Done()
}

func (a runes) print() string {
	var s string
	for _, c := range a {
		s += string(c)
	}
	return s
}
