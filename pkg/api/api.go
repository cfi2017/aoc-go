package api

type Solver interface {
	Solve(input string) (a, b int)
}
