package elements

import "math/rand"

//revive:disable:max-public-structs

type Element string

// Just one transformation
type Transformation struct {
	objectElement Element // May not be needed
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
	// Avoid nil map panic
	if t.transformations == nil {
		t.transformations = make(map[Element][]Transformation)
	}

	t.transformations[object] = append(t.transformations[object], transformation)
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
	// Append the leftmost parts of the original
	if index > 0 {
		newElements = append(newElements, m.Elements[:index]...)
	}
	// Append the new elements
	newElements = append(newElements, elements...)
	// Append the rightmost part of the
	if index < len(m.Elements)-1 {
		newElements = append(newElements, m.Elements[index+1:]...)
	}
	m.Elements = newElements
}

func (t *Transformations) Transform(molecule Molecule) []Molecule {
	newMolecules := []Molecule{}
	for k, element := range molecule.Elements {
		// Apply all transformations to thie element
		// Check for matching transformation
		_, ok := t.transformations[element]
		if !ok {
			continue // nothing to do
		}
		// A new moelcule for EACH translation
		for _, transformation := range t.transformations[element] {
			newMolecule := molecule
			newMolecule.swapElement(k, transformation.targets)
			newMolecules = append(newMolecules, newMolecule)
		}
	}
	return newMolecules
}

func CountUnique(molecules []Molecule) int {
	molMap := make(map[string]bool)
	for _, molecule := range molecules {
		molMap[molecule.signature()] = true
	}
	return len(molMap)
}

func Shuffle(slc Transforms) {
	cnt := len(slc)
	for i := 0; i < cnt; i++ {
		// choose index uniformly in [i, N-1]
		// r := i + rand.Intn(cnt-i)
		r := i + rand.Intn(cnt-i) // #nosec G404
		slc[r], slc[i] = slc[i], slc[r]
	}
}

type Transform struct {
	From string
	To   string
}

type Transforms []Transform

func (t Transforms) Len() int           { return len(t) }
func (t Transforms) Less(i, j int) bool { return len(t[i].To) < len(t[j].To) }
func (t Transforms) Swap(i, j int)      { t[i], t[j] = t[j], t[i] }

//revive:enable:max-public-structs
