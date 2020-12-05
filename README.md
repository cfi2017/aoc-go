[![GoDoc](https://godoc.org/github.com/cfi2017/aoc-go?status.svg)](https://godoc.org/github.com/cfi2017/bl3-save)
[![Go Report Card](https://goreportcard.com/badge/github.com/cfi2017/aoc-go)](https://goreportcard.com/report/github.com/cfi2017/bl3-save)
![GitHub All Releases](https://img.shields.io/github/downloads/cfi2017/aoc-go/total)

# Advent of Code - Go

Go implementation of Advent of Code 2020 Day 5

Heavily optimised for performance - see benchmarks below.

Credits to [@Stratege](https://github.com/Stratege)

## Benchmarks
```
/tmp/___gobench_github_com_cfi2017_aoc_go -test.v -test.bench . -test.run ^$
goos: linux
goarch: amd64
pkg: github.com/cfi2017/aoc-go
BenchmarkComputeInput
BenchmarkComputeInput-24    	    25584	        44730 ns/op
BenchmarkCompute
BenchmarkCompute-24         	   133684	         8707 ns/op
BenchmarkParse
BenchmarkParse-24           	123821754	         9.58 ns/op
PASS
```