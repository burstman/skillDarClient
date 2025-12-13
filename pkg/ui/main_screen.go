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

	// Create navigation buttons using custom NavButton
	homeBtn := skilltheme.NewNavButton("🏠\nHome", true, nil)
	ordersBtn := skilltheme.NewNavButton("📋\nOrders", false, nil)
	chatBtn := skilltheme.NewNavButton("💬\nChat", false, nil)
	profileBtn := skilltheme.NewNavButton("👤\nProfile", false, nil)

	// Set up button tap handlers
	homeBtn.OnTapped = func() {
		homeBtn.SetActive(true)
		ordersBtn.SetActive(false)
		chatBtn.SetActive(false)
		profileBtn.SetActive(false)
		contentContainer.Objects = []fyne.CanvasObject{createClientHomeContent(state)}
		contentContainer.Refresh()
	}

	ordersBtn.OnTapped = func() {
		homeBtn.SetActive(false)
		ordersBtn.SetActive(true)
		chatBtn.SetActive(false)
		profileBtn.SetActive(false)
		contentContainer.Objects = []fyne.CanvasObject{createOrdersContent(state)}
		contentContainer.Refresh()
	}

	chatBtn.OnTapped = func() {
		homeBtn.SetActive(false)
		ordersBtn.SetActive(false)
		chatBtn.SetActive(true)
		profileBtn.SetActive(false)
		contentContainer.Objects = []fyne.CanvasObject{createChatContent(state)}
		contentContainer.Refresh()
	}

	profileBtn.OnTapped = func() {
		homeBtn.SetActive(false)
		ordersBtn.SetActive(false)
		chatBtn.SetActive(false)
		profileBtn.SetActive(true)
		contentContainer.Objects = []fyne.CanvasObject{createProfileContent(state)}
		contentContainer.Refresh()
	}

	// Create navigation bar layout
	navItems := container.NewHBox(
		layout.NewSpacer(),
		homeBtn,
		layout.NewSpacer(),
		ordersBtn,
		layout.NewSpacer(),
		chatBtn,
		layout.NewSpacer(),
		profileBtn,
		layout.NewSpacer(),
	)

	// Stack the background and items
	navBarContent := container.NewStack(navBg, navItems)

	// Wrap in a fixed height container
	fixedNav := skilltheme.NewFixedHeightContainer(40, navBarContent)

	return fixedNav
}

// createClientHomeContent creates the home content for clients
func createClientHomeContent(state AppState) fyne.CanvasObject {
	title := widget.NewLabel("Available Workers")
	title.TextStyle = fyne.TextStyle{Bold: true}
	title.Alignment = fyne.TextAlignCenter

	// Search bar
	searchEntry := widget.NewEntry()
	searchEntry.SetPlaceHolder("Search for workers...")
	searchEntry.SetText("") // Fix Android bug with first character
	searchEntry.OnChanged = func(searchText string) {
		fmt.Println("Search text changed:", searchText)
		// TODO: Implement search filtering
	}

	// Shared state for category filtering
	selectedCategory := ""

	// Placeholder for filter callback
	var filterByCategory func(string)

	// Separator between sections
	separator1 := widget.NewSeparator()

	// Create workers section with pagination and filtering
	workersLabel, _, workersScroll, resetPagination := createWorkersSection(state, &selectedCategory)

	// Create map section (ignoring the mapLoadedChan since we don't need it here)
	mapSection, _ := createMapSection(state)

	// Define filter callback to reload workers
	filterByCategory = func(category string) {
		selectedCategory = category

		if category == "" {
			workersLabel.SetText("Loading all workers...")
		} else {
			workersLabel.SetText(fmt.Sprintf("Loading %s workers...", category))
		}

		// Reset pagination and reload workers (workers load async, no UI blocking)
		resetPagination()
	}

	// Create categories section with filter callback
	categoriesSection := createCategoriesSection(state, filterByCategory)

	// Combine everything in a VBox
	content := container.NewVBox(
		title,
		searchEntry,
		categoriesSection,
		separator1,
		workersLabel,
		workersScroll,
		mapSection,
	)

	return content
}

// createOrdersContent creates the orders/bookings content
func createOrdersContent(state AppState) fyne.CanvasObject {
	title := widget.NewLabel("My Orders")
	title.TextStyle = fyne.TextStyle{Bold: true}
	title.Alignment = fyne.TextAlignCenter

	noOrders := widget.NewLabel("No orders yet")
	noOrders.Alignment = fyne.TextAlignCenter

	browseBtn := widget.NewButton("Browse Workers", func() {
		// Switch back to home tab - not implemented yet
		fmt.Println("Browse workers clicked")
	})

	return container.NewVBox(
		title,
		layout.NewSpacer(),
		noOrders,
		browseBtn,
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
		title,
		layout.NewSpacer(),
		noMessages,
		layout.NewSpacer(),
	)
}

// createProfileContent creates the profile/settings content
func createProfileContent(state AppState) fyne.CanvasObject {
	title := widget.NewLabel("My Profile")
	title.TextStyle = fyne.TextStyle{Bold: true}
	title.Alignment = fyne.TextAlignCenter

	// Profile picture placeholder
	profileCircle := canvas.NewCircle(theme.PrimaryColor())
	profilePic := container.NewStack(profileCircle)
	profilePic.Resize(fyne.NewSize(80, 80))

	// Get user info from preferences
	prefs := state.GetPreferences()
	userName := prefs.GetUsername()
	userEmail := prefs.GetEmail()
	userPhone := "" // Phone number not stored in preferences yet

	nameLabel := widget.NewLabel(userName)
	nameLabel.TextStyle = fyne.TextStyle{Bold: true}
	nameLabel.Alignment = fyne.TextAlignCenter

	emailLabel := widget.NewLabel(userEmail)
	emailLabel.Alignment = fyne.TextAlignCenter

	phoneLabel := widget.NewLabel(userPhone)
	phoneLabel.Alignment = fyne.TextAlignCenter

	editBtn := widget.NewButton("Edit Profile", func() {
		fmt.Println("Edit profile clicked")
		state.ShowScreen("edit_profile_client")
	})

	settingsLabel := widget.NewLabel("Settings")
	settingsLabel.TextStyle = fyne.TextStyle{Bold: true}

	// Theme toggle button
	var themeToggle *widget.Button
	updateThemeButton := func() {
		if state.IsDarkTheme() {
			themeToggle.SetIcon(state.GetImage("darkTheme"))
			themeToggle.SetText("Light Mode")
			themeToggle.Alignment = widget.ButtonAlignLeading
		} else {
			themeToggle.SetIcon(state.GetImage("lightTheme"))
			themeToggle.SetText("Dark Mode")
			themeToggle.Alignment = widget.ButtonAlignLeading
		}
	}

	themeToggle = widget.NewButtonWithIcon("Dark Mode", state.GetImage("darkTheme"), func() {
		state.ToggleTheme()
		updateThemeButton()
	})
	themeToggle.Alignment = widget.ButtonAlignLeading
	notificationsBtn := widget.NewButton("Notifications", func() {
		fmt.Println("Notifications clicked")
	})
	notificationsBtn.Alignment = widget.ButtonAlignLeading
	languageBtn := widget.NewButton("Language", func() {
		fmt.Println("Language clicked")
	})
	languageBtn.Alignment = widget.ButtonAlignLeading

	helpBtn := widget.NewButton("Help & Support", func() {
		fmt.Println("Help clicked")
	})
	helpBtn.Alignment = widget.ButtonAlignLeading

	logoutBtn := widget.NewButton("Logout", func() {
		// Clear authentication data
		prefs := state.GetPreferences()
		prefs.ClearAuthData()

		// Clear API service token
		apiService := state.GetAPIService()
		apiService.SetToken("")

		fmt.Println("User logged out")

		// Navigate to welcome screen
		state.ShowScreen("welcome")
	})
	logoutBtn.Importance = widget.DangerImportance
	logoutBtn.Alignment = widget.ButtonAlignLeading
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
		themeToggle,
		notificationsBtn,
		languageBtn,
		helpBtn,
		layout.NewSpacer(),
		logoutBtn,
	)
}
