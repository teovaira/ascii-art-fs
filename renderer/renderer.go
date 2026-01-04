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
// Each character in the input is rendered as an 8-line ASCII block defined in the banner.
// Newline characters ('\n') in the input create separate ASCII-art blocks in the output.
//
// Input validation rules:
//   - Characters must be printable ASCII (range 32–126) or '\n'
//   - The banner map must contain all characters used in the input
//   - Each banner entry must consist of exactly 8 lines
//
// Parameters:
//   - input: The text to render. It may contain '\n' for multi-line ASCII output.
//   - banner: A map of runes to their 8-line ASCII art representations.
//
// Returns:
//   - A string containing the rendered ASCII art
//   - An error if the input contains invalid characters or the banner data is incomplete
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
	if len(parts) == 0 || len(parts) == 1 && parts[0] == "" {
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
