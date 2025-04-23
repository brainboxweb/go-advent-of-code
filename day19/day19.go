package day19

import (
	"bufio"
	"bytes"
	"math/rand"
	"regexp"
	"strings"
	"time"

	"github.com/brainboxweb/advent/day19/elements"
)

func Run(transformationString, elementString string) int {
	ee := parseElements(elementString)
	transformations := parseTransformations(transformationString)
	molecule := elements.Molecule{Elements: ee}
	newMolecules := transformations.Transform(molecule)
	return elements.CountUnique(newMolecules)
}

// type Element string

func parseElements(input string) []elements.Element {
	r, _ := regexp.Compile(`([A-Z][a-z]?)`)
	results := r.FindAllString(input, -1)
	ee := []elements.Element{}
	for _, result := range results {
		ee = append(ee, elements.Element(result))
	}
	return ee
}

func parseTransformations(input string) elements.Transformations {
	b := bytes.NewBufferString(input)
	scanner := bufio.NewScanner(b)
	transformations := elements.Transformations{}
	for scanner.Scan() {
		line := scanner.Text()
		items := strings.Split(line, " ")
		object := elements.Element(items[0])
		targetsString := items[2]
		targetStrings := parseElements(targetsString)
		targets := []elements.Element{}
		for _, target := range targetStrings {
			targets = append(targets, elements.Element(target))
		}
		trans := elements.NewTransformation(object, targets)
		transformations.AddTransformation(trans)
	}
	return transformations
}

// ----------------- Day 2 ---------------------
func ReverseEngineer(tranformations, shortMoleculeString, longMoleculeString string) int {
	//Prep transformations for regex
	b := bytes.NewBufferString(tranformations)
	scanner := bufio.NewScanner(b)
	transforms := elements.Transforms{}
	for scanner.Scan() {
		line := scanner.Text()
		items := strings.Split(line, " ")
		object := items[0]
		targetsString := items[2]
		transform := elements.Transform{object, targetsString}
		transforms = append(transforms, transform)
	}
	rand.Seed(time.Now().UTC().UnixNano())
	winningCount := 0

OuterLoop:
	for {
		elements.Shuffle(transforms)
		startString := longMoleculeString
		candidateWinningCount := 0
		for i := 1; i < 10; i++ { //Max attempts
			for _, transform := range transforms {
				r, _ := regexp.Compile(transform.To)
				matchCount := r.FindAllString(startString, -1)
				candidateWinningCount += len(matchCount)
				startString = r.ReplaceAllString(startString, transform.From)
				if startString == shortMoleculeString {
					winningCount = candidateWinningCount
					break OuterLoop
				}
			}
		}
	}
	return winningCount
}
