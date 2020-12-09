package day9

import (
	"io/ioutil"
	"testing"
)

const (
	data = "assets/data/day9/input.txt"
)

var (
	input   []uint64
	corrupt uint64
	pre     = 25
)

func init() {
	bs, err := ioutil.ReadFile(data)
	if err != nil {
		panic(err)
	}
	input = formatInput(string(bs))
	corrupt = firstCorrupt(input, pre)
}

func BenchmarkFirstCorrupt(b *testing.B) {
	for i := 0; i < b.N; i++ {
		firstCorrupt(input, pre)
	}
}

func BenchmarkCorruptWindow(b *testing.B) {
	for i := 0; i < b.N; i++ {
		corruptSumWindow(input, corrupt)
	}
}
