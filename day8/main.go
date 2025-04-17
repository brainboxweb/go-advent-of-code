package main

import (
	"bufio"
	"bytes"
	"strconv"
)

func Run(input string) int {
	total := 0
	b := bytes.NewBufferString(input)
	scanner := bufio.NewScanner(b)
	for scanner.Scan() {
		one, two := stringInfo(scanner.Text())
		total += one
		total -= two
	}
	return total
}

func stringInfo2(s string) (int, int) {
	q := strconv.Quote(s)
	return len(s), len(q)
}

func run2(input string) int {
	total := 0
	b := bytes.NewBufferString(input)
	scanner := bufio.NewScanner(b)
	for scanner.Scan() {
		one, two := stringInfo2(scanner.Text())
		total += two
		total -= one
	}
	return total
}

func stringInfo(s string) (int, int) {
	q, _ := strconv.Unquote(s)
	return len(s), len(q)
}
