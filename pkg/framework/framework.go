package framework

import (
	"errors"
	"github.com/cfi2017/aoc-go/pkg/api"
	"github.com/cfi2017/aoc-go/pkg/day5"
)

var (
	ErrNotImplemented = errors.New("error: solver not implemented")
)

var solverMap = map[int]api.Solver{
	5: &day5.Day5{},
}

func GetSolver(day int) (api.Solver, error) {
	solver, ok := solverMap[day]
	if !ok {
		return nil, ErrNotImplemented
	}
	return solver, nil
}
