package ui

import (
	"fmt"

	skilltheme "skillDar/pkg/theme"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

// CreateMainScreen builds the main app screen with bottom navigation
func CreateMainScreen(state AppState) fyne.CanvasObject {
	// Content container that will change based on selected tab
	currentContent := container.NewVBox(createClientHomeContent(state))

	// Bottom navigation bar
	bottomNav := createBottomNavigationBar(state, currentContent)

	// Main layout with bottom navigation
	mainLayout := container.NewBorder(
		nil,                                 // top
		bottomNav,                           // bottom
		nil,                                 // left
		nil,                                 // right
		container.NewScroll(currentContent), // center
	)

	return mainLayout
}

// createBottomNavigationBar creates the bottom navigation menu
func createBottomNavigationBar(state AppState, contentContainer *fyne.Container) fyne.CanvasObject {
	// Navigation bar background - same as content background
	navBg := canvas.NewRectangle(theme.Color(theme.ColorNameBackground))

	// Create navigation items
	homeItem := createNavItem("🏠", "Home", true)
	ordersItem := createNavItem("📋", "Orders", false)
	chatItem := createNavItem("💬", "Chat", false)
	profileItem := createNavItem("👤", "Profile", false)

	// Store references for updates
	var currentActive fyne.CanvasObject = homeItem

	// Wrap items in tappable containers
	homeBtn := widget.NewButton("", func() {
		currentActive = updateNavItems(homeItem, currentActive, ordersItem, chatItem, profileItem)
		contentContainer.Objects = []fyne.CanvasObject{createClientHomeContent(state)}
		contentContainer.Refresh()
	})

	ordersBtn := widget.NewButton("", func() {
		currentActive = updateNavItems(ordersItem, currentActive, homeItem, chatItem, profileItem)
		contentContainer.Objects = []fyne.CanvasObject{createOrdersContent(state)}
		contentContainer.Refresh()
	})

	chatBtn := widget.NewButton("", func() {
		currentActive = updateNavItems(chatItem, currentActive, homeItem, ordersItem, profileItem)
		contentContainer.Objects = []fyne.CanvasObject{createChatContent(state)}
		contentContainer.Refresh()
	})

	profileBtn := widget.NewButton("", func() {
		currentActive = updateNavItems(profileItem, currentActive, homeItem, ordersItem, chatItem)
		state.ShowScreen("edit_profile_client")
	})

	// Set button importance to low to make them less visible
	homeBtn.Importance = widget.LowImportance
	ordersBtn.Importance = widget.LowImportance
	chatBtn.Importance = widget.LowImportance
	profileBtn.Importance = widget.LowImportance

	navItems := container.NewHBox(
		layout.NewSpacer(),
		container.NewStack(homeBtn, homeItem),
		layout.NewSpacer(),
		container.NewStack(ordersBtn, ordersItem),
		layout.NewSpacer(),
		container.NewStack(chatBtn, chatItem),
		layout.NewSpacer(),
		container.NewStack(profileBtn, profileItem),
		layout.NewSpacer(),
	)

	return container.NewStack(
		navBg,
		container.NewPadded(navItems),
	)
}

// createNavItem creates a custom navigation item with background
func createNavItem(icon, label string, active bool) fyne.CanvasObject {
	iconLabel := widget.NewLabel(icon)
	iconLabel.Alignment = fyne.TextAlignCenter

	textLabel := widget.NewLabel(label)
	textLabel.Alignment = fyne.TextAlignCenter

	content := container.NewVBox(
		iconLabel,
		textLabel,
	)

	// Create background rectangle
	var bg *canvas.Rectangle
	if active {
		bg = canvas.NewRectangle(theme.Color(theme.ColorNamePrimary))
	} else {
		// Use a darker color for inactive buttons in dark mode
		bg = canvas.NewRectangle(theme.Color(theme.ColorNameBackground))
	}

	// Return stack without button widget - we'll add tap handling differently
	return container.NewStack(bg, container.NewPadded(content))
}

// updateNavItems updates the visual state of navigation items
func updateNavItems(active fyne.CanvasObject, oldActive fyne.CanvasObject, others ...fyne.CanvasObject) fyne.CanvasObject {
	// Update old active to inactive
	if oldActive != nil {
		oldBg := oldActive.(*fyne.Container).Objects[0].(*canvas.Rectangle)
		oldBg.FillColor = theme.Color(theme.ColorNameBackground)
		oldBg.Refresh()
	}

	// Update new active
	activeBg := active.(*fyne.Container).Objects[0].(*canvas.Rectangle)
	activeBg.FillColor = theme.Color(theme.ColorNamePrimary)
	activeBg.Refresh()

	return active
}

// createClientHomeContent creates the home content for clients
func createClientHomeContent(state AppState) fyne.CanvasObject {
	title := widget.NewLabel("Available Workers")
	title.TextStyle = fyne.TextStyle{Bold: true}
	title.Alignment = fyne.TextAlignCenter

	// Search bar
	searchEntry := widget.NewEntry()
	searchEntry.SetPlaceHolder("Search for workers...")
	searchEntry.OnChanged = func(searchText string) {
		fmt.Println("Search text changed:", searchText)
		// TODO: Implement search filtering
	}

	// Professional categories
	categoriesLabel := widget.NewLabel("Professional Categories")
	categoriesLabel.TextStyle = fyne.TextStyle{Bold: true}

	// Create category buttons with icons
	plumbingCard := createCategoryButton(state, "plumbing", "Plumbing")
	electricityCard := createCategoryButton(state, "electricity", "Electricity")
	paintingCard := createCategoryButton(state, "painting", "Painting")
	acFixingCard := createCategoryButton(state, "acFixing", "AC Fixing")
	homeCleaningCard := createCategoryButton(state, "homeCleaning", "Home Cleaning")
	smallRepairsCard := createCategoryButton(state, "smallRepairs", "Small Repairs")
	furnitureCard := createCategoryButton(state, "furnitureAssembly", "Furniture Assembly")
	waterLeakCard := createCategoryButton(state, "waterLeakage", "Water Leakage")
	applianceCard := createCategoryButton(state, "applianceRepair", "Appliance Repair")
	locksmithCard := createCategoryButton(state, "locksmith", "Locksmiths")

	// Use GridWrap with compact size for mobile
	categoriesContainer := container.NewGridWrap(
		fyne.NewSize(85, 85), // Smaller button size for mobile
		plumbingCard, electricityCard, paintingCard,
		acFixingCard, homeCleaningCard, smallRepairsCard,
		furnitureCard, waterLeakCard, applianceCard,
		locksmithCard,
	)

	// Available workers
	workersLabel := widget.NewLabel("Available Workers Near You (8)")
	workersLabel.TextStyle = fyne.TextStyle{Bold: true}

	worker1 := createSimpleWorkerCard(state, "Mohamed Hassan", "Plumber", "4.9", "0.8 km", "127", "180 TND/hr", true)
	worker2 := createSimpleWorkerCard(state, "Ahmed El-Sayed", "Electrician", "4.8", "1.2 km", "98", "200 TND/hr", true)
	worker3 := createSimpleWorkerCard(state, "Hossam Abid", "Tall", "4.5", "2.1 km", "55", "150 TND/hr", false)

	return container.NewVBox(
		title,
		searchEntry,
		categoriesLabel,
		categoriesContainer,
		workersLabel,
		worker1,
		worker2,
		worker3,
	)
}

// createOrdersContent creates the orders/bookings content
func createOrdersContent(state AppState) fyne.CanvasObject {
	title := widget.NewLabel("My Orders")
	title.TextStyle = fyne.TextStyle{Bold: true}
	title.Alignment = fyne.TextAlignCenter

	noOrders := widget.NewLabel("No orders yet")
	noOrders.Alignment = fyne.TextAlignCenter

	return container.NewVBox(
		layout.NewSpacer(),
		title,
		noOrders,
		layout.NewSpacer(),
	)
}

// createChatContent creates the chat/messages content
func createChatContent(state AppState) fyne.CanvasObject {
	title := widget.NewLabel("Messages")
	title.TextStyle = fyne.TextStyle{Bold: true}
	title.Alignment = fyne.TextAlignCenter

	noMessages := widget.NewLabel("No messages yet")
	noMessages.Alignment = fyne.TextAlignCenter

	return container.NewVBox(
		layout.NewSpacer(),
		title,
		noMessages,
		layout.NewSpacer(),
	)
}

// createSimpleWorkerCard creates a clickable worker card for clients
func createSimpleWorkerCard(state AppState, name, profession, rating, distance, reviewCount, price string, available bool) fyne.CanvasObject {
	// Profile picture placeholder
	profileCircle := canvas.NewCircle(theme.Color(skilltheme.ColorNameHighlight))
	profilePic := container.NewStack(profileCircle)
	profilePic.Resize(fyne.NewSize(50, 50))

	nameLabel := widget.NewLabel(name)
	nameLabel.TextStyle = fyne.TextStyle{Bold: true}

	// Verified badge
	verifiedLabel := widget.NewLabel("✓ Verified")
	verifiedBadge := container.NewHBox(
		widget.NewLabel(name),
		verifiedLabel,
	)

	professionLabel := widget.NewLabel(profession)

	ratingLabel := widget.NewLabel("⭐ " + rating)
	reviewLabel := widget.NewLabel("(" + reviewCount + ")")
	distanceLabel := widget.NewLabel("📍 " + distance)

	priceLabel := widget.NewLabel(price)
	priceLabel.TextStyle = fyne.TextStyle{Bold: true}

	statusLabel := widget.NewLabel("✅ Available")
	statusLabel.Importance = widget.SuccessImportance
	if !available {
		statusLabel.Text = "⏰ Busy"
		statusLabel.Importance = widget.WarningImportance
	}

	info := container.NewVBox(
		verifiedBadge,
		professionLabel,
		container.NewHBox(ratingLabel, reviewLabel, distanceLabel),
	)

	rightSide := container.NewVBox(
		priceLabel,
		statusLabel,
	)

	cardContent := container.NewBorder(
		nil, nil,
		container.NewHBox(profilePic, info),
		rightSide,
	)

	// Create a button that wraps the content
	btn := widget.NewButton("", func() {
		// Create worker profile and show screen
		worker := WorkerProfile{
			Name:            name,
			Profession:      profession,
			Rating:          4.9,
			ReviewCount:     127,
			Distance:        distance,
			HourlyRate:      180,
			CompletedJobs:   340,
			YearsExperience: 12,
			Available:       available,
			About:           "Professional installation and maintenance of electrical wiring, fixtures, and appliances.",
			Skills:          []string{"Plumbing", "Repair", "Installation"},
		}
		state.ShowWorkerProfile(worker)
	})

	// Stack content on button with minimal padding
	return container.NewStack(btn, cardContent)
}

// createCategoryButton creates a clickable category button with icon image
func createCategoryButton(state AppState, iconKey, name string) fyne.CanvasObject {
	// Create image from resource
	iconImage := canvas.NewImageFromResource(state.GetImage(iconKey))
	iconImage.FillMode = canvas.ImageFillContain
	iconImage.SetMinSize(fyne.NewSize(32, 32))

	nameLabel := widget.NewLabel(name)
	nameLabel.Alignment = fyne.TextAlignCenter
	nameLabel.Wrapping = fyne.TextWrapWord

	// Simple VBox with icon and text - DON'T wrap label in Center!
	innerContent := container.NewVBox(
		container.NewCenter(iconImage),
		nameLabel,
	)

	// Center everything vertically
	content := container.NewVBox(
		layout.NewSpacer(),
		innerContent,
		layout.NewSpacer(),
	)

	// Create a button that wraps the content
	btn := widget.NewButton("", func() {
		fmt.Println("Category clicked:", name)
		fmt.Println("Filtering workers by category:", name)
		// TODO: Filter workers by selected category
	})

	// Stack the content on top of the button
	return container.NewStack(btn, content)
}
