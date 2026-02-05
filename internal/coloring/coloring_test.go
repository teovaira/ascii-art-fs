package coloring_test

import (
	"ascii-art-color/internal/coloring"
	"testing"
)

func TestFindPositions(t *testing.T) {
	tests := []struct {
		name      string
		text      string
		substring string
		expected  []bool
	}{
		{
			name:      "single occurrence",
			text:      "hello",
			substring: "e",
			expected:  []bool{false, true, false, false, false},
		},
		{
			name:      "multi-character substring",
			text:      "kitten",
			substring: "kit",
			expected:  []bool{true, true, true, false, false, false},
		},
		{
			name:      "empty substring",
			text:      "kitten",
			substring: "",
			expected:  []bool{true, true, true, true, true, true},
		},
		{
			name:      "multiple occurrences",
			text:      "hello",
			substring: "ll",
			expected:  []bool{false, false, true, true, false},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			output := coloring.FindPositions(tt.text, tt.substring)

			if len(output) != len(tt.expected) {
				t.Fatalf(
					"expected length %d, got %d",
					len(tt.expected),
					len(output),
				)
			}

			for i := range tt.expected {
				if output[i] != tt.expected[i] {
					t.Errorf(
						"at index %d: expected %v, got %v",
						i,
						tt.expected[i],
						output[i],
					)
				}
			}
		})
	}
}
