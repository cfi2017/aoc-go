package day7

import (
	"strconv"
	"strings"
)

type Bags map[string]Bag
type Bag map[string]int

func (b Bag) containsAny(colors []string) bool {
	for _, color := range colors {
		if b.contains(color) {
			return true
		}
	}
	return false
}

func (b Bag) contains(color string) bool {
	_, ok := b[color]
	return ok
}

type Day7 struct {
}

func (Day7) Solve(input string) (a, b int) {
	bb := parseInput(input)
	color := "shiny gold"
	return bagsContainingBag(bb, color), countInnerBags(bb, color)

}

func parseInput(input string) Bags {
	lines := strings.Split(input, "\n")
	bb := make(Bags)
	for _, line := range lines {
		name, b := parse(line)
		bb[name] = b
	}
	return bb
}

func bagsContainingBag(bb Bags, color string) int {
	containingBags := make([]string, 1, len(bb))
	containingBags[0] = color
	last_length := 0
	for len(containingBags) > last_length {
		previous := last_length
		last_length = len(containingBags)
		for cc, bag := range bb {
			if bag.containsAny(containingBags[previous:]) {
				containingBags = append(containingBags, cc)
			}
		}
	}
	return len(containingBags)
}

func parseNthChild(tokens []string, n int) (string, int) {
	name := strings.Join(tokens[5+n*4:5+n*4+2], " ")
	val, _ := strconv.Atoi(tokens[4+n*4])
	return name, val
}

func parse(line string) (string, Bag) {
	tokens := strings.Split(line, " ")
	color := strings.Join(tokens[:2], " ")
	if strings.Contains(line, " no ") {
		return color, nil
	}
	childCount := strings.Count(line, ",") + 1
	children := make(Bag)
	for i := 0; i < childCount; i++ {
		name, val := parseNthChild(tokens, i)
		children[name] = val
	}
	return color, children
}

func countInnerBags(b Bags, color string) int {
	childSum := 0
	for k, v := range b[color] {
		childSum += v * (countInnerBags(b, k) + 1)
	}
	return childSum
}
