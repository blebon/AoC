package main

import (
	"fmt"
	"image"
	"strings"

	"github.com/blebon/AoC/2024/util"
)

const (
	FREE_SPACE rune = '.'
	WALL       rune = '#'
	START      rune = 'S'
	GOAL       rune = 'E'
)

var DIRECTIONS []image.Point = []image.Point{
	image.Pt(0, 1), image.Pt(-1, 0), image.Pt(0, -1), image.Pt(1, 0),
}

type Cheat struct {
	Start image.Point
	End   image.Point
}

type Track struct {
	Board     map[image.Point]rune
	Start     image.Point
	Goal      image.Point
	Distances map[image.Point]int
	Height    int
	Width     int
}

func newTrack(f string) Track {
	l, err := util.ReadFile(f)
	util.FileError(err)

	board := map[image.Point]rune{}
	height := len(l)
	var width int
	var start, goal image.Point
	for y, s := range l {
		if y == 0 {
			width = len(s)
		}
		for x, c := range s {
			p := image.Point{x, y}
			board[p] = c
			switch c {
			case START:
				start = p
			case GOAL:
				goal = p
			}
		}
	}

	s := Track{
		Board:     board,
		Start:     start,
		Goal:      goal,
		Distances: map[image.Point]int{start: 0},
		Height:    height,
		Width:     width,
	}
	return s
}

type Sprite struct {
	Position  image.Point
	Direction image.Point
}

type Path struct {
	Sprite Sprite
}

func (t *Track) findDistancesFromStart() {
	q := []Path{}
	sp := Sprite{Position: t.Start}
	q = append(q, Path{
		Sprite: sp,
	})

	for len(q) > 0 {
		p := q[0]
		q = q[1:]

		for _, d := range DIRECTIONS {
			np := p.Sprite.Position.Add(d)
			_, visited := t.Distances[np]
			if !visited && t.Board[np] != WALL {
				t.Distances[np] = t.Distances[p.Sprite.Position] + 1
				q = append(q, Path{
					Sprite: Sprite{Position: np},
				})
			}
		}
	}
}

func (t *Track) countCheatsSavingAtLeast(n int, maxCheat int) int {
	var ans int = 0
	for p1, d1 := range t.Distances {
		for p2, d2 := range t.Distances {
			cheat := util.Abs(p2.X-p1.X) + util.Abs(p2.Y-p1.Y)
			if cheat <= maxCheat && d2 >= d1+cheat+n {
				ans++
			}
		}
	}
	return ans
}

func (t *Track) String() string {
	sb := strings.Builder{}
	for y := range t.Height {
		for x := range t.Width {
			sb.WriteRune(t.Board[image.Point{x, y}])
		}
		sb.WriteString("\n")
	}
	return sb.String()
}

func (t *Track) DistancesPlot() string {
	sb := strings.Builder{}
	for y := range t.Height {
		for x := range t.Width {
			s := fmt.Sprintf("  %s", string(WALL))
			d, ok := t.Distances[image.Point{x, y}]
			if ok {
				s = fmt.Sprintf("%3d", d)
			}
			sb.WriteString(s)
		}
		sb.WriteString("\n")
	}
	return sb.String()
}
