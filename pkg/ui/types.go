package ui

import "fyne.io/fyne/v2"

// AppState defines the interface for app state management
// This allows screens to access navigation and app-level state
// AppState defines the interface for managing the application's global state and UI interactions.
// It provides methods for screen navigation, theme management, user role handling, and resource access.
// Implementations of this interface should ensure thread-safe operations when called from multiple goroutines.
type AppState interface {
	ShowScreen(screenName string)
	ShowWorkerProfile(worker WorkerProfile) //
	GetImage(name string) fyne.Resource
	SetUserRole(role string)
	GetUserRole() string
	ToggleTheme()
	GetThemeIcon() fyne.Resource
	IsDarkTheme() bool
	ShowConnectionError(status ConnectionStatus, message string)
	HideConnectionError()
}
