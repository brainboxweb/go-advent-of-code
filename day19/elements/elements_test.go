package elements_test

import (
	"testing"

	"github.com/brainboxweb/advent/day19/elements"
	"github.com/stretchr/testify/require"
)

func TestTransform(t *testing.T) {
	h := elements.Element("H")
	o := elements.Element("O")

	t1 := elements.NewTransformation(h, []elements.Element{h, o})
	t2 := elements.NewTransformation(h, []elements.Element{o, h})
	t3 := elements.NewTransformation(o, []elements.Element{h, h})

	transformations := elements.Transformations{}

	transformations.AddTransformation(t1)
	transformations.AddTransformation(t2)
	transformations.AddTransformation(t3)

	ee := []elements.Element{"H", "O", "H"}

	molecule := elements.Molecule{ee}

	result := transformations.Transform(molecule)

	//expected
	e1 := []elements.Element{"H", "O", "O", "H"}
	e2 := []elements.Element{"O", "H", "O", "H"}
	e3 := []elements.Element{"H", "H", "H", "H"}
	e4 := []elements.Element{"H", "O", "H", "O"}
	e5 := []elements.Element{"H", "O", "O", "H"}

	m1 := elements.Molecule{e1}
	m2 := elements.Molecule{e2}
	m3 := elements.Molecule{e3}
	m4 := elements.Molecule{e4}
	m5 := elements.Molecule{e5}

	expected := []elements.Molecule{m1, m2, m3, m4, m5}

	require.Equal(t, expected, result)
}
