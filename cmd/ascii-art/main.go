// Package main provides the ASCII art generator CLI application.
//
// The application orchestrates the parser and renderer packages to convert text input
// into graphical ASCII art representations. It handles command-line argument parsing,
// banner file selection, and error reporting with appropriate exit codes.
//
// Responsibilities of this package:
//   - Parse command-line arguments
//   - Validate and resolve banner file paths
//   - Coordinate between parser and renderer
//   - Handle errors with appropriate exit codes
//
// Any invalid input, missing files, or rendering errors are reported to stderr.
package main

import (
	"errors"
	"fmt"
	"os"
	"strings"

	"ascii-art-color/internal/parser"
	"ascii-art-color/internal/renderer"
)

const (
	// Exit codes for different error scenarios.
	exitCodeUsageError  = 1
	exitCodeBannerError = 2
	exitCodeRenderError = 3
	exitCodeColorError  = 4

	// Default banner style.
	defaultBanner = "standard"
)

func main() {
	text, banner, err := ParseArgs(os.Args)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(exitCodeUsageError)
	}

	bannerPath, err := GetBannerPath(banner)
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error:", err)
		os.Exit(exitCodeUsageError)
	}

	charMap, err := parser.LoadBanner(bannerPath)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error loading banner file: %v\n", err)
		os.Exit(exitCodeBannerError)
	}

	result, err := renderer.RendererASCII(text, charMap)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error rendering text: %v\n", err)
		os.Exit(exitCodeRenderError)
	}

	fmt.Print(result)
}

// hasColorFlag checks whether any argument contains the --color= flag.
func hasColorFlag(args []string) bool {
	for _, arg := range args {
		if strings.HasPrefix(arg, "--color=") {
			return true
		}
	}
	return false
}

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
func ParseArgs(args []string) (text string, banner string, err error) {
	if len(args) < 2 {
		return "", "", errors.New("usage: go run . \"text\" [banner]")
	}

	if len(args) > 3 {
		return "", "", errors.New("too many arguments\nusage: go run . \"text\" [banner]")
	}

	text = strings.ReplaceAll(args[1], "\\n", "\n")

	if len(args) == 3 {
		banner = args[2]
	} else {
		banner = defaultBanner
	}

	return text, banner, nil
}

// GetBannerPath converts a banner name to its corresponding file path.
//
// The function validates the banner name against a predefined map of valid banners
// (standard, shadow, thinkertoy) and returns the appropriate file path in the testdata
// directory.
//
// Parameters:
//   - banner: The banner name to resolve.
//
// Returns:
//   - The file path to the banner file.
//   - An error if the banner name is invalid.
func GetBannerPath(banner string) (string, error) {
	bannerPaths := map[string]string{
		"standard":   "testdata/standard.txt",
		"shadow":     "testdata/shadow.txt",
		"thinkertoy": "testdata/thinkertoy.txt",
	}

	path, exists := bannerPaths[banner]
	if !exists {
		return "", fmt.Errorf("invalid banner name: %q\nValid options: standard, shadow, thinkertoy", banner)
	}

	return path, nil
}
