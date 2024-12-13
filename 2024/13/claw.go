package main

import (
	"fmt"
	"image"
	"os"
	"strings"

	log "github.com/sirupsen/logrus"
)

const (
	A_COST      int    = 3
	B_COST      int    = 1
	MACHINE_FMT string = "Button A: X+%d, Y+%d\nButton B: X+%d, Y+%d\nPrize: X=%d, Y=%d"
	OFFSET      int    = 10000000000000
)

func countTokens(f string, offset bool) int {
	eqns := getEquations(f)
	r := 0
	for _, e := range eqns {
		if offset {
			e.b = e.b.Add(image.Point{OFFSET, OFFSET})
		}
		r += e.eval()
	}
	return r
}

type Equation struct {
	A image.Point
	B image.Point
	b image.Point
}

func getEquations(f string) []Equation {
	strEqns := getStrEquations(f)
	eqns := make([]Equation, 0, len(strEqns))
	for _, s := range strEqns {
		e := newEquation(s)
		eqns = append(eqns, e)
	}
	return eqns
}

func getStrEquations(f string) []string {
	records, err := os.ReadFile(f)
	if err != nil {
		log.Fatalf("error reading file: %v", err)
	}
	return strings.Split(string(records), "\n\n")
}

func newEquation(s string) Equation {
	var A, B, b image.Point
	fmt.Sscanf(s, MACHINE_FMT, &A.X, &A.Y, &B.X, &B.Y, &b.X, &b.Y)
	return Equation{A: A, B: B, b: b}
}

func (e *Equation) eval() int {
	det := e.A.X*e.B.Y - e.A.Y*e.B.X
	if det == 0 {
		return 0
	}
	a := (e.B.Y*e.b.X - e.B.X*e.b.Y) / det
	b := (e.A.X*e.b.Y - e.A.Y*e.b.X) / det
	if e.A.Mul(a).Add(e.B.Mul(b)) == e.b {
		// log.Debugf("%v %d %d %d", e, a, b, a*A_COST+b*B_COST)
		return a*A_COST + b*B_COST
	}
	return 0
}
