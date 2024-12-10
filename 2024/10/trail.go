package main

import (
	"image"
	"sync/atomic"

	"github.com/blebon/AoC/2024/util"
	log "github.com/sirupsen/logrus"
)

var directions []image.Point = []image.Point{
	{0, 1}, {1, 0}, {0, -1}, {-1, 0},
}

type search struct {
	Board      map[image.Point]rune
	Directions []image.Point
}

func getTrailCount(f string, rating bool) int {
	s := getBoard(f)
	return s.countTrails(rating)
}

func getBoard(f string) search {
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
		Board:      board,
		Directions: directions,
	}
}

func (s *search) countTrails(rating bool) int {
	ch := make(chan int64)
	go func() {
		defer close(ch)
		for p := range s.Board {
			if s.Board[p] == '0' {
				ch <- s.countTrailsFromPoint(p, rating)
			}
		}
	}()

	var ans atomic.Int64
	for v := range ch {
		ans.Add(v)
	}
	return int(ans.Load())
}

func (s *search) countTrailsFromPoint(p image.Point, rating bool) int64 {
	var ans int64 = 0

	q := []image.Point{p}
	visited := map[image.Point]bool{}
	for len(q) > 0 {
		p := q[0]
		q = q[1:]

		if s.Board[p] == '9' {
			ans++
			continue
		}

		for _, d := range s.Directions {
			nextPoint := p.Add(d)
			isGoodPath := s.Board[nextPoint] == s.Board[p]+1
			if isGoodPath {
				_, ok := visited[nextPoint]
				unvisited := !ok
				if unvisited || rating {
					q = append(q, nextPoint)
					visited[nextPoint] = true
				}
			}
		}
	}

	return ans
}
