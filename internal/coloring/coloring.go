// Package coloring provides utilities for colorizing matching substrings in ASCII art.
// It handles the translation between character indexes in plain text and
// column offsets in the rendered ASCII representation.
package coloring

import (
	"strings"
)

// Reset is the ANSI escape sequence to stop applying color.
const Reset = "\033[0m"

// ApplyColor applies ANSI color codes to occurrences of a substring within ASCII art.
// It maps text-based matching positions to ASCII art column widths and inserts
// the provided colorCode at the start of matches and the Reset code at the end.
func ApplyColor(asciiArt []string, text, substring, colorCode string, charWidths []int) []string {
	if len(asciiArt) == 0 || len(charWidths) == 0 || len(text) == 0 || substring == "" {
		return asciiArt
	}

	positions := findPositions(text, substring)
	result := make([]string, len(asciiArt))

	for i, line := range asciiArt {
		result[i] = colorLine(line, positions, charWidths, colorCode)
	}

	return result
}

// colorLine processes a single line of ASCII art to apply color codes.
func colorLine(line string, positions []bool, charWidths []int, colorCode string) string {
	var builder strings.Builder
	offset := 0

	for idx, width := range charWidths {
		if offset >= len(line) {
			break
		}

		end := offset + width
		if end > len(line) {
			end = len(line)
		}

		isStart := positions[idx] && (idx == 0 || !positions[idx-1])
		isEnd := positions[idx] && (idx == len(positions)-1 || !positions[idx+1])

		if isStart {
			builder.WriteString(colorCode)
		}

		builder.WriteString(line[offset:end])

		if isEnd {
			builder.WriteString(Reset)
		}

		offset = end
	}

	if offset < len(line) {
		builder.WriteString(line[offset:])
	}

	return builder.String()
}

// findPositions identifies which character indexes in the text match the substring.
// It returns a boolean slice of the same length as text where true indicates a match.
func findPositions(text, substring string) []bool {
	positions := make([]bool, len(text))
	subLen := len(substring)

	for i := 0; i <= len(text)-subLen; i++ {
		if text[i:i+subLen] == substring {
			for j := 0; j < subLen; j++ {
				positions[i+j] = true
			}
		}
	}
	return positions
}
