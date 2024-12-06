package main

import (
	"image"
	"sync/atomic"

	"github.com/blebon/AoC/2024/util"
	"github.com/schollz/progressbar/v3"
	"github.com/sirupsen/logrus"
)

var directions []image.Point = []image.Point{
	{0, -1}, {1, 0}, {0, 1}, {-1, 0},
}

type walk struct {
	File              string
	Board             map[image.Point]rune
	DirectionIndex    int
	Direction         image.Point
	VisitedDirections map[image.Point][]int
	X                 int
	Y                 int
	Init              image.Point
	Width             int
	Height            int
	Count             int
	// Steps             int
}

func getWalk(f string) walk {
	log := logrus.New()

	lines, err := util.ReadFile(f)
	if err != nil {
		log.Fatalf("error reading input file: %v", err)
	}

	w := walk{
		File:  f,
		Count: 0,
		// Steps: 0,
	}
	board := map[image.Point]rune{}
	visitedDir := map[image.Point][]int{}

	w.Height = len(lines)
	for j, l := range lines {
		if j == 0 {
			w.Width = len(l)
		}
		for i, c := range l {
			board[image.Point{i, j}] = c
			switch c {
			case '^':
				w.X = i
				w.Y = j
				w.Init = image.Point{i, j}
				board[w.Init] = 'X'
				w.Count += 1
				w.DirectionIndex = 0
				w.Direction = directions[w.DirectionIndex]
				visitedDir[w.Init] = append(visitedDir[w.Init], w.DirectionIndex)
			case '>':
				w.X = i
				w.Y = j
				w.Init = image.Point{i, j}
				board[w.Init] = 'X'
				w.Count += 1
				w.DirectionIndex = 1
				w.Direction = directions[w.DirectionIndex]
				visitedDir[w.Init] = append(visitedDir[w.Init], w.DirectionIndex)
			case 'v':
				w.X = i
				w.Y = j
				w.Init = image.Point{i, j}
				board[w.Init] = 'X'
				w.Count += 1
				w.DirectionIndex = 2
				w.Direction = directions[w.DirectionIndex]
				visitedDir[w.Init] = append(visitedDir[w.Init], w.DirectionIndex)
			case '<':
				w.X = i
				w.Y = j
				w.Init = image.Point{i, j}
				board[w.Init] = 'X'
				w.Count += 1
				w.DirectionIndex = 3
				w.Direction = directions[w.DirectionIndex]
				visitedDir[w.Init] = append(visitedDir[w.Init], w.DirectionIndex)
			}
		}
	}

	w.Board = board
	w.VisitedDirections = visitedDir
	if w.Count != 1 {
		log.Fatalf("wrong count: want 1, got %v", w.Count)
	}

	return w
}

// patrol walks the board. Returns true if a loop is detected.
// The loop is detected if a square marked with X has been
// previously visited in the same direction.
func (w *walk) patrol() bool {
	// log := logrus.New()
	for {
		p := image.Point{w.X, w.Y}
		p = p.Add(w.Direction)
		x := p.X
		y := p.Y
		// log.Debugf("At point %v, steps %d, count %d, rune %v, dir %v", p, w.Steps, w.Count, string(w.Board[p]), w.Direction)

		if x < 0 || x >= w.Width || y < 0 || y >= w.Height {
			return false
		} else if w.Board[p] == 'X' {
			w.X = x
			w.Y = y
			for i := range w.VisitedDirections[p] {
				if w.VisitedDirections[p][i] == w.DirectionIndex {
					return true
				}
			}
		} else if w.Board[p] == '.' {
			w.Board[p] = 'X'
			w.VisitedDirections[p] = append(w.VisitedDirections[p], w.DirectionIndex)
			w.Count += 1
			w.X = x
			w.Y = y
		} else if w.Board[p] == '#' {
			w.DirectionIndex = (w.DirectionIndex + 1) % 4
			w.Direction = directions[w.DirectionIndex]
			w.VisitedDirections[p] = append(w.VisitedDirections[p], w.DirectionIndex)
		}
	}
}

func getGuardPositions(f string) int {
	log := logrus.New()
	w := getWalk(f)
	log.Debugf("walk board obtained: %v", w.Board)
	w.patrol()
	return w.Count
}

func (w *walk) getObstaclePatrol(p *image.Point) bool {
	nw := getWalk(w.File)
	nw.Board[*p] = '#'
	return nw.patrol()
}

func getObstaclePositions(f string) int {
	// log := logrus.New()
	w := getWalk(f)
	// log.Debugf("walk board obtained: %v", w.Board)
	w.patrol()

	ch := make(chan *image.Point, 32)
	go func() {
		for p := range w.Board {
			if w.Board[p] == 'X' && p != w.Init {
				ch <- &p
			}
		}
		close(ch)
	}()

	var ans atomic.Int32
	bar := progressbar.Default(int64(w.Count-1), "finding unending loops")
	for p := range ch {
		if w.getObstaclePatrol(p) {
			ans.Add(1)
		}
		bar.Add(1)
		// log.Debugf("received %v, ans %v", p, ans.Load())
	}

	return int(ans.Load())
}
