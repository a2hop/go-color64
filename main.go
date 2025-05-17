package color64

import (
	"fmt"
)

// Int returns a Color value for a given color name
func Int(s string) Color {
	if color, ok := colorNames[s]; ok {
		return color
	}
	return CWhite // Default to white if not found
}

// String returns the color name for a given color value
func String(i int) string {
	for name, color := range colorNames {
		if int(color) == i {
			return name
		}
	}
	return ""
}

// Printf prints formatted text in the specified color code.
func Printf(color Color, format string, args ...any) {
	fmt.Printf("\033[38;5;%dm%s\033[0m\n", color, fmt.Sprintf(format, args...))
}
