package day8

import (
	"bufio"
	"bytes"
	"strconv"
)

func Part1(input string) int {
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

func stringInfo2(s string) (lenS, lenQ int) {
	q := strconv.Quote(s)
	return len(s), len(q)
}

func Part2(input string) int {
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

func stringInfo(s string) (lenS, lenQ int) {
	q, _ := strconv.Unquote(s)
	return len(s), len(q)
}
