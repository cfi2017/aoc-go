package main

import (
  "fmt"
  "io/ioutil"
  "strings"
  "time"
)

func main() {
  input := GetInput("input.txt")
  Compute(input)
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
    if line[i] == 'B' {
      r++
    }
  }
  for i := 7; i < 10; i++ {
    r = r << 1
    if line[i] == 'R' {
      r++
    }
  }
  return r
}

func t(start time.Time, function string) {
  r := time.Since(start)
  fmt.Printf("%s took %v\n", function, r)
}
