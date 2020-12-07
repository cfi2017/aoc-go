package day5

type Day5 struct {
}

func (Day5) Solve(input string) (a, b int) {
	max, id := compute(input)
	return int(max), int(id)
}

func compute(input string) (max, id uint16) {
	var seats [1025]bool
	for i := 0; i < len(input); i += 11 {
		seat := parse(input[i : i+10])
		seats[seat] = true
	}
	for seats[id] {
		id++
	}
	for seats[id] {
		id++
	}
	max = id + 1
	for seats[max] {
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
