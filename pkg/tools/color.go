package tools

import "github.com/fatih/color"

var (
	Yellow    = color.New(color.FgYellow).SprintFunc()
	Magenta   = color.New(color.FgMagenta).SprintFunc()
	Red       = color.New(color.FgRed).SprintFunc()
	HiMagenta = color.New(color.FgHiMagenta).SprintFunc()
	Green     = color.New(color.FgGreen).SprintFunc()
)
