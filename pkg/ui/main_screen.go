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
	// Create a theme-aware navbar background from theme package
	navBg := skilltheme.NewThemedNavBar()

	// Declare button variables (needed for cross-referencing in callbacks)
	var homeBtn, ordersBtn, chatBtn, profileBtn *widget.Button

	// Create navigation buttons with content
	homeBtn, homeContent := createNavButton("🏠", "Home", true, func() {
		// Update all buttons
		updateNavButton(homeBtn, true)
		updateNavButton(ordersBtn, false)
		updateNavButton(chatBtn, false)
		updateNavButton(profileBtn, false)
		// Update content
		contentContainer.Objects = []fyne.CanvasObject{createClientHomeContent(state)}
		contentContainer.Refresh()
	})

	ordersBtn, ordersContent := createNavButton("📋", "Orders", false, func() {
		updateNavButton(homeBtn, false)
		updateNavButton(ordersBtn, true)
		updateNavButton(chatBtn, false)
		updateNavButton(profileBtn, false)
		contentContainer.Objects = []fyne.CanvasObject{createOrdersContent(state)}
		contentContainer.Refresh()
	})

	chatBtn, chatContent := createNavButton("💬", "Chat", false, func() {
		updateNavButton(homeBtn, false)
		updateNavButton(ordersBtn, false)
		updateNavButton(chatBtn, true)
		updateNavButton(profileBtn, false)
		contentContainer.Objects = []fyne.CanvasObject{createChatContent(state)}
		contentContainer.Refresh()
	})

	profileBtn, profileContent := createNavButton("👤", "Profile", false, func() {
		updateNavButton(homeBtn, false)
		updateNavButton(ordersBtn, false)
		updateNavButton(chatBtn, false)
		updateNavButton(profileBtn, true)
		contentContainer.Objects = []fyne.CanvasObject{createProfileContent(state)}
		contentContainer.Refresh()
	})

	navItems := container.NewHBox(
		layout.NewSpacer(),
		container.NewStack(homeBtn, homeContent),
		layout.NewSpacer(),
		container.NewStack(ordersBtn, ordersContent),
		layout.NewSpacer(),
		container.NewStack(chatBtn, chatContent),
		layout.NewSpacer(),
		container.NewStack(profileBtn, profileContent),
		layout.NewSpacer(),
	)

	return container.NewStack(
		navBg,
		container.NewPadded(navItems),
	)
}

// createNavButton creates a navigation button similar to category buttons
func createNavButton(icon, label string, active bool, onTap func()) (*widget.Button, fyne.CanvasObject) {
	iconLabel := widget.NewLabel(icon)
	iconLabel.Alignment = fyne.TextAlignCenter

	textLabel := widget.NewLabel(label)
	textLabel.Alignment = fyne.TextAlignCenter

	content := container.NewVBox(
		iconLabel,
		textLabel,
	)

	// Create button
	btn := widget.NewButton("", onTap)

	// Set importance based on active state
	if active {
		btn.Importance = widget.HighImportance // Blue background like category buttons
	} else {
		btn.Importance = widget.LowImportance // Transparent like category buttons
	}

	return btn, content
}

// updateNavButton updates button importance (active/inactive state)
func updateNavButton(btn *widget.Button, active bool) {
	if active {
		btn.Importance = widget.HighImportance // Blue background
	} else {
		btn.Importance = widget.LowImportance // Transparent
	}
	btn.Refresh()
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

// createProfileContent creates the user profile content
func createProfileContent(state AppState) fyne.CanvasObject {
	title := widget.NewLabel("My Profile")
	title.TextStyle = fyne.TextStyle{Bold: true}
	title.Alignment = fyne.TextAlignCenter

	// Profile picture placeholder
	profileCircle := canvas.NewCircle(theme.PrimaryColor())
	profileCircle.Resize(fyne.NewSize(100, 100))
	profilePic := container.NewCenter(profileCircle)

	// User info
	nameLabel := widget.NewLabel("John Doe")
	nameLabel.TextStyle = fyne.TextStyle{Bold: true}
	nameLabel.Alignment = fyne.TextAlignCenter

	emailLabel := widget.NewLabel("john.doe@example.com")
	emailLabel.Alignment = fyne.TextAlignCenter

	phoneLabel := widget.NewLabel("+216 12 345 678")
	phoneLabel.Alignment = fyne.TextAlignCenter

	// Edit profile button
	editBtn := widget.NewButton("Edit Profile", func() {
		state.ShowScreen("edit_profile_client")
	})
	editBtn.Importance = widget.HighImportance

	// Settings options
	settingsLabel := widget.NewLabel("Settings")
	settingsLabel.TextStyle = fyne.TextStyle{Bold: true}

	notificationsBtn := widget.NewButton("Notifications", func() {
		fmt.Println("Notifications clicked")
	})

	languageBtn := widget.NewButton("Language", func() {
		fmt.Println("Language clicked")
	})

	helpBtn := widget.NewButton("Help & Support", func() {
		fmt.Println("Help clicked")
	})

	logoutBtn := widget.NewButton("Logout", func() {
		fmt.Println("Logout clicked")
	})
	logoutBtn.Importance = widget.DangerImportance

	return container.NewVBox(
		title,
		profilePic,
		nameLabel,
		emailLabel,
		phoneLabel,
		layout.NewSpacer(),
		editBtn,
		layout.NewSpacer(),
		settingsLabel,
		notificationsBtn,
		languageBtn,
		helpBtn,
		layout.NewSpacer(),
		logoutBtn,
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
