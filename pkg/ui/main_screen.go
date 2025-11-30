package ui

import (
	"fmt"
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
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
	// Create navigation buttons
	homeBtn := createNavButton("🏠", "Home", true)
	ordersBtn := createNavButton("📋", "Orders", false)
	chatBtn := createNavButton("💬", "Chat", false)
	profileBtn := createNavButton("👤", "Profile", false)

	// Navigation bar background
	navBg := canvas.NewRectangle(color.RGBA{245, 245, 245, 255})

	// Button handlers
	homeBtn.OnTapped = func() {
		updateNavButtons(homeBtn, ordersBtn, chatBtn, profileBtn)
		contentContainer.Objects = []fyne.CanvasObject{createClientHomeContent(state)}
		contentContainer.Refresh()
	}

	ordersBtn.OnTapped = func() {
		updateNavButtons(ordersBtn, homeBtn, chatBtn, profileBtn)
		contentContainer.Objects = []fyne.CanvasObject{createOrdersContent(state)}
		contentContainer.Refresh()
	}

	chatBtn.OnTapped = func() {
		updateNavButtons(chatBtn, homeBtn, ordersBtn, profileBtn)
		contentContainer.Objects = []fyne.CanvasObject{createChatContent(state)}
		contentContainer.Refresh()
	}

	profileBtn.OnTapped = func() {
		updateNavButtons(profileBtn, homeBtn, ordersBtn, chatBtn)
		state.ShowScreen("edit_profile_client")
	}

	navButtons := container.NewHBox(
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

	return container.NewStack(
		navBg,
		container.NewPadded(navButtons),
	)
}

// createNavButton creates a navigation button
func createNavButton(icon, label string, active bool) *widget.Button {
	btn := widget.NewButton(icon+"\n"+label, nil)
	if active {
		btn.Importance = widget.HighImportance
	}
	return btn
}

// updateNavButtons updates the active state of navigation buttons
func updateNavButtons(active *widget.Button, others ...*widget.Button) {
	active.Importance = widget.HighImportance
	active.Refresh()
	for _, btn := range others {
		btn.Importance = widget.MediumImportance
		btn.Refresh()
	}
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

	// Use GridWrap with increased size to fit text properly
	categoriesContainer := container.NewGridWrap(
		fyne.NewSize(115, 85), // button size
		plumbingCard, electricityCard, paintingCard,
		acFixingCard, homeCleaningCard, smallRepairsCard,
		furnitureCard, waterLeakCard, applianceCard,
		locksmithCard,
	)

	// Available workers
	workersLabel := widget.NewLabel("Available Workers")
	workersLabel.TextStyle = fyne.TextStyle{Bold: true}

	worker1 := createSimpleWorkerCard("Mohamed Hassan", "Plumber", "4.9", "0.8 km", "80 EGP/hr", true)
	worker2 := createSimpleWorkerCard("Karim Fathy", "Electrician", "4.5", "1.2 km", "75 EGP/hr", true)
	worker3 := createSimpleWorkerCard("Ahmed Ali", "Carpenter", "4.8", "2.1 km", "90 EGP/hr", false)

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

// createSimpleWorkerCard creates a worker card for clients (simplified version)
func createSimpleWorkerCard(name, profession, rating, distance, price string, available bool) fyne.CanvasObject {
	nameLabel := widget.NewLabel(name)
	nameLabel.TextStyle = fyne.TextStyle{Bold: true}

	professionLabel := widget.NewLabel(profession)

	ratingLabel := widget.NewLabel("⭐ " + rating)
	distanceLabel := widget.NewLabel("📍 " + distance)

	priceLabel := widget.NewLabel(price)

	statusLabel := widget.NewLabel("✅ Available")
	if !available {
		statusLabel.Text = "⏰ Busy"
	}

	contactBtn := widget.NewButton("Contact", func() {
		// Handle contact worker
	})
	if available {
		contactBtn.Importance = widget.SuccessImportance
	}

	info := container.NewVBox(
		nameLabel,
		professionLabel,
		container.NewHBox(ratingLabel, distanceLabel),
	)

	rightSide := container.NewVBox(
		priceLabel,
		statusLabel,
		contactBtn,
	)

	cardContent := container.NewBorder(
		nil, nil,
		info,
		rightSide,
	)

	return container.NewPadded(cardContent)
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
