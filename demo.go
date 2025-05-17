package color64

import (
	"fmt"
	"sort"
	"strings"
)

type demo struct{}

var Demo = &demo{}

func (d *demo) Run(args ...string) {
	if len(args) == 0 {
		d.PrintUsage()
		return
	}

	cmd := args[0]

	switch cmd {
	case "list", "colors":
		Print.AvailableColors()

	case "palette":
		d.PrintColorPalette()

	case "demo":
		d.PrintDemo()

	case "pick":
		d.DemoColorSelection()

	case "test":
		if len(args) > 1 {
			colorName := strings.ToLower(args[1])
			color := Int(colorName)
			Printf(color, "This is a test of the color '%s'", colorName)
		} else {
			fmt.Println("Usage: demo test <colorname>")
		}

	default:
		d.PrintUsage()
	}
}

// PrintColorPalette prints all the colors arranged in a grid
func (d *demo) PrintColorPalette() {
	fmt.Println("Color palette:")

	// Sort the colors for consistent output
	var colorList []string
	for name := range colorNames {
		colorList = append(colorList, name)
	}
	sort.Strings(colorList)

	// Print color codes in rows of 8
	for i := 0; i < len(colorList); i += 8 {
		for j := 0; j < 8 && i+j < len(colorList); j++ {
			colorName := colorList[i+j]
			color := colorNames[colorName]
			fmt.Printf("\033[38;5;%dm%-16s\033[0m", color, colorName)
		}
		fmt.Println()
	}
}

// PrintDemo demonstrates all the color functions
func (d *demo) PrintDemo() {
	fmt.Println("=== Color64 Demo ===")

	// Demonstrate basic colors using the Print struct
	Print.Black("This is Black text")
	Print.Red("This is Red text")
	Print.Green("This is Green text")
	Print.Yellow("This is Yellow text")
	Print.Blue("This is Blue text")
	Print.Magenta("This is Magenta text")
	Print.Cyan("This is Cyan text")
	Print.White("This is White text")

	// Demonstrate a few additional colors
	fmt.Println("\nSome additional colors:")
	Print.Orange("This is Orange text")
	Print.Purple("This is Purple text")
	Print.Pink("This is Pink text")
	Print.Turquoise("This is Turquoise text")
	Print.Gold("This is Gold text")
	Print.Lime("This is Lime text")

	// Demonstrate Printf with custom colors
	fmt.Println("\nCustom color examples:")
	Printf(CDeepPink, "This is %s with dynamic color", "text")
	Printf(CRoyalBlue, "Number: %d, String: %s", 42, "example")

	// Demonstrate Print.Color method
	fmt.Println("\nUsing Print.Color:")
	Print.Color("deeppink", "This is dynamic %s selection by name", "color")

	fmt.Println("\nFor a full list of colors, call Print.AvailableColors()")
	fmt.Println("To see the color palette, call Demo.PrintColorPalette()")
}

// DemoColorSelection demonstrates the color selection functionality
func (d *demo) DemoColorSelection() {
	fmt.Println("=== Color Selection Demo ===")

	// Demonstrate Bright color selection
	fmt.Println("\n• Bright colors by index:")
	for i := 0; i < 5; i++ {
		color := Pick.Bright(i)
		Printf(color, "Bright color at index %d", i)
	}

	// Demonstrate Dark color selection
	fmt.Println("\n• Dark colors by index:")
	for i := 0; i < 5; i++ {
		color := Pick.Dark(i)
		Printf(color, "Dark color at index %d", i)
	}

	// Demonstrate Random color selection
	fmt.Println("\n• Random colors:")
	for i := 0; i < 5; i++ {
		color := Pick.Random()
		Printf(color, "Random color #%d", i+1)
	}

	// Demonstrate RandomBright
	fmt.Println("\n• Random bright colors (5):")
	brightColors := Pick.RandomBright(5)
	for i, color := range brightColors {
		Printf(color, "Random bright color #%d", i+1)
	}

	// Demonstrate RandomDark
	fmt.Println("\n• Random dark colors (5):")
	darkColors := Pick.RandomDark(5)
	for i, color := range darkColors {
		Printf(color, "Random dark color #%d", i+1)
	}
}

func (d *demo) PrintUsage() {
	fmt.Println("Color64 Demo")
	fmt.Println("Usage:")
	fmt.Println("  demo             - Run the full demo")
	fmt.Println("  demo list        - List all available colors")
	fmt.Println("  demo colors      - Same as 'list'")
	fmt.Println("  demo palette     - Display the color palette")
	fmt.Println("  demo pick        - Demo color selection features")
	fmt.Println("  demo test <name> - Test a specific color by name")
}
