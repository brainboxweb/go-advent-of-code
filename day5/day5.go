package day5

import (
	"bufio"
	"bytes"

	"github.com/brainboxweb/advent/day5/naughty"
)

func Part1(input string) int {
	niceCount := 0
	b := bytes.NewBufferString(input)
	scanner := bufio.NewScanner(b)
	for scanner.Scan() {
		theText := scanner.Text()
		phr:= naughty.NewPhrase(theText)
		if phr.Nice() {
			niceCount++
		}
	}
	return niceCount
}


func Part2(input string) int {
	niceCount := 0
	b := bytes.NewBufferString(input)
	scanner := bufio.NewScanner(b)
	for scanner.Scan() {
		theText := scanner.Text()
		phr:= naughty.NewPhrase(theText)
		if phr.NiceTwo(){
			niceCount++
		}
	}
	return niceCount
}
