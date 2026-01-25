package color

import "testing"

func TestParseNamedColor(t *testing.T) {
	got, err := Parse("red")
	if err != nil {
		t.Fatalf("Parse(\"red\") returned error: %v", err)
	}

	want := RGB{R: 255, G: 0, B: 0}
	if got != want {
		t.Fatalf("Parse(\"red\") = %#v, want %#v", got, want)
	}
}
