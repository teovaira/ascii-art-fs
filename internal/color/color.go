package color

// RGB represents a 24-bit color.
type RGB struct {
	R, G, B uint8
}

// Parse converts a color specification string into RGB.
// For now supports only the named color "red".
func Parse(spec string) (RGB, error) {
	if spec == "red" {
		return RGB{255, 0, 0}, nil
	}
	return RGB{}, nil // temporary, we will add proper errors later
}
