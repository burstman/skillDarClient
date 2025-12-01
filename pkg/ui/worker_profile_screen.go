package ui

import (
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

// WorkerProfile represents a worker's profile data
type WorkerProfile struct {
	Name            string
	Profession      string
	Rating          float32
	ReviewCount     int
	Distance        string
	HourlyRate      int
	CompletedJobs   int
	YearsExperience int
	Available       bool
	About           string
	Skills          []string
}

// CreateWorkerProfileScreen builds a detailed worker profile screen
func CreateWorkerProfileScreen(state AppState, worker WorkerProfile) fyne.CanvasObject {
	// Create blue header background
	headerBg := canvas.NewRectangle(color.RGBA{R: 0x28, G: 0x7D, B: 0xF7, A: 0xFF})
	headerBg.SetMinSize(fyne.NewSize(0, 220))

	// Back button
	backBtn := widget.NewButton("←", func() {
		state.ShowScreen("main")
	})
	backBtn.Importance = widget.LowImportance

	// Verified badge
	verifiedBadge := widget.NewLabel("✓ Verified")
	verifiedBadge.TextStyle = fyne.TextStyle{Bold: true}

	topBar := container.NewBorder(nil, nil, backBtn, verifiedBadge)

	// Profile picture (circular)
	profileCircle := canvas.NewCircle(color.RGBA{R: 255, G: 200, B: 100, A: 255})
	profileCircle.StrokeColor = color.White
	profileCircle.StrokeWidth = 4

	profilePicContainer := container.NewStack(profileCircle)
	profilePicContainer.Resize(fyne.NewSize(100, 100))

	// Worker name
	nameLabel := canvas.NewText(worker.Name, color.White)
	nameLabel.Alignment = fyne.TextAlignCenter
	nameLabel.TextSize = 20
	nameLabel.TextStyle = fyne.TextStyle{Bold: true}

	// Profession
	professionLabel := canvas.NewText(worker.Profession, color.RGBA{R: 200, G: 220, B: 255, A: 255})
	professionLabel.Alignment = fyne.TextAlignCenter
	professionLabel.TextSize = 14

	// Rating and distance info
	ratingText := canvas.NewText("⭐ 4.9  (127 reviews)  📍 0.8 km", color.White)
	ratingText.Alignment = fyne.TextAlignCenter
	ratingText.TextSize = 12

	// Header content
	headerContent := container.NewVBox(
		topBar,
		layout.NewSpacer(),
		container.NewCenter(profilePicContainer),
		container.NewVBox(nameLabel),
		container.NewVBox(professionLabel),
		container.NewVBox(ratingText),
		layout.NewSpacer(),
	)

	header := container.NewStack(headerBg, container.NewPadded(headerContent))

	// Stats cards
	completedStat := createStatCard2("📦", "340", "Completed")
	experienceStat := createStatCard2("🏆", "12", "Years Exp.")
	ratingStat := createStatCard2("⭐", "4.9", "Rating")

	statsRow := container.NewGridWithColumns(3,
		completedStat,
		experienceStat,
		ratingStat,
	)

	// Action buttons
	callBtn := createRoundActionButton("📞", "Call", color.White)
	chatBtn := createRoundActionButton("💬", "Chat", color.White)
	hireBtn := createRoundActionButton("📅", "Hire", color.RGBA{R: 0x28, G: 0x7D, B: 0xF7, A: 0xFF})

	actionsRow := container.NewGridWithColumns(3, callBtn, chatBtn, hireBtn)

	// Price section
	priceCard := createPriceCard(worker.HourlyRate)

	// Tabs
	aboutTab := createTab("About", true)
	skillsTab := createTab("Skills", false)
	reviewsTab := createTab("Reviews", false)

	tabsRow := container.NewGridWithColumns(3, aboutTab, skillsTab, reviewsTab)

	// About section
	aboutTitle := widget.NewLabel("Summary")
	aboutTitle.TextStyle = fyne.TextStyle{Bold: true}

	aboutText := widget.NewLabel("Available Now for Booking ✓")
	aboutText.Importance = widget.SuccessImportance

	description := widget.NewLabel("Professional installation and maintenance of electrical wiring, fixtures, and appliances. Repair services, installation of air conditioners and appliances.")
	description.Wrapping = fyne.TextWrapWord

	aboutSection := container.NewVBox(
		aboutTitle,
		aboutText,
		description,
	)

	// Main content
	content := container.NewVBox(
		statsRow,
		widget.NewSeparator(),
		actionsRow,
		widget.NewSeparator(),
		priceCard,
		widget.NewSeparator(),
		tabsRow,
		widget.NewSeparator(),
		aboutSection,
	)

	// Full layout
	mainLayout := container.NewBorder(
		header,
		nil,
		nil,
		nil,
		container.NewVScroll(content),
	)

	return mainLayout
}

// createStatCard2 creates a stat card for the profile screen
func createStatCard2(icon, value, label string) fyne.CanvasObject {
	iconLabel := widget.NewLabel(icon)
	iconLabel.Alignment = fyne.TextAlignCenter
	iconLabel.TextStyle = fyne.TextStyle{Bold: true}

	valueLabel := widget.NewLabel(value)
	valueLabel.Alignment = fyne.TextAlignCenter
	valueLabel.TextStyle = fyne.TextStyle{Bold: true}

	textLabel := widget.NewLabel(label)
	textLabel.Alignment = fyne.TextAlignCenter

	content := container.NewVBox(
		iconLabel,
		valueLabel,
		textLabel,
	)

	bg := canvas.NewRectangle(color.RGBA{R: 250, G: 250, B: 250, A: 255})

	card := container.NewStack(bg, container.NewPadded(content))
	return card
}

// createRoundActionButton creates a rounded action button
func createRoundActionButton(icon, label string, bgColor color.Color) fyne.CanvasObject {
	iconLabel := widget.NewLabel(icon)
	iconLabel.Alignment = fyne.TextAlignCenter

	textLabel := widget.NewLabel(label)
	textLabel.Alignment = fyne.TextAlignCenter

	content := container.NewVBox(
		iconLabel,
		textLabel,
	)

	bg := canvas.NewRectangle(bgColor)

	card := container.NewStack(bg, container.NewPadded(content))

	// Make it tappable
	btn := widget.NewButton("", func() {
		// Handle action
	})

	return container.NewStack(card, btn)
}

// createPriceCard creates the pricing information card
func createPriceCard(hourlyRate int) fyne.CanvasObject {
	priceTitle := widget.NewLabel("Travel Price")
	priceTitle.Alignment = fyne.TextAlignCenter

	// Large price display
	priceText := canvas.NewText("TND 180", color.Black)
	priceText.Alignment = fyne.TextAlignCenter
	priceText.TextSize = 28
	priceText.TextStyle = fyne.TextStyle{Bold: true}

	perHourLabel := widget.NewLabel("Per Hour")
	perHourLabel.Alignment = fyne.TextAlignCenter

	minLabel := widget.NewLabel("(Minimum 2 hours)")
	minLabel.Alignment = fyne.TextAlignCenter

	content := container.NewVBox(
		priceTitle,
		container.NewVBox(priceText),
		perHourLabel,
		minLabel,
	)

	// Light beige/yellow background
	bg := canvas.NewRectangle(color.RGBA{R: 255, G: 248, B: 230, A: 255})

	return container.NewStack(bg, container.NewPadded(content))
}

// createTab creates a tab button
func createTab(label string, active bool) fyne.CanvasObject {
	tabLabel := widget.NewLabel(label)
	tabLabel.Alignment = fyne.TextAlignCenter

	if active {
		tabLabel.TextStyle = fyne.TextStyle{Bold: true}
		// Add underline indicator
		underline := canvas.NewRectangle(color.RGBA{R: 0x28, G: 0x7D, B: 0xF7, A: 0xFF})
		underline.SetMinSize(fyne.NewSize(0, 2))

		return container.NewBorder(nil, underline, nil, nil, container.NewPadded(tabLabel))
	}

	return container.NewPadded(tabLabel)
}
