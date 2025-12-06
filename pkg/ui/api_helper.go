package ui

import (
	"fmt"
	"net/http"
	"time"
)

// APIConfig holds API configuration
type APIConfig struct {
	BaseURL       string
	Timeout       time.Duration
	RetryAttempts int
	RetryDelay    time.Duration
}

// DefaultAPIConfig returns default API configuration
func DefaultAPIConfig() *APIConfig {
	return &APIConfig{
		BaseURL:       "https://developpement-skillkonnect.ngrok.app/api/v1",
		Timeout:       10 * time.Second,
		RetryAttempts: 3,
		RetryDelay:    2 * time.Second,
	}
}

// CheckConnection verifies if the API server is reachable
func CheckConnection(config *APIConfig) (bool, ConnectionStatus, string) {
	client := &http.Client{
		Timeout: config.Timeout,
	}

	// Try to reach the API health endpoint
	healthURL := config.BaseURL + "/health" // Adjust endpoint as needed

	resp, err := client.Get(healthURL)
	if err != nil {
		// Check if it's a network error or timeout
		if netErr, ok := err.(interface{ Timeout() bool }); ok && netErr.Timeout() {
			return false, StatusSlowConnection, "Connection timeout - slow network"
		}
		return false, StatusNoInternet, "No internet connection"
	}
	defer resp.Body.Close()

	if resp.StatusCode >= 500 {
		return false, StatusServerDown, "Server is currently down"
	}

	if resp.StatusCode >= 400 {
		return false, StatusServerDown, fmt.Sprintf("Server error: %d", resp.StatusCode)
	}

	return true, StatusConnected, "Connected"
}

// PeriodicConnectionCheck runs a periodic connection check
func PeriodicConnectionCheck(config *APIConfig, interval time.Duration, onStatusChange func(ConnectionStatus, string)) {
	ticker := time.NewTicker(interval)
	var lastStatus ConnectionStatus = StatusConnected

	go func() {
		for range ticker.C {
			_, status, message := CheckConnection(config)
			if status != lastStatus {
				lastStatus = status
				if onStatusChange != nil {
					onStatusChange(status, message)
				}
			}
		}
	}()
}

// APICall is a helper function to make API calls with error handling
func APICall(method, url string, onError func(ConnectionStatus, string)) error {
	client := &http.Client{
		Timeout: 10 * time.Second,
	}

	req, err := http.NewRequest(method, url, nil)
	if err != nil {
		if onError != nil {
			onError(StatusNoInternet, "Failed to create request")
		}
		return err
	}

	resp, err := client.Do(req)
	if err != nil {
		if netErr, ok := err.(interface{ Timeout() bool }); ok && netErr.Timeout() {
			if onError != nil {
				onError(StatusSlowConnection, "Request timeout")
			}
		} else {
			if onError != nil {
				onError(StatusNoInternet, "No internet connection")
			}
		}
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode >= 500 {
		if onError != nil {
			onError(StatusServerDown, "Server error")
		}
		return fmt.Errorf("server error: %d", resp.StatusCode)
	}

	return nil
}
