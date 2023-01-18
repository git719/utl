// utils.go

package utl

import (
	"github.com/fatih/color"
)

var (
	Gra = color.New(color.FgWhite).SprintFunc()
	Whi = color.New(color.FgWhite, color.Bold).SprintFunc()
	Red = color.New(color.FgRed, color.Bold).SprintFunc()
	Blu = color.New(color.FgBlue, color.Bold).SprintFunc()
	Pur = color.New(color.FgMagenta, color.Bold).SprintFunc()
	Gre = color.New(color.FgGreen).SprintFunc()
	Yel = color.New(color.FgYellow).SprintFunc()
	Cya = color.New(color.FgCyan).SprintFunc()
)
