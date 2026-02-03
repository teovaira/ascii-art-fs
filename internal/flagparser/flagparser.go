// Package flagparser validates command-line arguments for the ascii-art-color program.
//
// Its responsibility is limited to validating input format and values.
// It does NOT perform rendering or coloring logic.
//
// It ensures:
//   - the correct number of arguments is provided
//   - the --color flag (if present) appears in the correct position
//   - only one --color flag is used
//   - the provided color value is valid (named color, RGB, or HEX)
//
// Any invalid input results in a usage error.
package flagparser

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

// Argument count boundaries according to the project specification.
const (
	minimumArgs = 2
	maximumArgs = 5
)

// errUsage is the single user-facing error returned for any invalid CLI input.
// This keeps command-line output consistent and predictable.
var errUsage = errors.New("Usage: go run . [OPTION] [STRING]")

// allowedColors defines the set of supported named colors.
// The map uses struct{} values because only key existence matters.
var allowedColors = map[string]struct{}{
	"black":   {},
	"red":     {},
	"green":   {},
	"yellow":  {},
	"orange":  {},
	"blue":    {},
	"magenta": {},
	"cyan":    {},
	"white":   {},
	"purple":  {},
	"pink":    {},
	"brown":   {},
	"gray":    {},
}

// ParseArgs validates the provided command-line arguments.
//
// Rules enforced:
//   - the total number of arguments must be within allowed limits
//   - the --color flag, if present, must be the second argument
//   - only one --color flag is allowed
//   - the color value must be valid (named, RGB, or HEX)
//
// ParseArgs returns errUsage for any invalid argument format.
func ParseArgs(args []string) error {
	colorFlagCount := 0

	// Validate argument count boundaries.
	if len(args) < minimumArgs {
		return errUsage
	}
	if len(args) > maximumArgs {
		return errUsage
	}

	// Any flag-like argument must be the --color flag.
	if strings.HasPrefix(args[1], "-") && !strings.HasPrefix(args[1], "--color=") {
		return errUsage
	}

	// Scan arguments to detect the --color flag and enforce its position.
	for i := 1; i < len(args); i++ {
		if strings.HasPrefix(args[i], "--color=") {
			colorFlagCount++

			// Only one --color flag is allowed.
			if colorFlagCount > 1 {
				return errUsage
			}

			// The --color flag must appear as the second argument.
			if i != 1 {
				return errUsage
			}
		}
	}

	// If --color is provided, a string to color must follow.
	if strings.HasPrefix(args[1], "--color=") && len(args) < 3 {
		return errUsage
	}

	// Extract and validate the color value from the --color flag.
	if strings.HasPrefix(args[1], "--color=") {
		_, color, found := strings.Cut(args[1], "=")
		if !found || color == "" {
			return errUsage
		}

		// Delegate color format validation to a helper function.
		if err := validateColor(color); err != nil {
			return fmt.Errorf("%v\n%s", err, errUsage)
		}
	}

	return nil
}

// validateColor validates the color value provided to the --color flag.
//
// Supported formats:
//   - predefined named colors (e.g. red, green, blue)
//   - RGB format: rgb(r,g,b) where each value is between 0 and 255
//   - HEX format: #RRGGBB
//
// It returns a descriptive error for any invalid color format.
func validateColor(color string) error {

	// Check if the color is a supported named color.
	if _, exists := allowedColors[color]; exists {
		return nil
	}

	// Validate RGB format: rgb(r,g,b)
	if strings.HasPrefix(color, "rgb(") {
		inner := strings.TrimSuffix(strings.TrimPrefix(color, "rgb("), ")")
		parts := strings.Split(inner, ",")
		if len(parts) != 3 {
			return errors.New("invalid RGB format: expected rgb(r,g,b)")
		}

		// Validate each RGB component.
		for i := 0; i < len(parts); i++ {
			digits, err := strconv.Atoi(parts[i])
			if err != nil {
				return errors.New("invalid RGB value: must be a number")
			}
			if digits < 0 || digits > 255 {
				return errors.New("invalid RGB value: must be between 0 and 255")
			}
		}
		return nil
	}

	// Validate HEX format: #RRGGBB
	colorValue, checkHex := strings.CutPrefix(color, "#")
	if checkHex {
		if colorValue == "" {
			return errors.New("invalid HEX color: missing hexadecimal value")
		}
		if len(colorValue) != 6 {
			return errors.New("invalid HEX color: expected 6 hexadecimal characters")
		}

		// Ensure all characters are valid hexadecimal digits.
		for _, ch := range colorValue {
			if !(ch >= '0' && ch <= '9') &&
				!(ch >= 'a' && ch <= 'f') &&
				!(ch >= 'A' && ch <= 'F') {
				return errors.New("invalid HEX color: contains non-hexadecimal character")
			}
		}
		return nil
	}

	// Any other format is unsupported.
	return errors.New("unsupported color format")
}
