package main

import (
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
	"strings"

	"github.com/blebon/AoC/2024/util"
)

const (
	DEBUGGER_FMT string = `Register A: %d
Register B: %d
Register C: %d

Program: %s`
)

type Debugger struct {
	A                  int
	B                  int
	C                  int
	Program            string
	ProgramSlice       []string
	Instructions       []string
	Operands           []int
	InstructionPointer int
	Output             []string
}

func newDebugger(f string) Debugger {
	l, err := os.ReadFile(f)
	util.FileError(err)

	s := string(l)
	var A, B, C int
	var p string
	fmt.Sscanf(s, DEBUGGER_FMT, &A, &B, &C, &p)

	sl := strings.Split(p, ",")
	i := make([]string, 0, len(sl)/2)
	op := make([]int, 0, len(sl)/2)
	isI := true
	for _, s := range sl {
		if isI {
			i = append(i, s)
		} else {
			o, err := strconv.Atoi(s)
			if err != nil {
				log.Fatal(err)
			}
			op = append(op, o)
		}
		isI = !isI
	}

	return Debugger{
		A:                  A,
		B:                  B,
		C:                  C,
		Program:            p,
		ProgramSlice:       sl,
		Instructions:       i,
		Operands:           op,
		InstructionPointer: 0,
	}
}

func (d *Debugger) getComboOperandValue(op int) int {
	switch op {
	case 0, 1, 2, 3:
		return op
	case 4:
		return d.A
	case 5:
		return d.B
	case 6:
		return d.C
	case 7:
		fallthrough
	default:
		log.Fatalf("invalid operand %d", op)
	}
	return -1
}

func (d *Debugger) adv(op int) {
	o := d.getComboOperandValue(op)
	num := d.A
	den := int(math.Pow(2, float64(o)))
	d.A = num / den
	d.InstructionPointer += 1
}

func (d *Debugger) bxl(op int) {
	d.B = d.B ^ op
	d.InstructionPointer += 1
}

func (d *Debugger) bst(op int) {
	o := d.getComboOperandValue(op)
	d.B = o % 8
	d.InstructionPointer += 1
}

func (d *Debugger) jnz(op int) {
	if d.A == 0 {
		d.InstructionPointer += 1
	} else {
		d.InstructionPointer = op
	}
}

func (d *Debugger) bxc(op int) {
	d.B = d.B ^ d.C
	d.InstructionPointer += 1
}

func (d *Debugger) out(op int) {
	r := d.getComboOperandValue(op) % 8
	s := fmt.Sprintf("%d", r)
	d.Output = append(d.Output, strings.Split(s, "")...)
	d.InstructionPointer += 1
}

func (d *Debugger) bdv(op int) {
	o := d.getComboOperandValue(op)
	num := d.A
	den := int(math.Pow(2, float64(o)))
	d.B = num / den
	d.InstructionPointer += 1
}

func (d *Debugger) cdv(op int) {
	o := d.getComboOperandValue(op)
	num := d.A
	den := int(math.Pow(2, float64(o)))
	d.C = num / den
	d.InstructionPointer += 1
}

func (d *Debugger) run() {
	d.InstructionPointer = 0
	for d.InstructionPointer < len(d.Instructions) {
		op := d.Operands[d.InstructionPointer]
		switch d.Instructions[d.InstructionPointer] {
		case "0":
			d.adv(op)
		case "1":
			d.bxl(op)
		case "2":
			d.bst(op)
		case "3":
			d.jnz(op)
		case "4":
			d.bxc(op)
		case "5":
			d.out(op)
		case "6":
			d.bdv(op)
		case "7":
			d.cdv(op)
		default:
			d.InstructionPointer += 1
		}
	}
}

func (d *Debugger) PrintOutput() string {
	return strings.Join(d.Output, ",")
}

func (d *Debugger) reset(i int) {
	d.A = i
	d.B = 0
	d.C = 0
	d.Output = []string{}
}

func (d *Debugger) findUncorruptedA() int {
	a := int(math.Pow(2, float64(len(d.Output))))

	for {
		d.reset(a)
		d.run()

		if d.PrintOutput() == d.Program {
			return a
		}

		if len(d.ProgramSlice) == len(d.Output) {
			for i := len(d.ProgramSlice) - 1; i >= 0; i-- {
				if d.ProgramSlice[i] != d.Output[i] {
					a += int(math.Pow(8, float64(i)))
					break
				}
			}
		} else if len(d.ProgramSlice) > len(d.Output) {
			a *= 2
		} else if len(d.ProgramSlice) < len(d.Output) {
			a /= 2
		}
	}
}
