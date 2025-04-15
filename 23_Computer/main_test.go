package main

import (
	"github.com/stretchr/testify/require"
	"testing"
)

var tests = []struct {
	input    string
	expected Instruction
}{
	{
		"jio a, +19",
		Instruction{"jio", "a", 19},
	},
	{
		"inc a",
		Instruction{"inc", "a", 0},
	},
	{
		"jmp +23",
		Instruction{"jmp", "", 23},
	},
	{
		"jmp -7",
		Instruction{"jmp", "", -7},
	},
}

func TestParse(t *testing.T) {
	for _, test := range tests {
		if actual := parse(test.input); actual != test.expected {
			t.Errorf("Convert(%s) = %v, expected %v.",
				test.input, actual, test.expected)
		}
	}
}

func TestParseInstructions(t *testing.T) {

	//expected
	inst1 := Instruction{"inc", "a", 0}
	inst2 := Instruction{"jio", "a", 2}
	inst3 := Instruction{"tpl", "a", 0}
	inst4 := Instruction{"inc", "a", 0}

	expected := InstructionSet{}

	expected.Add(inst1)
	expected.Add(inst2)
	expected.Add(inst3)
	expected.Add(inst4)

	input := `inc a
jio a, +2
tpl a
inc a`

	instructionSet := parseInstructions(input)

	require.Equal(t, expected, instructionSet)
}

var tests2 = []struct {
	input           string
	intialValueForA int
	expectedA       int
	expectedB       int
}{
	{
		`inc a
jio a, +2
tpl a
inc a`,
		0,
		2,
		0,
	},
	{
		day22data,
		0,
		1,
		184,
	},
	{
		day22data,
		1,
		1,
		231,
	},
}

func TestRun(t *testing.T) {
	for _, test := range tests2 {

		a, b := Run(test.input, test.intialValueForA)

		if a != test.expectedA {
			t.Errorf("Convert(%s) = %d, expected %d for (a).",
				test.input, a, test.expectedA)
		}

		if b != test.expectedB {
			t.Errorf("Convert(%s) = %d, expected %d for (b).",
				test.input, b, test.expectedB)
		}
	}
}

//tpl r sets register r to triple its current value, then continues with the next instruction.
func TestHlf(t *testing.T) {

	reg := Register{"a", 10}
	hlf(&reg)

	exectedValue := 5

	require.Equal(t, exectedValue, reg.Value)
}

//tpl r sets register r to triple its current value, then continues with the next instruction.
func TestTpl(t *testing.T) {

	reg := Register{"a", 10}
	tpl(&reg)

	exectedValue := 30

	require.Equal(t, exectedValue, reg.Value)
}

//inc r increments register r, adding 1 to it, then continues with the next instruction.
func TestInc(t *testing.T) {

	reg := Register{"a", 7}
	inc(&reg)

	exectedValue := 8

	require.Equal(t, exectedValue, reg.Value)
}

//jmp offset is a jump; it continues with the instruction offset away relative to itself.
func TestJump(t *testing.T) {

	inst1 := Instruction{"inc", "a", 0}
	inst2 := Instruction{"jmp", "a", 2}
	inst3 := Instruction{"tpl", "a", 0}
	inst4 := Instruction{"inc", "a", 0}

	set := InstructionSet{}

	set.Add(inst1)
	set.Add(inst2)
	set.Add(inst3)
	set.Add(inst4)

	set.Fetch() //the increment
	set.Fetch() //The jump itself
	set.Jump(2) //Jump two (manually)

	result, _ := set.Fetch() //expect the second inc
	require.Equal(t, inst4, result)

}

//jie r, offset is like jmp, but only jumps if register r is even ("jump if even").
func TestJie(t *testing.T) {

	inst1 := Instruction{"inc", "a", 0}
	inst2 := Instruction{"jie", "a", 2}
	inst3 := Instruction{"tpl", "a", 0}
	inst4 := Instruction{"inc", "a", 0}

	set := InstructionSet{}

	set.Add(inst1)
	set.Add(inst2)
	set.Add(inst3)
	set.Add(inst4)

	set.Fetch() //the increment
	set.Fetch() //The jump itself

	set.Jie(2, 100) //Jump if even

	result, _ := set.Fetch() //expect the second inc
	require.Equal(t, inst4, result)

}

//jie r, offset is like jmp, but only jumps if register r is even ("jump if even").
func TestJieAgain(t *testing.T) {

	inst1 := Instruction{"inc", "a", 0}
	inst2 := Instruction{"jie", "a", 2}
	inst3 := Instruction{"tpl", "a", 0}
	inst4 := Instruction{"inc", "a", 0}

	set := InstructionSet{}

	set.Add(inst1)
	set.Add(inst2)
	set.Add(inst3)
	set.Add(inst4)

	set.Fetch() //the increment
	set.Fetch() //The jump itself

	set.Jie(2, 99) //Jump if even. It's odd... so nothing really happens

	result, _ := set.Fetch() //expect the second inc
	require.Equal(t, inst3, result)
}

//jio r, offset is like jmp, but only jumps if register r is 1 ("jump if one", not odd).
func TestJio(t *testing.T) {

	inst1 := Instruction{"inc", "a", 0}
	inst2 := Instruction{"jie", "a", 2}
	inst3 := Instruction{"tpl", "a", 0}
	inst4 := Instruction{"inc", "a", 0}

	set := InstructionSet{}

	set.Add(inst1)
	set.Add(inst2)
	set.Add(inst3)
	set.Add(inst4)

	set.Fetch() //the increment
	set.Fetch() // the jump itself

	set.Jio(2, 1) //Jump if ONE (manuanlly)

	result, _ := set.Fetch()
	require.Equal(t, inst4, result)
}

//jio r, offset is like jmp, but only jumps if register r is 1 ("jump if one", not odd).
func TestJioAgain(t *testing.T) {

	inst1 := Instruction{"inc", "a", 0}
	inst2 := Instruction{"jie", "a", 2}
	inst3 := Instruction{"tpl", "a", 0}
	inst4 := Instruction{"inc", "a", 0}

	set := InstructionSet{}

	set.Add(inst1)
	set.Add(inst2)
	set.Add(inst3)
	set.Add(inst4)

	set.Fetch() //the increment

	set.Jio(2, 11) //Jump if ONE

	result, _ := set.Fetch()
	require.Equal(t, inst2, result)
}

func TestFetch(t *testing.T) {

	inst1 := Instruction{"inc", "a", 0}
	inst2 := Instruction{"hlf", "a", 0}

	set := InstructionSet{}

	set.Add(inst1)
	set.Add(inst2)

	result, _ := set.Fetch()
	require.Equal(t, inst1, result)

	result2, _ := set.Fetch()
	require.Equal(t, inst2, result2)
}

func TestOffset(t *testing.T) {

	inst1 := Instruction{"inc", "a", 0}
	inst2 := Instruction{"hlf", "a", 0}
	inst3 := Instruction{"tpl", "a", 0}
	inst4 := Instruction{"xxx", "a", 0}

	set := InstructionSet{}

	set.Add(inst1)
	set.Add(inst2)
	set.Add(inst3)
	set.Add(inst4)

	set.Fetch()

	set.offset(2) //Jump over inst2 and inst3

	result2, _ := set.Fetch()
	require.Equal(t, inst4, result2)
}

func TestExecute(t *testing.T) {

	//	inc a
	//	jio a, +2
	//	tpl a
	//	inc a

	inst1 := Instruction{"inc", "a", 0}
	inst2 := Instruction{"jio", "a", 2}
	inst3 := Instruction{"tpl", "a", 0}
	inst4 := Instruction{"inc", "a", 0}

	set := InstructionSet{}

	set.Add(inst1)
	set.Add(inst2)
	set.Add(inst3)
	set.Add(inst4)

	registers := make(map[string]*Register)
	registers["a"] = &Register{"a", 0}
	registers["b"] = &Register{"b", 0}

	finalRegisters := execute(&set, registers)

	expected := 2

	require.Equal(t, expected, finalRegisters["a"].Value)
}

const day22data = `jio a, +19
inc a
tpl a
inc a
tpl a
inc a
tpl a
tpl a
inc a
inc a
tpl a
tpl a
inc a
inc a
tpl a
inc a
inc a
tpl a
jmp +23
tpl a
tpl a
inc a
inc a
tpl a
inc a
inc a
tpl a
inc a
tpl a
inc a
tpl a
inc a
tpl a
inc a
inc a
tpl a
inc a
inc a
tpl a
tpl a
inc a
jio a, +8
inc b
jie a, +4
tpl a
inc a
jmp +2
hlf a
jmp -7`
