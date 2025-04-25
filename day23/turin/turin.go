package turin

import (
	"bufio"
	"bytes"
	"regexp"
	"strconv"
)

func Run(input string, initialValueForA int) (valA, valB int) {
	instructionSet := parseInstructions(input)

	registers := make(map[string]*Register)
	registers["a"] = &Register{"a", initialValueForA}
	registers["b"] = &Register{"b", 0}

	finalRegisters := execute(&instructionSet, registers)

	return finalRegisters["a"].Value, finalRegisters["b"].Value
}

func execute(set *InstructionSet, registers map[string]*Register) map[string]*Register {
	for {
		inst, ok := set.Fetch()

		if !ok {
			break
		}

		switch inst.action {
		case "hlf":
			hlf(registers[inst.registerName])
		case "inc":
			inc(registers[inst.registerName])
		case "tpl":
			tpl(registers[inst.registerName])
		case "jmp":
			set.Jump(inst.increment)
		case "jie":
			set.Jie(inst.increment, registers[inst.registerName].Value)
		case "jio":
			set.Jio(inst.increment, registers[inst.registerName].Value)
		default:
			panic("Unexpected " + inst.action)
		}
	}
	return registers
}

func parseInstructions(input string) InstructionSet {
	b := bytes.NewBufferString(input)
	scanner := bufio.NewScanner(b)

	set := InstructionSet{}

	for scanner.Scan() {
		text := scanner.Text()
		instruction := parse(text)
		set.Add(instruction)
	}
	return set
}

func parse(phrase string) Instruction {
	// Action
	var validAction = regexp.MustCompile(`([a-z]{3})`)
	action := validAction.FindString(phrase)

	// name of the register
	registerName := ""
	var validName = regexp.MustCompile(`\s([a-z])`)
	matches := validName.FindStringSubmatch(phrase)
	if len(matches) > 0 {
		registerName = matches[1]
	}

	// Incrememnt
	increment := 0
	var validIncrement = regexp.MustCompile(`([+-]\d+)`)
	match2 := validIncrement.FindString(phrase)
	if match2 != "" {
		i64, _ := strconv.ParseInt(match2, 10, 32)
		increment = int(i64)
	}

	instruction := Instruction{action, registerName, increment}

	return instruction
}

type Register struct {
	Name  string
	Value int
}

func hlf(reg *Register) {
	reg.Value /= 2
}

func inc(reg *Register) {
	reg.Value++
}

func tpl(reg *Register) {
	reg.Value *= 3
}

type Instruction struct {
	action       string
	registerName string
	increment    int
}

type InstructionSet struct {
	instructions map[int]Instruction
	currentIndex int
}

func (s *InstructionSet) Add(inst Instruction) {
	// nil map
	if s.instructions == nil {
		s.instructions = make(map[int]Instruction)
	}
	s.instructions[len(s.instructions)] = inst
}

func (s *InstructionSet) Fetch() (Instruction, bool) {
	_, ok := s.instructions[s.currentIndex]
	if !ok {
		return Instruction{}, false
	}
	inst := s.instructions[s.currentIndex]
	s.currentIndex++
	return inst, true
}

func (s *InstructionSet) Jump(count int) {
	s.offset(count - 1)
}

func (s *InstructionSet) Jie(count int, registerValue int) {
	if registerValue%2 == 0 {
		s.offset(count - 1)
	}
}

func (s *InstructionSet) Jio(count int, registerValue int) {
	if registerValue == 1 {
		s.offset(count - 1)
	}
}

func (s *InstructionSet) offset(count int) {
	s.currentIndex += count
}
