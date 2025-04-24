package main

import (
	//	"fmt"
	"bufio"
	"bytes"
	"regexp"
	"strings"
	//	"sort"
	"math/rand"
	"time"
)

func run(transformationString, elementString string) int {
	elements := parseElements(elementString)
	transformations := parseTransformations(transformationString)
	molecule := Molecule{elements}
	newMolecules := transformations.transform(molecule)
	return countUnique(newMolecules)
}

type Element string

func parseElements(input string) []Element {
	r, _ := regexp.Compile(`([A-Z][a-z]?)`)
	results := r.FindAllString(input, -1)
	elements := []Element{}
	for _, result := range results {
		elements = append(elements, Element(result))
	}
	return elements
}

func parseTransformations(input string) Transformations {
	b := bytes.NewBufferString(input)
	scanner := bufio.NewScanner(b)
	transformations := Transformations{}
	for scanner.Scan() {
		line := scanner.Text()
		items := strings.Split(line, " ")
		object := Element(items[0])
		targetsString := items[2]
		targetStrings := parseElements(targetsString)
		targets := []Element{}
		for _, target := range targetStrings {
			targets = append(targets, Element(target))
		}
		trans := NewTransformation(object, targets)
		transformations.AddTransformation(trans)
	}
	return transformations
}

// Just one transformation
type Transformation struct {
	objectElement Element //May not be needed
	targets       []Element
}

// Just one transformation
func NewTransformation(object Element, targets []Element) Transformation {
	return Transformation{object, targets}
}

type Transformations struct {
	transformations map[Element][]Transformation
}

func (t *Transformations) AddTransformation(transformation Transformation) {
	object := transformation.objectElement
	//Avoid nil map panic
	if t.transformations == nil {
		t.transformations = make(map[Element][]Transformation)
	}
	_, ok := t.transformations[object]
	if !ok {
		t.transformations[object] = append(t.transformations[object], transformation)
	} else {
		t.transformations[object] = append(t.transformations[object], transformation)
	}
}

type Molecule struct {
	Elements []Element
}

func (m Molecule) signature() string {
	signature := ""
	for _, element := range m.Elements {
		signature += string(element)
	}
	return signature
}

func (m *Molecule) swapElement(index int, elements []Element) {
	newElements := []Element{}
	//Append the leftmost parts of the original
	if index > 0 {
		newElements = append(newElements, m.Elements[:index]...)
	}
	//Append the new elements
	newElements = append(newElements, elements...)
	//Append the rightmost part of the
	if index < len(m.Elements)-1 {
		newElements = append(newElements, m.Elements[index+1:]...)
	}
	m.Elements = newElements
}

func (t *Transformations) transform(molecule Molecule) []Molecule {
	newMolecules := []Molecule{}
	for k, element := range molecule.Elements {
		//Apply all transformations to thie elemnt
		//Check for matching transformation
		_, ok := t.transformations[element]
		if !ok {
			continue //nothing to do
		}
		//A new moelcule for EACH translation
		for _, transformation := range t.transformations[element] {
			newMolecule := molecule
			newMolecule.swapElement(k, transformation.targets)
			newMolecules = append(newMolecules, newMolecule)
		}
	}
	return newMolecules
}

func countUnique(molecules []Molecule) int {
	molMap := make(map[string]bool)
	for _, molecule := range molecules {
		molMap[molecule.signature()] = true
	}
	return len(molMap)
}

// ----------------- Day 2 ---------------------
func reverseEngineer(tranformations, shortMoleculeString, longMoleculeString string) int {
	//Prep transformations for regex
	b := bytes.NewBufferString(tranformations)
	scanner := bufio.NewScanner(b)
	transforms := Transforms{}
	for scanner.Scan() {
		line := scanner.Text()
		items := strings.Split(line, " ")
		object := items[0]
		targetsString := items[2]
		transform := Transform{object, targetsString}
		transforms = append(transforms, transform)
	}
	rand.Seed(time.Now().UTC().UnixNano())
	winningCount := 0

OuterLoop:
	for {
		Shuffle(transforms)
		startString := longMoleculeString
		candidateWinningCount := 0
		for i := 1; i < 10; i++ { //Max attempts
			for _, transform := range transforms {
				r, _ := regexp.Compile(transform.to)
				matchCount := r.FindAllString(startString, -1)
				candidateWinningCount += len(matchCount)
				startString = r.ReplaceAllString(startString, transform.from)
				if startString == shortMoleculeString {
					winningCount = candidateWinningCount
					break OuterLoop
				}
			}
		}
	}
	return winningCount
}

func Shuffle(slc Transforms) {
	N := len(slc)
	for i := 0; i < N; i++ {
		// choose index uniformly in [i, N-1]
		r := i + rand.Intn(N-i)
		slc[r], slc[i] = slc[i], slc[r]
	}
}

type Transform struct {
	from string
	to   string
}

type Transforms []Transform

func (t Transforms) Len() int           { return len(t) }
func (t Transforms) Less(i, j int) bool { return len(t[i].to) < len(t[j].to) }
func (t Transforms) Swap(i, j int)      { t[i], t[j] = t[j], t[i] }
