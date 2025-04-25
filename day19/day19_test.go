package day19_test

import (
	"testing"

	"github.com/brainboxweb/advent/day19"
)

// ----- Day One -----
func TestRun(t *testing.T) {
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
			calibData,
			509,
		},
	}

	for _, test := range tests {
		if actual := day19.Run(test.trans, test.elements); actual != test.expected {
			t.Errorf("Convert(%q) = %d, expected %d.",
				test.trans, actual, test.expected)
		}
	}
}

// ----- Day Two -----
func TestReverseEngineer(t *testing.T) {
	var tests2 = []struct {
		trans    string
		elements string
		target   string
		expected int
	}{
		{
			day19data,
			"e",
			calibData,
			195,
		},
	}

	for _, test := range tests2 {
		if actual := day19.ReverseEngineer(test.trans, test.elements, test.target); actual != test.expected {
			t.Errorf("Convert(%q) = %d, expected %d.",
				test.trans, actual, test.expected)
		}
	}
}

const calibData = `CRnCaSiRnBSiRnFArTiBPTiTiBFArPBCaSiThSiRnTiBPBPMgArCaSiRnTiMgArCaSiThCaSiRnFArRnSiRnFArTiTiBFArCaCaSiRnSiThCaCaSiRnMgArFYSiRnFYCaFArSiThCaSiThPBPTiMgArCaPRnSiAlArPBCaCaSiRnFYSiThCaRnFArArCaCaSiRnPBSiRnFArMgYCaCaCaCaSiThCaCaSiAlArCaCaSiRnPBSiAlArBCaCaCaCaSiThCaPBSiThPBPBCaSiRnFYFArSiThCaSiRnFArBCaCaSiRnFYFArSiThCaPBSiThCaSiRnPMgArRnFArPTiBCaPRnFArCaCaCaCaSiRnCaCaSiRnFYFArFArBCaSiThFArThSiThSiRnTiRnPMgArFArCaSiThCaPBCaSiRnBFArCaCaPRnCaCaPMgArSiRnFYFArCaSiThRnPBPMgAr`

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
