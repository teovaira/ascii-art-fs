// Package flagparser validates command-line arguments for the ascii-art-color program.
//
// It ensures:
//   - correct number of arguments
//   - correct position and syntax of the --color flag
//   - valid color formats (named colors, RGB, or HEX)
//
// Any invalid input results in a usage error.
package flagparser

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

const (
	minimumArgs = 2
	maximumArgs = 5
)

var errUsage = errors.New("Usage: go run . [OPTION] [STRING]")

// ParseArgs validates the provided command-line arguments.
//
// Rules enforced:
//   - arguments count must be within allowed limits
//   - the --color flag, if present, must be the second argument
//   - only one --color flag is allowed
//   - the color value must be valid (named, RGB, or HEX)
//
// It returns an error if the arguments do not follow the expected format.
func ParseArgs(args []string) error {
	colorFlagCount := 0
	if len(args) < minimumArgs {
		return errUsage
	}
	if len(args) > maximumArgs {
		return errUsage
	}
	if err := validateColorFlagSyntax(args); err != nil {
		return errUsage
	}
	for i, arg := range args {
		if strings.HasPrefix(arg, "--color=") {
			colorFlagCount++
			if i != 1 {
				return errUsage
			}
			if colorFlagCount > 1 {
				return errUsage
			}
		}
	}
	if strings.HasPrefix(args[1], "--color=") && len(args) < 3 {
		return errUsage
	}
	if strings.HasPrefix(args[1], "--color=") {
		checkColorInTheFlag := strings.Split(args[1], "=")
		if len(checkColorInTheFlag) > 1 && checkColorInTheFlag[1] == "" {
			return errUsage
		}
		color := checkColorInTheFlag[1]
		if color != "" {
			if err := validColors(color); err != nil {
				return fmt.Errorf("%v\n%s", err, errUsage)
			}
		}
	}
	return nil

}

// validateColorFlagSyntax checks the syntactic correctness of the --color flag.
//
// It validates that:
//   - the flag starts with '--'
//   - the flag contains an '=' separator
//
// This function does not validate the color value itself.
func validateColorFlagSyntax(args []string) error {
	isFlag := strings.HasPrefix(args[1], "-")
	if isFlag {
		hasDoubleDash := strings.HasPrefix(args[1], "--")

		if !hasDoubleDash {
			return errUsage
		}

		hasEqual := strings.Contains(args[1], "=")
		if !hasEqual {
			return errUsage
		}

	}
	return nil
}

// validColors validates the color value provided to the --color flag.
//
// Supported formats:
//   - predefined color names (red, green, blue, etc.)
//   - RGB format: rgb(r,g,b) where each value is 0–255
//   - HEX format: #RRGGBB
//
// It returns a descriptive error for any invalid format.
func validColors(color string) error {
	allowedColors := map[string]bool{
		"red":     true,
		"green":   true,
		"yellow":  true,
		"orange":  true,
		"blue":    true,
		"magenta": true,
	}
	if _, exists := allowedColors[color]; exists {
		return nil
	}

	if strings.HasPrefix(color, "rgb(") {
		inner := strings.TrimSuffix(strings.TrimPrefix(color, "rgb("), ")")
		separatedText := strings.Split(inner, ",")
		if len(separatedText) != 3 {
			return errors.New("invalid RGB format: expected rgb(r,g,b)")
		}
		for i := 0; i < len(separatedText); i++ {
			digits, err := strconv.Atoi(separatedText[i])
			if err != nil {
				return errors.New("invalid RGB value: must be a number")
			}
			if digits < 0 || digits > 255 {

				return errors.New("invalid RGB value: must be between 0 and 255")

			}
		}
		return nil
	}
	colorValue, checkHex := strings.CutPrefix(color, "#")

	if checkHex {
		if colorValue == "" {
			return errors.New("invalid HEX color: missing hexadecimal value")
		}
		if len(colorValue) != 6 {
			return errors.New("invalid HEX color: expected 6 hexadecimal characters")
		}
		for _, ch := range colorValue {
			if !(ch >= '0' && ch <= '9') && !(ch >= 'a' && ch <= 'f') && !(ch >= 'A' && ch <= 'F') {
				return errors.New("invalid HEX color: contains non-hexadecimal character")
			}
		}

		return nil
	}
	return errors.New("unsupported color format")
}
