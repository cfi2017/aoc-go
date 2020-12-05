package day5

import (
	"strings"
)

type Day5 struct {
}

func (Day5) Solve(input string) (a, b int) {
	max, id := compute(strings.Split(input, "\n"))
	return int(max), int(id)
}

func compute(lines []string) (max, id uint16) {
	var seats [1025]uint8
	for _, line := range lines {
		seat := parse(line)
		seats[seat] = 1
	}
	for seats[id] == 0 {
		id++
	}
	for seats[id] == 1 {
		id++
	}
	max = id + 1
	for seats[max] == 1 {
		max++
	}
	max--
	return
}

func parse(line string) (r uint16) {
	for i := 0; i < 7; i++ {
		r = r << 1
		r += 1 - uint16((line[i]&4)>>2)
	}
	for i := 7; i < 10; i++ {
		r = r << 1
		r += uint16((line[i] & 2) >> 1)
	}
	return r
}
