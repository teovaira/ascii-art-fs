package flagparser_test

import (
	"ascii-art-color/internal/flagparser"
	"testing"
)

func TestParseArgs_NoArguments(t *testing.T) {
	args := []string{"program"}
	err := flagparser.ParseArgs(args)
	if err == nil {
		t.Errorf("expected usage error when no arguments are provided")
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
		t.Errorf("expected usage error when too many arguments are provided")
	}
}
func TestParseArgs_InvalidColorPrefix(t *testing.T) {
	args := []string{"program",
		"-color:black",
		"some text",
	}
	err := flagparser.ParseArgs(args)
	if err == nil {
		t.Errorf("expected usage error for invalid flag format")
	}
}
func TestParseArgs_FormatColor(t *testing.T) {
	args := []string{
		"program",
		"--color:red",
		"some text",
	}
	err := flagparser.ParseArgs(args)
	if err == nil {
		t.Errorf("exprected usage error for invalid --color format")
	}
}
func TestParseArgs_SingleStringAllowed(t *testing.T) {
	args := []string{
		"program",
		"text",
	}
	err := flagparser.ParseArgs(args)
	if err != nil {
		t.Errorf("unexpected error for valid single string input: %v", err)
	}
}
func TestParseArgs_FlagAndStringAllowed(t *testing.T) {
	args := []string{"program", "--color=red", "text"}
	err := flagparser.ParseArgs(args)
	if err != nil {
		t.Errorf("unexpected  error for valid color flag and string:%v", err)
	}
}

func TestParseArgs_ColorSubstringAllowed(t *testing.T) {
	args := []string{"program", "--color=red", "text", "substring"}
	err := flagparser.ParseArgs(args)
	if err != nil {
		t.Errorf("unexpected error for valid color flag,string and substring:%v", err)
	}
}
func TestParseArgs_MissingString(t *testing.T) {
	args := []string{"program", "--color=red"}
	err := flagparser.ParseArgs(args)
	if err == nil {
		t.Errorf("expected usage error when color flag is provided without a string")
	}
}
func TestParseArgs_MultipleFlags(t *testing.T) {
	args := []string{"program", "--color=red", "--color=blue", "text"}
	err := flagparser.ParseArgs(args)
	if err == nil {
		t.Errorf("expected usage error when multiple color flags are provided")
	}
}
func TestParseArgs_InvalidPositionForColorFlag(t *testing.T) {
	args := []string{"program", "text", "--color=red"}
	err := flagparser.ParseArgs(args)
	if err == nil {
		t.Errorf("expected usage error when color flag is in an invalid position")
	}
}
func TestParseArgs_SubstringMissingWhileStringExists(t *testing.T) {
	args := []string{"program", "--color=red", "text"}
	err := flagparser.ParseArgs(args)
	if err != nil {
		t.Errorf("unexpected error when substring is optional")
	}
}
func TestParseArgs_ValidRGBColor(t *testing.T) {
	args := []string{"program", "--color=rgb(255,0,0)", "text"}
	err := flagparser.ParseArgs(args)
	if err != nil {
		t.Errorf("unexpected error for valid RGB color:%v", err)
	}
}
func TestParseArgs_ValidHexColor(t *testing.T) {
	args := []string{"program", "--color=#ff0000", "text"}
	err := flagparser.ParseArgs(args)
	if err != nil {
		t.Errorf("unexpected error for valid HEX color:%v", err)
	}
}
func TestParseArgs_InvalidRGB_OutOfRange(t *testing.T) {
	args := []string{"program", "--color=rgb(300,0,0)", "text"}
	err := flagparser.ParseArgs(args)
	if err == nil {
		t.Errorf("expected error for RGB value out of range")
	}
}
func TestParseArgs_InvalidHexLength(t *testing.T) {
	args := []string{"program", "--color=#123", "text"}
	err := flagparser.ParseArgs(args)
	if err == nil {
		t.Errorf("expected error for invalid HEX length")
	}
}
func TestParseArgs_InvalidColorName(t *testing.T) {
	args := []string{"program", "--color=pink", "text"}
	err := flagparser.ParseArgs(args)
	if err == nil {
		t.Errorf("expected error for unsupported color name")
	}
}
