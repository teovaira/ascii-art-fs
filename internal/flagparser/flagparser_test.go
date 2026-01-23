package flagparser_test

import (
	"ascii-art-color/internal/flagparser"
	"testing"
)

func TestParseArgs_NoArguments(t *testing.T) {
	args := []string{"program"}
	err := flagparser.ParseArgs(args)
	if err == nil {
		t.Errorf("Error was expected")
	}

}
func TestParseArgs_TooManyArgs(t *testing.T) {
	args := []string{"program",
		"banner",
		"--color=red",
		"substring",
		"some text",
		"EXTRA"}
	err := flagparser.ParseArgs(args)
	if err == nil {
		t.Errorf("Error too many args")
	}
}
