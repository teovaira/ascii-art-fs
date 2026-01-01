package renderer

import (
	"fmt"
	"strings"
)

const bannerHeight = 8

func RendererASCII(input string, banner map[rune][]string) (string, error) {
	result := ""
	parts := strings.Split(input, "\n")
	// Remove trailing empty string only if input ends with \n
	if len(parts) > 0 && parts[len(parts)-1] == "" {
		parts = parts[0 : len(parts)-1]
	}
	// Handle special case: input is just empty or just "\n"
	if len(parts) == 0 || len(parts) == 1 && parts[0] == "" {
		return result, nil
	}
	if len(banner) == 0 {
		return result, fmt.Errorf("banner is empty")
	}
	for p, line := range parts {
		// Handle empty lines(from consecutive \n\n)
		if line == "" {
			result += "\n"
			continue
		}

		// Render each line of the ASCII art
		for i := 0; i < bannerHeight; i++ {
			for _, ch := range line {
				value, err := characterValidation(ch, banner)
				if err != nil {
					return "", err
				}
				result += value[i]
			}
			result += "\n"
		}
		// Don't add extra newline after the last part
		if p == len(parts)-1 {
			// Remove the last newline character
			result = result[:len(result)-1]
		}
	}
	return result, nil
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
