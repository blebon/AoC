package main

import (
	"image"
	"strings"

	"github.com/blebon/AoC/2024/util"
	"github.com/sirupsen/logrus"
)

var diagonals []image.Point = []image.Point{
	{-1, -1}, {1, -1}, {1, 1}, {-1, 1},
}

var directions []image.Point = []image.Point{
	{-1, -1}, {0, -1}, {1, -1}, {1, 0},
	{1, 1}, {0, 1}, {-1, 1}, {-1, 0},
}

type search struct {
	Board      map[image.Point]rune
	Directions []image.Point
}

func getBoard(f string) search {
	log := logrus.New()

	lines, err := util.ReadFile(f)
	if err != nil {
		log.Fatalf("error reading input file: %v", err)
	}

	board := map[image.Point]rune{}

	for j, l := range lines {
		for i, c := range l {
			board[image.Point{i, j}] = c
		}
	}

	return search{
		Board: board,
	}
}

// getBoardWords gets all word permutations of length l
// at point p.
func (s *search) getBoardWords(p image.Point, l int) []string {
	words := make([]string, len(s.Directions))
	for i, d := range s.Directions {
		for k := range l {
			words[i] += string(s.Board[p.Add(d.Mul(k))])
		}
	}
	return words
}

func countXmas(f string) int {
	return countPattern(f, "XMAS")
}

func countPattern(f, pattern string) int {
	log := logrus.New()

	s := getBoard(f)
	s.Directions = directions

	var ans int = 0
	for p := range s.Board {
		pointWords := strings.Join(s.getBoardWords(p, len(pattern)), " ")
		log.Debugf("Point %v Words %v", p, pointWords)
		ans += strings.Count(pointWords, pattern)
	}

	return ans
}

func countX_Mas(f string) int {
	log := logrus.New()

	s := getBoard(f)
	s.Directions = diagonals

	var ans int = 0
	for p := range s.Board {
		if s.Board[p] == 'A' {
			pointSubStr := strings.Join(s.getBoardWords(p, 2), "")
			log.Debugf("Point %v Diagonal sub string %v", p, pointSubStr)
			if strings.Contains("ASASAMAMASASAM", pointSubStr) {
				ans += 1
			}
		}
	}

	return ans
}
