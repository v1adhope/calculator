package colorize

import "github.com/fatih/color"

type ColorFMT struct { //only used colors
	Green, Magenta, Red, Yellow *color.Color
}

func New() *ColorFMT {
	return &ColorFMT{
		Green:   color.New(color.FgGreen, color.Bold),
		Magenta: color.New(color.FgHiMagenta, color.Bold),
		Red:     color.New(color.FgRed, color.Bold),
		Yellow:  color.New(color.FgYellow, color.Bold),
	}
}
