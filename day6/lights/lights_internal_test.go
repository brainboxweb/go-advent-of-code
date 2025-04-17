package lights

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestParseInstruction(t *testing.T) {
	var tests = []struct {
		input    string
		expected instruction
	}{
		{"toggle 461,550 through 564,900", instruction{
			Action: "toggle",
			Start:  "461,550",
			End:    "564,900",
		}},
		{"turn off 461,550 through 564,900", instruction{
			Action: "off",
			Start:  "461,550",
			End:    "564,900",
		}},
		{"turn on 461,550 through 564,900", instruction{
			Action: "on",
			Start:  "461,550",
			End:    "564,900",
		}},
	}
	for _, test := range tests {
		if actual := parseInstruction(test.input); actual != test.expected {
			t.Errorf("Convert(%s) = %s, expected %s.",
				test.input, actual, test.expected)
		}
	}
}

func TestSwitchOn(t *testing.T) {
	ll := New(nil)
	key := "0,0"
	ll.switchOn(key)

	_, ok := ll.lights[key]

	expected := true
	require.Equal(t, expected, ok)
}

func TestSwitchOnOff(t *testing.T) {
	// var lights = make(map[string]light)
	ll := New(nil)
	key := "0,0"
	ll.lights[key] = light{}

	ll.switchOff(key)

	_, ok := ll.lights[key]

	expected := false
	require.Equal(t, expected, ok)
}

func TestToggle(t *testing.T) {
	ll := New(nil)
	key := "0,0"
	ll.lights[key] = light{}

	ll.toggle(key) //Should switch off
	_, ok := ll.lights[key]
	expected := false
	require.Equal(t, expected, ok)

	ll.toggle(key) //Should switch on
	_, ok2 := ll.lights[key]
	expected2 := true
	require.Equal(t, expected2, ok2)
}

func TestProcesslight(t *testing.T) {
	ll := New(nil)
	key := "0,0"
	ll.lights[key] = light{}

	ll.processlight(key, "off") //Should switch off
	_, ok := ll.lights[key]
	expected := false
	require.Equal(t, expected, ok)

	ll.processlight(key, "on") //Should switch on
	_, ok2 := ll.lights[key]
	expected2 := true
	require.Equal(t, expected2, ok2)

	ll.processlight(key, "toggle") //Should switch off
	_, ok3 := ll.lights[key]
	expected3 := false
	require.Equal(t, expected3, ok3)
}

func TestInstruction(t *testing.T) {
	ll := New(nil)
	input := `toggle 0,0 through 1,1`
	ll.processInstruction(input)

	_, ok := ll.lights["0,0"]
	expected := true
	require.Equal(t, expected, ok)

	_, ok2 := ll.lights["1,1"]
	expected2 := true
	require.Equal(t, expected2, ok2)
}

//---  Advanced Light

func TestSwitchOnAdjusted(t *testing.T) {
	ll := NewAdvanced(nil)
	key := "0,0"

	ll.switchOn(key)
	expected := light{Level: 1}
	require.Equal(t, expected, ll.lights[key])

	ll.switchOn(key) //Switch on again
	expected2 := light{Level: 2}
	require.Equal(t, expected2, ll.lights[key])
}

func TestSwitchOffAdjusted(t *testing.T) {
	ll := NewAdvanced(nil)
	key := "0,0"
	ll.lights[key] = light{Level: 2}

	ll.switchOff(key)
	expected := light{Level: 1}
	require.Equal(t, expected, ll.lights[key])

	ll.switchOff(key) //should delete it

	_, ok2 := ll.lights[key]
	expected2 := false
	require.Equal(t, expected2, ok2)
}

func TestToggleAdjusted(t *testing.T) {
	ll := NewAdvanced(nil)
	key := "0,0"

	ll.toggle(key)
	expected := light{Level: 2}
	require.Equal(t, expected, ll.lights[key])

	ll.toggle(key)
	expected2 := light{Level: 4}
	require.Equal(t, expected2, ll.lights[key])
}
