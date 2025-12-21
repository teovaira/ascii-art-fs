package renderer

func rendererASCII(input string, banner map[rune][]string) string {
	result := ""
	for i := 0; i < 8; i++ {
		for _, ch := range input {
			rows := banner[ch]

			result += rows[i]
		}
		if i != 7 {
			result += "\n"
		}
	}

	return result
}
