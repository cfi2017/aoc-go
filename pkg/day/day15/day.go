package day15

import (
	"strconv"
	"strings"
)

type Day15 struct {
}

func (d Day15) Solve(input string) (a, b int) {
	mem := parse(input)
	a = Solve(mem, 2020)
	b = Solve(mem, 30000000)
	return
}

func parse(input string) []int {
	arr := strings.Split(input, ",")
	nums := make([]int, len(arr))
	for i := range arr {
		num, _ := strconv.Atoi(arr[i])
		nums[i] = num
	}
	return nums
}

func Solve(in []int, target int) int {
	last := 0
	mem := make(map[int][]int)
	for i := 0; i <= target; i++ {
		if i < len(in) {
			last = in[i]
		} else if len(mem[last]) == 1 {
			last = 0
		} else {
			last = mem[last][len(mem[last])-1] - mem[last][len(mem[last])-2]
		}

		if _, ok := mem[last]; ok {
			mem[last] = append(mem[last], i)
		} else {
			mem[last] = []int{i}
		}
	}
	return last
}
