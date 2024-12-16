package main

import (
	"image"
	"maps"
	"sort"
	"strings"

	"github.com/blebon/AoC/2024/util"
	log "github.com/sirupsen/logrus"
)

var DEBUG bool = false

type Deer struct {
	Position  image.Point
	Direction image.Point
}

type Maze struct {
	Board     map[image.Point]rune
	Deer      Deer
	Goal      image.Point
	MinScore  int
	InMinPath map[image.Point]bool
	Height    int
	Width     int
}

type Path struct {
	Deer  Deer
	Moves int
	Turns int
	Tiles map[image.Point]bool
}

func getMaze(f string) Maze {
	l, err := util.ReadFile(f)
	util.FileError(err)

	board := map[image.Point]rune{}
	deer := Deer{Direction: image.Point{1, 0}}
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
		Deer:      deer,
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
	return 1000*p.Turns + p.Moves
}

func (m *Maze) searchShortestPath() {
	q := []Path{{
		Deer:  m.Deer,
		Moves: 0,
		Turns: 0,
		Tiles: map[image.Point]bool{m.Deer.Position: true},
	}}

	distance := map[Deer]int{}

	for len(q) > 0 {
		sort.Slice(q, func(i, j int) bool {
			return q[i].getScore() < q[j].getScore()
		})

		p := q[0]
		q = q[1:]

		score := p.getScore()

		v, ok := distance[p.Deer]
		if ok && v < score {
			continue
		}
		distance[p.Deer] = score

		if p.Deer.Position == m.Goal && score <= m.MinScore {
			m.MinScore = score
			maps.Copy(m.InMinPath, p.Tiles)
			continue
		}

		if m.Board[p.Deer.Position] == '#' {
			continue
		}

		for dir, addTurn := range map[image.Point]int{
			p.Deer.Direction: 0,
			{-p.Deer.Direction.Y, p.Deer.Direction.X}: 1,
			{p.Deer.Direction.Y, -p.Deer.Direction.X}: 1,
		} {
			np := p.Deer.Position.Add(dir)
			newTiles := maps.Clone(p.Tiles)
			newTiles[np] = true
			q = append(q, Path{
				Deer:  Deer{Position: np, Direction: dir},
				Moves: p.Moves + 1,
				Turns: p.Turns + addTurn,
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
