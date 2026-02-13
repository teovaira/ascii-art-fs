// Package main provides the ASCII art generator CLI application.
//
// The application orchestrates the parser, renderer, color, flagparser, and coloring
// packages to convert text input into graphical ASCII art representations, optionally
// with ANSI color support for full text or specific substrings.
//
// Usage:
//
//	go run . "text" [banner]
//	go run . --color=<color> "text" [banner]
//	go run . --color=<color> <substring> "text" [banner]
//
// Responsibilities of this package:
//   - Parse and validate command-line arguments
//   - Route between normal mode and color mode
//   - Validate and resolve banner file paths
//   - Coordinate between parser, renderer, and coloring
//   - Handle errors with appropriate exit codes
//
// Any invalid input, missing files, or rendering errors are reported to stderr.
package main

import (
	"ascii-art-color/internal/parser"
	"ascii-art-color/internal/renderer"
	"fmt"
	"os"
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
	if hasColorFlag(os.Args) {
		runColorMode(os.Args)
		return
	}

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
