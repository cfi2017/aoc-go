package day5

import (
	"io/ioutil"
	"strings"
	"testing"
)

const (
	data = "assets/data/day5/input.txt"
)

var lines string

var first string

func init() {
	bs, err := ioutil.ReadFile(data)
	if err != nil {
		panic(err)
	}
	lines = string(bs)
	first = strings.Split(lines, "\n")[0]
}

func BenchmarkCompute(b *testing.B) {
	for n := 0; n < b.N; n++ {
		compute(lines)
	}
}

func BenchmarkParse(b *testing.B) {
	for n := 0; n < b.N; n++ {
		parse(first)
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
