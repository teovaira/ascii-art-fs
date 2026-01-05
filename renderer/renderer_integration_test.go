// Package renderer_test contains integration tests for the renderer package.
// These tests verify ASCII art rendering using real banner files.
package renderer_test

import (
	"ascii-art/parser"
	"ascii-art/renderer"
	"strings"
	"testing"
)

// TestWithRealStandardBanner verifies that RendererASCII correctly renders
// input using the standard banner with various input cases.
func TestWithRealStandardBanner(t *testing.T) {
	banner, err := parser.LoadBanner("../testdata/standard.txt")
	if err != nil {
		t.Skipf("skipping integration test: %v", err)
	}

	tests := []struct {
		name      string
		input     string
		wantLines int
	}{
		{"simple word", "Hello", 9},
		{"with space", "Hello World", 9},
		{"with numbers", "Hello123", 9},
		{"single newline", "Hello\nWorld", 17},
		{"double newline", "A\n\nB", 18},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			output, err := renderer.RendererASCII(tt.input, banner)
			if err != nil {
				t.Fatalf("RendererASCII failed: %v", err)
			}

			lines := strings.Split(output, "\n")
			if len(lines) != tt.wantLines {
				t.Errorf("expected %d lines, got %d", tt.wantLines, len(lines))
			}
		})
	}
}

// TestWithRealShadowBanner verifies that RendererASCII renders correctly
// when using the shadow banner.
func TestWithRealShadowBanner(t *testing.T) {
	banner, err := parser.LoadBanner("../testdata/shadow.txt")
	if err != nil {
		t.Skipf("skipping integration test: %v", err)
	}

	output, err := renderer.RendererASCII("A", banner)
	if err != nil {
		t.Fatalf("RendererASCII failed: %v", err)
	}

	lines := strings.Split(output, "\n")
	if len(lines) != 9 {
		t.Errorf("expected 9 lines (8 content + empty from trailing newline), got %d", len(lines))
	}
}

// TestWithRealThinkertoyBanner verifies that RendererASCII renders correctly
// when using the thinkertoy banner.
func TestWithRealThinkertoyBanner(t *testing.T) {
	banner, err := parser.LoadBanner("../testdata/thinkertoy.txt")
	if err != nil {
		t.Skipf("skipping integration test: %v", err)
	}

	output, err := renderer.RendererASCII("Hello", banner)
	if err != nil {
		t.Fatalf("RendererASCII failed: %v", err)
	}

	lines := strings.Split(output, "\n")
	if len(lines) != 9 {
		t.Errorf("expected 9 lines (8 content + empty from trailing newline), got %d", len(lines))
	}
}
