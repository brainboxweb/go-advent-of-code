package elements

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestSwapElement(t *testing.T) {
	// Starting Molecule
	elementH := Element("H")
	elementO := Element("O")
	elementP := Element("P")

	ee := []Element{elementH, elementO, elementP}
	mol := Molecule{ee}

	// New elements
	elementFe := Element("Fe")
	elementSi := Element("Si")

	// Perform the swap
	mol.swapElement(0, []Element{elementFe, elementSi})

	// Expected molecule
	expectedElements := []Element{elementFe, elementSi, elementO, elementP}
	expectedMol := Molecule{expectedElements}

	require.Equal(t, expectedMol, mol)
}
