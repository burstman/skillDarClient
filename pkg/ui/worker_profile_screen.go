package ui

import (
	"image/color"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"

	skilltheme "skillDar/pkg/theme"
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
	headerBg := canvas.NewRectangle(theme.Color(theme.ColorNamePrimary))
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
	profileCircle := canvas.NewCircle(theme.Color(skilltheme.ColorNameHighlight))
	profileCircle.StrokeColor = theme.Color(theme.ColorNameBackground)
	profileCircle.StrokeWidth = 4

	profilePicContainer := container.NewStack(profileCircle)
	profilePicContainer.Resize(fyne.NewSize(100, 100))

	// Worker name
	nameLabel := canvas.NewText(worker.Name, theme.Color(theme.ColorNameBackground))
	nameLabel.Alignment = fyne.TextAlignCenter
	nameLabel.TextSize = 20
	nameLabel.TextStyle = fyne.TextStyle{Bold: true}

	// Profession
	professionLabel := canvas.NewText(worker.Profession, theme.Color(theme.ColorNameBackground))
	professionLabel.Alignment = fyne.TextAlignCenter
	professionLabel.TextSize = 14

	// Rating and distance info
	ratingText := canvas.NewText("⭐ 4.9  (127 reviews)  📍 0.8 km", theme.Color(theme.ColorNameBackground))
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
	callBtn := createRoundActionButton("📞", "Call", theme.Color(theme.ColorNameBackground))
	chatBtn := createRoundActionButton("💬", "Chat", theme.Color(theme.ColorNameBackground))
	hireBtn := createRoundActionButton("📅", "Hire", theme.Color(theme.ColorNameBackground))

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

	bg := canvas.NewRectangle(theme.Color(theme.ColorNameBackground))

	card := container.NewStack(bg, container.NewPadded(content))
	return card
}

// createRoundActionButton creates a rounded action button
func createRoundActionButton(icon, label string, bgColor color.Color) fyne.CanvasObject {
	// Set text color based on background - white for primary (blue), dark for others
	var textColor color.Color
	if bgColor == theme.Color(theme.ColorNamePrimary) {
		textColor = theme.Color(theme.ColorNameBackground) // White text on blue background
	} else {
		textColor = theme.Color(theme.ColorNameForeground) // Dark text on light background
	}

	bg := canvas.NewCircle(bgColor)

	// Create text overlays with proper colors
	iconText := canvas.NewText(icon, textColor)
	iconText.Alignment = fyne.TextAlignCenter
	iconText.TextSize = 20

	labelText := canvas.NewText(label, textColor)
	labelText.Alignment = fyne.TextAlignCenter
	labelText.TextSize = 12

	textContent := container.NewVBox(
		layout.NewSpacer(),
		container.NewCenter(iconText),
		container.NewCenter(labelText),
		layout.NewSpacer(),
	)

	card := container.NewStack(bg, textContent)
	card.Resize(fyne.NewSize(80, 80))

	// Wrap in a tappable button with animation
	btn := &tappableContainer{
		content: card,
		bg:      bg,
		bgColor: bgColor,
		onTap: func() {
			// Handle action based on label
			println("===============================")
			println("BUTTON TAPPED:", label)
			println("===============================")
		},
	}
	btn.ExtendBaseWidget(btn)

	return btn
}

// tappableContainer is a custom widget that makes any content tappable
type tappableContainer struct {
	widget.BaseWidget
	content fyne.CanvasObject
	bg      *canvas.Circle
	bgColor color.Color
	onTap   func()
}

func (t *tappableContainer) CreateRenderer() fyne.WidgetRenderer {
	return &tappableRenderer{content: t.content}
}

func (t *tappableContainer) Tapped(*fyne.PointEvent) {
	if t.onTap != nil {
		// Visual feedback: show gray background briefly for all buttons
		if t.bg != nil {
			// Save original color
			originalColor := t.bgColor

			// Use gray for tap feedback (visible in both light and dark themes)
			tapColor := color.RGBA{R: 180, G: 180, B: 180, A: 255}

			t.bg.FillColor = tapColor
			canvas.Refresh(t.bg)

			// Restore original color after a short delay
			go func() {
				time.Sleep(100 * time.Millisecond)
				// Use fyne.Do to safely update UI from goroutine
				fyne.Do(func() {
					t.bg.FillColor = originalColor
					canvas.Refresh(t.bg)
				})
			}()
		}

		t.onTap()
	}
}

type tappableRenderer struct {
	content fyne.CanvasObject
}

func (r *tappableRenderer) Layout(size fyne.Size) {
	r.content.Resize(size)
}

func (r *tappableRenderer) MinSize() fyne.Size {
	return r.content.MinSize()
}

func (r *tappableRenderer) Refresh() {
	canvas.Refresh(r.content)
}

func (r *tappableRenderer) Objects() []fyne.CanvasObject {
	return []fyne.CanvasObject{r.content}
}

func (r *tappableRenderer) Destroy() {}

// createPriceCard creates the pricing information card
func createPriceCard(hourlyRate int) fyne.CanvasObject {
	priceTitle := widget.NewLabel("Travel Price")
	priceTitle.Alignment = fyne.TextAlignCenter

	// Large price display
	priceText := canvas.NewText("TND 180", theme.Color(theme.ColorNameForeground))
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
	bg := canvas.NewRectangle(theme.Color(skilltheme.ColorNameHighlight))

	return container.NewStack(bg, container.NewPadded(content))
}

// createTab creates a tab button
func createTab(label string, active bool) fyne.CanvasObject {
	tabLabel := widget.NewLabel(label)
	tabLabel.Alignment = fyne.TextAlignCenter

	if active {
		tabLabel.TextStyle = fyne.TextStyle{Bold: true}
		// Add underline indicator
		underline := canvas.NewRectangle(theme.Color(theme.ColorNamePrimary))
		underline.SetMinSize(fyne.NewSize(0, 2))

		return container.NewBorder(nil, underline, nil, nil, container.NewPadded(tabLabel))
	}

	return container.NewPadded(tabLabel)
}
