package main

import (
  "io/ioutil"
  "strings"
  "testing"
)

var lines []string

func init() {
  bs, err := ioutil.ReadFile("input.txt")
  if err != nil {
    panic(err)
  }
  lines = strings.Split(string(bs), "\n")
}

func BenchmarkComputeInput(b *testing.B) {
  for n := 0; n < b.N; n++ {
    ComputeInput("input.txt")
  }
}

func BenchmarkCompute(b *testing.B) {
  for n := 0; n < b.N; n++ {
    Compute(lines)
  }
}

func BenchmarkParse(b *testing.B) {
  for n := 0; n < b.N; n++ {
    Parse(lines[0])
  }
}

func TestCompute(t *testing.T) {
  max, id := Compute(lines)
  if max != 883 {
    panic("invalid max value")
  }
  if id != 532 {
    panic("invalid id value")
  }
}

func TestComputeInput(t *testing.T) {
  max, id := ComputeInput("input.txt")
  if max != 883 {
    panic("invalid max value")
  }
  if id != 532 {
    panic("invalid id value")
  }
}
