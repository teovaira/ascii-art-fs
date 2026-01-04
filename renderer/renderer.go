// Package renderer provides functionality for converting input text into ASCII art
// using predefined banner character definitions.
//
// The renderer processes printable ASCII characters (range 32–126) and renders each
// character as an 8-line ASCII representation based on the provided banner map.
// Newlines in the input are preserved and rendered as separate ASCII-art blocks.
//
// The package validates input characters and banner integrity, returning errors for
// unsupported characters or invalid banner data.
package renderer

import (
	"fmt"
	"strings"
)

const bannerHeight = 8

// RendererASCII converts an input string into ASCII art using the provided banner map.
//
// The input may contain printable ASCII characters (codes 32–126) and newline characters ('\n').
// Newlines are treated as line separators and are not rendered as characters.
//
// Behavior:
//   - Empty input or input consisting only of a single newline returns an empty result.
//   - Consecutive newline characters produce empty output lines, preserving input structure.
//   - Each non-empty input line is rendered as a block of bannerHeight ASCII-art rows.
//   - A trailing newline in the input does not produce an extra ASCII-art block.
//
// Validation:
//   - Returns an error if the input contains non-printable characters (excluding '\n').
//   - Returns an error if the banner map is empty.
//   - Returns an error if a character is missing from the banner or has an invalid height.
//
// Parameters:
//   - input: The text to render as ASCII art.
//   - banner: A map associating each rune with its ASCII-art representation,
//     where each value must contain exactly bannerHeight rows.
//
// Returns:
//   - The rendered ASCII-art string.
//   - An error if input validation or banner validation fails.
func RendererASCII(input string, banner map[rune][]string) (string, error) {
	var result strings.Builder
	for _, ch := range input {
		if ch == '\n' {
			continue
		}
		if ch < 32 || ch > 126 {
			return "", fmt.Errorf("not printable characters")
		}
	}
	parts := strings.Split(input, "\n")
	// Remove trailing empty string only if input ends with \n
	if len(parts) > 0 && parts[len(parts)-1] == "" {
		parts = parts[0 : len(parts)-1]
	}
	// Handle special case: input is just empty or just "\n"
	if input == "" || input == "\n" {
		return "", nil
	}
	if len(banner) == 0 {
		return "", fmt.Errorf("banner is empty")
	}

	for _, line := range parts {
		// Handle empty lines(from consecutive \n\n)
		if line == "" {
			result.WriteString("\n")
			continue
		}

		// Render each line of the ASCII art
		for i := 0; i < bannerHeight; i++ {
			for _, ch := range line {

				value, err := characterValidation(ch, banner)
				if err != nil {
					return "", err
				}
				result.WriteString(value[i])
			}
			result.WriteString("\n")
		}

	}
	output := result.String()
	// Don't add extra newline after the last part
	if output != "" && output[len(output)-1] == 10 {
		// Remove the last newline character
		output = output[:len(output)-1]
	}
	return output, nil
}

// characterValidation validates that a character exists in the banner map
// and that its ASCII-art representation has the correct height.
//
// Parameters:
//   - ch: The character to validate.
//   - banner: The banner map containing ASCII-art definitions.
//
// Returns:
//   - The ASCII-art rows corresponding to the character.
//   - An error if the character does not exist in the banner
//     or if it does not contain exactly bannerHeight rows.

func characterValidation(ch rune, banner map[rune][]string) ([]string, error) {

	value, exists := banner[ch]
	if exists == false {
		return []string{}, fmt.Errorf("the character does not exist in the banner")
	}
	if len(value) != bannerHeight {
		return []string{}, fmt.Errorf("The character does not have correct number of rows")
	}
	return value, nil
}
