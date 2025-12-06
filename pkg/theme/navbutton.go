package theme

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

// NavButton is a custom navigation button with text (can include icon with newlines)
type NavButton struct {
	widget.BaseWidget
	Text     string
	Active   bool
	OnTapped func()
}

// NewNavButton creates a new navigation button
// text can include icon and label with newline, e.g., "🏠\nHome"
func NewNavButton(text string, active bool, onTapped func()) *NavButton {
	btn := &NavButton{
		Text:     text,
		Active:   active,
		OnTapped: onTapped,
	}
	btn.ExtendBaseWidget(btn)
	return btn
}

// Tapped handles tap events
func (b *NavButton) Tapped(_ *fyne.PointEvent) {
	if b.OnTapped != nil {
		b.OnTapped()
	}
}

// SetActive updates the active state of the button
func (b *NavButton) SetActive(active bool) {
	b.Active = active
	b.Refresh()
}

// CreateRenderer creates the renderer for the navigation button
func (b *NavButton) CreateRenderer() fyne.WidgetRenderer {
	label := widget.NewLabel(b.Text)
	label.Alignment = fyne.TextAlignCenter
	label.TextStyle = fyne.TextStyle{Bold: false}

	content := container.NewCenter(label)

	// Create background with proper color
	var bgColor fyne.ThemeColorName
	if b.Active {
		bgColor = theme.ColorNamePrimary
	} else {
		bgColor = ColorNameNavBar
	}
	bg := canvas.NewRectangle(fyne.CurrentApp().Settings().Theme().Color(bgColor, fyne.CurrentApp().Settings().ThemeVariant()))

	return &navButtonRenderer{
		button:  b,
		bg:      bg,
		label:   label,
		content: content,
	}
}

type navButtonRenderer struct {
	button  *NavButton
	bg      *canvas.Rectangle
	label   *widget.Label
	content *fyne.Container
}

func (r *navButtonRenderer) Layout(size fyne.Size) {
	r.bg.Resize(size)
	r.content.Resize(size)
}

func (r *navButtonRenderer) MinSize() fyne.Size {
	return r.content.MinSize()
}

func (r *navButtonRenderer) Refresh() {
	// Update text in case it changed
	r.label.SetText(r.button.Text)

	// Update background color based on active state
	if r.button.Active {
		r.bg.FillColor = fyne.CurrentApp().Settings().Theme().Color(theme.ColorNamePrimary, fyne.CurrentApp().Settings().ThemeVariant())
	} else {
		r.bg.FillColor = fyne.CurrentApp().Settings().Theme().Color(ColorNameNavBar, fyne.CurrentApp().Settings().ThemeVariant())
	}

	r.bg.Refresh()
	r.label.Refresh()
}

func (r *navButtonRenderer) Objects() []fyne.CanvasObject {
	return []fyne.CanvasObject{r.bg, r.content}
}

func (r *navButtonRenderer) Destroy() {}
