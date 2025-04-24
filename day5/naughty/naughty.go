package naughty

import (
	"regexp"
)

type Phrase struct {
	phrase string
}

func NewPhrase(phrase string) Phrase {
	return Phrase{phrase}
}

func (phr Phrase) Nice() bool {
	if phr.blacklist() {
		return false
	}
	if phr.vowels() && phr.doubleLetter() {
		return true
	}
	return false
}

func (phr Phrase) NiceTwo() bool {
	if phr.doublesNoOverlap() && phr.repeatWithGap() {
		return true
	}
	return false
}

func (phr Phrase) vowels() bool {
	r := regexp.MustCompile(`[aeiou]`)
	result := r.FindAllString(phr.phrase, -1)
	return len(result) > 2
}

func (phr Phrase) doubleLetter() bool {
	current := " "
	for _, char := range phr.phrase {
		c := string(char)
		if current == c {
			return true
		}
		current = c
	}
	return false
}

func (phr Phrase) blacklist() bool {
	r := regexp.MustCompile("ab|cd|pq|xy")
	result := r.MatchString(phr.phrase)
	return result
}

func (phr Phrase) doublesNoOverlap() bool {
	tokens := tokenize(phr.phrase)
	matchCount := 0
	for i := 0; i < len(tokens); i++ {
		subject := tokens[i]
		// find match
		for j := 0; j < len(tokens); j++ {
			if i == j { // don't compare with self
				continue
			}
			if i+1 == j { // overlap match
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

func (phr Phrase) repeatWithGap() bool {
	count := len(phr.phrase)
	matchCount := 0
	for i := 0; i < count-2; i++ {
		if string(phr.phrase[i]) == string(phr.phrase[i+2]) {
			matchCount++
		}
		if matchCount > 0 {
			return true
		}
	}
	return false
}
