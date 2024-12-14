package main

import (
	"fmt"
	"image"
	"image/color"
	"image/png"
	"os"

	"github.com/blebon/AoC/2024/util"
	log "github.com/sirupsen/logrus"
)

const (
	ROBOT_FMT = "p=%d,%d v=%d,%d"
)

var (
	tiles_x int = 101
	tiles_y int = 103
)

func getSafetyNumber(f string, sec int) int {
	r := getRobots(f)
	r.move(sec)
	return r.safetyNumber()
}

type Robot struct {
	Position image.Point
	Velocity image.Point
	Quadrant image.Point
}

type Robots []*Robot
type Quadrants map[image.Point]int

func getRobots(f string) Robots {
	l, err := util.ReadFile(f)
	if err != nil {
		log.Fatalf("error reading input file: %v", err)
	}
	robots := make(Robots, 0, len(l))
	for _, s := range l {
		p, v, q := image.Point{}, image.Point{}, image.Point{}
		fmt.Sscanf(s, ROBOT_FMT, &p.X, &p.Y, &v.X, &v.Y)
		r := &Robot{Position: p, Velocity: v, Quadrant: q}
		log.Debugf("%v", *r)
		robots = append(robots, r)
	}
	return robots
}

func (robots Robots) move(sec int) {
	area := image.Rectangle{image.Point{0, 0}, image.Point{tiles_x, tiles_y}}
	for _, r := range robots {
		r.Position = r.Position.Add(r.Velocity.Mul(sec)).Mod(area)
		r.checkQuadrant(area)
	}
}

func (r *Robot) checkQuadrant(area image.Rectangle) {
	getCoord := func(x int) int {
		if x < 0 {
			return 0
		} else if x > 0 {
			return 1
		}
		return -1
	}
	r.Quadrant = image.Point{getCoord(r.Position.X - area.Dx()/2), getCoord(r.Position.Y - area.Dy()/2)}
}

func (robots Robots) safetyNumber() int {
	q := Quadrants{}
	for _, r := range robots {
		log.Debugf("%v", *r)
		q[r.Quadrant] += 1
	}
	log.Debugf("%v", q)
	return q[image.Point{0, 0}] * q[image.Point{0, 1}] * q[image.Point{1, 0}] * q[image.Point{1, 1}]
}

func getXmasTree(f string, print bool) int {
	robots := getRobots(f)
	area := image.Rectangle{image.Point{0, 0}, image.Point{tiles_x, tiles_y}}
	i := 0
	for {
		i++
		visited := map[image.Point]bool{}
		for _, r := range robots {
			r.Position = r.Position.Add(r.Velocity).Mod(area)
			visited[r.Position] = true
		}
		if len(visited) == len(robots) {
			break
		}
	}
	if print {
		robots.writePng()
	}
	return i
}

func (robots Robots) writePng() {
	img := image.NewRGBA(image.Rect(0, 0, tiles_x, tiles_y))
	for y := range tiles_y {
		for x := range tiles_x {
			img.Set(x, y, color.Black)
		}
	}
	for _, r := range robots {
		img.Set(r.Position.X, r.Position.Y, color.White)
	}

	f, err := os.Create("XMas.png")
	if err != nil {
		log.Warnf("error creating image file: %v", err)
	}
	err = png.Encode(f, img)
	if err != nil {
		log.Warnf("error encoding image file: %v", err)
	}
}
