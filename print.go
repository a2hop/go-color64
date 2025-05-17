package color64

import (
	"fmt"
	"sort"
	"strings"
)

type print struct{}

var Print = &print{}

// Named print methods for convenience.
func (p *print) Black(format string, args ...any)     { Printf(CBlack, format, args...) }
func (p *print) Red(format string, args ...any)       { Printf(CRed, format, args...) }
func (p *print) Green(format string, args ...any)     { Printf(CGreen, format, args...) }
func (p *print) Yellow(format string, args ...any)    { Printf(CYellow, format, args...) }
func (p *print) Blue(format string, args ...any)      { Printf(CBlue, format, args...) }
func (p *print) Magenta(format string, args ...any)   { Printf(CMagenta, format, args...) }
func (p *print) Cyan(format string, args ...any)      { Printf(CCyan, format, args...) }
func (p *print) White(format string, args ...any)     { Printf(CWhite, format, args...) }
func (p *print) Navy(format string, args ...any)      { Printf(CNavy, format, args...) }
func (p *print) DarkGreen(format string, args ...any) { Printf(CDarkGreen, format, args...) }
func (p *print) Teal(format string, args ...any)      { Printf(CTeal, format, args...) }
func (p *print) DarkRed(format string, args ...any)   { Printf(CDarkRed, format, args...) }
func (p *print) Purple(format string, args ...any)    { Printf(CPurple, format, args...) }
func (p *print) Olive(format string, args ...any)     { Printf(COlive, format, args...) }
func (p *print) Silver(format string, args ...any)    { Printf(CSilver, format, args...) }
func (p *print) Gray(format string, args ...any)      { Printf(CGray, format, args...) }
func (p *print) Maroon(format string, args ...any)    { Printf(CMaroon, format, args...) }
func (p *print) Lime(format string, args ...any)      { Printf(CLime, format, args...) }
func (p *print) Aqua(format string, args ...any)      { Printf(CAqua, format, args...) }
func (p *print) Fuchsia(format string, args ...any)   { Printf(CFuchsia, format, args...) }
func (p *print) Brown(format string, args ...any)     { Printf(CBrown, format, args...) }
func (p *print) Orange(format string, args ...any)    { Printf(COrange, format, args...) }
func (p *print) Pink(format string, args ...any)      { Printf(CPink, format, args...) }
func (p *print) Lavender(format string, args ...any)  { Printf(CLavender, format, args...) }

// Additional print methods for the remaining colors (32-63)
func (p *print) Indigo(format string, args ...any)        { Printf(CIndigo, format, args...) }
func (p *print) Turquoise(format string, args ...any)     { Printf(CTurquoise, format, args...) }
func (p *print) Coral(format string, args ...any)         { Printf(CCoral, format, args...) }
func (p *print) Salmon(format string, args ...any)        { Printf(CSalmon, format, args...) }
func (p *print) Khaki(format string, args ...any)         { Printf(CKhaki, format, args...) }
func (p *print) Peach(format string, args ...any)         { Printf(CPeach, format, args...) }
func (p *print) Beige(format string, args ...any)         { Printf(CBeige, format, args...) }
func (p *print) Mint(format string, args ...any)          { Printf(CMint, format, args...) }
func (p *print) SeaGreen(format string, args ...any)      { Printf(CSeaGreen, format, args...) }
func (p *print) SkyBlue(format string, args ...any)       { Printf(CSkyBlue, format, args...) }
func (p *print) RoyalBlue(format string, args ...any)     { Printf(CRoyalBlue, format, args...) }
func (p *print) IndianRed(format string, args ...any)     { Printf(CIndianRed, format, args...) }
func (p *print) Plum(format string, args ...any)          { Printf(CPlum, format, args...) }
func (p *print) Crimson(format string, args ...any)       { Printf(CCrimson, format, args...) }
func (p *print) Gold(format string, args ...any)          { Printf(CGold, format, args...) }
func (p *print) ForestGreen(format string, args ...any)   { Printf(CForestGreen, format, args...) }
func (p *print) Violet(format string, args ...any)        { Printf(CViolet, format, args...) }
func (p *print) SlateBlue(format string, args ...any)     { Printf(CSlateBlue, format, args...) }
func (p *print) PaleGreen(format string, args ...any)     { Printf(CPaleGreen, format, args...) }
func (p *print) PaleTurquoise(format string, args ...any) { Printf(CPaleTurquoise, format, args...) }
func (p *print) PaleViolet(format string, args ...any)    { Printf(CPaleViolet, format, args...) }
func (p *print) SteelBlue(format string, args ...any)     { Printf(CSteelBlue, format, args...) }
func (p *print) CadetBlue(format string, args ...any)     { Printf(CCadetBlue, format, args...) }
func (p *print) PowderBlue(format string, args ...any)    { Printf(CPowderBlue, format, args...) }
func (p *print) Firebrick(format string, args ...any)     { Printf(CFirebrick, format, args...) }
func (p *print) DarkViolet(format string, args ...any)    { Printf(CDarkViolet, format, args...) }
func (p *print) DarkOrange(format string, args ...any)    { Printf(CDarkOrange, format, args...) }
func (p *print) DarkCyan(format string, args ...any)      { Printf(CDarkCyan, format, args...) }
func (p *print) DarkMagenta(format string, args ...any)   { Printf(CDarkMagenta, format, args...) }
func (p *print) DeepPink(format string, args ...any)      { Printf(CDeepPink, format, args...) }
func (p *print) MediumBlue(format string, args ...any)    { Printf(CMediumBlue, format, args...) }
func (p *print) DodgerBlue(format string, args ...any)    { Printf(CDodgerBlue, format, args...) }

// PrintAvailableColors prints a list of all available color names
func (p *print) AvailableColors() {
	fmt.Println("Available colors:")

	// Sort the colors for consistent output
	var colorList []string
	for name := range colorNames {
		colorList = append(colorList, name)
	}
	sort.Strings(colorList)

	// Print each color using its own color
	for _, name := range colorList {
		Printf(colorNames[name], "• %s", name)
	}
}

// Also add the method to print a specific color by name
func (p *print) Color(name string, format string, args ...any) {
	color, ok := colorNames[strings.ToLower(name)]
	if !ok {
		color = CWhite // Default to white if not found
	}
	Printf(color, format, args...)
}
