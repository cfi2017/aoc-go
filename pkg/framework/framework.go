package framework

import (
	"errors"
	"github.com/cfi2017/aoc-go/pkg/api"
	"github.com/cfi2017/aoc-go/pkg/day/day5"
	"github.com/cfi2017/aoc-go/pkg/day/day9"
)

var (
	ErrNotImplemented = errors.New("error: solver not implemented")
)

var solverMap = map[int]api.Solver{
	5: &day5.Day5{},
	9: &day9.Day9{},
}

func GetSolver(day int) (api.Solver, error) {
	solver, ok := solverMap[day]
	if !ok {
		return nil, ErrNotImplemented
	}
	return solver, nil
}
