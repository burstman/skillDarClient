package ui

import (
	"fmt"
	"image/color"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

// ConnectionStatus represents the current connection state
type ConnectionStatus int

const (
	StatusConnected ConnectionStatus = iota
	StatusNoInternet
	StatusServerDown
	StatusSlowConnection
)

// ConnectionNotification is a custom widget that shows connection status
type ConnectionNotification struct {
	widget.BaseWidget
	status    ConnectionStatus
	message   string
	container *fyne.Container
	onDismiss func()
	app       fyne.App
}

// NewConnectionNotification creates a new connection status notification
func NewConnectionNotification(app fyne.App, status ConnectionStatus, message string, onDismiss func()) *ConnectionNotification {
	cn := &ConnectionNotification{
		app:       app,
		status:    status,
		message:   message,
		onDismiss: onDismiss,
	}
	cn.ExtendBaseWidget(cn)
	return cn
}

// CreateRenderer implements fyne.Widget
func (cn *ConnectionNotification) CreateRenderer() fyne.WidgetRenderer {
	cn.container = cn.createContent()
	return widget.NewSimpleRenderer(cn.container)
}

func (cn *ConnectionNotification) createContent() *fyne.Container {
	prefs := cn.app.Preferences()
	var bgColor, textColor, iconText string

	switch cn.status {
	case StatusNoInternet:
		bgColor = prefs.StringWithFallback("color.no_internet.bg", "#FF5252")     // Red
		textColor = prefs.StringWithFallback("color.no_internet.text", "#FFFFFF") // White
		iconText = "📡"
	case StatusServerDown:
		bgColor = prefs.StringWithFallback("color.server_down.bg", "#FF6F00")     // Orange
		textColor = prefs.StringWithFallback("color.server_down.text", "#FFFFFF") // White
		iconText = "⚠️"
	case StatusSlowConnection:
		bgColor = prefs.StringWithFallback("color.slow_connection.bg", "#FFC107")     // Amber
		textColor = prefs.StringWithFallback("color.slow_connection.text", "#000000") // Black
		iconText = "🐌"
	default: // StatusConnected
		bgColor = prefs.StringWithFallback("color.connected.bg", "#4CAF50")     // Green
		textColor = prefs.StringWithFallback("color.connected.text", "#FFFFFF") // White
		iconText = "✓"
	}

	// Background
	bgCol, _ := parseHexColor(bgColor)
	bg := canvas.NewRectangle(bgCol)

	textCol, _ := parseHexColor(textColor)
	// Icon
	icon := canvas.NewText(iconText, textCol)
	icon.TextSize = 18

	// Message
	messageLabel := canvas.NewText(cn.message, textCol)
	messageLabel.TextSize = 14

	// Close button (X)
	closeIcon := widget.NewButton("✕", func() {
		if cn.onDismiss != nil {
			cn.onDismiss()
		}
	})

	// Content layout
	contentWithBorder := container.NewBorder(nil, nil, icon, closeIcon, messageLabel)

	// Stack: background, content, invisible tappable overlay
	return container.NewStack(bg, contentWithBorder)
}

// SetStatus updates the notification status
func (cn *ConnectionNotification) SetStatus(status ConnectionStatus, message string) {
	cn.status = status
	cn.message = message
	cn.Refresh()
}

// ConnectionManager manages connection notifications
type ConnectionManager struct {
	notification *ConnectionNotification
	container    *fyne.Container
	isShowing    bool
	app          fyne.App
}

// NewConnectionManager creates a new connection manager
func NewConnectionManager(app fyne.App) *ConnectionManager {
	return &ConnectionManager{
		container: container.NewVBox(),
		isShowing: false,
		app:       app,
	}
}

// ShowNotification displays a connection status notification
func (cm *ConnectionManager) ShowNotification(status ConnectionStatus, message string, duration time.Duration) {
	// Create notification with dismiss callback
	dismissFunc := func() {
		cm.HideNotification()
	}

	if cm.notification == nil {
		cm.notification = NewConnectionNotification(cm.app, status, message, dismissFunc)
	} else {
		cm.notification.onDismiss = dismissFunc
		cm.notification.SetStatus(status, message)
	}

	if !cm.isShowing {
		cm.container.Objects = []fyne.CanvasObject{cm.notification}
		cm.container.Refresh()
		cm.isShowing = true
	}

	// Note: Auto-dismiss removed to avoid threading issues
	// User must click the notification to dismiss it
}

// HideNotification hides the current notification
func (cm *ConnectionManager) HideNotification() {
	cm.container.Objects = []fyne.CanvasObject{}
	cm.container.Refresh()
	cm.isShowing = false
}

// GetContainer returns the notification container to be added to the UI
func (cm *ConnectionManager) GetContainer() *fyne.Container {
	return cm.container
}

// Helper function to parse hex colors
func parseHexColor(s string) (color.Color, error) {
	c := color.RGBA{A: 255}
	var err error
	if len(s) == 7 && s[0] == '#' { // #RRGGBB
		_, err = fmt.Sscanf(s, "#%02x%02x%02x", &c.R, &c.G, &c.B)
	} else if len(s) == 9 && s[0] == '#' { // #RRGGBBAA
		_, err = fmt.Sscanf(s, "#%02x%02x%02x%02x", &c.R, &c.G, &c.B, &c.A)
	} else {
		return color.Black, fmt.Errorf("invalid hex color format: %s", s)
	}
	return c, err
}
