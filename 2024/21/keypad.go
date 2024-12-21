package main

import (
	"fmt"
	"image"
	"strings"

	"github.com/blebon/AoC/2024/util"
)

var numericKeypad map[rune]image.Point = map[rune]image.Point{
	'7': {0, 3}, '8': {1, 3}, '9': {2, 3},
	'4': {0, 2}, '5': {1, 2}, '6': {2, 2},
	'1': {0, 1}, '2': {1, 1}, '3': {2, 1},
	' ': {0, 0}, '0': {1, 0}, 'A': {2, 0},
}

var directionalKeypad map[rune]image.Point = map[rune]image.Point{
	' ': {1, 0}, '^': {1, 1}, 'A': {2, 1},
	'<': {0, 0}, 'v': {1, 0}, '>': {2, 0},
}

const (
	NUMERIC     string = "Numeric"
	DIRECTIONAL string = "Directional"
)

type Keypad struct {
	Layout    map[rune]image.Point
	Start     rune
	Name      string
	PathCache map[string]string
	SeqCache  map[string]string
}

var NumericKeypad Keypad = Keypad{
	Layout:    numericKeypad,
	Start:     'A',
	Name:      NUMERIC,
	PathCache: map[string]string{},
	SeqCache:  map[string]string{},
}

var DirectionalKeypad Keypad = Keypad{
	Layout:    directionalKeypad,
	Start:     'A',
	Name:      DIRECTIONAL,
	PathCache: map[string]string{},
	SeqCache:  map[string]string{},
}

func getComplexitySum(f string, robots int) int {
	codes, err := util.ReadFile(f)
	util.FileError(err)

	ans := 0
	for _, code := range codes {
		ans += getComplexity(code, robots)
	}
	return ans
}

func getComplexity(code string, robots int) int {
	var num int
	fmt.Sscanf(code, "%dA", &num)
	seqLen := getSequenceLength(code, robots)
	return num * seqLen
}

func getSequenceLength(code string, robots int) int {
	seq := NumericKeypad.generateSequence(code)
	c := len(seq)
	if robots > 0 {
		numericSeqCache := map[string][]int{}
		c = getNumericSequenceLength(seq, robots, 1, numericSeqCache)
	}
	return c
}

func getNumericSequenceLength(code string, robots int, robot int, numericSeqCache map[string][]int) int {
	ans, ok := numericSeqCache[code]
	if ok && robot <= len(ans) && ans[robot-1] != 0 {
		return ans[robot-1]
	}
	if !ok {
		numericSeqCache[code] = make([]int, robots)
	}

	seq := DirectionalKeypad.generateSequence(code)
	if robot == robots {
		return len(seq)
	}

	seqs := splitSequence(seq)
	l := 0
	for _, seq := range seqs {
		c := getNumericSequenceLength(seq, robots, robot+1, numericSeqCache)
		l += c
	}

	numericSeqCache[code][robot-1] = l
	return l
}

func splitSequence(seq string) []string {
	var seqs []string
	var s string

	for _, c := range seq {
		s += string(c)
		if c == 'A' {
			seqs = append(seqs, s)
			s = ""
		}
	}

	return seqs
}

func (k *Keypad) generateSequence(code string) string {
	if ans, ok := k.SeqCache[code]; ok {
		return ans
	}

	var sequence []string
	start := k.Start
	for _, target := range code {
		path := k.getPath(target, start)
		sequence = append(sequence, path)
		start = target
	}

	ans := strings.Join(sequence, "")
	k.SeqCache[code] = ans
	return ans
}

func (k *Keypad) getPath(target rune, start rune) string {
	cacheKey := fmt.Sprintf("%v-%v", string(start), string(target))
	if path, ok := k.PathCache[cacheKey]; ok {
		return path
	}

	paths := k.getDirections(start, target)
	path := strings.Join(paths, "")
	path += "A"

	k.PathCache[cacheKey] = path

	return path
}

func (k *Keypad) getDirections(start, target rune) []string {
	from := k.Layout[start]
	to := k.Layout[target]
	delta := to.Sub(from)

	paths := []string{}

	switch k.Name {
	case NUMERIC:
		if from.Y == 0 && to.X == 0 {
			verticalFirst(&paths, delta)
		} else if from.X == 0 && to.Y == 0 {
			horizontalFirst(&paths, delta)
		} else if delta.X < 0 {
			horizontalFirst(&paths, delta)
		} else {
			verticalFirst(&paths, delta)
		}
	case DIRECTIONAL:
		if from.X == 0 && to.Y == 1 {
			horizontalFirst(&paths, delta)
		} else if from.Y == 1 && to.X == 0 {
			verticalFirst(&paths, delta)
		} else if delta.X < 0 {
			horizontalFirst(&paths, delta)
		} else {
			verticalFirst(&paths, delta)
		}
	}

	return paths
}

func verticalFirst(paths *[]string, delta image.Point) {
	addVertical(paths, delta)
	addHorizontal(paths, delta)
}

func horizontalFirst(paths *[]string, delta image.Point) {
	addHorizontal(paths, delta)
	addVertical(paths, delta)
}

func addHorizontal(paths *[]string, delta image.Point) {
	for range util.Abs(delta.X) {
		if delta.X > 0 {
			*paths = append(*paths, ">")
		} else {
			*paths = append(*paths, "<")
		}
	}
}

func addVertical(paths *[]string, delta image.Point) {
	for range util.Abs(delta.Y) {
		if delta.Y > 0 {
			*paths = append(*paths, "^")
		} else {
			*paths = append(*paths, "v")
		}
	}
}
