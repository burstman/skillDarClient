package ui

import (
	"context"
	"fmt"
	"image/color"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"fyne.io/x/fyne/mousemappan"
)

// createMapSection creates the map section with lazy loading
// Returns: mapSection and a channel that signals when map is loaded
// The map is created immediately and the channel sends when ready
func createMapSection(state AppState) (fyne.CanvasObject, <-chan bool) {
	mapLabel := widget.NewLabel("Workers Near You")
	mapLabel.TextStyle = fyne.TextStyle{Bold: true}

	// Spacer to reserve space for the map
	mapSpacer := canvas.NewRectangle(color.Color(color.RGBA{0, 0, 0, 0}))
	mapSpacer.SetMinSize(fyne.NewSize(0, 350))

	// Track if map is loaded and loading
	var mapLoaded bool = false
	var mapLoading bool = false
	mapLoadedChan := make(chan bool, 1) // Channel to signal when map is loaded

	// Create a scroll-aware container that loads the map when scrolled into view
	mapSection := container.NewVBox(
		widget.NewSeparator(),
		mapLabel,
		mapSpacer,
	)

	// Load map on demand using a goroutine
	loadMapAsync := func() {
		if mapLoaded || mapLoading {
			return
		}

		mapLoading = true
		fmt.Printf("[DEBUG] [%s] Starting map widget creation...\n", time.Now().Format("15:04:05.000"))
		go func() {
			// Create a context with timeout for map widget creation
			ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
			defer cancel()

			fmt.Printf("[DEBUG] [%s] Creating map widget asynchronously (10s timeout)...\n", time.Now().Format("15:04:05.000"))

			// Create the map widget (this can be slow on first load)
			mapWidget := mousemappan.NewMapWithOptions(
				mousemappan.WithOsmTiles(),
				mousemappan.WithZoomButtons(true),
				mousemappan.WithScrollButtons(false),
			)

			// Check if context expired or cancelled
			if ctx.Err() != nil {
				fmt.Printf("[DEBUG] [%s] Map widget creation timed out\n", time.Now().Format("15:04:05.000"))
				mapLoading = false
				return
			}

			fmt.Printf("[DEBUG] [%s] Map widget created, configuring zoom and location...\n", time.Now().Format("15:04:05.000"))
			// Set zoom level and center on Tunis, Tunisia
			mapWidget.Zoom(13)
			mapWidget.CenterOnLocation(36.8065, 10.1815)
			fmt.Printf("[DEBUG] [%s] Map configured - Zoom: 13, Location: Tunis (36.8065, 10.1815)\n", time.Now().Format("15:04:05.000"))

			// Update UI on main thread
			fyne.Do(func() {

				// Create a spacer with the desired height and stack it with the map
				mapHeightSpacer := canvas.NewRectangle(color.Color(color.RGBA{0, 0, 0, 0}))
				mapHeightSpacer.SetMinSize(fyne.NewSize(0, 350))

				// Stack the spacer (for height) with the map widget (for content)
				mapStack := container.NewStack(mapHeightSpacer, mapWidget)

				// Replace the spacer with the stacked map
				mapSection.Objects[2] = mapStack
				mapSection.Refresh()
				mapLoaded = true
				mapLoading = false
				fmt.Printf("[DEBUG] [%s] Map successfully loaded and displayed\n", time.Now().Format("15:04:05.000"))
				// Signal that map is loaded
				mapLoadedChan <- true
			})
		}()
	}

	// Load map immediately (no delay for welcome screen)
	go func() {
		fmt.Printf("[DEBUG] [%s] Map load scheduled: starting immediately for welcome screen...\n", time.Now().Format("15:04:05.000"))
		if !mapLoaded && !mapLoading {
			loadMapAsync()
		}
	}()

	return mapSection, mapLoadedChan
}
