--- Day 23: Opening the Turing Lock ---
==



Tests
---
        go test


Part One
---
        input := instructionSet // The multiline instruction set
        initialValueForRegisterA := 0
        registerA, registerB := Run(tinput, initialValueForRegisterA)


Part One
---
        input := instructionSet // The multi-line instruction set
        initialValueForRegisterA := 0
        _, registerB := Run(tinput, initialValueForRegisterA)  //expect 184

Part Two
---
        input := instructionSet // The multi-line instruction set
        initialValueForRegisterA := 1
        _, registerB := Run(tinput, initialValueForRegisterA)  //expect 231

