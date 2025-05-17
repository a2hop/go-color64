package color64

import (
	"math/rand"
	"sort"
	"time"
)

type pick struct {
	rng *rand.Rand
}

// Create a new pick instance with its own random number generator
var Pick = &pick{
	rng: rand.New(rand.NewSource(time.Now().UnixNano())),
}

// getBrightness returns a brightness value for a color (0-100)
// Higher values mean brighter colors
func getBrightness(c Color) int {
	intColor := int(c)

	// Basic colors (0-15)
	if intColor <= 7 {
		// The first 8 colors are the darker variants
		return 25
	} else if intColor <= 15 {
		// Colors 8-15 are brighter variants of 0-7
		return 60
	}

	// Color cube (16-231) - decompose into RGB components
	if intColor >= 16 && intColor <= 231 {
		// Convert color number to RGB (each 0-5)
		i := intColor - 16
		r := i / 36
		g := (i / 6) % 6
		b := i % 6

		// Calculate perceived brightness (human eyes are more sensitive to green)
		// Using a weighted formula: 0.299*R + 0.587*G + 0.114*B
		brightness := int((0.299*float64(r) + 0.587*float64(g) + 0.114*float64(b)) * 20)
		return brightness
	}

	// Grayscale (232-255)
	if intColor >= 232 {
		// Convert to brightness (0-23 → 0-100)
		return (intColor - 232) * 100 / 23
	}

	return 50 // Default middle brightness
}

// isBright determines if a color is considered bright
func isBright(c Color) bool {
	// Consider colors with brightness > 50 as bright
	return getBrightness(c) > 50
}

// Bright picks a bright color by index
func (p *pick) Bright(index int) Color {
	brightColors := []Color{}

	// Collect all bright colors
	for _, color := range colorNames {
		if isBright(color) {
			brightColors = append(brightColors, color)
		}
	}

	// Sort colors in REVERSE order of brightness (brightest first)
	sort.Slice(brightColors, func(i, j int) bool {
		return getBrightness(brightColors[i]) > getBrightness(brightColors[j])
	})

	// Return the indexed color, wrapping around if needed
	if len(brightColors) > 0 {
		return brightColors[index%len(brightColors)]
	}
	return CWhite // Default
}

// Dark picks a dark color by index
func (p *pick) Dark(index int) Color {
	darkColors := []Color{}

	// Collect all dark colors
	for _, color := range colorNames {
		if !isBright(color) {
			darkColors = append(darkColors, color)
		}
	}

	// Sort colors in order of darkness (darkest first)
	sort.Slice(darkColors, func(i, j int) bool {
		return getBrightness(darkColors[i]) < getBrightness(darkColors[j])
	})

	// Return the indexed color, wrapping around if needed
	if len(darkColors) > 0 {
		return darkColors[index%len(darkColors)]
	}
	return CBlack // Default
}

// Random picks a random color from all available colors
func (p *pick) Random() Color {
	colors := colorValues()
	if len(colors) == 0 {
		return CWhite // Default
	}
	return colors[p.rng.Intn(len(colors))]
}

// RandomBright returns a slice of randomly selected bright colors
func (p *pick) RandomBright(count int) []Color {
	brightColors := []Color{}

	// Collect all bright colors
	for _, color := range colorNames {
		if isBright(color) {
			brightColors = append(brightColors, color)
		}
	}

	// Shuffle the bright colors
	p.rng.Shuffle(len(brightColors), func(i, j int) {
		brightColors[i], brightColors[j] = brightColors[j], brightColors[i]
	})

	// Return requested number of colors, or all if count is larger
	if count > len(brightColors) {
		count = len(brightColors)
	}

	return brightColors[:count]
}

// RandomDark returns a slice of randomly selected dark colors
func (p *pick) RandomDark(count int) []Color {
	darkColors := []Color{}

	// Collect all dark colors
	for _, color := range colorNames {
		if !isBright(color) {
			darkColors = append(darkColors, color)
		}
	}

	// Shuffle the dark colors
	p.rng.Shuffle(len(darkColors), func(i, j int) {
		darkColors[i], darkColors[j] = darkColors[j], darkColors[i]
	})

	// Return requested number of colors, or all if count is larger
	if count > len(darkColors) {
		count = len(darkColors)
	}

	return darkColors[:count]
}

// colorValues returns a slice of all available colors
func colorValues() []Color {
	colors := make([]Color, 0, len(colorNames))
	for _, color := range colorNames {
		colors = append(colors, color)
	}
	return colors
}
