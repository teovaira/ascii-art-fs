package main

import (
	"os/exec"
	"strings"
	"testing"
)

// Integration test: Run actual program and check output
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
				// Should have 8 lines
				lines := strings.Count(output, "\n")
				return lines == 8
			},
		},
		{
			name:        "Empty string",
			args:        []string{""},
			expectError: false,
			checkOutput: func(output string) bool {
				// Should have 8 empty lines
				return output == "\n\n\n\n\n\n\n\n"
			},
		},
		{
			name:        "With shadow banner",
			args:        []string{"Hi", "shadow"},
			expectError: false,
			checkOutput: func(output string) bool {
				// Should have 8 lines
				lines := strings.Count(output, "\n")
				return lines == 8
			},
		},
		{
			name:        "No arguments - usage error",
			args:        []string{},
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
			// Build command
			args := append([]string{"run", "main.go"}, tt.args...)
			cmd := exec.Command("go", args...)
			
			// Run command
			output, err := cmd.CombinedOutput()
			
			// Check error expectation
			if tt.expectError && err == nil {
				t.Errorf("Expected error but got none")
			}
			if !tt.expectError && err != nil {
				t.Errorf("Unexpected error: %v\nOutput: %s", err, output)
			}
			
			// Check output if provided
			if !tt.expectError && tt.checkOutput != nil {
				if !tt.checkOutput(string(output)) {
					t.Errorf("Output check failed.\nOutput:\n%s", output)
				}
			}
		})
	}
}