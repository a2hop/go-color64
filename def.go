package color64

// Color is an integer type for color codes.
type Color int

// Named color constants (using standard 256-color codes).
const (
	CBlack   Color = 16
	CRed     Color = 196
	CGreen   Color = 46
	CYellow  Color = 226
	CBlue    Color = 21
	CMagenta Color = 201
	CCyan    Color = 51
	CWhite   Color = 231

	// Additional colors to reach 64 total
	// Basic colors (16-31)
	CNavy      Color = 17
	CDarkGreen Color = 22
	CTeal      Color = 23
	CDarkRed   Color = 52
	CPurple    Color = 53
	COlive     Color = 58
	CSilver    Color = 145
	CGray      Color = 240
	CMaroon    Color = 88
	CLime      Color = 118
	CAqua      Color = 45
	CFuchsia   Color = 165
	CBrown     Color = 130
	COrange    Color = 208
	CPink      Color = 175
	CLavender  Color = 183

	// More colors (32-47)
	CIndigo      Color = 54
	CTurquoise   Color = 30
	CCoral       Color = 209
	CSalmon      Color = 173
	CKhaki       Color = 143
	CPeach       Color = 215
	CBeige       Color = 180
	CMint        Color = 121
	CSeaGreen    Color = 85
	CSkyBlue     Color = 117
	CRoyalBlue   Color = 62
	CIndianRed   Color = 167
	CPlum        Color = 96
	CCrimson     Color = 160
	CGold        Color = 220
	CForestGreen Color = 28

	// More colors (48-63)
	CViolet        Color = 99
	CSlateBlue     Color = 61
	CPaleGreen     Color = 120
	CPaleTurquoise Color = 159
	CPaleViolet    Color = 183
	CSteelBlue     Color = 67
	CCadetBlue     Color = 73
	CPowderBlue    Color = 152
	CFirebrick     Color = 124
	CDarkViolet    Color = 92
	CDarkOrange    Color = 166
	CDarkCyan      Color = 36
	CDarkMagenta   Color = 91
	CDeepPink      Color = 198
	CMediumBlue    Color = 20
	CDodgerBlue    Color = 33
)

// ColorNames maps color names to their Color value.
var colorNames = map[string]Color{
	"black":         CBlack,
	"red":           CRed,
	"green":         CGreen,
	"yellow":        CYellow,
	"blue":          CBlue,
	"magenta":       CMagenta,
	"cyan":          CCyan,
	"white":         CWhite,
	"navy":          CNavy,
	"darkgreen":     CDarkGreen,
	"teal":          CTeal,
	"darkred":       CDarkRed,
	"purple":        CPurple,
	"olive":         COlive,
	"silver":        CSilver,
	"gray":          CGray,
	"maroon":        CMaroon,
	"lime":          CLime,
	"aqua":          CAqua,
	"fuchsia":       CFuchsia,
	"brown":         CBrown,
	"orange":        COrange,
	"pink":          CPink,
	"lavender":      CLavender,
	"indigo":        CIndigo,
	"turquoise":     CTurquoise,
	"coral":         CCoral,
	"salmon":        CSalmon,
	"khaki":         CKhaki,
	"peach":         CPeach,
	"beige":         CBeige,
	"mint":          CMint,
	"seagreen":      CSeaGreen,
	"skyblue":       CSkyBlue,
	"royalblue":     CRoyalBlue,
	"indianred":     CIndianRed,
	"plum":          CPlum,
	"crimson":       CCrimson,
	"gold":          CGold,
	"forestgreen":   CForestGreen,
	"violet":        CViolet,
	"slateblue":     CSlateBlue,
	"palegreen":     CPaleGreen,
	"paleturquoise": CPaleTurquoise,
	"paleviolet":    CPaleViolet,
	"steelblue":     CSteelBlue,
	"cadetblue":     CCadetBlue,
	"powderblue":    CPowderBlue,
	"firebrick":     CFirebrick,
	"darkviolet":    CDarkViolet,
	"darkorange":    CDarkOrange,
	"darkcyan":      CDarkCyan,
	"darkmagenta":   CDarkMagenta,
	"deeppink":      CDeepPink,
	"mediumblue":    CMediumBlue,
	"dodgerblue":    CDodgerBlue,
}
