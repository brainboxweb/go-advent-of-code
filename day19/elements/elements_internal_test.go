package elements

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestSwapElement(t *testing.T) {

	//Starting Molecule
	H := Element("H")
	O := Element("O")
	P := Element("P")

	ee := []Element{H, O, P}
	mol := Molecule{ee}

	//New elements
	Fe := Element("Fe")
	Si := Element("Si")

	//Perform the swap
	mol.swapElement(0, []Element{Fe, Si})

	//Expected molecule
	expectedElements := []Element{Fe, Si, O, P}
	expectedMol := Molecule{expectedElements}

	require.Equal(t, expectedMol, mol)
}
