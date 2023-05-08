// utils.go

package utl

import (
	"fmt"
	"github.com/gookit/color"
)

var (
	Red = color.FgLightRed.Render
	Blu = color.FgLightBlue.Render
	Gre = color.FgGreen.Render
	Yel = color.FgYellow.Render
	Whi = color.FgWhite.Render
	Cya = color.FgCyan.Render
	Mag = color.FgLightMagenta.Render
	Gra = color.FgDarkGray.Render

	Red2 = color.FgRed.Render
	Blu2 = color.FgBlue.Render
	Gre2 = color.FgLightGreen.Render
	Yel2 = color.FgLightYellow.Render
	Whi2 = color.FgLightWhite.Render
	Cya2 = color.FgLightCyan.Render
	Mag2 = color.FgMagenta.Render
)

func PrintColorSamples() {
	fmt.Println("Red  " + Red("!@#$%#$%&#%*^2314897123589"))
	fmt.Println("Blu  " + Blu("!@#$%#$%&#%*^2314897123589"))
	fmt.Println("Gre  " + Gre("!@#$%#$%&#%*^2314897123589"))
	fmt.Println("Yel  " + Yel("!@#$%#$%&#%*^2314897123589"))
	fmt.Println("Whi  " + Whi("!@#$%#$%&#%*^2314897123589"))
	fmt.Println("Cya  " + Cya("!@#$%#$%&#%*^2314897123589"))
	fmt.Println("Mag  " + Mag("!@#$%#$%&#%*^2314897123589"))
	fmt.Println("Gra  " + Gra("!@#$%#$%&#%*^2314897123589"))

	fmt.Println("Red2 " + Red2("!@#$%#$%&#%*^2314897123589"))
	fmt.Println("Blu2 " + Blu2("!@#$%#$%&#%*^2314897123589"))
	fmt.Println("Gre2 " + Gre2("!@#$%#$%&#%*^2314897123589"))
	fmt.Println("Yel2 " + Yel2("!@#$%#$%&#%*^2314897123589"))
	fmt.Println("Whi2 " + Whi2("!@#$%#$%&#%*^2314897123589"))
	fmt.Println("Cya2 " + Cya2("!@#$%#$%&#%*^2314897123589"))
	fmt.Println("Mag2 " + Mag2("!@#$%#$%&#%*^2314897123589"))
}
