package parser

import (
	"bufio"
	"os"
)

// Banner represents the ASCII-art data for all supported characters.
type Banner map[rune][]string

// LoadBanner will read a banner file (e.g. standard.txt) and return its parsed representation.
func LoadBanner(path string) (Banner, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var lines []string
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		return nil, err
	}

	banner := make(Banner)

	// ASCII printable characters from space (32) to tilde (126)
	runeCode := rune(32)
	i := 0
	for i+8 <= len(lines) && runeCode <= 126 {
		// take 8 lines for this character
		block := lines[i : i+8]
		banner[runeCode] = block
		runeCode++
		i += 9 // 8 lines + 1 empty separator line
	}
	return banner, nil
}
