package main

import (
	"testing"
)

const TEST string = "test.txt"
const TEST2 string = "test2.txt"
const TEST_FMT string = "wrong %s: want %v, got %v"

func TestC9(t *testing.T) {
	d := Debugger{
		C:            9,
		Instructions: []string{"2"},
		Operands:     []int{6},
	}
	d.run()
	want := 1
	got := d.B
	if want != got {
		t.Errorf(TEST_FMT, "value in B register", want, got)
	}
}

func TestA10(t *testing.T) {
	d := Debugger{
		A:            10,
		Instructions: []string{"5", "5", "5"},
		Operands:     []int{0, 1, 4},
	}
	d.run()
	want := "0,1,2"
	got := d.PrintOutput()
	if want != got {
		t.Errorf(TEST_FMT, "output", want, got)
	}
}

func TestA2024(t *testing.T) {
	d := Debugger{
		A:            2024,
		Instructions: []string{"0", "5", "3"},
		Operands:     []int{1, 4, 0},
	}
	d.run()
	want := "4,2,5,6,7,7,7,7,3,1,0"
	got := d.PrintOutput()
	if d.A != 0 {
		t.Errorf(TEST_FMT, "value in A register", 0, d.A)
	}
	if want != got {
		t.Errorf(TEST_FMT, "output", want, got)
	}
}

func TestB29(t *testing.T) {
	d := Debugger{
		B:            29,
		Instructions: []string{"1"},
		Operands:     []int{7},
	}
	d.run()
	want := 26
	got := d.B
	if want != got {
		t.Errorf(TEST_FMT, "value in B register", want, got)
	}
}

func TestB2024(t *testing.T) {
	d := Debugger{
		B:            2024,
		C:            43690,
		Instructions: []string{"4"},
		Operands:     []int{0},
	}
	d.run()
	want := 44354
	got := d.B
	if want != got {
		t.Errorf(TEST_FMT, "value in B register", want, got)
	}
}

func TestPart1(t *testing.T) {
	d := newDebugger(TEST)
	d.run()
	var want string = "4,6,3,5,6,3,5,2,1,0"
	got := d.PrintOutput()
	if want != got {
		t.Errorf(TEST_FMT, "output", want, got)
	}
}

func TestPart2(t *testing.T) {
	d := newDebugger(TEST2)
	A := d.findUncorruptedA()
	if A != 117440 {
		t.Errorf(TEST_FMT, "initial value in A register", 117440, A)
	}
	var want string = d.Program
	got := d.PrintOutput()
	if want != got {
		t.Errorf(TEST_FMT, "output", want, got)
	}
}
