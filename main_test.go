package main

import (
	"testing"
)

// ============================================
// UNIT TESTS FOR ParseArgs FUNCTION
// ============================================

// Test 1: ParseArgs with no arguments (just program name)
func TestParseArgs_NoArguments(t *testing.T) {
	// Arrange
	args := []string{"./ascii-art"} // Just program name
	
	// Act
	_, _, err := ParseArgs(args)
	
	// Assert
	if err == nil {
		t.Error("Expected error for no arguments, got nil")
	}
	
	expectedMsg := "Usage: go run . \"text\" [banner]"
	if err.Error() != expectedMsg {
		t.Errorf("Expected error message: %q, got: %q", expectedMsg, err.Error())
	}
}

// Test 2: ParseArgs with just text (should default to standard)
func TestParseArgs_TextOnly(t *testing.T) {
	// Arrange
	args := []string{"./ascii-art", "Hello"}
	
	// Act
	text, banner, err := ParseArgs(args)
	
	// Assert
	if err != nil {
		t.Errorf("Expected no error, got: %v", err)
	}
	
	if text != "Hello" {
		t.Errorf("Expected text: 'Hello', got: %q", text)
	}
	
	if banner != "standard" {
		t.Errorf("Expected banner: 'standard', got: %q", banner)
	}
}

// Test 3: ParseArgs with text and banner name
func TestParseArgs_TextAndBanner(t *testing.T) {
	// Arrange
	args := []string{"./ascii-art", "Hello", "shadow"}
	
	// Act
	text, banner, err := ParseArgs(args)
	
	// Assert
	if err != nil {
		t.Errorf("Expected no error, got: %v", err)
	}
	
	if text != "Hello" {
		t.Errorf("Expected text: 'Hello', got: %q", text)
	}
	
	if banner != "shadow" {
		t.Errorf("Expected banner: 'shadow', got: %q", banner)
	}
}

// Test 4: ParseArgs with too many arguments
func TestParseArgs_TooManyArguments(t *testing.T) {
	// Arrange
	args := []string{"./ascii-art", "Hello", "shadow", "extra"}
	
	// Act
	_, _, err := ParseArgs(args)
	
	// Assert
	if err == nil {
		t.Error("Expected error for too many arguments, got nil")
	}
}

// Test 5: ParseArgs with all valid banner types
func TestParseArgs_AllBannerTypes(t *testing.T) {
	testCases := []struct {
		args           []string
		expectedBanner string
	}{
		{[]string{"prog", "Hi", "standard"}, "standard"},
		{[]string{"prog", "Hi", "shadow"}, "shadow"},
		{[]string{"prog", "Hi", "thinkertoy"}, "thinkertoy"},
	}
	
	for _, tc := range testCases {
		// Act
		_, banner, err := ParseArgs(tc.args)
		
		// Assert
		if err != nil {
			t.Errorf("Args %v: expected no error, got: %v", tc.args, err)
		}
		
		if banner != tc.expectedBanner {
			t.Errorf("Args %v: expected banner %q, got: %q", 
				tc.args, tc.expectedBanner, banner)
		}
	}
}

// Test 6: ParseArgs with empty string text
func TestParseArgs_EmptyStringText(t *testing.T) {
	// Arrange
	args := []string{"./ascii-art", ""}
	
	// Act
	text, banner, err := ParseArgs(args)
	
	// Assert
	if err != nil {
		t.Errorf("Expected no error for empty string, got: %v", err)
	}
	
	if text != "" {
		t.Errorf("Expected empty text, got: %q", text)
	}
	
	if banner != "standard" {
		t.Errorf("Expected banner: 'standard', got: %q", banner)
	}
}

// ============================================
// UNIT TESTS FOR GetBannerPath FUNCTION
// ============================================

// Test 7: GetBannerPath converts banner name to file path
func TestGetBannerPath_ValidBanners(t *testing.T) {
	testCases := []struct {
		banner       string
		expectedPath string
	}{
		{"standard", "testdata/standard.txt"},
		{"shadow", "testdata/shadow.txt"},
		{"thinkertoy", "testdata/thinkertoy.txt"},
	}
	
	for _, tc := range testCases {
		// Act
		path, err := GetBannerPath(tc.banner)
		
		// Assert
		if err != nil {
			t.Errorf("Banner %q: expected no error, got: %v", tc.banner, err)
		}
		
		if path != tc.expectedPath {
			t.Errorf("Banner %q: expected path %q, got: %q",
				tc.banner, tc.expectedPath, path)
		}
	}
}

// Test 8: GetBannerPath with invalid banner name
func TestGetBannerPath_InvalidBanner(t *testing.T) {
	// Arrange
	banner := "invalid"
	
	// Act
	_, err := GetBannerPath(banner)
	
	// Assert
	if err == nil {
		t.Error("Expected error for invalid banner, got nil")
	}
}