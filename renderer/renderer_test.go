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
