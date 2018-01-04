package goyam

import "github.com/fatih/color"

/*
	Color definitions for the pretty output!
*/
var (
	Red       = color.New(color.FgRed).SprintFunc()
	Yellow    = color.New(color.FgYellow).SprintFunc()
	Green     = color.New(color.FgGreen).SprintFunc()
	GreenBold = color.New(color.Bold, color.FgGreen).SprintFunc()
	Cyan      = color.New(color.FgCyan).SprintFunc()
	Magenta   = color.New(color.FgMagenta).SprintFunc()
	Bold      = color.New(color.Bold).SprintFunc()
)
