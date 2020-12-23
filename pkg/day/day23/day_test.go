package day23

import (
	"fmt"
	"io/ioutil"
	"testing"
)

const (
	data   = "assets/data/day23/input.txt"
	sample = "389125467"
)

var (
	input       []uint32
	sampleInput []uint32
)

func init() {
	bs, err := ioutil.ReadFile(data)
	if err != nil {
		panic(err)
	}
	input = parse(string(bs))
	sampleInput = parse(sample)
}

func TestDay23_SolveSampleA(t *testing.T) {
	a := solve(sampleInput, uint32(len(sampleInput)), 100, false)
	if a != 67384529 {
		panic(fmt.Sprintf("assertion error, %d", a))
	}
}

func TestDay23_SolveInputA(t *testing.T) {
	a := solve(input, uint32(len(input)), 100, false)
	if a != 72496583 {
		panic(fmt.Sprintf("assertion error, %d", a))
	}
}

func TestDay23_SolveSampleB(t *testing.T) {
	a := solve(sampleInput, 1000000, 10000000, true)
	if a != 149245887792 {
		panic(fmt.Sprintf("assertion error, %d", a))
	}
}

func TestDay23_SolveInputB(t *testing.T) {
	a := solve(input, 1000000, 10000000, true)
	if a != 41785843847 {
		panic(fmt.Sprintf("assertion error, %d", a))
	}
}
