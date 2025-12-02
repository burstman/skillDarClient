package theme

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/widget"
)

// ThemedNavBar is a widget that automatically updates its color based on theme
type ThemedNavBar struct {
	widget.BaseWidget
	bg *canvas.Rectangle
}

// NewThemedNavBar creates a new theme-aware navigation bar background
func NewThemedNavBar() *ThemedNavBar {
	nav := &ThemedNavBar{}
	nav.ExtendBaseWidget(nav) //embed base widget functionality
	return nav
}

func (n *ThemedNavBar) CreateRenderer() fyne.WidgetRenderer {
	// Create background with current theme color
	n.bg = canvas.NewRectangle(fyne.CurrentApp().Settings().Theme().Color(ColorNameNavBar, fyne.CurrentApp().Settings().ThemeVariant()))
	return &themedNavBarRenderer{nav: n, bg: n.bg}
}

type themedNavBarRenderer struct {
	nav *ThemedNavBar
	bg  *canvas.Rectangle
}

func (r *themedNavBarRenderer) Layout(size fyne.Size) {
	r.bg.Resize(size)
}

func (r *themedNavBarRenderer) MinSize() fyne.Size {
	// Set a fixed height for the navigation bar
	return fyne.NewSize(0, 10) // Match the height set in main_screen.go
}

func (r *themedNavBarRenderer) Refresh() {
	// Update color when theme changes
	r.bg.FillColor = fyne.CurrentApp().Settings().Theme().Color(ColorNameNavBar, fyne.CurrentApp().Settings().ThemeVariant())
	canvas.Refresh(r.bg)
}

func (r *themedNavBarRenderer) Objects() []fyne.CanvasObject {
	return []fyne.CanvasObject{r.bg}
}

func (r *themedNavBarRenderer) Destroy() {}
