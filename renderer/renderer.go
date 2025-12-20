package renderer

func rendererASCII(input string, banner map[rune][]string) string {
	result := ""
	for _, ch := range input {
		rows := banner[ch]
		for i, row := range rows {
			result += row
			if i != len(rows)-1 {
				result += "\n"
			}
		}
	}
	return result
}
