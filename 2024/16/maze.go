package main

import (
	"image"
	"maps"
	"sort"
	"strings"

	"github.com/blebon/AoC/2024/util"
	log "github.com/sirupsen/logrus"
)

const (
	TURN int = 1000
	MOVE int = 1
)

var DEBUG bool = false

type Sprite struct {
	Position  image.Point
	Direction image.Point
}

type Maze struct {
	Board     map[image.Point]rune
	Start     Sprite
	Goal      image.Point
	MinScore  int
	InMinPath map[image.Point]bool
	Height    int
	Width     int
}

type Path struct {
	Deer  Sprite
	Score int
	Tiles map[image.Point]bool
}

func getMaze(f string) Maze {
	l, err := util.ReadFile(f)
	util.FileError(err)

	board := map[image.Point]rune{}
	deer := Sprite{Direction: image.Point{1, 0}}
	goal := image.Point{}

	height := len(l)
	width := 0
	for j, s := range l {
		if j == 0 {
			width = len(s)
		}
		for i, c := range s {
			p := image.Point{i, j}
			board[p] = c
			switch c {
			case 'S':
				deer.Position = p
			case 'E':
				goal = p
			}
		}
	}

	m := Maze{
		Board:     board,
		Start:     deer,
		Goal:      goal,
		MinScore:  int(^uint(0) >> 1),
		InMinPath: map[image.Point]bool{},
		Height:    height,
		Width:     width,
	}

	if DEBUG {
		log.Infof("\n%v", m.String())
	}

	return m
}

func (p *Path) getScore() int {
	return p.Score
}

func (m *Maze) searchShortestPath() {
	q := []Path{{
		Deer:  m.Start,
		Score: 0,
		Tiles: map[image.Point]bool{m.Start.Position: true},
	}}

	distance := map[Sprite]int{}

	for len(q) > 0 {
		sort.Slice(q, func(i, j int) bool {
			return q[i].getScore() < q[j].getScore()
		})

		p := q[0]
		q = q[1:]

		v, ok := distance[p.Deer]
		if ok && v < p.getScore() {
			continue
		}
		distance[p.Deer] = p.getScore()

		if p.Deer.Position == m.Goal && p.getScore() <= m.MinScore {
			m.MinScore = p.getScore()
			maps.Copy(m.InMinPath, p.Tiles)
			continue
		}

		if m.Board[p.Deer.Position] == '#' {
			continue
		}

		for dir, addScore := range map[image.Point]int{
			p.Deer.Direction: MOVE,
			{-p.Deer.Direction.Y, p.Deer.Direction.X}: TURN + MOVE,
			{p.Deer.Direction.Y, -p.Deer.Direction.X}: TURN + MOVE,
		} {
			np := p.Deer.Position.Add(dir)
			newTiles := maps.Clone(p.Tiles)
			newTiles[np] = true
			q = append(q, Path{
				Deer:  Sprite{Position: np, Direction: dir},
				Score: p.getScore() + addScore,
				Tiles: newTiles,
			})
		}
	}
}

func (m *Maze) getScore() int {
	return m.MinScore
}

func (m *Maze) getTiles() int {
	return len(m.InMinPath)
}

func (m *Maze) String() string {
	s := strings.Builder{}
	for y := range m.Height {
		for x := range m.Width {
			s.WriteRune(m.Board[image.Point{x, y}])
		}
		s.WriteString("\n")
	}
	return s.String()
}
