package ui

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

// CreateMainWindow creates and configures the main SkillKonnect window
func CreateMainWindow(a fyne.App) fyne.Window {
	// Main window
	w := a.NewWindow("SkillKonnect")
	// w.SetMaster() // Temporarily commented out to test window visibility

	// Simulate phone size on desktop (ignored on real phone)
	// Common Android phone sizes:
	// - iPhone SE: 375x667
	// - iPhone 12/13: 390x844
	// - Samsung Galaxy S21: 360x800
	// - Pixel 5: 393x851
	// - Average Android phone: ~360-410 width, ~800-900 height
	w.Resize(fyne.NewSize(390, 844)) // iPhone 12/13 size - good Android approximation

	// UI elements
	title := widget.NewLabel("Welcome to SkillKonnect")
	title.Alignment = fyne.TextAlignCenter
	title.TextStyle = fyne.TextStyle{Bold: true}

	subtitle := widget.NewLabel("Connect skills, build networks")
	subtitle.Alignment = fyne.TextAlignCenter

	getStartedBtn := widget.NewButton("Please Login to Begin", func() {
		pop := widget.NewPopUp(widget.NewLabel("Hello from mobile!"), w.Canvas())
		pop.Show()
	})

	// Layout
	content := container.NewVBox(
		layout.NewSpacer(),
		title,
		subtitle,
		layout.NewSpacer(),
		getStartedBtn,
		layout.NewSpacer(),
	)

	// Center everything with padding (looks great on phones)
	w.SetContent(container.NewPadded(container.NewCenter(content)))

	// Make sure window is visible
	w.Show()

	// Center the window on screen
	w.CenterOnScreen()

	return w
}
