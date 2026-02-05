package coloring_test

import (
	"ascii-art-color/internal/coloring"
	"testing"
)

func TestFindPositions_SingleOccurrence(t *testing.T) {
	text := "hello"
	substring := "e"
	expected := []bool{false, true, false, false, false}
	output := coloring.FindPositions(text, substring)
	if len(output) != len(expected) {
		t.Fatalf("expected length %d,got %d", len(expected), len(output))
	}
	for i := range expected {
		if output[i] != expected[i] {
			t.Errorf(
				"at index %d: expected %v,got %v",
				i,
				expected[i],
				output[i],
			)
		}
	}
}
func TestFindPositions_MultiCharacterSubstring(t *testing.T) {
	text := "kitten"
	substring := "kit"
	expected := []bool{true, true, true, false, false, false}
	output := coloring.FindPositions(text, substring)
	if len(output) != len(expected) {
		t.Fatalf("expected length %d,got %d", len(expected), len(output))
	}
	for i := range expected {
		if output[i] != expected[i] {
			t.Errorf(
				"at index %d: expected %v,got %v",
				i,
				expected[i],
				output[i],
			)
		}
	}
}
func TestFindPositions_EmptySubstring(t *testing.T) {
	text := "kitten"
	substring := ""
	expected := []bool{true, true, true, true, true, true}
	output := coloring.FindPositions(text, substring)
	if len(output) != len(expected) {
		t.Fatalf("expected length %d,got %d", len(expected), len(output))
	}
	for i := range expected {
		if output[i] != expected[i] {
			t.Errorf("at index %d: expected %v,got %v",
				i,
				expected[i],
				output[i],
			)
		}
	}
}
