package main

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestParseElements(t *testing.T) {
	input := `AlThFSSi`
	expected := []Element{"Al", "Th", "F", "S", "Si"}
	result := parseElements(input)
	require.Equal(t, expected, result)
}

func TestParseTransformations(t *testing.T) {

	input := `H => HO
H => OH
O => HH`

	h := Element("H")
	o := Element("O")

	t1 := NewTransformation(h, []Element{h, o})
	t2 := NewTransformation(h, []Element{o, h})
	t3 := NewTransformation(o, []Element{h, h})

	transformations := Transformations{}
	transformations.AddTransformation(t1)
	transformations.AddTransformation(t2)
	transformations.AddTransformation(t3)

	expected := transformations

	result := parseTransformations(input)
	require.Equal(t, expected, result)
}

func TestTransform(t *testing.T) {

	h := Element("H")
	o := Element("O")

	t1 := NewTransformation(h, []Element{h, o})
	t2 := NewTransformation(h, []Element{o, h})
	t3 := NewTransformation(o, []Element{h, h})

	transformations := Transformations{}

	transformations.AddTransformation(t1)
	transformations.AddTransformation(t2)
	transformations.AddTransformation(t3)

	elements := []Element{"H", "O", "H"}

	molecule := Molecule{elements}

	result := transformations.transform(molecule)

	//expected
	e1 := []Element{"H", "O", "O", "H"}
	e2 := []Element{"O", "H", "O", "H"}
	e3 := []Element{"H", "H", "H", "H"}
	e4 := []Element{"H", "O", "H", "O"}
	e5 := []Element{"H", "O", "O", "H"}

	m1 := Molecule{e1}
	m2 := Molecule{e2}
	m3 := Molecule{e3}
	m4 := Molecule{e4}
	m5 := Molecule{e5}

	expected := []Molecule{m1, m2, m3, m4, m5}

	require.Equal(t, expected, result)
}

func TestSwapElement(t *testing.T) {

	//Starting Molecule
	H := Element("H")
	O := Element("O")
	P := Element("P")

	elements := []Element{H, O, P}
	mol := Molecule{elements}

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

//----- Day One -----
var tests = []struct {
	trans    string
	elements string
	expected int
}{
	{
		`H => HO
H => OH
O => HH`,
		`HOH`,
		4,
	},
	{
		day19data,
		`ORnPBPMgArCaCaCaSiThCaCaSiThCaCaPBSiRnFArRnFArCaCaSiThCaCaSiThCaCaCaCaCaCaSiRnFYFArSiRnMgArCaSiRnPTiTiBFYPBFArSiRnCaSiRnTiRnFArSiAlArPTiBPTiRnCaSiAlArCaPTiTiBPMgYFArPTiRnFArSiRnCaCaFArRnCaFArCaSiRnSiRnMgArFYCaSiRnMgArCaCaSiThPRnFArPBCaSiRnMgArCaCaSiThCaSiRnTiMgArFArSiThSiThCaCaSiRnMgArCaCaSiRnFArTiBPTiRnCaSiAlArCaPTiRnFArPBPBCaCaSiThCaPBSiThPRnFArSiThCaSiThCaSiThCaPTiBSiRnFYFArCaCaPRnFArPBCaCaPBSiRnTiRnFArCaPRnFArSiRnCaCaCaSiThCaRnCaFArYCaSiRnFArBCaCaCaSiThFArPBFArCaSiRnFArRnCaCaCaFArSiRnFArTiRnPMgArF`,
		576,
	},
}

func TestRun(t *testing.T) {
	for _, test := range tests {
		if actual := run(test.trans, test.elements); actual != test.expected {
			t.Errorf("Convert(%q) = %d, expected %d.",
				test.trans, actual, test.expected)
		}
	}
}

//----- Day Two -----
var tests2 = []struct {
	trans    string
	elements string
	target   string
	expected int
}{
	{
		day19data,
		"e",
		`ORnPBPMgArCaCaCaSiThCaCaSiThCaCaPBSiRnFArRnFArCaCaSiThCaCaSiThCaCaCaCaCaCaSiRnFYFArSiRnMgArCaSiRnPTiTiBFYPBFArSiRnCaSiRnTiRnFArSiAlArPTiBPTiRnCaSiAlArCaPTiTiBPMgYFArPTiRnFArSiRnCaCaFArRnCaFArCaSiRnSiRnMgArFYCaSiRnMgArCaCaSiThPRnFArPBCaSiRnMgArCaCaSiThCaSiRnTiMgArFArSiThSiThCaCaSiRnMgArCaCaSiRnFArTiBPTiRnCaSiAlArCaPTiRnFArPBPBCaCaSiThCaPBSiThPRnFArSiThCaSiThCaSiThCaPTiBSiRnFYFArCaCaPRnFArPBCaCaPBSiRnTiRnFArCaPRnFArSiRnCaCaCaSiThCaRnCaFArYCaSiRnFArBCaCaCaSiThFArPBFArCaSiRnFArRnCaCaCaFArSiRnFArTiRnPMgArF`,
		207,
	},
}

func TestReverseEngineer(t *testing.T) {
	for _, test := range tests2 {
		if actual := reverseEngineer(test.trans, test.elements, test.target); actual != test.expected {
			t.Errorf("Convert(%q) = %d, expected %d.",
				test.trans, actual, test.expected)
		}
	}
}

const day19data = `Al => ThF
Al => ThRnFAr
B => BCa
B => TiB
B => TiRnFAr
Ca => CaCa
Ca => PB
Ca => PRnFAr
Ca => SiRnFYFAr
Ca => SiRnMgAr
Ca => SiTh
F => CaF
F => PMg
F => SiAl
H => CRnAlAr
H => CRnFYFYFAr
H => CRnFYMgAr
H => CRnMgYFAr
H => HCa
H => NRnFYFAr
H => NRnMgAr
H => NTh
H => OB
H => ORnFAr
Mg => BF
Mg => TiMg
N => CRnFAr
N => HSi
O => CRnFYFAr
O => CRnMgAr
O => HP
O => NRnFAr
O => OTi
P => CaP
P => PTi
P => SiRnFAr
Si => CaSi
Th => ThCa
Ti => BP
Ti => TiTi
e => HF
e => NAl
e => OMg`
