package ui

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	xwidget "fyne.io/x/fyne/widget"
)

// CreateMapScreen builds a map screen showing worker locations
func CreateMapScreen(state AppState) fyne.CanvasObject {
	// Create the map widget
	mapWidget := xwidget.NewMap()

	// Wrap in a container that expands to fill all available space
	// This ensures the map takes up the full content area
	mapContainer := container.NewStack(mapWidget)

	return mapContainer
}
