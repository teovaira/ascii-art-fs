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

func TestParseNamedColors(t *testing.T) {
	tests := []struct {
		name    string
		spec    string
		want    RGB
		wantErr bool
	}{
		{"red", "red", RGB{255, 0, 0}, false},
		{"green", "green", RGB{0, 255, 0}, false},
		{"blue", "blue", RGB{0, 0, 255}, false},
		{"case insensitive", "RED", RGB{255, 0, 0}, false},
		{"unknown color", "blurple", RGB{}, true},
		{"hex red", "#ff0000", RGB{255, 0, 0}, false},
		{"orange", "orange", RGB{255, 165, 0}, false},
		{"purple", "purple", RGB{128, 0, 128}, false},
		{"pink", "pink", RGB{255, 192, 203}, false},
		{"brown", "brown", RGB{165, 42, 42}, false},
		{"gray", "gray", RGB{128, 128, 128}, false},
		{"rgb red", "rgb(255, 0, 0)", RGB{255, 0, 0}, false},
		{"rgb invalid count", "rgb(255)", RGB{}, true},
		{"rgb out of range", "rgb(300, 0, 0)", RGB{}, true},
		{"rgb red", "rgb(255, 0, 0)", RGB{255, 0, 0}, false},
		{"rgb spaces", "rgb( 255 , 0 , 0 )", RGB{255, 0, 0}, false},
		{"rgb invalid count", "rgb(255)", RGB{}, true},
		{"rgb out of range", "rgb(300, 0, 0)", RGB{}, true},
		{"rgb non-number", "rgb(a, 0, 0)", RGB{}, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Parse(tt.spec)
			if (err != nil) != tt.wantErr {
				t.Fatalf(`Parse(%q) error = %v, wantErr %t`, tt.spec, err, tt.wantErr)
			}
			if !tt.wantErr && got != tt.want {
				t.Fatalf(`Parse(%q) = %#v, want %#v`, tt.spec, got, tt.want)
			}
		})
	}
}

func TestANSI(t *testing.T) {
	tests := []struct {
		name string
		rgb  RGB
		want string
	}{
		{
			name: "red",
			rgb:  RGB{255, 0, 0},
			want: "\033[38;2;255;0;0m",
		},
		{
			name: "green",
			rgb:  RGB{0, 255, 0},
			want: "\033[38;2;0;255;0m",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := ANSI(tt.rgb)
			if got != tt.want {
				t.Fatalf("ANSI(%#v) = %q, want %q", tt.rgb, got, tt.want)
			}

		})
	}
}

func TestWrap(t *testing.T) {
	tests := []struct {
		name    string
		text    string
		spec    string
		want    string
		wantErr bool
	}{
		{
			name: "red text",
			text: "Hello",
			spec: "red",
			want: "\033[38;2;255;0;0mHello\033[0m",
		},
		{
			name: "hex red",
			text: "World",
			spec: "#ff0000",
			want: "\033[38;2;255;0;0mWorld\033[0m",
		},
		{
			name:    "unknown color error",
			text:    "Fail",
			spec:    "blurple",
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Wrap(tt.text, tt.spec)

			if (err != nil) != tt.wantErr {
				t.Fatalf("Wrap(%q, %q) error = %v, wantErr %t",
					tt.text, tt.spec, err, tt.wantErr)
			}

			if !tt.wantErr && got != tt.want {
				t.Fatalf("Wrap(%q, %q) = %q, want %q", tt.text,
					tt.spec, got, tt.want)
			}

		})
	}
}
