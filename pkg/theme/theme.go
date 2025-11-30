package theme

import (
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/theme"
)

type SkillKonnectTheme struct {
	variant fyne.ThemeVariant
}

// NewSkillKonnectTheme creates a new theme with the specified variant (light or dark)
// Usage examples:
//
//	NewSkillKonnectTheme(theme.VariantLight)  // Light theme
//	NewSkillKonnectTheme(theme.VariantDark)   // Dark theme
func NewSkillKonnectTheme(variant fyne.ThemeVariant) fyne.Theme {
	return SkillKonnectTheme{variant: variant}
}

// Color lets you override specific named colors.
func (t SkillKonnectTheme) Color(name fyne.ThemeColorName, variant fyne.ThemeVariant) color.Color {
	switch name {
	case theme.ColorNamePrimary:
		return color.RGBA{R: 0x28, G: 0x7D, B: 0xF7, A: 0xFF} // brand blue
	case theme.ColorNameBackground:
		if t.variant == theme.VariantLight {
			return color.RGBA{R: 0xF5, G: 0xF5, B: 0xF5, A: 0xFF} // Light gray background
		}
		return color.RGBA{R: 0x1A, G: 0x1A, B: 0x1A, A: 0xFF} // Dark background
	case theme.ColorNameForeground:
		if t.variant == theme.VariantLight {
			return color.RGBA{R: 0x20, G: 0x20, B: 0x20, A: 0xFF} // Dark text for light mode
		}
		return color.RGBA{R: 0xE0, G: 0xE0, B: 0xE0, A: 0xFF} // Light text for dark mode
	}
	// Use the theme's variant (can be light or dark)
	return theme.DefaultTheme().Color(name, t.variant)
}

func (t SkillKonnectTheme) Font(style fyne.TextStyle) fyne.Resource {
	return theme.DefaultTheme().Font(style)
}

func (t SkillKonnectTheme) Size(name fyne.ThemeSizeName) float32 {
	return theme.DefaultTheme().Size(name)
}

func (t SkillKonnectTheme) Icon(name fyne.ThemeIconName) fyne.Resource {
	return theme.DefaultTheme().Icon(name)
}
