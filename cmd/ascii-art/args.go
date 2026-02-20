package main

import (
	"errors"
	"strings"
)

// ParseArgs parses command-line arguments and extracts text and banner name.
//
// The function validates argument count, extracts the text argument, interprets
// escape sequences (like \n), and determines the banner name (defaulting to "standard"
// if not provided).
//
// Parameters:
//   - args: Command-line arguments slice (args[0] is program name).
//
// Returns:
//   - text: The text to render (with escape sequences interpreted).
//   - banner: The banner name to use.
//   - err: An error if argument validation fails.
const usageMsg = "Usage: go run . [STRING] [BANNER]\n\nEX: go run . something standard"

func ParseArgs(args []string) (text string, banner string, err error) {
	if len(args) < 2 {
		return "", "", errors.New(usageMsg)
	}

	if len(args) > 3 {
		return "", "", errors.New(usageMsg)
	}

	text = strings.ReplaceAll(args[1], "\\n", "\n")

	if len(args) == 3 {
		banner = args[2]
	} else {
		banner = defaultBanner
	}

	return text, banner, nil
}
