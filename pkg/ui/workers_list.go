package ui

import (
	"fmt"
	skilltheme "skillDar/pkg/theme"
	"strings"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

// createWorkersSection creates the workers list section with pagination
func createWorkersSection(state AppState, selectedCategory *string) (*widget.Label, *fyne.Container, *container.Scroll, func()) {
	workersContainer := container.NewVBox()
	workersLabel := widget.NewLabel("Loading workers...")
	workersLabel.TextStyle = fyne.TextStyle{Bold: true}

	// Pagination state - use pointers to allow modification
	paginationState := struct {
		currentPage    int
		pageLimit      int
		isLoading      bool
		hasMoreWorkers bool
	}{
		currentPage:    1,
		pageLimit:      10,
		isLoading:      false,
		hasMoreWorkers: true,
	}

	// Worker loading queue to prevent simultaneous requests
	type loadingQueue struct {
		isProcessing bool
		queue        chan struct{}
	}

	queue := &loadingQueue{
		isProcessing: false,
		queue:        make(chan struct{}, 1), // Single worker queue
	}

	// Separate worker pool goroutine that processes loading sequentially
	go func() {
		for range queue.queue {
			queue.isProcessing = true
			// Signal to retry after queue empties
			<-time.After(100 * time.Millisecond)
		}
	}()

	// Function to load workers from API
	var loadWorkers = func() {
		if paginationState.isLoading || !paginationState.hasMoreWorkers {
			fmt.Printf("Skipping load - isLoading: %v, hasMoreWorkers: %v\n", paginationState.isLoading, paginationState.hasMoreWorkers)
			return
		}

		// Skip if already processing
		if queue.isProcessing {
			fmt.Println("Load already in progress, skipping...")
			return
		}

		paginationState.isLoading = true

		// Submit load task to queue
		go func() {
			// Get API service
			apiService := state.GetAPIService()

			fmt.Printf("[DEBUG] [%s] Starting worker load - Page: %d, Limit: %d, Category: %s\n", time.Now().Format("15:04:05.000"), paginationState.currentPage, paginationState.pageLimit, *selectedCategory)

			// Fetch workers from API with category filter
			workersResp, err := apiService.GetWorkersByCategory(paginationState.currentPage, paginationState.pageLimit, *selectedCategory)

			// Handle error first (no async needed for errors)
			if err != nil {
				fmt.Printf("Failed to load workers (Page %d): %v\n", paginationState.currentPage, err)

				// Check if error is due to token expiration (401 Unauthorized)
				errMsg := err.Error()
				if strings.Contains(errMsg, "401") || strings.Contains(errMsg, "Unauthorized") || strings.Contains(errMsg, "no authentication token") {
					fmt.Println("Token expired or invalid - redirecting to login")
					fyne.Do(func() {
						// Clear authentication data
						prefs := state.GetPreferences()
						prefs.ClearAuthData()
						// Clear API service token
						apiService.SetToken("")
						// Redirect to login screen
						state.ShowScreen("login")
					})
					return
				}

				// Show user-friendly error message for other errors
				fyne.Do(func() {
					errorMsg := "Unable to load workers. Please check your connection and try again."
					workersLabel.SetText(errorMsg)
					paginationState.isLoading = false
					queue.isProcessing = false
				})
				return
			}

			fmt.Printf("[DEBUG] [%s] Worker load completed: %d workers received, total in category: %d\n", time.Now().Format("15:04:05.000"), len(workersResp.Workers), workersResp.Total)

			// Update label with count (quick UI update)
			fyne.Do(func() {
				workersLabel.SetText(fmt.Sprintf("Available Workers Near You (%d)", workersResp.Total))
			})

			// Render workers asynchronously in a separate goroutine for better responsiveness
			go func() {
				visibleCount := 5 // Show 5 workers initially
				deferredCount := 0
				visibleAmount := visibleCount
				if len(workersResp.Workers) < visibleCount {
					visibleAmount = len(workersResp.Workers)
				}
				if len(workersResp.Workers) > visibleCount {
					deferredCount = len(workersResp.Workers) - visibleCount
				}
				fmt.Printf("[DEBUG] [%s] Starting async render of %d visible workers, deferring %d workers\n", time.Now().Format("15:04:05.000"), visibleAmount, deferredCount)

				// Render all workers instantly (no delay)
				for i, worker := range workersResp.Workers {
					workerCopy := worker
					workerIndex := i

					// Render immediately in separate goroutine
					go func(idx int, w Worker) {
						fyne.Do(func() {
							// Only add if not already rendered
							if idx >= len(workersContainer.Objects) {
								rating := fmt.Sprintf("%.1f", w.Rating)
								reviews := fmt.Sprintf("%d", w.Reviews)
								price := w.Price

								workerCard := createSimpleWorkerCard(
									state,
									w.Name,
									w.Profession,
									rating,
									w.Distance,
									reviews,
									price,
									w.Available,
								)
								workersContainer.Add(workerCard)
								workersContainer.Refresh()
								fmt.Printf("[DEBUG] [%s] Rendered worker #%d\n", time.Now().Format("15:04:05.000"), idx+1)
							}
						})
					}(workerIndex, workerCopy)
				}

				// After a short delay, finalize state
				time.Sleep(time.Duration(len(workersResp.Workers)*20) * time.Millisecond)
				fyne.Do(func() {
					workersContainer.Refresh()
					fmt.Printf("[DEBUG] [%s] Async render complete: %d workers rendered\n", time.Now().Format("15:04:05.000"), len(workersResp.Workers))

					// Check if there are more workers
					if len(workersResp.Workers) < paginationState.pageLimit {
						paginationState.hasMoreWorkers = false
					} else {
						paginationState.currentPage++
					}

					paginationState.isLoading = false
					queue.isProcessing = false
				})
			}()
		}()
	}

	// Make workers scrollable with minimum height
	workersScroll := container.NewVScroll(workersContainer)
	workersScroll.SetMinSize(fyne.NewSize(0, 300)) // Fixed height only, width adapts to screen
	workersScroll.OnScrolled = func(pos fyne.Position) {
		fmt.Printf("Workers scrolled to position: X=%.2f, Y=%.2f\n", pos.X, pos.Y)

		// Check if we're near the bottom - load more workers
		if pos.Y > 40 && paginationState.hasMoreWorkers {
			fmt.Println(">>> Loading more workers...")
			loadWorkers()
		}
	}

	// Function to reset pagination when filtering
	resetPagination := func() {
		paginationState.currentPage = 1
		paginationState.pageLimit = 10
		paginationState.isLoading = false
		paginationState.hasMoreWorkers = true
		queue.isProcessing = false // Reset queue state
		workersContainer.Objects = nil
		workersContainer.Refresh()
		// Reset scroll to top so user sees new workers immediately
		workersScroll.ScrollToTop()
		fmt.Println("Pagination reset, now loading workers with category:", *selectedCategory)
		// Now that loadWorkers is defined, we can call it
		loadWorkers()
	}

	// Load initial workers
	loadWorkers()

	return workersLabel, workersContainer, workersScroll, resetPagination
}

// createSimpleWorkerCard creates a clickable worker card for clients
func createSimpleWorkerCard(state AppState, name, profession, rating, distance, reviewCount, price string, available bool) fyne.CanvasObject {
	// Profile picture placeholder
	profileCircle := canvas.NewCircle(theme.Color(skilltheme.ColorNameHighlight))
	profilePic := container.NewStack(profileCircle)
	profilePic.Resize(fyne.NewSize(50, 50))

	nameLabel := widget.NewLabel(name)
	nameLabel.TextStyle = fyne.TextStyle{Bold: true}

	// Verified badge
	verifiedLabel := widget.NewLabel("✓ Verified")
	verifiedBadge := container.NewHBox(
		widget.NewLabel(name),
		verifiedLabel,
	)

	professionLabel := widget.NewLabel(profession)

	ratingLabel := widget.NewLabel("⭐ " + rating)
	reviewLabel := widget.NewLabel("(" + reviewCount + ")")
	distanceLabel := widget.NewLabel("📍 " + distance)

	priceLabel := widget.NewLabel(price)
	priceLabel.TextStyle = fyne.TextStyle{Bold: true}

	statusLabel := widget.NewLabel("✅ Available")
	statusLabel.Importance = widget.SuccessImportance
	if !available {
		statusLabel.Text = "⏰ Busy"
		statusLabel.Importance = widget.WarningImportance
	}

	info := container.NewVBox(
		verifiedBadge,
		professionLabel,
		container.NewHBox(ratingLabel, reviewLabel, distanceLabel),
	)

	rightSide := container.NewVBox(
		priceLabel,
		statusLabel,
	)

	cardContent := container.NewBorder(
		nil, nil,
		container.NewHBox(profilePic, info),
		rightSide,
	)

	// Create a button that wraps the content
	btn := widget.NewButton("", func() {
		// Create worker profile and show screen
		worker := WorkerProfile{
			Name:            name,
			Profession:      profession,
			Rating:          4.9,
			ReviewCount:     127,
			Distance:        distance,
			HourlyRate:      180,
			CompletedJobs:   340,
			YearsExperience: 12,
			Available:       available,
			About:           "Professional installation and maintenance of electrical wiring, fixtures, and appliances.",
			Skills:          []string{"Plumbing", "Repair", "Installation"},
		}
		state.ShowWorkerProfile(worker)
	})

	// Stack content on button with minimal padding
	return container.NewStack(btn, cardContent)
}
