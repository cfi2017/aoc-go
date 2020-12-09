package day7

import (
	"io/ioutil"
	"testing"
)

const (
	data = "assets/data/day7/input.txt"
)

var (
	bags  Bags
	color string
)

func init() {
	bs, err := ioutil.ReadFile(data)
	if err != nil {
		panic(err)
	}
	bags = parseInput(string(bs))
	color = "shiny gold"

}

func BenchmarkCountInnerBags(b *testing.B) {
	for i := 0; i < b.N; i++ {
		countInnerBags(bags, color)
	}
}

func BenchmarkPart1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		bagsContainingBag(bags, color)
	}
}
