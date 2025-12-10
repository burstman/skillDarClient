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
	searchEntry.SetText("") // Fix Android bug with first character
	searchEntry.OnChanged = func(searchText string) {
		fmt.Println("Search text changed:", searchText)
		// TODO: Implement search filtering
	}

	// Professional categories
	categoriesLabel := widget.NewLabel("Professional Categories")
	categoriesLabel.TextStyle = fyne.TextStyle{Bold: true}

	// Workers container (declare early so category buttons can reference it)
	workersContainer := container.NewVBox()
	workersLabel := widget.NewLabel("Loading workers...")
	workersLabel.TextStyle = fyne.TextStyle{Bold: true}

	// Pagination state
	currentPage := 1
	pageLimit := 10
	isLoading := false
	hasMoreWorkers := true
	selectedCategory := "" // Track selected category

	// Forward declaration for loadWorkers
	var loadWorkers func()

	// Function to filter workers by category
	filterByCategory := func(category string) {
		selectedCategory = category
		currentPage = 1
		hasMoreWorkers = true
		isLoading = false
		workersContainer.Objects = nil // Clear current workers
		workersContainer.Refresh()

		if category == "" {
			workersLabel.SetText("Loading all workers...")
		} else {
			workersLabel.SetText(fmt.Sprintf("Loading %s workers...", category))
		}

		// Trigger loading workers with new category
		loadWorkers()
	}

	// Create category buttons with icons and filter callback
	plumbingCard := createCategoryButtonWithFilter(state, "plumbing", "Plumbing", "plumber", filterByCategory)
	electricityCard := createCategoryButtonWithFilter(state, "electricity", "Electricity", "electrician", filterByCategory)
	paintingCard := createCategoryButtonWithFilter(state, "painting", "Painting", "painter", filterByCategory)
	acFixingCard := createCategoryButtonWithFilter(state, "acFixing", "AC Fixing", "ac-technician", filterByCategory)
	homeCleaningCard := createCategoryButtonWithFilter(state, "homeCleaning", "Home Cleaning", "cleaner", filterByCategory)
	smallRepairsCard := createCategoryButtonWithFilter(state, "smallRepairs", "Small Repairs", "handyman", filterByCategory)
	furnitureCard := createCategoryButtonWithFilter(state, "furnitureAssembly", "Furniture Assembly", "furniture-assembler", filterByCategory)
	waterLeakCard := createCategoryButtonWithFilter(state, "waterLeakage", "Water Leakage", "plumber", filterByCategory)
	applianceCard := createCategoryButtonWithFilter(state, "applianceRepair", "Appliance Repair", "appliance-repair", filterByCategory)
	locksmithCard := createCategoryButtonWithFilter(state, "locksmith", "Locksmiths", "locksmith", filterByCategory)

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
	categoriesScroll.SetMinSize(fyne.NewSize(0, 250)) // Fixed height only, width adapts to screen
	categoriesScroll.OnScrolled = func(pos fyne.Position) {
		fmt.Printf("Categories scrolled to position: X=%.2f, Y=%.2f\n", pos.X, pos.Y)
	}

	// Separator between sections
	separator1 := widget.NewSeparator()

	// Function to load workers from API
	loadWorkers = func() {
		if isLoading || !hasMoreWorkers {
			return
		}
		isLoading = true

		// Run API call in goroutine to avoid blocking
		go func() {
			// Get API service
			apiService := state.GetAPIService()

			fmt.Printf("Loading workers - Page: %d, Limit: %d, Category: %s\n", currentPage, pageLimit, selectedCategory)

			// Fetch workers from API with category filter
			workersResp, err := apiService.GetWorkersByCategory(currentPage, pageLimit, selectedCategory)

			// All UI updates must be in fyne.Do()
			fyne.Do(func() {
				if err != nil {
					fmt.Printf("Failed to load workers (Page %d): %v\n", currentPage, err)
					workersLabel.SetText(fmt.Sprintf("Failed to load workers: %v", err))
					isLoading = false
					return
				}

				fmt.Printf("Successfully loaded %d workers (Total: %d)\n", len(workersResp.Workers), workersResp.Total)

				// Update label with count
				workersLabel.SetText(fmt.Sprintf("Available Workers Near You (%d)", workersResp.Total))

				// Add worker cards
				for _, worker := range workersResp.Workers {
					rating := fmt.Sprintf("%.1f", worker.Rating)
					reviews := fmt.Sprintf("%d", worker.Reviews)
					// Price comes as string from API, use it directly
					price := worker.Price

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
			})
		}()
	}
	// Load initial workers
	loadWorkers() // Make workers scrollable with minimum height
	workersScroll := container.NewVScroll(workersContainer)
	workersScroll.SetMinSize(fyne.NewSize(0, 300)) // Fixed height only, width adapts to screen
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
func createCategoryButton(state AppState, iconKey, name string, textsize uint8) fyne.CanvasObject {
	// Create image from resource
	iconImage := canvas.NewImageFromResource(state.GetImage(iconKey))
	iconImage.FillMode = canvas.ImageFillContain
	iconImage.SetMinSize(fyne.NewSize(32, 32))

	// Split name into words to handle multi-word categories
	words := []string{}
	currentWord := ""
	for _, char := range name {
		if char == ' ' {
			if currentWord != "" {
				words = append(words, currentWord)
				currentWord = ""
			}
		} else {
			currentWord += string(char)
		}
	}
	if currentWord != "" {
		words = append(words, currentWord)
	}

	// Create text labels based on number of words
	var textContainer *fyne.Container
	if len(words) == 1 {
		// Single word - one line
		nameLabel := canvas.NewText(name, theme.Color(theme.ColorNameForeground))
		nameLabel.Alignment = fyne.TextAlignCenter
		nameLabel.TextSize = float32(textsize)
		nameLabel.TextStyle = fyne.TextStyle{Bold: false}
		textContainer = container.NewCenter(nameLabel)
	} else {
		// Multiple words - split into two lines
		line1 := words[0]
		line2 := ""
		if len(words) > 1 {
			for i := 1; i < len(words); i++ {
				if i > 1 {
					line2 += " "
				}
				line2 += words[i]
			}
		}

		label1 := canvas.NewText(line1, theme.Color(theme.ColorNameForeground))
		label1.Alignment = fyne.TextAlignCenter
		label1.TextSize = 11
		label1.TextStyle = fyne.TextStyle{Bold: false}

		label2 := canvas.NewText(line2, theme.Color(theme.ColorNameForeground))
		label2.Alignment = fyne.TextAlignCenter
		label2.TextSize = 11
		label2.TextStyle = fyne.TextStyle{Bold: false}

		textContainer = container.NewVBox(
			container.NewCenter(label1),
			container.NewCenter(label2),
		)
	}

	// Simple VBox with icon and text
	innerContent := container.NewVBox(
		container.NewCenter(iconImage),
		textContainer,
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

// createCategoryButtonWithFilter creates a clickable category button with icon image and filter callback
func createCategoryButtonWithFilter(state AppState, iconKey, displayName, apiCategory string, onFilter func(string)) fyne.CanvasObject {
	// Create image from resource
	iconImage := canvas.NewImageFromResource(state.GetImage(iconKey))
	iconImage.FillMode = canvas.ImageFillContain
	iconImage.SetMinSize(fyne.NewSize(32, 32))

	// Split name into words to handle multi-word categories
	words := []string{}
	currentWord := ""
	for _, char := range displayName {
		if char == ' ' {
			if currentWord != "" {
				words = append(words, currentWord)
				currentWord = ""
			}
		} else {
			currentWord += string(char)
		}
	}
	if currentWord != "" {
		words = append(words, currentWord)
	}

	// Create text labels based on number of words
	var textContainer *fyne.Container
	if len(words) == 1 {
		// Single word - one line
		nameLabel := canvas.NewText(displayName, theme.Color(theme.ColorNameForeground))
		nameLabel.Alignment = fyne.TextAlignCenter
		nameLabel.TextSize = 11
		nameLabel.TextStyle = fyne.TextStyle{Bold: false}
		textContainer = container.NewCenter(nameLabel)
	} else {
		// Multiple words - split into two lines
		line1 := words[0]
		line2 := ""
		if len(words) > 1 {
			for i := 1; i < len(words); i++ {
				if i > 1 {
					line2 += " "
				}
				line2 += words[i]
			}
		}

		label1 := canvas.NewText(line1, theme.Color(theme.ColorNameForeground))
		label1.Alignment = fyne.TextAlignCenter
		label1.TextSize = 11
		label1.TextStyle = fyne.TextStyle{Bold: false}

		label2 := canvas.NewText(line2, theme.Color(theme.ColorNameForeground))
		label2.Alignment = fyne.TextAlignCenter
		label2.TextSize = 11
		label2.TextStyle = fyne.TextStyle{Bold: false}

		textContainer = container.NewVBox(
			container.NewCenter(label1),
			container.NewCenter(label2),
		)
	}

	// Simple VBox with icon and text
	innerContent := container.NewVBox(
		container.NewCenter(iconImage),
		textContainer,
	)

	// Center everything vertically
	content := container.NewVBox(
		layout.NewSpacer(),
		innerContent,
		layout.NewSpacer(),
	)

	// Create a button that wraps the content
	btn := widget.NewButton("", func() {
		fmt.Printf("Category clicked: %s (API: %s)\n", displayName, apiCategory)
		// Call the filter callback with the API category value
		onFilter(apiCategory)
	})

	// Stack the content on top of the button
	return container.NewStack(btn, content)
}
