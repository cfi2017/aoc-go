package day15

import (
	"io/ioutil"
	"testing"
)

const (
	data   = "assets/data/day15/input.txt"
	sample = "0,3,6"
)

var (
	input       []int
	sampleInput []int
)

func init() {
	bs, err := ioutil.ReadFile(data)
	if err != nil {
		panic(err)
	}
	input = parse(string(bs))
	sampleInput = parse(sample)
}

func TestSolveSample(t *testing.T) {
	// sample = 0,3,6
	res := Solve(sampleInput, 2020)
	if res != 436 {
		t.Fatal("invalid output", res)
	}
}

func TestSolve2020(t *testing.T) {
	res := Solve(input, 2020)
	if res != 206 {
		t.Fatal("invalid output", res)
	}
}

func TestSolve30000000(t *testing.T) {
	res := Solve(input, 30000000)
	if res != 955 {
		t.Fatal("invalid output", res)
	}
}

func BenchmarkSolve2020(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Solve(input, 2020)
	}
}

func BenchmarkSolve30000000(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Solve(input, 30000000)
	}
}
