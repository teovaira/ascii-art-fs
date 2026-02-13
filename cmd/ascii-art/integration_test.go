package main

import (
	"os/exec"
	"strings"
	"testing"
)

func TestMainProgram_Integration(t *testing.T) {
	tests := []struct {
		name        string
		args        []string
		expectError bool
		checkOutput func(string) bool
	}{
		{
			name:        "Hello with standard banner",
			args:        []string{"Hello"},
			expectError: false,
			checkOutput: func(output string) bool {
				return strings.Count(output, "\n") == 8
			},
		},
		{
			name:        "Empty string",
			args:        []string{""},
			expectError: false,
			checkOutput: func(output string) bool {
				return output == ""
			},
		},
		{
			name:        "With shadow banner",
			args:        []string{"Hi", "shadow"},
			expectError: false,
			checkOutput: func(output string) bool {
				return strings.Count(output, "\n") == 8
			},
		},
		{
			name:        "With thinkertoy banner",
			args:        []string{"Go", "thinkertoy"},
			expectError: false,
			checkOutput: func(output string) bool {
				return strings.Count(output, "\n") == 8
			},
		},
		{
			name:        "Multiple words with spaces",
			args:        []string{"Hello World"},
			expectError: false,
			checkOutput: func(output string) bool {
				return strings.Count(output, "\n") == 8 && len(output) > 0
			},
		},
		{
			name:        "Text with newline",
			args:        []string{"Hello\nWorld"},
			expectError: false,
			checkOutput: func(output string) bool {
				return strings.Count(output, "\n") == 16
			},
		},
		{
			name:        "No arguments - usage error",
			args:        []string{},
			expectError: true,
			checkOutput: nil,
		},
		{
			name:        "Too many arguments",
			args:        []string{"Hello", "standard", "extra"},
			expectError: true,
			checkOutput: nil,
		},
		{
			name:        "Invalid banner",
			args:        []string{"Hello", "invalid"},
			expectError: true,
			checkOutput: nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			args := append([]string{"run", "main.go"}, tt.args...)
			cmd := exec.Command("go", args...)
			output, err := cmd.CombinedOutput()

			if tt.expectError && err == nil {
				t.Errorf("Expected error but got none")
			}
			if !tt.expectError && err != nil {
				t.Errorf("Unexpected error: %v\nOutput: %s", err, output)
			}

			if !tt.expectError && tt.checkOutput != nil {
				if !tt.checkOutput(string(output)) {
					t.Errorf("Output check failed.\nOutput:\n%s", output)
				}
			}
		})
	}
}

func TestRunColorMode(t *testing.T) {
	tests := []struct {
		name        string
		args        []string
		expectError bool
		checkOutput func(string) bool
	}{
		{
			name: "full text colored red",
			args: []string{"--color=red", "hello"},
			checkOutput: func(output string) bool {
				return strings.Contains(output, "\033[38;2;255;0;0m") &&
					strings.Contains(output, "\033[0m") &&
					strings.Count(output, "\n") == 8
			},
		},
		{
			name: "substring colored",
			args: []string{"--color=red", "He", "Hello"},
			checkOutput: func(output string) bool {
				return strings.Contains(output, "\033[38;2;255;0;0m") &&
					strings.Count(output, "\n") == 8
			},
		},
		{
			name: "full text with shadow banner",
			args: []string{"--color=blue", "Hi", "shadow"},
			checkOutput: func(output string) bool {
				return strings.Contains(output, "\033[38;2;0;0;255m") &&
					strings.Count(output, "\n") == 8
			},
		},
		{
			name: "full text with thinkertoy banner",
			args: []string{"--color=green", "Go", "thinkertoy"},
			checkOutput: func(output string) bool {
				return strings.Contains(output, "\033[38;2;0;255;0m") &&
					strings.Count(output, "\n") == 8
			},
		},
		{
			name: "substring with banner",
			args: []string{"--color=green", "Go", "Hello Go", "thinkertoy"},
			checkOutput: func(output string) bool {
				return strings.Contains(output, "\033[38;2;0;255;0m") &&
					strings.Count(output, "\n") == 8
			},
		},
		{
			name: "hex color format",
			args: []string{"--color=#ff0000", "hello"},
			checkOutput: func(output string) bool {
				return strings.Contains(output, "\033[38;2;255;0;0m") &&
					strings.Count(output, "\n") == 8
			},
		},
		{
			name: "rgb color format",
			args: []string{"--color=rgb(0,255,0)", "hello"},
			checkOutput: func(output string) bool {
				return strings.Contains(output, "\033[38;2;0;255;0m") &&
					strings.Count(output, "\n") == 8
			},
		},
		{
			name: "multiline text",
			args: []string{"--color=red", "hello\\nworld"},
			checkOutput: func(output string) bool {
				return strings.Contains(output, "\033[38;2;255;0;0m") &&
					strings.Count(output, "\n") == 16
			},
		},
		{
			name: "special characters",
			args: []string{"--color=yellow", "(%&) ??"},
			checkOutput: func(output string) bool {
				return strings.Contains(output, "\033[38;2;255;255;0m") &&
					strings.Count(output, "\n") == 8
			},
		},
		{
			name: "substring not found in text",
			args: []string{"--color=red", "xyz", "Hello"},
			checkOutput: func(output string) bool {
				return !strings.Contains(output, "\033[38;2;") &&
					strings.Count(output, "\n") == 8
			},
		},
		{
			name: "single character text",
			args: []string{"--color=red", "A"},
			checkOutput: func(output string) bool {
				return strings.Contains(output, "\033[38;2;255;0;0m") &&
					strings.Count(output, "\n") == 8
			},
		},
		{
			name: "text with spaces",
			args: []string{"--color=red", "Hello World"},
			checkOutput: func(output string) bool {
				return strings.Contains(output, "\033[38;2;255;0;0m") &&
					strings.Count(output, "\n") == 8
			},
		},
		{
			name: "single character substring",
			args: []string{"--color=blue", "B", "RGB()"},
			checkOutput: func(output string) bool {
				return strings.Contains(output, "\033[38;2;0;0;255m") &&
					strings.Count(output, "\n") == 8
			},
		},
		{
			name:        "invalid color name",
			args:        []string{"--color=notacolor", "hello"},
			expectError: true,
		},
		{
			name:        "wrong flag format missing equals",
			args:        []string{"--color", "red", "hello"},
			expectError: true,
		},
		{
			name:        "empty color value",
			args:        []string{"--color=", "hello"},
			expectError: true,
		},
		{
			name:        "invalid hex format",
			args:        []string{"--color=#zzzzzz", "hello"},
			expectError: true,
		},
		{
			name:        "invalid rgb format",
			args:        []string{"--color=rgb(999,0,0)", "hello"},
			expectError: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			args := append([]string{"run", "main.go"}, tt.args...)
			cmd := exec.Command("go", args...)
			output, err := cmd.CombinedOutput()

			if tt.expectError {
				if err == nil {
					t.Errorf("expected error but got none\nOutput: %s", output)
				}
				return
			}
			if err != nil {
				t.Fatalf("unexpected error: %v\nOutput: %s", err, output)
			}
			if tt.checkOutput != nil && !tt.checkOutput(string(output)) {
				t.Errorf("output check failed\nOutput:\n%s", output)
			}
		})
	}
}

func TestMainProgram_RealBannerFiles(t *testing.T) {
	banners := []string{"standard", "shadow", "thinkertoy"}

	for _, banner := range banners {
		t.Run("Banner_"+banner, func(t *testing.T) {
			cmd := exec.Command("go", "run", "main.go", "ABC", banner)
			output, err := cmd.CombinedOutput()

			if err != nil {
				t.Errorf("Failed to run with %s banner: %v\nOutput: %s",
					banner, err, output)
			}

			if len(output) == 0 {
				t.Errorf("Expected output for banner %s, got empty", banner)
			}

			lines := strings.Count(string(output), "\n")
			if lines != 8 {
				t.Errorf("Expected 8 lines for banner %s, got %d", banner, lines)
			}
		})
	}
}

func TestMainProgram_ErrorHandling(t *testing.T) {
	errorTests := []struct {
		name     string
		args     []string
		errorMsg string
	}{
		{
			name:     "No arguments",
			args:     []string{},
			errorMsg: "usage:",
		},
		{
			name:     "Invalid banner",
			args:     []string{"Hello", "notexist"},
			errorMsg: "invalid banner",
		},
	}

	for _, tt := range errorTests {
		t.Run(tt.name, func(t *testing.T) {
			cmd := exec.Command("go", append([]string{"run", "main.go"}, tt.args...)...)
			output, err := cmd.CombinedOutput()

			if err == nil {
				t.Errorf("Expected error for %s, got none", tt.name)
			}

			if !strings.Contains(string(output), tt.errorMsg) {
				t.Errorf("Expected error message containing %q, got: %s",
					tt.errorMsg, output)
			}
		})
	}
}
