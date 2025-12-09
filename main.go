//go:generate fyne bundle -o bundle.go assets

package main

import (
	"fmt"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"

	skilltheme "skillDar/pkg/theme"
	uiscreen "skillDar/pkg/ui"
)

// AppState manages navigation and theme across screens
type AppState struct {
	app               fyne.App
	window            fyne.Window
	isDarkTheme       bool
	preferences       *uiscreen.PreferencesManager // Preferences manager
	apiService        *uiscreen.APIService         // API service
	screens           map[string]fyne.CanvasObject
	icons             map[string]fyne.Resource    // Map of all app icons
	userRole          string                      // "client" or "worker"
	screenHistory     []string                    // Navigation history
	currentWorker     *uiscreen.WorkerProfile     // Current worker being viewed
	connectionManager *uiscreen.ConnectionManager // Connection status manager
	currentContent    *fyne.Container             // Current screen content container
}

// ShowScreen displays a screen by name with the top bar
func (as *AppState) ShowScreen(screenName string) {
	if screen, exists := as.screens[screenName]; exists {
		// Add to history (avoid duplicates)
		if len(as.screenHistory) == 0 || as.screenHistory[len(as.screenHistory)-1] != screenName {
			as.screenHistory = append(as.screenHistory, screenName)
		}

		// Create screen content with notification area
		as.currentContent = container.NewBorder(
			nil,    // Top
			nil,    // Bottom
			nil,    // Left
			nil,    // Right
			screen, // Center
		)

		// Wrap screen with top bar and notification area
		layout := container.NewBorder(
			container.NewVBox(
				as.connectionManager.GetContainer(), // Connection notifications
				as.createTopBar(),                   // Top (back button)
			),
			nil,               // Bottom
			nil,               // Left
			nil,               // Right
			as.currentContent, // Center (screen content)
		)
		as.window.SetContent(layout)
	}
}

// ShowWorkerProfile displays a worker's profile screen
func (as *AppState) ShowWorkerProfile(worker uiscreen.WorkerProfile) {
	as.currentWorker = &worker

	// Create the worker profile screen
	profileScreen := uiscreen.CreateWorkerProfileScreen(as, worker)

	// Add to history
	screenName := "worker_profile"
	if len(as.screenHistory) == 0 || as.screenHistory[len(as.screenHistory)-1] != screenName {
		as.screenHistory = append(as.screenHistory, screenName)
	}

	// Create content with notification area
	as.currentContent = container.NewBorder(
		nil,           // Top
		nil,           // Bottom
		nil,           // Left
		nil,           // Right
		profileScreen, // Center
	)

	// Show with notification area at top
	layout := container.NewBorder(
		as.connectionManager.GetContainer(), // Connection notifications
		nil,                                 // Bottom
		nil,                                 // Left
		nil,                                 // Right
		as.currentContent,                   // Center
	)
	as.window.SetContent(layout)
}

// createTopBar builds the top navigation bar with back button only
func (as *AppState) createTopBar() *fyne.Container {
	// Back button (only show if we're past the main screen)
	var backBtn *widget.Button
	currentScreen := ""
	if len(as.screenHistory) > 0 {
		currentScreen = as.screenHistory[len(as.screenHistory)-1]
	}

	// Show back button only if we're past the main screen (not on welcome, login, or main)
	showBackButton := len(as.screenHistory) > 1 &&
		currentScreen != "welcome" &&
		currentScreen != "login" &&
		currentScreen != "main"

	if showBackButton {
		backBtn = widget.NewButton("←", func() {
			as.goBack()
		})
	} else {
		// Empty placeholder if no back button needed
		backBtn = widget.NewButton("", nil)
		backBtn.Disable()
		backBtn.Hide() // Hide it completely
	}

	return container.NewBorder(
		nil, nil,
		backBtn, // Left
		nil,     // Right (theme toggle moved to profile screen)
		nil,     // Center
	)
}

// goBack navigates to the previous screen
func (as *AppState) goBack() {
	if len(as.screenHistory) > 1 {
		// Remove current screen
		as.screenHistory = as.screenHistory[:len(as.screenHistory)-1]
		// Get previous screen
		previousScreen := as.screenHistory[len(as.screenHistory)-1]
		// Remove it from history before showing (ShowScreen will re-add it)
		as.screenHistory = as.screenHistory[:len(as.screenHistory)-1]
		// Show previous screen
		as.ShowScreen(previousScreen)
	}
}

// toggleTheme switches between light and dark theme
func (as *AppState) ToggleTheme() {
	as.isDarkTheme = !as.isDarkTheme
	fmt.Println("Theme toggled. isDarkTheme:", as.isDarkTheme)

	// Save theme preference
	if as.isDarkTheme {
		as.preferences.SetTheme("dark")
	} else {
		as.preferences.SetTheme("light")
	}

	variant := theme.VariantLight
	if as.isDarkTheme {
		variant = theme.VariantDark
	}
	as.app.Settings().SetTheme(skilltheme.NewSkillKonnectTheme(variant))
	as.window.Content().Refresh()
}

// getThemeIcon returns the appropriate icon for current theme
func (as *AppState) GetThemeIcon() fyne.Resource {
	if as.isDarkTheme {
		return as.icons["lightTheme"]
	}
	return as.icons["darkTheme"]
}

// IsDarkTheme returns whether dark theme is currently active
func (as *AppState) IsDarkTheme() bool {
	return as.isDarkTheme
}

// GetImage returns an image resource by name
func (as *AppState) GetImage(name string) fyne.Resource {
	return as.icons[name]
}

// SetUserRole sets the user role (client or worker)
func (as *AppState) SetUserRole(role string) {
	as.userRole = role
	as.preferences.SetUserRole(role) // Save to preferences
	fmt.Println("User role set to:", role)
}

// GetUserRole returns the current user role
func (as *AppState) GetUserRole() string {
	return as.userRole
}

// ShowConnectionError displays a connection error notification
func (as *AppState) ShowConnectionError(status uiscreen.ConnectionStatus, message string) {
	// Show notification (manual dismiss required)
	as.connectionManager.ShowNotification(status, message, 0)
}

// HideConnectionError hides the connection error notification
func (as *AppState) HideConnectionError() {
	as.connectionManager.HideNotification()
}

// GetAPIService returns the API service
func (as *AppState) GetAPIService() *uiscreen.APIService {
	return as.apiService
}

// GetPreferences returns the preferences manager
func (as *AppState) GetPreferences() *uiscreen.PreferencesManager {
	return as.preferences
}

// initializeIcons creates and returns the map of all app icons
func initializeIcons() map[string]fyne.Resource {
	return map[string]fyne.Resource{
		"lightTheme":        resourceActiverPng,    // Light theme = active/on icon
		"darkTheme":         resourceDeactivatePng, // Dark theme = deactivate/off icon
		"logoImage":         resourceSkillDarLogoPng,
		"logoInCircle":      resourceSkillDarLogoInCerclePng,
		"plumbing":          resourcePlumberIcoPng,
		"electricity":       resourceElectricienIcoPng,
		"painting":          resourcePaintingIcoPng,
		"acFixing":          resourceACrepairePng,
		"homeCleaning":      resourceHomeCleaningPng,
		"smallRepairs":      resourceSmallrepairehandymanPng,
		"furnitureAssembly": resourceFurnitureassemblyPng,
		"waterLeakage":      resourceWaterleakagePng,
		"applianceRepair":   resourceAppliancerepairePng,
		"locksmith":         resourceLocksmithPng,
	}
}

func main() {
	// Create the app with unique ID for preferences
	a := app.NewWithID("com.skilldar.client")
	w := a.NewWindow("SkillDar")
	w.SetMaster()
	w.Resize(fyne.NewSize(390, 844)) // iPhone 12/13 size

	// Initialize preferences manager
	prefManager := uiscreen.NewPreferencesManager(a)

	// Initialize API service
	apiService := uiscreen.NewAPIService()
	// Set token if user is logged in
	if prefManager.IsLoggedIn() {
		apiService.SetToken(prefManager.GetAuthToken())
	}

	// Initialize app state
	state := &AppState{
		app:               a,
		window:            w,
		isDarkTheme:       prefManager.IsDarkTheme(), // Load theme preference
		preferences:       prefManager,
		apiService:        apiService,
		screens:           make(map[string]fyne.CanvasObject),
		icons:             initializeIcons(),
		userRole:          prefManager.GetUserRole(), // Load user role preference
		connectionManager: uiscreen.NewConnectionManager(a),
	}

	// Set initial theme based on saved preference
	variant := theme.VariantLight
	if state.isDarkTheme {
		variant = theme.VariantDark
	}
	a.Settings().SetTheme(skilltheme.NewSkillKonnectTheme(variant))

	// Register screens
	state.screens["welcome"] = uiscreen.CreateWelcomeScreen(state)
	state.screens["login"] = uiscreen.CreateLoginScreen(state)
	state.screens["main"] = uiscreen.CreateMainScreen(state)
	state.screens["profile"] = uiscreen.CreateProfileScreen(state)
	state.screens["edit_profile_client"] = uiscreen.CreateEditProfileClientScreen(state)

	// Always show welcome screen first
	state.ShowScreen("welcome")

	// Make sure window is visible
	w.Show()
	w.CenterOnScreen()

	// Show and run
	w.ShowAndRun()
}
