package main

import (
	"io/ioutil"
	"strings"
	"testing"
)

var lines []string

func init() {
	bs, err := ioutil.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}
	lines = strings.Split(string(bs), "\n")
}

func BenchmarkComputeInput(b *testing.B) {
	for n := 0; n < b.N; n++ {
		computeInput("input.txt")
	}
}

func BenchmarkCompute(b *testing.B) {
	for n := 0; n < b.N; n++ {
		compute(lines)
	}
}

func BenchmarkParse(b *testing.B) {
	for n := 0; n < b.N; n++ {
		parse(lines[0])
	}
}

func TestCompute(t *testing.T) {
	max, id := compute(lines)
	if max != 883 {
		panic("invalid max value")
	}
	if id != 532 {
		panic("invalid id value")
	}
}

func TestComputeInput(t *testing.T) {
	max, id := computeInput("input.txt")
	if max != 883 {
		panic("invalid max value")
	}
	if id != 532 {
		panic("invalid id value")
	}
}
