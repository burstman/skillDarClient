package ui

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

// CreateWelcomeScreen creates the welcome screen UI
// onGetStarted is a callback function that runs when "Get Started" button is clicked
func CreateWelcomeScreen(state AppState) fyne.CanvasObject {
	// Top centered logo
	topLogo := canvas.NewImageFromResource(state.GetImage("logoImage"))
	topLogo.FillMode = canvas.ImageFillContain
	topLogo.SetMinSize(fyne.NewSize(120, 40))

	// Header with centered logo
	header := container.NewCenter(topLogo)

	// Welcome text
	title := canvas.NewText("Welcome to SkillDar", nil)
	title.Alignment = fyne.TextAlignCenter
	title.TextSize = 24
	title.TextStyle = fyne.TextStyle{Bold: true}

	subtitle := canvas.NewText("Your Home, Our Expertise", nil)
	subtitle.Alignment = fyne.TextAlignCenter
	subtitle.TextSize = 14

	// Center logo with circles (using the pre-made image)
	centerLogoWithCircles := canvas.NewImageFromResource(state.GetImage("logoInCircle"))
	centerLogoWithCircles.FillMode = canvas.ImageFillContain
	centerLogoWithCircles.SetMinSize(fyne.NewSize(200, 200))

	circleContainer := container.NewCenter(centerLogoWithCircles)

	// Get Started button
	getStartedBtn := widget.NewButton("Get Started", func() {
		state.ShowScreen("login")
	})
	getStartedBtn.Importance = widget.HighImportance

	// Main content layout
	content := container.NewVBox(
		header,
		layout.NewSpacer(),
		container.NewCenter(title),
		container.NewCenter(subtitle),
		layout.NewSpacer(),
		container.NewCenter(circleContainer),
		layout.NewSpacer(),
		layout.NewSpacer(),
		container.NewPadded(getStartedBtn),
		layout.NewSpacer(),
	)

	// Return padded content
	return content
}
