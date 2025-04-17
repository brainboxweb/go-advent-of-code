package wiring

import (
	"bufio"
	"bytes"
	"log"
	"regexp"
	"strconv"
	"strings"
)

const maxSignal = 65535

type Wire struct {
	ID     string
	Signal int
}

var wires = make(map[string]Wire)

func Run(input, wireName string) int {
	for k := range wires {
		delete(wires, k)
	}
	queue := Queue{}
	queue.Clear()
	b := bytes.NewBufferString(input)
	scanner := bufio.NewScanner(b)
	for scanner.Scan() {
		queue.Add(scanner.Text())
	}
OuterLoop:
	for {
		queueLen := queue.Len()
		if queueLen == 0 {
			break
		}
		for _, instruction := range queue.Instructions {
			result := parseInstruction(instruction)
			if result == true {
				queue.DeleteItem(instruction)
				continue OuterLoop
			}
		}
	}
	return wires[wireName].Signal
}

// --- Queue
type Queue struct {
	Instructions map[string]string
}

func (q *Queue) Add(s string) {

	if s == "" {
		return
	}
	//Avoid nil map panic
	if q.Instructions == nil {
		q.Instructions = make(map[string]string)
	}
	q.Instructions[s] = s
}

func (q *Queue) DeleteItem(k string) {
	delete(q.Instructions, k)
}

func (q *Queue) Clear() {
	q.Instructions = make(map[string]string)
}

func (q *Queue) Len() int {
	return len(q.Instructions)
}

func parse(phrase string) []string {
	tokens := strings.Split(phrase, " ")
	return tokens
}

func getOperator(tokens []string) string {
	switch {
	case tokens[0] == "NOT":
		return "NOT"
	case tokens[1] == "->":
		return "ASSIGN"
	case tokens[1] == "AND":
		return "AND"
	case tokens[1] == "OR":
		return "OR"
	case tokens[1] == "LSHIFT":
		return "LSHIFT"
	case tokens[1] == "RSHIFT":
		return "RSHIFT"
	}
	panic(tokens)
}

func parseInstruction(phrase string) bool {

	tokens := parse(phrase)
	operator := getOperator(tokens)

	switch operator {
	//"123 -> a"
	//b -> a"
	case "ASSIGN":
		return ASSIGN(tokens[0], tokens[2])
	//"x AND y -> d"
	//3 AND y -> z"
	case "AND":
		return AND(tokens[0], tokens[2], tokens[4])
	//"x OR y -> d"
	case "OR":
		return OR(tokens[0], tokens[2], tokens[4])
	case "LSHIFT":
		shiftCount, err := strconv.ParseUint(tokens[2], 10, 16)
		if err != nil {
			log.Fatal(err)
			panic("did not see that coming")
		}
		return LSHIFT(tokens[0], shiftCount, tokens[4])
	case "RSHIFT":
		shiftCount, err := strconv.ParseUint(tokens[2], 10, 16)
		if err != nil {
			panic("did not see that coming")
		}
		return RSHIFT(tokens[0], shiftCount, tokens[4])
	case "NOT":
		return NOT(tokens[1], tokens[3])
	default:
		panic("unknown operator")
	}
}

func NOT(inputKey string, targetKey string) bool {
	if _, ok := wires[inputKey]; !ok {
		return false
	}
	input := wires[inputKey].Signal
	signal := (maxSignal - input)
	wire := Wire{targetKey, signal}
	wires[targetKey] = wire
	return true
}

func ASSIGN(input, targetKey string) bool {
	//input can be wire (letters)  OR value (number)
	match, _ := regexp.MatchString("[a-z]{1,2}", input)
	if match == true {
		//its a wire
		inputKey := input
		if _, ok := wires[inputKey]; !ok {
			return false
		}
		inputSignal := wires[inputKey].Signal
		wire := Wire{targetKey, inputSignal}
		wires[targetKey] = wire
		return true
	}

	//its a signal. Assign it
	signal, err := strconv.Atoi(input)
	if err != nil { //it's a value.
		panic("Not expected")
	}
	wire := Wire{targetKey, signal}
	wires[targetKey] = wire
	return true
}

func AND(inputKey1 string, inputKey2 string, targetKey string) bool {
	var input1, input2 int

	//Keys can be numbers.
	//input can be wire (letters)  OR value (number)
	match, _ := regexp.MatchString("[a-z]{1,2}", inputKey1)
	if match == true {
		if _, ok := wires[inputKey1]; !ok {
			//		fmt.Println("token  missing ", inputKey1)
			return false
		}
		input1 = wires[inputKey1].Signal
	} else {

		inputOne, err := strconv.Atoi(inputKey1)
		if err != nil { //it's a value.
			panic("Not expected")
		}
		input1 = inputOne
	}

	//Keys can be numbers.
	//input can be wire (letters)  OR value (number)
	match2, _ := regexp.MatchString("[a-z]{1,2}", inputKey2)
	if match2 == true {
		if _, ok := wires[inputKey2]; !ok {
			//		fmt.Println("token  missing ", inputKey1)
			return false
		}
		input2 = wires[inputKey2].Signal

	} else {
		inputTwo, err := strconv.Atoi(inputKey2)
		if err != nil { //it's a value.
			panic("Not expected")
		}
		input2 = inputTwo
	}

	signal := input1 & input2 //bitwise and

	wire := Wire{targetKey, signal}
	wires[targetKey] = wire
	return true
}

func OR(inputKey1 string, inputKey2 string, targetKey string) bool {
	if _, ok := wires[inputKey1]; !ok {
		return false
	}
	if _, ok := wires[inputKey2]; !ok {
		return false
	}
	input1 := wires[inputKey1].Signal
	input2 := wires[inputKey2].Signal
	signal := input1 | input2 //bitwise and
	wire := Wire{targetKey, signal}
	wires[targetKey] = wire
	return true
}

func LSHIFT(inputKey string, shiftCount uint64, targetKey string) bool {
	if _, ok := wires[inputKey]; !ok {
		return false
	}
	input := wires[inputKey].Signal
	signal := input << shiftCount
	wire := Wire{targetKey, signal}
	wires[targetKey] = wire
	return true
}

func RSHIFT(inputKey string, shiftCount uint64, targetKey string) bool {
	if _, ok := wires[inputKey]; !ok {
		return false
	}
	input := wires[inputKey].Signal
	signal := input >> shiftCount
	wire := Wire{targetKey, signal}
	wires[targetKey] = wire
	return true
}
