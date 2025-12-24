package renderer

import "strings"

const bannerHeight = 8

func RendererASCII(input string, banner map[rune][]string) string {
	result := ""
	parts := strings.Split(input, "\n")
	if len(parts) > 0 && parts[len(parts)-1] == "" {
		parts = parts[0 : len(parts)-1]
	}
	if len(parts) == 1 && parts[0] == "" {
		return result
	}
	for p, line := range parts {

		for i := 0; i < bannerHeight; i++ {
			for _, ch := range line {
				rows := banner[ch]

				result += rows[i]

			}
			if i != bannerHeight-1 {
				result += "\n"
			}

		}
		if p != len(parts)-1 {
			result += "\n"
		}
	}
	return result
}
