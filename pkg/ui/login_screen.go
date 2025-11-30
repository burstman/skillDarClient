package ui

import (
	"fmt"
	"net/url"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

// CreateLoginScreen builds the login/welcome screen
func CreateLoginScreen(state AppState) fyne.CanvasObject {
	title := widget.NewLabel("Welcome to SkillKonnect")
	title.Alignment = fyne.TextAlignCenter
	title.TextStyle = fyne.TextStyle{Bold: true}

	subtitle := widget.NewLabel("Connect skills, build networks")
	subtitle.Alignment = fyne.TextAlignCenter

	emailEntry := widget.NewEntry()
	emailEntry.SetPlaceHolder("Email or Username")

	passwordEntry := widget.NewPasswordEntry()
	passwordEntry.SetPlaceHolder("Password")

	loginBtn := widget.NewButton("Login", func() {
		email := emailEntry.Text
		password := passwordEntry.Text
		if email == "" || password == "" {
			fmt.Println("Please fill in all fields")
		} else {
			fmt.Println("Logged in as:", email)
			// Navigate to main screen
			state.ShowScreen("main")
		}
	})
	loginBtn.Importance = widget.HighImportance

	// Divider
	orLabel := widget.NewLabel("────── OR ──────")
	orLabel.Alignment = fyne.TextAlignCenter

	// Facebook login button
	facebookBtn := widget.NewButton("Continue with Facebook", func() {
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
	googleBtn := widget.NewButton("Continue with Google", func() {
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

	content := container.NewVBox(
		layout.NewSpacer(),
		title,
		subtitle,
		emailEntry,
		passwordEntry,
		loginBtn,
		orLabel,
		facebookBtn,
		googleBtn,
		layout.NewSpacer(),
	)

	return container.NewPadded(container.NewCenter(content))
}
