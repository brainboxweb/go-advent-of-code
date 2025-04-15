package main

import (
	"bufio"
	"bytes"
	"regexp"
)

func Run(input string) int {

	niceCount := 0
	b := bytes.NewBufferString(input)
	scanner := bufio.NewScanner(b)
	for scanner.Scan() {
		theText := scanner.Text()
		if nice(theText) {
			niceCount++
		}
	}
	return niceCount
}

func nice(phrase string) bool {
	if blacklist(phrase) == true {
		return false
	}
	if vowels(phrase) && doubleLetter(phrase) {
		return true
	}
	return false
}

func vowels(phrase string) bool {
	r, _ := regexp.Compile(`[aeiou]`)
	// Will print 'true'.
	result := r.FindAllString(phrase, -1)
	if len(result) > 2 {
		return true
	}
	return false
}

func doubleLetter(phrase string) bool {
	current := " "
	for _, char := range phrase {
		c := string(char)
		if current == c {
			return true
		}
		current = c
	}
	return false
}

func blacklist(phrase string) bool {
	r, _ := regexp.Compile("ab|cd|pq|xy")
	result := r.MatchString(phrase)
	return result
}

func RunDayTwo(input string) int {
	niceCount := 0
	b := bytes.NewBufferString(input)
	scanner := bufio.NewScanner(b)
	for scanner.Scan() {
		theText := scanner.Text()
		if niceTwo(theText) {
			niceCount++
		}
	}
	return niceCount
}

func niceTwo(phrase string) bool {
	if doublesNoOverlap(phrase) && repeatWithGap(phrase) {
		return true
	}
	return false
}

func doublesNoOverlap(phrase string) bool {
	tokens := tokenize(phrase)
	matchCount := 0
	for i := 0; i < len(tokens); i++ {
		subject := tokens[i]
		//find match
		for j := 0; j < len(tokens); j++ {
			if i == j { //don't compare with self
				continue
			}
			if i+1 == j { //overlap match
				continue
			}
			object := tokens[j]
			if object == subject {
				matchCount++
			}
			if matchCount > 1 {
				return true
			}
		}
	}
	return false
}

func tokenize(phrase string) map[int]string {
	tokenLength := len(phrase) - 1
	tokens := make(map[int]string) // len(a)=5[int]string
	for i := 0; i < tokenLength; i++ {
		tokens[i] = string(phrase[i]) + string(phrase[i+1])
	}
	return tokens
}

func repeatWithGap(phrase string) bool {
	count := len(phrase)
	matchCount := 0
	for i := 0; i < count-2; i++ {
		if string(phrase[i]) == string(phrase[i+2]) {
			matchCount++
		}
		if matchCount > 0 {
			return true
		}
	}
	return false
}
