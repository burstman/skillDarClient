package ui

import (
	"fmt"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

// createCategoriesSection creates the professional categories section
func createCategoriesSection(state AppState, filterByCategory func(string)) fyne.CanvasObject {
	categoriesLabel := widget.NewLabel("Professional Categories")
	categoriesLabel.TextStyle = fyne.TextStyle{Bold: true}

	// Create category buttons with icons and filter callback
	plumbingCard := createCategoryButtonWithFilter(state, "plumbing", "Plumbing", "plumber", filterByCategory)
	electricityCard := createCategoryButtonWithFilter(state, "electricity", "Electricity", "electrician", filterByCategory)
	paintingCard := createCategoryButtonWithFilter(state, "painting", "Painting", "painter", filterByCategory)
	acFixingCard := createCategoryButtonWithFilter(state, "acFixing", "AC Fixing", "ac-technician", filterByCategory)
	homeCleaningCard := createCategoryButtonWithFilter(state, "homeCleaning", "Home Cleaning", "cleaner", filterByCategory)
	smallRepairsCard := createCategoryButtonWithFilter(state, "smallRepairs", "Small Repairs", "handyman", filterByCategory)
	furnitureCard := createCategoryButtonWithFilter(state, "furnitureAssembly", "Furniture Assembly", "furniture-assembler", filterByCategory)
	waterLeakCard := createCategoryButtonWithFilter(state, "waterLeakage", "Water Leakage", "plumber", filterByCategory)
	applianceCard := createCategoryButtonWithFilter(state, "applianceRepair", "Appliance Repair", "appliance-repair", filterByCategory)
	locksmithCard := createCategoryButtonWithFilter(state, "locksmith", "Locksmiths", "locksmith", filterByCategory)

	// Use GridWrap with compact size for mobile
	categoriesGrid := container.NewGridWrap(
		fyne.NewSize(85, 85), // Smaller button size for mobile
		plumbingCard, electricityCard, paintingCard,
		acFixingCard, homeCleaningCard, smallRepairsCard,
		furnitureCard, waterLeakCard, applianceCard,
		locksmithCard,
	)

	// Make categories scrollable in a fixed height container
	categoriesScroll := container.NewVScroll(categoriesGrid)
	categoriesScroll.SetMinSize(fyne.NewSize(0, 250)) // Fixed height only, width adapts to screen
	categoriesScroll.OnScrolled = func(pos fyne.Position) {
		fmt.Printf("Categories scrolled to position: X=%.2f, Y=%.2f\n", pos.X, pos.Y)
	}

	return container.NewVBox(categoriesLabel, categoriesScroll)
}

// createCategoryButton creates a clickable category button with icon image
func createCategoryButton(state AppState, iconKey, name string, textsize uint8) fyne.CanvasObject {
	// Create image from resource
	iconImage := canvas.NewImageFromResource(state.GetImage(iconKey))
	iconImage.FillMode = canvas.ImageFillContain
	iconImage.SetMinSize(fyne.NewSize(32, 32))

	// Split name into words to handle multi-word categories
	words := []string{}
	currentWord := ""
	for _, char := range name {
		if char == ' ' {
			if currentWord != "" {
				words = append(words, currentWord)
				currentWord = ""
			}
		} else {
			currentWord += string(char)
		}
	}
	if currentWord != "" {
		words = append(words, currentWord)
	}

	// Create text labels based on number of words
	var textContainer *fyne.Container
	if len(words) == 1 {
		// Single word - one line
		nameLabel := canvas.NewText(name, theme.Color(theme.ColorNameForeground))
		nameLabel.Alignment = fyne.TextAlignCenter
		nameLabel.TextSize = float32(textsize)
		nameLabel.TextStyle = fyne.TextStyle{Bold: false}
		textContainer = container.NewCenter(nameLabel)
	} else {
		// Multiple words - split into two lines
		line1 := words[0]
		line2 := ""
		if len(words) > 1 {
			for i := 1; i < len(words); i++ {
				if i > 1 {
					line2 += " "
				}
				line2 += words[i]
			}
		}

		label1 := canvas.NewText(line1, theme.Color(theme.ColorNameForeground))
		label1.Alignment = fyne.TextAlignCenter
		label1.TextSize = 11
		label1.TextStyle = fyne.TextStyle{Bold: false}

		label2 := canvas.NewText(line2, theme.Color(theme.ColorNameForeground))
		label2.Alignment = fyne.TextAlignCenter
		label2.TextSize = 11
		label2.TextStyle = fyne.TextStyle{Bold: false}

		textContainer = container.NewVBox(
			container.NewCenter(label1),
			container.NewCenter(label2),
		)
	}

	// Simple VBox with icon and text
	innerContent := container.NewVBox(
		container.NewCenter(iconImage),
		textContainer,
	)

	// Center everything vertically
	content := container.NewVBox(
		layout.NewSpacer(),
		innerContent,
		layout.NewSpacer(),
	)

	// Create a button that wraps the content
	btn := widget.NewButton("", func() {
		fmt.Println("Category clicked:", name)
		fmt.Println("Filtering workers by category:", name)
		// TODO: Filter workers by selected category
	})

	// Stack the content on top of the button
	return container.NewStack(btn, content)
}

// createCategoryButtonWithFilter creates a clickable category button with icon image and filter callback
func createCategoryButtonWithFilter(state AppState, iconKey, displayName, apiCategory string, onFilter func(string)) fyne.CanvasObject {
	// Create image from resource
	iconImage := canvas.NewImageFromResource(state.GetImage(iconKey))
	iconImage.FillMode = canvas.ImageFillContain
	iconImage.SetMinSize(fyne.NewSize(32, 32))

	// Split name into words to handle multi-word categories
	words := []string{}
	currentWord := ""
	for _, char := range displayName {
		if char == ' ' {
			if currentWord != "" {
				words = append(words, currentWord)
				currentWord = ""
			}
		} else {
			currentWord += string(char)
		}
	}
	if currentWord != "" {
		words = append(words, currentWord)
	}

	// Create text labels based on number of words
	var textContainer *fyne.Container
	if len(words) == 1 {
		// Single word - one line
		nameLabel := canvas.NewText(displayName, theme.Color(theme.ColorNameForeground))
		nameLabel.Alignment = fyne.TextAlignCenter
		nameLabel.TextSize = 11
		nameLabel.TextStyle = fyne.TextStyle{Bold: false}
		textContainer = container.NewCenter(nameLabel)
	} else {
		// Multiple words - split into two lines
		line1 := words[0]
		line2 := ""
		if len(words) > 1 {
			for i := 1; i < len(words); i++ {
				if i > 1 {
					line2 += " "
				}
				line2 += words[i]
			}
		}

		label1 := canvas.NewText(line1, theme.Color(theme.ColorNameForeground))
		label1.Alignment = fyne.TextAlignCenter
		label1.TextSize = 11
		label1.TextStyle = fyne.TextStyle{Bold: false}

		label2 := canvas.NewText(line2, theme.Color(theme.ColorNameForeground))
		label2.Alignment = fyne.TextAlignCenter
		label2.TextSize = 11
		label2.TextStyle = fyne.TextStyle{Bold: false}

		textContainer = container.NewVBox(
			container.NewCenter(label1),
			container.NewCenter(label2),
		)
	}

	// Simple VBox with icon and text
	innerContent := container.NewVBox(
		container.NewCenter(iconImage),
		textContainer,
	)

	// Center everything vertically
	content := container.NewVBox(
		layout.NewSpacer(),
		innerContent,
		layout.NewSpacer(),
	)

	// Create a button that wraps the content
	btn := widget.NewButton("", func() {
		fmt.Printf("[DEBUG] [%s] Category button tapped: %s (API: %s)\n", time.Now().Format("15:04:05.000"), displayName, apiCategory)
		fmt.Println("[DEBUG] Triggering worker reload with new category...")
		// Call the filter callback with the API category value
		onFilter(apiCategory)
	})

	// Stack the content on top of the button
	return container.NewStack(btn, content)
}
