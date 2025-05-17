# go-color64

A Go package for printing colored text to the terminal using 64 named colors.

## Features

- Print text in 64 different named colors.
- Print using color names or by calling color-specific methods.
- List all available colors.

## Installation

```sh
go get github.com/a2hop/go-color64
```

## Usage

```go
import "github.com/a2hop/go-color64"
```

### Print with a specific color

```go
color64.Print.Red("This is red: %d\n", 123)
color64.Print.Blue("This is blue!\n")
color64.Print.Color("gold", "This is gold!\n")
```

### Format string support

All print methods support format strings similar to `fmt.Printf`:

```go
color64.Print.Green("Score: %d points", 100)
color64.Print.Yellow("Hello, %s! You have %.2f coins.", "Player", 45.75)
```

### Using a color by name

You can dynamically select colors at runtime using the `Color` method:

```go
// User-selected color
userColor := "magenta"
color64.Print.Color(userColor, "This text will be %s\n", userColor)

// Map status to colors
status := "error"
statusColors := map[string]string{
    "success": "green",
    "warning": "yellow",
    "error":   "red",
}
color64.Print.Color(statusColors[status], "Operation %s!\n", status)
```

### List all available colors

```go
color64.Print.AvailableColors()
```

## Available Colors

The package provides 64 named colors, accessible through their respective methods:

### Basic Colors
- `Print.Black()`
- `Print.Red()`
- `Print.Green()`
- `Print.Yellow()`
- `Print.Blue()`
- `Print.Magenta()`
- `Print.Cyan()`
- `Print.White()`

### Extended Colors
- `Print.Navy()`
- `Print.DarkGreen()`
- `Print.Teal()`
- `Print.DarkRed()`
- `Print.Purple()`
- `Print.Olive()`
- `Print.Silver()`
- `Print.Gray()`
- `Print.Maroon()`
- `Print.Lime()`
- `Print.Aqua()`
- `Print.Fuchsia()`
- `Print.Brown()`
- `Print.Orange()`
- `Print.Pink()`
- `Print.Lavender()`
- `Print.Indigo()`
- `Print.Turquoise()`
- `Print.Coral()`
- `Print.Salmon()`
- `Print.Khaki()`
- `Print.Peach()`
- `Print.Beige()`
- `Print.Mint()`
- `Print.SeaGreen()`
- `Print.SkyBlue()`
- `Print.RoyalBlue()`
- `Print.IndianRed()`
- `Print.Plum()`
- `Print.Crimson()`
- `Print.Gold()`
- `Print.ForestGreen()`
- `Print.Violet()`
- `Print.SlateBlue()`
- `Print.PaleGreen()`
- `Print.PaleTurquoise()`
- `Print.PaleViolet()`
- `Print.SteelBlue()`
- `Print.CadetBlue()`
- `Print.PowderBlue()`
- `Print.Firebrick()`
- `Print.DarkViolet()`
- `Print.DarkOrange()`
- `Print.DarkCyan()`
- `Print.DarkMagenta()`
- `Print.DeepPink()`
- `Print.MediumBlue()`
- `Print.DodgerBlue()`

## API

- `Print.<Color>(format string, args ...any)` — Print using a named color (e.g., `Print.Red`, `Print.Blue`, etc.).
- `Print.Color(name string, format string, args ...any)` — Print using a color by name (case-insensitive).
- `Print.AvailableColors()` — Print a list of all available color names.

## Example Program

```go
package main

import (
    "github.com/a2hop/go-color64"
)

func main() {
    // Display all available colors
    color64.Print.AvailableColors()
    
    // Simple examples
    color64.Print.Red("This is red text\n")
    color64.Print.Blue("This is %s text\n", "blue")
    
    // Using a different color for each word
    colors := []string{"Red", "Orange", "Yellow", "Green", "Blue", "Indigo", "Violet"}
    for _, color := range colors {
        color64.Print.Color(color, "%s ", color)
    }
    println() // New line
}
```

## License

MIT