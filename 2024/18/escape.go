package main

import (
	"fmt"
	"image"
	"sort"
	"strings"

	"github.com/blebon/AoC/2024/util"
)

const (
	FREE_SPACE  rune = '.'
	FALLEN_BYTE rune = '#'
)

var DIRECTIONS []image.Point = []image.Point{
	image.Pt(0, 1), image.Pt(-1, 0), image.Pt(0, -1), image.Pt(1, 0),
}

type Space struct {
	Board             map[image.Point]rune
	Falls             []image.Point
	Start             image.Point
	Goal              image.Point
	Steps             int
	FirstBlockingByte string
}

func newSpace(f string, goalOrdinate int) Space {
	l, err := util.ReadFile(f)
	util.FileError(err)

	board := map[image.Point]rune{}
	for y := range goalOrdinate + 1 {
		for x := range goalOrdinate + 1 {
			board[image.Point{x, y}] = FREE_SPACE
		}
	}

	falls := []image.Point{}
	for _, s := range l {
		var x, y int
		fmt.Sscanf(s, "%d,%d", &x, &y)
		falls = append(falls, image.Point{x, y})
	}

	s := Space{
		Board: board,
		Falls: falls,
		Start: image.Pt(0, 0),
		Goal:  image.Pt(goalOrdinate, goalOrdinate),
		Steps: 0,
	}
	return s
}

func (s *Space) fall(n int) {
	for i := range n {
		s.Board[s.Falls[i]] = FALLEN_BYTE
	}
}

type Sprite struct {
	Position  image.Point
	Direction image.Point
}

type Path struct {
	Sprite Sprite
	Steps  int
}

func (s *Space) move() {
	q := []Path{
		{
			Sprite: Sprite{Position: s.Start, Direction: DIRECTIONS[0]},
			Steps:  0,
		},
		{
			Sprite: Sprite{Position: s.Start, Direction: DIRECTIONS[3]},
			Steps:  0,
		},
	}

	visited := map[Sprite]bool{}
	for i := range DIRECTIONS {
		visited[Sprite{Position: s.Start, Direction: DIRECTIONS[i]}] = true
	}

	for len(q) > 0 {
		sort.Slice(q, func(i, j int) bool {
			return q[i].Steps < q[j].Steps
		})

		p := q[0]
		q = q[1:]

		if p.Sprite.Position == s.Goal {
			s.Steps = p.Steps
			break
		}

		for _, d := range DIRECTIONS {
			np := p.Sprite.Position.Add(d)
			if s.Board[np] == FREE_SPACE {
				sp := Sprite{Position: np, Direction: d}
				_, ok := visited[sp]
				if ok {
					continue
				}
				visited[sp] = true
				q = append(q, Path{
					Sprite: sp,
					Steps:  p.Steps + 1,
				})
			}
		}
	}
}

func (s *Space) findFirstBlockByteAfter(n int) {
	for {
		n++
		s.Steps = 0
		s.fall(n)
		s.move()
		if s.Steps == 0 {
			b := s.Falls[n-1]
			s.FirstBlockingByte = fmt.Sprintf("%d,%d", b.X, b.Y)
			break
		}
	}
}

func (s *Space) String() string {
	sb := strings.Builder{}
	for y := range s.Goal.Y + 1 {
		for x := range s.Goal.X + 1 {
			sb.WriteRune(s.Board[image.Point{x, y}])
		}
		sb.WriteString("\n")
	}
	return sb.String()
}
