package main

import (
	"image"
	"strings"

	"github.com/blebon/AoC/2024/util"
	log "github.com/sirupsen/logrus"
)

func getGPSSum(f string, scaleUp bool) int {
	w := getWarehouse(f, scaleUp)
	w.move()
	return w.getGPSSum()
}

type Warehouse struct {
	Board   map[image.Point]rune
	Moves   []image.Point
	Robot   image.Point
	Width   int
	Height  int
	ScaleUp bool
}

func getWarehouse(f string, scaleUp bool) Warehouse {
	l, err := util.ReadFile(f)
	util.FileError(err)

	board := map[image.Point]rune{}
	robot := image.Point{}

	var width int = 0
	var height int = 0
	for j, s := range l {
		if j == 0 {
			if scaleUp {
				width = 2 * len(s)
			} else {
				width = len(s)
			}
		}
		if s == "" {
			height = j
			break
		}
		for i, c := range s {
			if scaleUp {
				p1 := image.Point{2 * i, j}
				p2 := image.Point{2*i + 1, j}
				switch c {
				case '.', '#':
					board[p1] = c
					board[p2] = c
				case 'O':
					board[p1] = '['
					board[p2] = ']'
				case '@':
					board[p1] = c
					robot = p1
					board[p2] = '.'
				}
			} else {
				p := image.Point{i, j}
				board[p] = c
				if c == '@' {
					robot = p
				}
			}
		}
	}

	moves := []image.Point{}
	for _, s := range l[height+1:] {
		for _, c := range s {
			switch c {
			case '^':
				moves = append(moves, image.Point{0, -1})
			case '<':
				moves = append(moves, image.Point{-1, 0})
			case 'v':
				moves = append(moves, image.Point{0, 1})
			case '>':
				moves = append(moves, image.Point{1, 0})
			}
		}
	}

	log.Debugf("%v %v", board, moves)
	log.Debugf("%d %d", width, height)

	return Warehouse{
		Board:   board,
		Moves:   moves,
		Robot:   robot,
		Width:   width,
		Height:  height,
		ScaleUp: scaleUp,
	}
}

func (w *Warehouse) move() {
	for _, d := range w.Moves {
		// log.Infof("\n%v", w.String())
		np := w.Robot.Add(d)
		c := w.Board[np]
		switch c {
		case '.':
			w.moveRobot(np)
		case '#':
			continue
		case 'O':
			w.pushBlocks(np, d)
		case '[', ']':
			w.pushCrates(d)
		}
	}
}

func (w *Warehouse) moveRobot(np image.Point) {
	w.Board[w.Robot] = '.'
	w.Board[np] = '@'
	w.Robot = np
}

func (w *Warehouse) pushBlocks(np, d image.Point) {
	for i := range max(w.Height, w.Width) {
		next := np.Add(d.Mul(i + 1))
		c, ok := w.Board[next]
		if !ok || c == '#' {
			return
		} else if c == 'O' {
			continue
		} else if c == '.' {
			w.moveRobot(np)
			w.Board[next] = 'O'
			return
		} else {
			return
		}
	}
}

func (w *Warehouse) pushCrates(d image.Point) {
	push := w.getCratesAndRobot(d)

	if push != nil {
		for c := range push {
			w.Board[c] = '.'
		}
		for c := range push {
			w.Board[c.Add(d)] = push[c]
		}
		w.Robot = w.Robot.Add(d)
	}
}

func (w *Warehouse) getCratesAndRobot(d image.Point) map[image.Point]rune {
	q := []image.Point{w.Robot}
	crates := map[image.Point]rune{}
	for len(q) > 0 {
		p := q[0]
		q = q[1:]

		_, ok := crates[p]
		if ok {
			continue
		}
		crates[p] = w.Board[p]

		np := p.Add(d)
		switch w.Board[np] {
		case '#':
			return nil
		case '[':
			q = append(q, np)
			q = append(q, np.Add(image.Point{1, 0}))
		case ']':
			q = append(q, np)
			q = append(q, np.Add(image.Point{-1, 0}))
		}
	}
	return crates
}

func (w *Warehouse) getGPSSum() int {
	var ans int = 0
	for p, c := range w.Board {
		if c == 'O' || c == '[' {
			ans += 100*p.Y + p.X
		}
	}
	return ans
}

func (w *Warehouse) String() string {
	s := strings.Builder{}
	for y := range w.Height {
		for x := range w.Width {
			s.WriteRune(w.Board[image.Point{x, y}])
		}
		s.WriteString("\n")
	}
	return s.String()
}
