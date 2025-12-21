package renderer

import (
	"testing"
)

func TestEmptyInput(t *testing.T) {
	input := ""
	banner := map[rune][]string{}
	output := rendererASCII(input, banner)
	if input != output {
		t.Errorf("Error")
	}
}
func TestSingleCharacter(t *testing.T) {
	input := "A"
	expected := `A1
A2
A3
A4
A5
A6
A7
A8`
	banner := map[rune][]string{
		'A': {
			"A1",
			"A2",
			"A3",
			"A4",
			"A5",
			"A6",
			"A7",
			"A8",
		},
	}
	output := rendererASCII(input, banner)
	if expected != output {
		t.Error("Error")
	}
}
func TestMultipleCharacters(t *testing.T) {
	input := "AB"
	expected := `A1B1
A2B2
A3B3
A4B4
A5B5
A6B6
A7B7
A8B8`
	banner := map[rune][]string{
		'A': {
			"A1",
			"A2",
			"A3",
			"A4",
			"A5",
			"A6",
			"A7",
			"A8",
		},
		'B': {
			"B1",
			"B2",
			"B3",
			"B4",
			"B5",
			"B6",
			"B7",
			"B8",
		},
	}
	output := rendererASCII(input, banner)
	if expected != output {
		t.Error("Error")
	}

}
