package ui

import (
	"fmt"
	"net/url"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

// CreateLoginScreen builds the login/welcome screen
func CreateLoginScreen(state AppState) fyne.CanvasObject {
	// Title
	title := canvas.NewText("Welcome to SkillDar", nil)
	title.Alignment = fyne.TextAlignCenter
	title.TextSize = 22
	title.TextStyle = fyne.TextStyle{Bold: true}

	// Subtitle
	subtitle := canvas.NewText("Connect skills, build networks", nil)
	subtitle.Alignment = fyne.TextAlignCenter
	subtitle.TextSize = 13

	// Email/Username entry
	emailEntry := widget.NewEntry()
	emailEntry.SetPlaceHolder("Email or Username")

	// Password entry with eye icon
	passwordEntry := widget.NewPasswordEntry()
	passwordEntry.SetPlaceHolder("Password")

	// Login button
	loginBtn := widget.NewButton("Login", func() {
		email := emailEntry.Text
		password := passwordEntry.Text
		if email == "" || password == "" {
			fmt.Println("Please fill in all fields")
			state.ShowConnectionError(StatusNoInternet, "Please fill in all fields")
		} else {
			// Simulate API call - check connection first
			apiConfig := DefaultAPIConfig()

			// Example: Check connection before login
			isConnected, status, message := CheckConnection(apiConfig)

			if !isConnected {
				// Show connection error
				state.ShowConnectionError(status, message)
				return
			}

			// Simulate successful login
			fmt.Println("Logged in as:", email)
			state.HideConnectionError()

			// Navigate to main screen
			state.ShowScreen("main")
		}
	})
	loginBtn.Importance = widget.HighImportance

	// Divider with OR in the middle using ASCII lines
	orDivider := widget.NewLabel("────────────── OR ──────────────")
	orDivider.Alignment = fyne.TextAlignCenter

	// Facebook login button
	facebookBtn := widget.NewButton("  Continue with Facebook", func() {
		fmt.Println("Opening Facebook OAuth...")
		// Facebook OAuth URL (replace with your app credentials)
		clientID := "YOUR_FACEBOOK_APP_ID"
		redirectURI := "http://localhost:8080/auth/facebook/callback"

		facebookAuthURL := fmt.Sprintf(
			"https://www.facebook.com/v12.0/dialog/oauth?client_id=%s&redirect_uri=%s&scope=email,public_profile",
			clientID,
			url.QueryEscape(redirectURI),
		)

		authURL, _ := url.Parse(facebookAuthURL)
		fyne.CurrentApp().OpenURL(authURL)

		// TODO: Set up callback server to receive the auth code
		// For now, navigate to main screen for testing
		state.ShowScreen("main")
	})

	// Google login button
	googleBtn := widget.NewButton("  Continue with Google", func() {
		fmt.Println("Opening Google OAuth...")
		// Google OAuth URL (replace with your app credentials)
		clientID := "YOUR_GOOGLE_CLIENT_ID"
		redirectURI := "http://localhost:8080/auth/google/callback"

		googleAuthURL := fmt.Sprintf(
			"https://accounts.google.com/o/oauth2/v2/auth?client_id=%s&redirect_uri=%s&response_type=code&scope=email%%20profile",
			clientID,
			url.QueryEscape(redirectURI),
		)

		authURL, _ := url.Parse(googleAuthURL)
		fyne.CurrentApp().OpenURL(authURL)

		// TODO: Set up callback server to receive the auth code
		// For now, navigate to main screen for testing
		state.ShowScreen("main")
	})

	// Main content with proper spacing
	content := container.NewVBox(
		layout.NewSpacer(),
		container.NewCenter(title),
		container.NewCenter(subtitle),
		layout.NewSpacer(),
		emailEntry,
		passwordEntry,
		loginBtn,
		layout.NewSpacer(),
		orDivider,
		layout.NewSpacer(),
		facebookBtn,
		googleBtn,
		layout.NewSpacer(),
		layout.NewSpacer(),
	)

	return container.NewPadded(content)
}
