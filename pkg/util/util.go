package util

import (
	"io/ioutil"
	"strings"
)

func SanitizeInput(input string) string {
	s := strings.ReplaceAll(input, "\r\n", "\n")
	if s[len(s)-1] == '\n' {
		return s[:len(s)-1]
	}
	return s
}

func GetInput(path string) string {
	bs, err := ioutil.ReadFile(path)
	if err != nil {
		panic(err)
	}
	return string(bs)
}

func GetInputArray(path string) []string {
	bs, err := ioutil.ReadFile(path)
	if err != nil {
		panic(err)
	}
	return strings.Split(string(bs), "\n")
}
