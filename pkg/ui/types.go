package ui

import "fyne.io/fyne/v2"

// AppState defines the interface for app state management
// This allows screens to access navigation and app-level state
type AppState interface {
	ShowScreen(screenName string)
	GetImage(name string) fyne.Resource
	SetUserRole(role string)
	GetUserRole() string
}
