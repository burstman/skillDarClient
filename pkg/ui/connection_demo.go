package ui

import (
	"time"
)

// Example usage of connection notifications

// Example 1: Show error when API call fails
func ExampleAPICallWithNotification(state AppState) {
	// Simulate an API call
	apiConfig := DefaultAPIConfig()

	err := APICall("GET", apiConfig.BaseURL+"/workers", func(status ConnectionStatus, message string) {
		// This callback is called when there's a connection error
		state.ShowConnectionError(status, message)
	})

	if err != nil {
		// Error notification already shown via callback
		return
	}

	// Success - hide any existing error
	state.HideConnectionError()
}

// Example 2: Manual connection check
func ExampleManualConnectionCheck(state AppState) {
	apiConfig := DefaultAPIConfig()

	isConnected, status, message := CheckConnection(apiConfig)

	if !isConnected {
		// Show error notification
		state.ShowConnectionError(status, message)
	} else {
		// Hide notification if connection is restored
		state.HideConnectionError()
	}
}

// Example 3: Periodic connection monitoring
func ExamplePeriodicMonitoring(state AppState) {
	apiConfig := DefaultAPIConfig()

	// Check connection every 30 seconds
	PeriodicConnectionCheck(apiConfig, 30*time.Second, func(status ConnectionStatus, message string) {
		if status != StatusConnected {
			state.ShowConnectionError(status, message)
		} else {
			state.HideConnectionError()
		}
	})
}

// Example 4: Show different types of notifications
func ExampleShowDifferentNotifications(state AppState) {
	// No internet
	state.ShowConnectionError(StatusNoInternet, "No internet connection. Please check your network.")

	time.Sleep(3 * time.Second)

	// Server down
	state.ShowConnectionError(StatusServerDown, "Server is temporarily unavailable. Please try again later.")

	time.Sleep(3 * time.Second)

	// Slow connection
	state.ShowConnectionError(StatusSlowConnection, "Slow connection detected. Loading may take longer.")

	time.Sleep(3 * time.Second)

	// Hide notification
	state.HideConnectionError()
}
