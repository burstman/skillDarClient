package ui

import (
	"fmt"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

// CreateWelcomeScreen creates the welcome screen UI with map loading
// The "Get Started" button is only shown when the map is loaded
func CreateWelcomeScreen(state AppState) fyne.CanvasObject {
	// Top centered logo
	topLogo := canvas.NewImageFromResource(state.GetImage("logoImage"))
	topLogo.FillMode = canvas.ImageFillContain
	topLogo.SetMinSize(fyne.NewSize(120, 40))

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

	// Create button container (starts empty, button added when map loads)
	buttonContainer := container.NewVBox()

	// Create a loading label
	loadingLabel := widget.NewLabel("Loading map...")
	loadingLabel.Alignment = fyne.TextAlignCenter

	// Load map and wait for it to be ready
	go func() {
		fmt.Printf("[DEBUG] [%s] Welcome screen: Starting map load...\n", time.Now().Format("15:04:05.000"))
		mapSection, mapLoadedChan := createMapSection(state)

		// Wait for map to load
		<-mapLoadedChan
		fmt.Printf("[DEBUG] [%s] Welcome screen: Map loaded, showing Get Started button\n", time.Now().Format("15:04:05.000"))

		// Update UI on main thread
		fyne.Do(func() {
			// Remove loading label
			buttonContainer.Objects = nil

			// Create and add Get Started button
			getStartedBtn := widget.NewButton("Get Started", func() {
				// Check if user is already authenticated
				if state.GetPreferences().IsLoggedIn() {
					// User has token, skip login and go to main screen
					state.ShowScreen("main")
				} else {
					// No token, show login screen
					state.ShowScreen("login")
				}
			})
			getStartedBtn.Importance = widget.HighImportance

			buttonContainer.Add(container.NewPadded(getStartedBtn))
			buttonContainer.Refresh()
		})

		// Keep the map reference (prevent garbage collection during loading)
		_ = mapSection
	}()

	// Main content layout with map and button sections
	content := container.NewVBox(
		layout.NewSpacer(),
		container.NewCenter(title),
		container.NewCenter(subtitle),
		layout.NewSpacer(),
		container.NewCenter(circleContainer),
		layout.NewSpacer(),
		container.NewCenter(loadingLabel), // Initially shows loading message
		buttonContainer,                   // Will contain button when map is ready
		layout.NewSpacer(),
	)

	// Return padded content
	return content
}
