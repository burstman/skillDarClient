package main

import (
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/theme"
)

// SkillKonnectTheme is a minimal custom theme that forces a light appearance for key elements
// while delegating all other values to Fyne's default theme. Keep it small so future Fyne
// changes won't require large rewrites.
type SkillKonnectTheme struct{}

func (SkillKonnectTheme) Color(name fyne.ThemeColorName, variant fyne.ThemeVariant) color.Color {
	switch name {
	case theme.ColorNamePrimary:
		return color.RGBA{R: 0x28, G: 0x7D, B: 0xF7, A: 0xFF} // brand blue
	case theme.ColorNameBackground:
		return color.RGBA{R: 0xFF, G: 0xFF, B: 0xFF, A: 0xFF} // white background
	case theme.ColorNameInputBackground:
		return color.RGBA{R: 0xFF, G: 0xFF, B: 0xFF, A: 0xFF}
	case theme.ColorNamePlaceHolder:
		return color.RGBA{R: 0x88, G: 0x88, B: 0x88, A: 0xFF}
	case theme.ColorNameButton:
		return color.RGBA{R: 0xF0, G: 0xF0, B: 0xF0, A: 0xFF}
	case theme.ColorNameDisabled:
		return color.RGBA{R: 0x99, G: 0x99, B: 0x99, A: 0xFF}
	}

	// delegate everything else to default theme
	return theme.DefaultTheme().Color(name, variant)
}

func (SkillKonnectTheme) Font(style fyne.TextStyle) fyne.Resource {
	return theme.DefaultTheme().Font(style)
}

func (SkillKonnectTheme) Size(name fyne.ThemeSizeName) float32 {
	return theme.DefaultTheme().Size(name)
}

func (SkillKonnectTheme) Icon(name fyne.ThemeIconName) fyne.Resource {
	return theme.DefaultTheme().Icon(name)
}
