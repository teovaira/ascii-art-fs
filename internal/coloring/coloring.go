// Package coloring provides utilities for detecting
// character positions in text that should be colorized.
//
// The coloring package is responsible for identifying which
// characters in an input string belong to occurrences of a
// specified substring. The resulting position data is later
// consumed by the renderer to apply color formatting.
package coloring

// FindPositions determines which character positions in a text
// string belong to occurrences of a given substring.
//
// The function scans the input text and marks all character
// positions that are part of any substring match. Multiple and
// overlapping occurrences are supported. If the substring is
// empty, all character positions are marked as true.
//
// Parameters:
//   - text: The input string in which substring occurrences
//     are searched.
//   - substring: The substring to locate within the input text.
//     Matching is case-sensitive.
//
// Returns:
//   - A boolean slice with the same length as text, where each
//     index is true if the corresponding character belongs to
//     a substring occurrence, and false otherwise.
func FindPositions(text string, substring string) []bool {
	result := make([]bool, len(text))
	if len(substring) == 0 {
		for i := range result {
			result[i] = true
		}
	} else {
		for i := 0; i <= len(text)-len(substring); i++ {
			match := true

			for p := 0; p < len(substring); p++ {
				if text[i+p] != substring[p] {
					match = false
					break
				}
			}

			if match {
				for p := 0; p < len(substring); p++ {
					result[i+p] = true
				}
			}
		}
	}

	return result
}
