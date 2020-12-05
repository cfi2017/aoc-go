package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

func main() {
	input := GetInput(os.Args[1])
	max, id := Compute(input)
	fmt.Printf("max is %d, id is %d", max, id)
}

func ComputeInput(input string) (max, id uint16) {
	return Compute(GetInput(input))
}

func GetInput(input string) []string {
	bs, err := ioutil.ReadFile(input)
	if err != nil {
		panic(err)
	}
	return strings.Split(string(bs), "\n")
}

func Compute(lines []string) (max, id uint16) {
	var seats [1025]uint8
	for _, line := range lines {
		seat := Parse(line)
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

func Parse(line string) (r uint16) {
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
