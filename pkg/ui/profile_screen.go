package ui

import (
	"fmt"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"

	skilltheme "skillDar/pkg/theme"
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
	profilePic := canvas.NewCircle(theme.Color(theme.ColorNamePrimary))
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

	// Action buttons - remove importance to use default text color
	callBtn := widget.NewButton("📞\nCall", func() {
		// TODO: Implement action
	})

	chatBtn := widget.NewButton("💬\nChat", func() {
		// TODO: Implement action
	})

	hireBtn := widget.NewButton("💼\nHire", func() {
		// TODO: Implement action
	})

	buttonsGrid := container.NewGridWithColumns(3, callBtn, chatBtn, hireBtn)

	// Use ThemedNavBar as background for buttons
	navBarBg := skilltheme.NewThemedNavBar()
	actionsRow := container.NewStack(navBarBg, buttonsGrid)

	// Price card with background
	priceTitle := widget.NewLabel("Per Hour")
	priceTitle.Alignment = fyne.TextAlignCenter

	priceAmount := canvas.NewText("$80", theme.Color(skilltheme.ColorNameHighlight))
	priceAmount.Alignment = fyne.TextAlignCenter
	priceAmount.TextSize = 24
	priceAmount.TextStyle = fyne.TextStyle{Bold: true}

	priceNote := widget.NewLabel("(Minimum 2 hours)")
	priceNote.Alignment = fyne.TextAlignCenter

	priceContent := container.NewVBox(priceTitle, priceAmount, priceNote)
	priceBackground := canvas.NewRectangle(theme.Color(skilltheme.ColorNameHighlight))
	priceCard := container.NewStack(priceBackground, container.NewPadded(priceContent))

	// About section
	aboutContent := widget.NewLabel("Professional plumber with 12 years of experience in all plumbing work...")
	aboutContent.Wrapping = fyne.TextWrapWord

	skillContent := widget.NewLabel("• Pipe Installation\n• Leak Repairs\n• Drain Cleaning\n• Water Heater Services")

	reviewsContent := widget.NewLabel("⭐⭐⭐⭐⭐\n\"Excellent service! Highly recommend.\"\n\n⭐⭐⭐⭐⭐\n\"Very professional and timely.\"")

	tabContentContainer := container.NewStack()
	switchTab := func(tabIndex int) {

		tabContentContainer.Objects = []fyne.CanvasObject{}

		switch tabIndex {
		case 0:
			tabContentContainer.Objects = []fyne.CanvasObject{aboutContent}
		case 1:
			tabContentContainer.Objects = []fyne.CanvasObject{skillContent}
		case 2:
			tabContentContainer.Objects = []fyne.CanvasObject{reviewsContent}
		}
		// Refresh to show new content
		tabContentContainer.Refresh()

	}

	// Update styling
	tab1Btn := widget.NewButton("About", func() {
		fmt.Println("about tab")
		switchTab(0)
	})
	tab2Btn := widget.NewButton("Skills", func() {
		fmt.Println("skill tab")
		switchTab(1)
	})
	tab3Btn := widget.NewButton("Reviews", func() {
		fmt.Println("reviews tab")
		switchTab(2)
	})

	tabsRow := container.NewGridWithColumns(3, tab1Btn, tab2Btn, tab3Btn)

	switchTab(0) // Default to first tab

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
		tabContentContainer,
		//aboutText,
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
