package ui

import (
	"fmt"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

// CreateEditProfileClientScreen builds the client profile edit screen
func CreateEditProfileClientScreen(state AppState) fyne.CanvasObject {
	// Header
	title := widget.NewLabel("Edit Client Profile")
	title.Alignment = fyne.TextAlignCenter
	title.TextStyle = fyne.TextStyle{Bold: true}

	// Profile picture section
	profilePicBtn := widget.NewButton("Change Profile Picture", func() {
		fmt.Println("Change profile picture clicked")
		// TODO: Implement image picker
	})

	// Form fields
	nameEntry := widget.NewEntry()
	nameEntry.SetPlaceHolder("Full Name")
	nameEntry.SetText("John Doe") // Pre-filled example

	emailEntry := widget.NewEntry()
	emailEntry.SetPlaceHolder("Email")
	emailEntry.SetText("john@example.com")

	phoneEntry := widget.NewEntry()
	phoneEntry.SetPlaceHolder("Phone Number")
	phoneEntry.SetText("+2164567890")

	locationEntry := widget.NewEntry()
	locationEntry.SetPlaceHolder("Location/Address")
	locationEntry.SetText("New York, USA")

	bioEntry := widget.NewMultiLineEntry()
	bioEntry.SetPlaceHolder("Tell us about yourself...")
	bioEntry.SetMinRowsVisible(4)

	// Save button
	saveBtn := widget.NewButton("Save Changes", func() {
		fmt.Println("Saving client profile...")
		fmt.Println("Name:", nameEntry.Text)
		fmt.Println("Email:", emailEntry.Text)
		fmt.Println("Phone:", phoneEntry.Text)
		fmt.Println("Location:", locationEntry.Text)
		fmt.Println("Bio:", bioEntry.Text)
		// TODO: Send to API
		state.ShowScreen("main")
	})
	saveBtn.Importance = widget.HighImportance

	// Cancel button
	cancelBtn := widget.NewButton("Cancel", func() {
		state.ShowScreen("main")
	})

	// Layout
	content := container.NewVBox(
		title,
		layout.NewSpacer(),
		profilePicBtn,
		widget.NewLabel("Personal Information"),
		nameEntry,
		emailEntry,
		phoneEntry,
		locationEntry,
		widget.NewLabel("Bio"),
		bioEntry,
		layout.NewSpacer(),
		saveBtn,
		cancelBtn,
	)

	scroll := container.NewVScroll(content)
	return scroll
}
