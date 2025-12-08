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

	// Wrap in a fixed height container (adjust the height value as needed)
	fixedNav := skilltheme.NewFixedHeightContainer(40, navBarContent) // Change 50 to your desired height

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
	categoriesGrid := container.NewGridWrap(
		fyne.NewSize(85, 85), // Smaller button size for mobile
		plumbingCard, electricityCard, paintingCard,
		acFixingCard, homeCleaningCard, smallRepairsCard,
		furnitureCard, waterLeakCard, applianceCard,
		locksmithCard,
	)

	// Make categories scrollable in a fixed height container
	categoriesScroll := container.NewVScroll(categoriesGrid)
	categoriesScroll.SetMinSize(fyne.NewSize(400, 250)) // Fixed height for categories section
	categoriesScroll.OnScrolled = func(pos fyne.Position) {
		fmt.Printf("Categories scrolled to position: X=%.2f, Y=%.2f\n", pos.X, pos.Y)
	}

	// Separator between sections
	separator1 := widget.NewSeparator()

	// Available workers
	workersLabel := widget.NewLabel("Loading workers...")
	workersLabel.TextStyle = fyne.TextStyle{Bold: true}

	// Workers container
	workersContainer := container.NewVBox()

	// Pagination state
	currentPage := 1
	pageLimit := 10
	isLoading := false
	hasMoreWorkers := true

	// Function to load workers from API
	loadWorkers := func() {
		if isLoading || !hasMoreWorkers {
			return
		}
		isLoading = true

		// Run API call in goroutine to avoid blocking
		go func() {
			// Get API service
			apiService := state.GetAPIService()

			// Fetch workers from API
			workersResp, err := apiService.GetWorkers(currentPage, pageLimit)

			if err != nil {
				fmt.Println("Failed to load workers:", err)
				workersLabel.SetText("Failed to load workers")
				isLoading = false
				return
			}

			// Update label with count
			workersLabel.SetText(fmt.Sprintf("Available Workers Near You (%d)", workersResp.TotalCount))

			// Add worker cards
			for _, worker := range workersResp.Workers {
				rating := fmt.Sprintf("%.1f", worker.Rating)
				reviews := fmt.Sprintf("%d", worker.ReviewCount)
				price := fmt.Sprintf("%.0f TND/hr", worker.HourlyRate)

				workersContainer.Add(createSimpleWorkerCard(
					state,
					worker.Name,
					worker.Profession,
					rating,
					worker.Distance,
					reviews,
					price,
					worker.Available,
				))
			}

			workersContainer.Refresh()

			// Check if there are more workers
			if len(workersResp.Workers) < pageLimit {
				hasMoreWorkers = false
			} else {
				currentPage++
			}

			isLoading = false
		}()
	}

	// Load initial workers
	loadWorkers()

	// Make workers scrollable with minimum height
	workersScroll := container.NewVScroll(workersContainer)
	workersScroll.SetMinSize(fyne.NewSize(400, 300)) // Give workers section proper height
	workersScroll.OnScrolled = func(pos fyne.Position) {
		fmt.Printf("Workers scrolled to position: X=%.2f, Y=%.2f\n", pos.X, pos.Y)

		// Check if we're near the bottom - load more workers
		if pos.Y > 40 && hasMoreWorkers {
			fmt.Println(">>> Loading more workers...")
			loadWorkers()
		}
	}

	// Combine everything in a VBox
	content := container.NewVBox(
		title,
		searchEntry,
		categoriesLabel,
		categoriesScroll,
		separator1,
		workersLabel,
		workersScroll,
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
	nameLabel.Alignment = fyne.TextAlignLeading

	emailLabel := widget.NewLabel("john.doe@example.com")
	emailLabel.Alignment = fyne.TextAlignLeading

	phoneLabel := widget.NewLabel("+216 12 345 678")
	phoneLabel.Alignment = fyne.TextAlignLeading

	// Edit profile button
	editBtn := widget.NewButton("Edit Profile", func() {
		state.ShowScreen("edit_profile_client")
	})
	editBtn.Importance = widget.HighImportance
	editBtn.Alignment = widget.ButtonAlignLeading

	// Settings options
	settingsLabel := widget.NewLabel("Settings")
	settingsLabel.TextStyle = fyne.TextStyle{Bold: true}
	settingsLabel.Alignment = fyne.TextAlignLeading

	// Theme toggle with custom icons (dynamic button)
	var themeToggle *widget.Button
	updateThemeButton := func() {
		if state.IsDarkTheme() {
			themeToggle.SetText("Light Mode")
			themeToggle.SetIcon(state.GetImage("lightTheme"))
			themeToggle.Alignment = widget.ButtonAlignLeading
		} else {
			themeToggle.SetText("Dark Mode")
			themeToggle.SetIcon(state.GetImage("darkTheme"))
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
		fmt.Println("Logout clicked")
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
