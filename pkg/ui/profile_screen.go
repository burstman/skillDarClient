package ui

import (
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

// CreateProfileScreen builds the profile screen with stats and info cards
func CreateProfileScreen(state AppState) fyne.CanvasObject {
	// Header with back button and user name
	backBtn := widget.NewButton("←", func() {
		state.ShowScreen("main")
	})
	userName := widget.NewLabel("Mohamed Hassan")
	userName.TextStyle = fyne.TextStyle{Bold: true}

	header := container.NewBorder(
		nil, nil,
		backBtn,
		nil,
		userName,
	)

	// Profile picture (circular placeholder)
	profilePic := canvas.NewCircle(color.RGBA{R: 100, G: 150, B: 200, A: 255})
	profilePicContainer := container.NewPadded(profilePic)
	profilePicContainer.Resize(fyne.NewSize(80, 80))

	// User name and verification badge
	userNameLabel := widget.NewLabel("Mohamed Hassan ✓")
	userNameLabel.Alignment = fyne.TextAlignCenter
	userNameLabel.TextStyle = fyne.TextStyle{Bold: true}

	userType := widget.NewLabel("Plumber")
	userType.Alignment = fyne.TextAlignCenter

	// Stats row (experience, rating, etc.)
	statsLabel := widget.NewLabel("⭐ 4.9 (127)  🔥 0.9 km")
	statsLabel.Alignment = fyne.TextAlignCenter

	// Stats cards
	stat1 := createStatCard("📦", "340", "Certificates")
	stat2 := createStatCard("🏆", "12", "Years Experience")
	stat3 := createStatCard("⭐", "4.9", "Rating")

	statsRow := container.NewGridWithColumns(3, stat1, stat2, stat3)

	// Action buttons
	callBtn := createActionButton("📞", "Call")
	chatBtn := createActionButton("💬", "Chat")
	hireBtn := createActionButton("💼", "Hire")
	hireBtn.Importance = widget.HighImportance

	actionsRow := container.NewGridWithColumns(3, callBtn, chatBtn, hireBtn)

	// Price card with background
	priceTitle := widget.NewLabel("Per Hour")
	priceTitle.Alignment = fyne.TextAlignCenter

	priceAmount := canvas.NewText("$80", color.Black)
	priceAmount.Alignment = fyne.TextAlignCenter
	priceAmount.TextSize = 24
	priceAmount.TextStyle = fyne.TextStyle{Bold: true}

	priceNote := widget.NewLabel("(Minimum 2 hours)")
	priceNote.Alignment = fyne.TextAlignCenter

	priceContent := container.NewVBox(priceTitle, priceAmount, priceNote)
	priceBackground := canvas.NewRectangle(color.RGBA{R: 255, G: 248, B: 220, A: 255})
	priceCard := container.NewStack(priceBackground, container.NewPadded(priceContent))

	// Tabs
	tab1 := widget.NewLabel("About")
	tab1.TextStyle = fyne.TextStyle{Bold: true}
	tab2 := widget.NewLabel("Skills")
	tab3 := widget.NewLabel("Reviews")

	tabsRow := container.NewGridWithColumns(3, tab1, tab2, tab3)

	// About section
	aboutText := widget.NewLabel("Professional plumber with 12 years of experience in all plumbing work...")
	aboutText.Wrapping = fyne.TextWrapWord

	// Available button
	availableBtn := widget.NewButton("Available Now for Booking", func() {
		// TODO: Implement booking
	})
	availableBtn.Importance = widget.SuccessImportance

	// Main content
	content := container.NewVBox(
		header,
		//layout.NewSpacer(),
		container.NewCenter(profilePicContainer),
		userNameLabel,
		userType,
		statsLabel,
		statsRow,
		actionsRow,
		priceCard,
		tabsRow,
		aboutText,
		availableBtn,
	)

	scroll := container.NewVScroll(content)
	return scroll
}

// createStatCard creates a card with icon, number, and label
func createStatCard(icon, number, label string) *fyne.Container {
	iconLabel := widget.NewLabel(icon)
	iconLabel.Alignment = fyne.TextAlignCenter
	iconLabel.TextStyle = fyne.TextStyle{Bold: true}

	numLabel := widget.NewLabel(number)
	numLabel.Alignment = fyne.TextAlignCenter
	numLabel.TextStyle = fyne.TextStyle{Bold: true}

	textLabel := widget.NewLabel(label)
	textLabel.Alignment = fyne.TextAlignCenter

	card := container.NewVBox(iconLabel, numLabel, textLabel)
	return container.NewPadded(card)
}

// createActionButton creates a button with icon and text
func createActionButton(icon, text string) *widget.Button {
	return widget.NewButton(icon+" "+text, func() {
		// TODO: Implement action
	})
}
