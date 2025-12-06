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
	workersLabel := widget.NewLabel("Available Workers Near You (5)")
	workersLabel.TextStyle = fyne.TextStyle{Bold: true}

	// Dummy worker data - start with first 5 workers
	allWorkers := []struct {
		name       string
		profession string
		rating     string
		distance   string
		reviews    string
		price      string
		available  bool
	}{
		{"Mohamed Hassan", "Plumber", "4.9", "0.8 km", "127", "180 TND/hr", true},
		{"Ahmed El-Sayed", "Electrician", "4.8", "1.2 km", "98", "200 TND/hr", true},
		{"Hossam Abid", "Painter", "4.5", "2.1 km", "55", "150 TND/hr", false},
		{"Youssef Mansour", "AC Technician", "4.7", "1.5 km", "89", "190 TND/hr", true},
		{"Karim Saidi", "Cleaner", "4.6", "0.5 km", "112", "120 TND/hr", true},
		// Next batch
		{"Ali Ben Salem", "Plumber", "4.8", "3.2 km", "145", "170 TND/hr", true},
		{"Sofiane Gharbi", "Electrician", "4.9", "2.8 km", "203", "210 TND/hr", false},
		{"Mehdi Trabelsi", "Carpenter", "4.7", "1.9 km", "78", "165 TND/hr", true},
		{"Nabil Chebbi", "Locksmith", "4.6", "2.5 km", "92", "140 TND/hr", true},
		{"Rami Bouazizi", "Painter", "4.5", "3.0 km", "67", "155 TND/hr", true},
		// Third batch
		{"Farid Jelassi", "AC Technician", "4.8", "1.8 km", "134", "195 TND/hr", true},
		{"Tarek Maatoug", "Plumber", "4.7", "2.2 km", "101", "175 TND/hr", false},
		{"Walid Hamdi", "Electrician", "4.9", "0.9 km", "187", "205 TND/hr", true},
		{"Sami Ayari", "Cleaner", "4.6", "1.7 km", "88", "125 TND/hr", true},
		{"Bassem Jribi", "Carpenter", "4.5", "3.5 km", "72", "160 TND/hr", true},
	}

	currentDisplayCount := 5
	isLoading := false

	// Workers container - start with first 5
	workersContainer := container.NewVBox()
	for i := 0; i < currentDisplayCount && i < len(allWorkers); i++ {
		w := allWorkers[i]
		workersContainer.Add(createSimpleWorkerCard(state, w.name, w.profession, w.rating, w.distance, w.reviews, w.price, w.available))
	}

	// Make workers scrollable with minimum height
	workersScroll := container.NewVScroll(workersContainer)
	workersScroll.SetMinSize(fyne.NewSize(400, 300)) // Give workers section proper height
	workersScroll.OnScrolled = func(pos fyne.Position) {
		fmt.Printf("Workers scrolled to position: X=%.2f, Y=%.2f\n", pos.X, pos.Y)

		// Check if we're near the bottom (Y position > 40 means scrolled down significantly)
		if pos.Y > 40 && !isLoading && currentDisplayCount < len(allWorkers) {
			isLoading = true
			fmt.Println(">>> Loading more workers...")

			// Load next 5 workers
			oldCount := currentDisplayCount
			currentDisplayCount += 5
			if currentDisplayCount > len(allWorkers) {
				currentDisplayCount = len(allWorkers)
			}

			// Add new workers to container
			for i := oldCount; i < currentDisplayCount; i++ {
				w := allWorkers[i]
				workersContainer.Add(createSimpleWorkerCard(state, w.name, w.profession, w.rating, w.distance, w.reviews, w.price, w.available))
			}

			// Update label
			workersLabel.SetText(fmt.Sprintf("Available Workers Near You (%d)", currentDisplayCount))
			workersContainer.Refresh()

			fmt.Printf(">>> Loaded %d more workers. Total: %d\n", currentDisplayCount-oldCount, currentDisplayCount)
			isLoading = false
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
