package day9

import (
	"strconv"
	"strings"
)

type Day9 struct {
}

func (d Day9) Solve(input string) (a, b int) {
	in := formatInput(input)
	corrupt := firstCorrupt(in, 25)
	window := corruptSumWindow(in, corrupt)
	return int(corrupt), int(minIntSlice(window) + maxIntSlice(window))
}

func formatInput(input string) []uint64 {
	ss := strings.Split(input, "\n")
	out := make([]uint64, len(ss))
	for index, s := range ss {
		i, _ := strconv.Atoi(s)
		out[index] = uint64(i)
	}
	return out
}

func firstCorrupt(nums []uint64, pre int) uint64 {
	for i := pre; i < len(nums); i++ {
		if !checkWindow(nums[i-pre:i], nums[i]) {
			return nums[i]
		}
	}
	return 0
}

func checkWindow(nums []uint64, num uint64) bool {
	for i, x := range nums {
		for _, y := range nums[i:] {
			if x != y && x+y == num {
				return true
			}
		}
	}
	return false
}

func corruptSumWindow(nums []uint64, corrupt uint64) []uint64 {
	for i, x := range nums {
		sum := x
		for i2, y := range nums[i+1:] {
			sum += y
			if sum == corrupt {
				return nums[i : i+1+i2]
			} else if sum > corrupt {
				break
			}
		}
	}
	return nil
}

func corruptSumRecursive(nums []uint64, target uint64) []uint64 {
	return recurse(nums, 0, 0, 0, target)
}

func recurse(nums []uint64, start, end, sum, target uint64) []uint64 {
	sum += nums[start+end]
	if sum == target {
		return nums[start : start+end]
	} else if sum < target {
		return recurse(nums, start, end+1, sum, target)
	} else { // sum > target
		return recurse(nums, start+1, 0, sum, target)
	}
}

func minIntSlice(nums []uint64) (m uint64) {
	for i, e := range nums {
		if i == 0 || e < m {
			m = e
		}
	}
	return
}

func maxIntSlice(nums []uint64) (m uint64) {
	for i, e := range nums {
		if i == 0 || e > m {
			m = e
		}
	}
	return
}
