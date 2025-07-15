package text

import "fmt"

func FormatToGreen(s string) string {
	return Format(s, FormatOptions{Color: "green"})
}

type FormatOptions struct {
	Color string // green, red, yellow
	Bold  bool
}

func Format(s string, options FormatOptions) string {
	colorCodes := map[string]string{
		"black":   "30",
		"red":     "31",
		"green":   "32",
		"yellow":  "33",
		"blue":    "34",
		"magenta": "35",
		"cyan":    "36",
		"white":   "37",
	}
	colorCode, ok := colorCodes[options.Color]
	if !ok {
		// no color
		colorCode = "0"
	}
	style := ""
	if options.Bold {
		style = "1;"
	}
	return fmt.Sprintf("\033[%s%sm%s\033[0m", style, colorCode, s)
}
