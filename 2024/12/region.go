package main

import (
	"fmt"
	"image"

	"github.com/blebon/AoC/2024/util"
	log "github.com/sirupsen/logrus"
)

var directions []image.Point = []image.Point{
	{0, 1}, {1, 0}, {0, -1}, {-1, 0},
}

type region struct {
	r         rune
	area      int
	perimeter int
	sides     int
}

type regions []*region

type search struct {
	Board      map[image.Point]rune
	Directions []image.Point
	Visited    map[image.Point]bool
	Regions    regions
}

func getPrice(f string, sides bool) int {
	s := getBoard(f)
	s.findRegions()
	return s.getCost(sides)
}

func getBoard(f string) search {
	lines, err := util.ReadFile(f)
	if err != nil {
		log.Fatalf("error reading input file: %v", err)
	}

	board := map[image.Point]rune{}

	for j, l := range lines {
		for i, r := range l {
			board[image.Point{i, j}] = r
		}
	}

	return search{
		Board:      board,
		Directions: directions,
		Visited:    map[image.Point]bool{},
		Regions:    regions{},
	}
}

func (s *search) findRegions() {
	for p := range s.Board {
		s.addRegionFromPoint(p)
	}
}

func (s *search) addRegionFromPoint(p image.Point) {
	if s.Visited[p] {
		return
	}
	s.Visited[p] = true

	reg := region{
		r:         s.Board[p],
		area:      1,
		perimeter: 0,
		sides:     0,
	}

	q := []image.Point{p}
	for len(q) > 0 {
		p := q[0]
		q = q[1:]

		for _, d := range s.Directions {
			nextPoint := p.Add(d)
			if s.Board[nextPoint] != reg.r {
				reg.perimeter += 1
				rot := p.Add(image.Point{d.Y, d.X})
				if s.Board[rot] != reg.r || s.Board[rot.Add(d)] == reg.r {
					reg.sides += 1
				}
			} else if !s.Visited[nextPoint] {
				reg.area += 1
				s.Visited[nextPoint] = true
				q = append(q, nextPoint)
			}
		}
	}

	s.Regions = append(s.Regions, &reg)
}

func (s *search) getCost(sides bool) int {
	log.Debugf("regions %v", s.Regions.print())
	var ans int = 0
	for _, r := range s.Regions {
		if sides {
			ans += r.area * r.sides
		} else {
			ans += r.area * r.perimeter
		}
	}
	return ans
}

func (r *region) String() string {
	return fmt.Sprintf("{%v %v %v %v}", string(r.r), r.area, r.perimeter, r.sides)
}

func (r regions) print() string {
	return fmt.Sprintf("%v", r)
}
