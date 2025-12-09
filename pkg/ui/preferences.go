package ui

import (
	"fyne.io/fyne/v2"
)

// PreferenceKeys defines all preference keys used in the app
const (
	PrefKeyTheme    = "app.theme"       // "light" or "dark"
	PrefKeyUserRole = "user.role"       // "client" or "worker"
	PrefKeyUserID   = "user.id"         // User ID
	PrefKeyUsername = "user.username"   // Username
	PrefKeyEmail    = "user.email"      // Email
	PrefKeyToken    = "user.auth_token" // Auth token
	PrefKeyLanguage = "app.language"    // "en", "fr", "ar"
)

// PreferencesManager handles all app preferences
type PreferencesManager struct {
	prefs fyne.Preferences
}

// NewPreferencesManager creates a new preferences manager
func NewPreferencesManager(app fyne.App) *PreferencesManager {
	return &PreferencesManager{
		prefs: app.Preferences(),
	}
}

// Theme preferences
func (pm *PreferencesManager) GetTheme() string {
	return pm.prefs.StringWithFallback(PrefKeyTheme, "light")
}

func (pm *PreferencesManager) SetTheme(theme string) {
	pm.prefs.SetString(PrefKeyTheme, theme)
}

func (pm *PreferencesManager) IsDarkTheme() bool {
	return pm.GetTheme() == "dark"
}

func (pm *PreferencesManager) ToggleTheme() string {
	if pm.IsDarkTheme() {
		pm.SetTheme("light")
		return "light"
	}
	pm.SetTheme("dark")
	return "dark"
}

// User role preferences
func (pm *PreferencesManager) GetUserRole() string {
	return pm.prefs.StringWithFallback(PrefKeyUserRole, "client")
}

func (pm *PreferencesManager) SetUserRole(role string) {
	pm.prefs.SetString(PrefKeyUserRole, role)
}

// User authentication preferences
func (pm *PreferencesManager) GetUserID() int {
	return pm.prefs.Int(PrefKeyUserID)
}

func (pm *PreferencesManager) SetUserID(id int) {
	pm.prefs.SetInt(PrefKeyUserID, id)
}

func (pm *PreferencesManager) GetUsername() string {
	return pm.prefs.String(PrefKeyUsername)
}

func (pm *PreferencesManager) SetUsername(username string) {
	pm.prefs.SetString(PrefKeyUsername, username)
}

func (pm *PreferencesManager) GetEmail() string {
	return pm.prefs.String(PrefKeyEmail)
}

func (pm *PreferencesManager) SetEmail(email string) {
	pm.prefs.SetString(PrefKeyEmail, email)
}

func (pm *PreferencesManager) GetAuthToken() string {
	return pm.prefs.String(PrefKeyToken)
}

func (pm *PreferencesManager) SetAuthToken(token string) {
	pm.prefs.SetString(PrefKeyToken, token)
}

func (pm *PreferencesManager) IsLoggedIn() bool {
	return pm.GetAuthToken() != ""
}

func (pm *PreferencesManager) ClearAuthData() {
	pm.prefs.SetString(PrefKeyUserID, "")
	pm.prefs.SetString(PrefKeyUsername, "")
	pm.prefs.SetString(PrefKeyEmail, "")
	pm.prefs.SetString(PrefKeyToken, "")
}

// Language preferences
func (pm *PreferencesManager) GetLanguage() string {
	return pm.prefs.StringWithFallback(PrefKeyLanguage, "en")
}

func (pm *PreferencesManager) SetLanguage(lang string) {
	pm.prefs.SetString(PrefKeyLanguage, lang)
}

// Generic get/set methods for custom preferences
func (pm *PreferencesManager) GetString(key string, fallback string) string {
	return pm.prefs.StringWithFallback(key, fallback)
}

func (pm *PreferencesManager) SetString(key, value string) {
	pm.prefs.SetString(key, value)
}

func (pm *PreferencesManager) GetBool(key string, fallback bool) bool {
	return pm.prefs.BoolWithFallback(key, fallback)
}

func (pm *PreferencesManager) SetBool(key string, value bool) {
	pm.prefs.SetBool(key, value)
}

func (pm *PreferencesManager) GetInt(key string, fallback int) int {
	return pm.prefs.IntWithFallback(key, fallback)
}

func (pm *PreferencesManager) SetInt(key string, value int) {
	pm.prefs.SetInt(key, value)
}

func (pm *PreferencesManager) GetFloat(key string, fallback float64) float64 {
	return pm.prefs.FloatWithFallback(key, fallback)
}

func (pm *PreferencesManager) SetFloat(key string, value float64) {
	pm.prefs.SetFloat(key, value)
}
