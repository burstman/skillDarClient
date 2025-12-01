# Fyne Framework - Complete Guide

**Author:** GitHub Copilot  
**Date:** December 1, 2025  
**Project:** SkillDar Client Application

---

## Table of Contents

1. [Introduction](#introduction)
2. [Core Concepts](#core-concepts)
3. [Architecture Overview](#architecture-overview)
4. [Critical Rules](#critical-rules)
5. [Best Practices](#best-practices)
6. [Common Patterns](#common-patterns)
7. [Debugging Guide](#debugging-guide)
8. [Performance Tips](#performance-tips)
9. [Your App Architecture](#your-app-architecture)
10. [Quick Reference](#quick-reference)

---

## Introduction

Fyne is a modern, cross-platform GUI toolkit for Go. It provides a declarative approach to building user interfaces with native performance and a consistent look across all platforms.

**Key Features:**

- Cross-platform (Windows, macOS, Linux, iOS, Android)
- Material Design inspired
- Built-in theming support
- Canvas-based rendering
- Data binding support
- Easy to learn and use

---

## Core Concepts

### 1. App & Window Hierarchy

```
fyne.App (Application Instance)
  │
  ├── fyne.Window (Main Window)
  │     └── fyne.CanvasObject (Content)
  │           ├── Widgets (Button, Label, Entry, etc.)
  │           ├── Containers (VBox, HBox, Border, etc.)
  │           └── Canvas Primitives (Rectangle, Circle, Image, etc.)
  │
  └── fyne.Window (Additional Windows)
        └── fyne.CanvasObject (Content)
```

**Important Points:**

- One `App` instance per application
- Multiple `Window` instances possible
- Each `Window` displays exactly ONE `CanvasObject`
- To show multiple widgets, use a `Container`

**Example:**

```go
// Create application
app := app.New()

// Create window
window := app.NewWindow("My App")

// Set content (only ONE CanvasObject)
window.SetContent(myContent)

// Show window
window.ShowAndRun()
```

---

### 2. CanvasObject - The Base Interface

Everything visible in Fyne implements the `CanvasObject` interface:

**Three Main Types:**

1. **Widgets** - Interactive UI components

   ```go
   button := widget.NewButton("Click Me", func() {
       fmt.Println("Clicked!")
   })
   ```

2. **Containers** - Layout managers

   ```go
   vbox := container.NewVBox(widget1, widget2, widget3)
   ```

3. **Canvas Primitives** - Low-level drawing elements
   ```go
   rect := canvas.NewRectangle(color.RGBA{R: 255, G: 0, B: 0, A: 255})
   text := canvas.NewText("Hello", color.White)
   ```

**Key Methods:**

- `Show()` - Make visible
- `Hide()` - Make invisible
- `Refresh()` - Redraw the object
- `Size()` - Get current size
- `Position()` - Get current position
- `MinSize()` - Get minimum required size

---

### 3. Containers (Layouts)

Containers are the foundation of Fyne layouts. They automatically manage the positioning and sizing of child widgets.

#### Common Container Types

**Vertical Box (VBox)**

```go
container.NewVBox(
    widget.NewLabel("First"),
    widget.NewLabel("Second"),
    widget.NewLabel("Third"),
)
```

**Horizontal Box (HBox)**

```go
container.NewHBox(
    widget.NewButton("Left", nil),
    widget.NewButton("Center", nil),
    widget.NewButton("Right", nil),
)
```

**Border Layout**

```go
container.NewBorder(
    topWidget,    // Top
    bottomWidget, // Bottom
    leftWidget,   // Left
    rightWidget,  // Right
    centerWidget, // Center (fills remaining space)
)
```

**Grid Layout**

```go
// Fixed columns
container.NewGridWithColumns(3,
    widget1, widget2, widget3,
    widget4, widget5, widget6,
)

// Fixed rows
container.NewGridWithRows(2,
    widget1, widget2, widget3,
    widget4, widget5, widget6,
)
```

**Other Useful Containers**

```go
container.NewCenter(widget)           // Center content
container.NewScroll(content)          // Scrollable
container.NewPadded(widget)           // Add padding
container.NewMax(widget)              // Fill available space
container.NewStack(layer1, layer2)    // Layered (z-index)
```

---

### 4. Widgets

#### Standard Widgets

**Text Display**

```go
label := widget.NewLabel("Simple text")
label.SetText("New text")

// With data binding
data := binding.NewString()
label := widget.NewLabelWithData(data)
data.Set("Updates automatically!")
```

**Buttons**

```go
// Standard button
btn := widget.NewButton("Click", func() {
    fmt.Println("Clicked")
})

// With icon
btn := widget.NewButtonWithIcon("Save", theme.DocumentSaveIcon(), func() {
    // Save action
})

// Importance levels
btn.Importance = widget.HighImportance  // Primary action
btn.Importance = widget.MediumImportance // Default
btn.Importance = widget.LowImportance   // Secondary action
```

**Text Input**

```go
entry := widget.NewEntry()
entry.SetPlaceHolder("Enter text...")
entry.OnChanged = func(text string) {
    fmt.Println("Changed:", text)
}

// Password entry
password := widget.NewPasswordEntry()

// Multiline entry
multiline := widget.NewMultiLineEntry()
```

**Forms**

```go
form := widget.NewForm(
    widget.NewFormItem("Name", widget.NewEntry()),
    widget.NewFormItem("Email", widget.NewEntry()),
    widget.NewFormItem("Password", widget.NewPasswordEntry()),
)

form.OnSubmit = func() {
    fmt.Println("Form submitted")
}

form.OnCancel = func() {
    fmt.Println("Form cancelled")
}
```

**Lists**

```go
data := []string{"Item 1", "Item 2", "Item 3"}

list := widget.NewList(
    func() int {
        return len(data)
    },
    func() fyne.CanvasObject {
        return widget.NewLabel("Template")
    },
    func(id widget.ListItemID, obj fyne.CanvasObject) {
        obj.(*widget.Label).SetText(data[id])
    },
)

list.OnSelected = func(id widget.ListItemID) {
    fmt.Println("Selected:", data[id])
}
```

**Progress Indicators**

```go
// Infinite progress
progress := widget.NewProgressBarInfinite()

// Determinate progress
progress := widget.NewProgressBar()
progress.SetValue(0.5) // 50%
```

---

### 5. Data Binding

Data binding automatically synchronizes data with UI elements.

**String Binding**

```go
// Create binding
data := binding.NewString()

// Bind to widget
label := widget.NewLabelWithData(data)
entry := widget.NewEntryWithData(data)

// Update (both widgets update automatically)
data.Set("New Value")

// Listen for changes
data.AddListener(binding.NewDataListener(func() {
    val, _ := data.Get()
    fmt.Println("Data changed:", val)
}))
```

**Other Binding Types**

```go
intData := binding.NewInt()
floatData := binding.NewFloat()
boolData := binding.NewBool()
listData := binding.NewStringList()

// Bind to widgets
check := widget.NewCheckWithData("Enabled", boolData)
slider := widget.NewSliderWithData(0, 100, floatData)
```

**Key Benefits:**

- Automatic UI updates
- Two-way synchronization
- Cleaner code
- Less manual refresh calls

---

### 6. Theming

#### Built-in Themes

```go
// Light theme
app.Settings().SetTheme(theme.LightTheme())

// Dark theme
app.Settings().SetTheme(theme.DarkTheme())

// Get current variant
variant := fyne.CurrentApp().Settings().ThemeVariant()
```

#### Custom Themes

```go
type MyCustomTheme struct {
    variant fyne.ThemeVariant
}

func NewMyCustomTheme(variant fyne.ThemeVariant) fyne.Theme {
    return &MyCustomTheme{variant: variant}
}

func (m *MyCustomTheme) Color(name fyne.ThemeColorName, variant fyne.ThemeVariant) color.Color {
    switch name {
    case theme.ColorNameBackground:
        if variant == theme.VariantDark {
            return color.RGBA{R: 20, G: 20, B: 20, A: 255}
        }
        return color.RGBA{R: 255, G: 255, B: 255, A: 255}
    case theme.ColorNamePrimary:
        return color.RGBA{R: 0, G: 122, B: 255, A: 255}
    // ... other colors
    default:
        return theme.DefaultTheme().Color(name, variant)
    }
}

func (m *MyCustomTheme) Size(name fyne.ThemeSizeName) float32 {
    switch name {
    case theme.SizeNamePadding:
        return 8
    case theme.SizeNameText:
        return 14
    default:
        return theme.DefaultTheme().Size(name)
    }
}

func (m *MyCustomTheme) Font(style fyne.TextStyle) fyne.Resource {
    return theme.DefaultTheme().Font(style)
}

func (m *MyCustomTheme) Icon(name fyne.ThemeIconName) fyne.Resource {
    return theme.DefaultTheme().Icon(name)
}

// Apply theme
app.Settings().SetTheme(NewMyCustomTheme(theme.VariantLight))
```

---

### 7. Resources (Images, Icons, Fonts)

#### Runtime Loading

```go
// Load image from file
img, err := fyne.LoadResourceFromPath("assets/logo.png")
if err != nil {
    log.Fatal(err)
}

// Use in widget
image := canvas.NewImageFromResource(img)
```

#### Bundled Resources (Recommended)

**Step 1: Add go:generate directive**

```go
//go:generate fyne bundle -o bundle.go assets
```

**Step 2: Run bundle command**

```bash
go generate
```

**Step 3: Use bundled resources**

```go
// Creates variables like resourceLogoPng, resourceIconJpg, etc.
icon := resourceLogoPng
image := canvas.NewImageFromResource(icon)
```

**Benefits:**

- No external file dependencies
- Faster loading
- Better for distribution
- Works in packaged apps

---

### 8. Navigation & Screen Management

#### Single Window, Multiple Screens Approach

```go
type AppState struct {
    app     fyne.App
    window  fyne.Window
    screens map[string]fyne.CanvasObject
    history []string
}

func (as *AppState) ShowScreen(name string) {
    if screen, exists := as.screens[name]; exists {
        as.history = append(as.history, name)
        as.window.SetContent(screen)
    }
}

func (as *AppState) GoBack() {
    if len(as.history) > 1 {
        as.history = as.history[:len(as.history)-1]
        prevScreen := as.history[len(as.history)-1]
        as.window.SetContent(as.screens[prevScreen])
    }
}

// Usage
state.screens["home"] = CreateHomeScreen(state)
state.screens["settings"] = CreateSettingsScreen(state)
state.ShowScreen("home")
```

---

## Critical Rules

### ⚠️ Rule 1: Window Has ONE Content

```go
// ❌ WRONG - Second SetContent replaces the first
window.SetContent(button1)
window.SetContent(button2) // button1 is gone!

// ✅ CORRECT - Use a container
window.SetContent(container.NewVBox(button1, button2))
```

### ⚠️ Rule 2: Always Refresh After Changes

```go
// ❌ WRONG - Change won't appear
label.SetText("New Text")

// ✅ CORRECT - Refresh to update display
label.SetText("New Text")
label.Refresh()

// ✅ OR - Use data binding (auto-refresh)
data := binding.NewString()
label := widget.NewLabelWithData(data)
data.Set("New Text") // Automatically refreshes!
```

### ⚠️ Rule 3: Don't Fight the Layout System

```go
// ❌ WRONG - Manual positioning doesn't work reliably
widget.Move(fyne.NewPos(100, 200))
widget.Resize(fyne.NewSize(300, 50))

// ✅ CORRECT - Use containers for layout
container.NewBorder(nil, nil, nil, widget, nil) // Right side
container.NewPadded(widget) // With padding
```

### ⚠️ Rule 4: Respect Minimum Sizes

```go
// ❌ WRONG - Forcing size smaller than MinSize
widget.Resize(fyne.NewSize(10, 10)) // Will be overridden

// ✅ CORRECT - Check and respect MinSize
minSize := widget.MinSize()
widget.Resize(fyne.NewSize(
    fyne.Max(minSize.Width, desiredWidth),
    fyne.Max(minSize.Height, desiredHeight),
))
```

### ⚠️ Rule 5: Containers Own Their Children

```go
// ❌ WRONG - Modifying container children directly
box := container.NewVBox()
box.Objects = append(box.Objects, newWidget) // Don't do this!

// ✅ CORRECT - Use container methods
box.Add(newWidget)
box.Remove(oldWidget)
box.Refresh()
```

---

## Best Practices

### 1. State Management

**Centralized State Pattern**

```go
type AppState struct {
    // App references
    app    fyne.App
    window fyne.Window

    // UI state
    isDarkTheme bool
    screens     map[string]fyne.CanvasObject

    // Business state
    currentUser *User
    userRole    string

    // Navigation
    screenHistory []string
}

// Pass state to screen creators
func CreateLoginScreen(state *AppState) fyne.CanvasObject {
    // Screen can access and modify state
    button := widget.NewButton("Login", func() {
        state.currentUser = &User{...}
        state.ShowScreen("home")
    })
    return container.NewVBox(button)
}
```

### 2. Screen Organization

**Pre-create vs Lazy Load**

```go
// Pre-create (faster switching, more memory)
func (as *AppState) InitializeScreens() {
    as.screens["home"] = CreateHomeScreen(as)
    as.screens["settings"] = CreateSettingsScreen(as)
    as.screens["profile"] = CreateProfileScreen(as)
}

// Lazy load (slower switching, less memory)
func (as *AppState) ShowScreen(name string) {
    if _, exists := as.screens[name]; !exists {
        as.screens[name] = as.createScreen(name)
    }
    as.window.SetContent(as.screens[name])
}
```

### 3. Resource Management

**Centralized Resource Map**

```go
type AppState struct {
    icons map[string]fyne.Resource
}

func initializeIcons() map[string]fyne.Resource {
    return map[string]fyne.Resource{
        "logo":     resourceLogoPng,
        "settings": resourceSettingsIcoPng,
        "user":     resourceUserIcoPng,
    }
}

func (as *AppState) GetIcon(name string) fyne.Resource {
    if icon, ok := as.icons[name]; ok {
        return icon
    }
    return theme.QuestionIcon() // Fallback
}
```

### 4. Event Handling

**Callback Pattern**

```go
// Define clear callbacks
type OnLoginFunc func(username, password string) error

func CreateLoginScreen(onLogin OnLoginFunc) fyne.CanvasObject {
    usernameEntry := widget.NewEntry()
    passwordEntry := widget.NewPasswordEntry()

    loginBtn := widget.NewButton("Login", func() {
        err := onLogin(usernameEntry.Text, passwordEntry.Text)
        if err != nil {
            // Show error
        }
    })

    return container.NewVBox(usernameEntry, passwordEntry, loginBtn)
}
```

### 5. Error Handling UI

```go
func ShowError(window fyne.Window, err error) {
    dialog.ShowError(err, window)
}

func ShowConfirmation(window fyne.Window, message string, callback func(bool)) {
    dialog.ShowConfirm("Confirm", message, callback, window)
}

func ShowCustomDialog(window fyne.Window, title, message string) {
    content := widget.NewLabel(message)
    dialog := dialog.NewCustom(title, "OK", content, window)
    dialog.Show()
}
```

---

## Common Patterns

### Pattern 1: Card/Panel Layout

```go
func CreateCard(title, content string) *fyne.Container {
    titleLabel := widget.NewLabelWithStyle(title,
        fyne.TextAlignLeading,
        fyne.TextStyle{Bold: true})

    contentLabel := widget.NewLabel(content)
    contentLabel.Wrapping = fyne.TextWrapWord

    card := container.NewVBox(
        titleLabel,
        widget.NewSeparator(),
        contentLabel,
    )

    return container.NewPadded(card)
}
```

### Pattern 2: Scrollable List

```go
func CreateScrollableList(items []string) *fyne.Container {
    vbox := container.NewVBox()

    for _, item := range items {
        // Capture loop variable
        itemText := item

        itemWidget := widget.NewButton(itemText, func() {
            fmt.Println("Clicked:", itemText)
        })

        vbox.Add(itemWidget)
    }

    return container.NewScroll(vbox)
}
```

### Pattern 3: Form with Validation

```go
func CreateValidatedForm() *fyne.Container {
    nameEntry := widget.NewEntry()
    emailEntry := widget.NewEntry()
    errorLabel := widget.NewLabel("")
    errorLabel.Hide()

    submitBtn := widget.NewButton("Submit", func() {
        // Validation
        if nameEntry.Text == "" {
            errorLabel.SetText("Name is required")
            errorLabel.Show()
            return
        }

        if !isValidEmail(emailEntry.Text) {
            errorLabel.SetText("Invalid email")
            errorLabel.Show()
            return
        }

        errorLabel.Hide()
        // Process form...
    })

    return container.NewVBox(
        widget.NewLabel("Name:"),
        nameEntry,
        widget.NewLabel("Email:"),
        emailEntry,
        errorLabel,
        submitBtn,
    )
}
```

### Pattern 4: Loading State

```go
type Screen struct {
    content   *fyne.Container
    loading   *widget.ProgressBarInfinite
    mainView  *fyne.Container
}

func (s *Screen) ShowLoading() {
    s.mainView.Hide()
    s.loading.Show()
    s.content.Refresh()
}

func (s *Screen) HideLoading() {
    s.loading.Hide()
    s.mainView.Show()
    s.content.Refresh()
}

func (s *Screen) LoadData() {
    s.ShowLoading()

    go func() {
        // Fetch data...
        time.Sleep(2 * time.Second)

        // Update UI on main thread
        s.HideLoading()
    }()
}
```

### Pattern 5: Modal Dialog

```go
func ShowCustomModal(window fyne.Window, title string, content fyne.CanvasObject, onConfirm func()) {
    confirmBtn := widget.NewButton("Confirm", func() {
        onConfirm()
    })

    cancelBtn := widget.NewButton("Cancel", func() {
        // Dialog closes automatically
    })

    buttons := container.NewHBox(
        layout.NewSpacer(),
        cancelBtn,
        confirmBtn,
    )

    dialogContent := container.NewVBox(
        content,
        widget.NewSeparator(),
        buttons,
    )

    dialog := dialog.NewCustom(title, "", dialogContent, window)
    dialog.Show()
}
```

---

## Debugging Guide

### Issue 1: Widget Not Showing

**Checklist:**

1. ✅ Is widget added to a container?
2. ✅ Is container set as window content?
3. ✅ Did you call `.Show()` (if you previously hid it)?
4. ✅ Is widget visible (not hidden by another widget)?
5. ✅ Check MinSize - might be 0x0

**Debug:**

```go
fmt.Println("Widget size:", widget.Size())
fmt.Println("Widget min size:", widget.MinSize())
fmt.Println("Widget visible:", widget.Visible())
```

### Issue 2: Layout Incorrect

**Common Causes:**

- Wrong container type
- Missing spacers
- Size constraints

**Solutions:**

```go
// Add flexible space
container.NewHBox(
    widget1,
    layout.NewSpacer(), // Pushes widgets apart
    widget2,
)

// Try different containers
container.NewBorder(top, bottom, left, right, center)
container.NewPadded(widget) // Add breathing room
container.NewMax(widget)    // Fill available space
```

### Issue 3: Changes Not Appearing

**Fix:**

```go
// Option 1: Refresh specific widget
widget.SetText("New")
widget.Refresh()

// Option 2: Refresh container
container.Refresh()

// Option 3: Refresh entire window
window.Content().Refresh()

// Option 4: Use data binding (auto-refresh)
data := binding.NewString()
widget := widget.NewLabelWithData(data)
data.Set("New") // Automatically refreshes
```

### Issue 4: Memory Leaks

**Common Causes:**

- Not removing listeners
- Circular references
- Goroutines not stopped

**Prevention:**

```go
// Remove listeners when done
listener := binding.NewDataListener(func() {...})
data.AddListener(listener)
defer data.RemoveListener(listener)

// Stop goroutines
ctx, cancel := context.WithCancel(context.Background())
defer cancel()

go func() {
    for {
        select {
        case <-ctx.Done():
            return
        default:
            // Work...
        }
    }
}()
```

### Issue 5: Performance Problems

**Symptoms:**

- Slow rendering
- High CPU usage
- Laggy interactions

**Solutions:**

```go
// 1. Use List instead of VBox for many items
list := widget.NewList(
    func() int { return len(data) },
    func() fyne.CanvasObject { return widget.NewLabel("") },
    func(id widget.ListItemID, obj fyne.CanvasObject) {
        obj.(*widget.Label).SetText(data[id])
    },
)

// 2. Lazy load images
var image *canvas.Image
go func() {
    resource := loadImage()
    image = canvas.NewImageFromResource(resource)
}()

// 3. Debounce updates
var timer *time.Timer
entry.OnChanged = func(text string) {
    if timer != nil {
        timer.Stop()
    }
    timer = time.AfterFunc(300*time.Millisecond, func() {
        // Process text
    })
}
```

---

## Performance Tips

### 1. Widget Reuse

```go
// ❌ BAD - Creates new widgets on every update
func updateList(data []string) {
    container.Objects = nil
    for _, item := range data {
        container.Add(widget.NewLabel(item))
    }
}

// ✅ GOOD - Reuse existing widgets
func updateList(data []string) {
    for i, item := range data {
        if i < len(container.Objects) {
            container.Objects[i].(*widget.Label).SetText(item)
        } else {
            container.Add(widget.NewLabel(item))
        }
    }
}
```

### 2. Batch Updates

```go
// ❌ BAD - Multiple refreshes
for i := 0; i < 100; i++ {
    label.SetText(fmt.Sprintf("Item %d", i))
    label.Refresh()
}

// ✅ GOOD - Single refresh
for i := 0; i < 100; i++ {
    label.SetText(fmt.Sprintf("Item %d", i))
}
label.Refresh()
```

### 3. Use Data Binding for Lists

```go
// ❌ BAD - Manual list management
func updateItems(items []string) {
    container.RemoveAll()
    for _, item := range items {
        container.Add(widget.NewLabel(item))
    }
    container.Refresh()
}

// ✅ GOOD - Data binding
data := binding.NewStringList()
list := widget.NewListWithData(data,
    func() fyne.CanvasObject {
        return widget.NewLabel("")
    },
    func(item binding.DataItem, obj fyne.CanvasObject) {
        obj.(*widget.Label).Bind(item.(binding.String))
    },
)

// Update automatically refreshes
data.Set(newItems)
```

### 4. Lazy Loading

```go
// Pre-create only what's needed
func (as *AppState) ShowScreen(name string) {
    if _, exists := as.screens[name]; !exists {
        as.screens[name] = as.createScreen(name)
    }
    as.window.SetContent(as.screens[name])
}
```

### 5. Image Optimization

```go
// Load images asynchronously
func loadImageAsync(path string) *canvas.Image {
    placeholder := canvas.NewImageFromResource(theme.FileImageIcon())

    go func() {
        resource, err := fyne.LoadResourceFromPath(path)
        if err == nil {
            placeholder.Resource = resource
            placeholder.Refresh()
        }
    }()

    return placeholder
}
```

---

## Your App Architecture

### Current Structure Analysis

```go
type AppState struct {
    app           fyne.App
    window        fyne.Window
    isDarkTheme   bool
    screens       map[string]fyne.CanvasObject  // ✅ Good: Pre-created screens
    icons         map[string]fyne.Resource      // ✅ Good: Centralized resources
    userRole      string
    screenHistory []string                       // ✅ Good: Navigation history
    currentWorker *uiscreen.WorkerProfile
}
```

### Strengths

1. ✅ **Centralized State Management**

   - Single source of truth
   - Easy to pass to screens
   - Clean navigation

2. ✅ **Pre-created Screens**

   - Fast screen switching
   - No recreation overhead
   - Consistent state

3. ✅ **Resource Management**

   - Bundled resources
   - Centralized icon map
   - Helper methods (GetImage)

4. ✅ **Navigation System**

   - History tracking
   - Back button support
   - Flexible screen switching

5. ✅ **Theme Management**
   - Custom theme support
   - Toggle functionality
   - Persistent across screens

### Recommendations

**1. Add Context for Cancellation**

```go
type AppState struct {
    // ... existing fields
    ctx    context.Context
    cancel context.CancelFunc
}

func NewAppState() *AppState {
    ctx, cancel := context.WithCancel(context.Background())
    return &AppState{
        ctx:    ctx,
        cancel: cancel,
        // ... other fields
    }
}
```

**2. Add Loading States**

```go
type AppState struct {
    // ... existing fields
    isLoading bool
    loadingMessage string
}

func (as *AppState) ShowLoading(message string) {
    as.isLoading = true
    as.loadingMessage = message
    // Show loading overlay
}

func (as *AppState) HideLoading() {
    as.isLoading = false
    // Hide loading overlay
}
```

**3. Add Error Handling**

```go
func (as *AppState) ShowError(err error) {
    dialog.ShowError(err, as.window)
}

func (as *AppState) ShowInfo(title, message string) {
    dialog.ShowInformation(title, message, as.window)
}
```

**4. Screen Lifecycle Hooks**

```go
type Screen interface {
    OnShow()  // Called when screen becomes visible
    OnHide()  // Called when screen is hidden
    OnDestroy() // Called when screen is removed
}

func (as *AppState) ShowScreen(name string) {
    // Hide current screen
    if current, ok := as.getCurrentScreen().(Screen); ok {
        current.OnHide()
    }

    // Show new screen
    if screen, exists := as.screens[name]; exists {
        as.window.SetContent(screen)
        if s, ok := screen.(Screen); ok {
            s.OnShow()
        }
    }
}
```

---

## Quick Reference

### Common Widgets

```go
widget.NewLabel("Text")
widget.NewButton("Click", callback)
widget.NewEntry()
widget.NewPasswordEntry()
widget.NewCheck("Checkbox", callback)
widget.NewRadioGroup(options, callback)
widget.NewSelect(options, callback)
widget.NewSlider(min, max)
widget.NewProgressBar()
widget.NewProgressBarInfinite()
```

### Common Containers

```go
container.NewVBox(items...)
container.NewHBox(items...)
container.NewBorder(top, bottom, left, right, center)
container.NewGrid(columns, items...)
container.NewCenter(item)
container.NewScroll(item)
container.NewPadded(item)
container.NewMax(item)
```

### Common Dialogs

```go
dialog.ShowInformation(title, message, window)
dialog.ShowError(err, window)
dialog.ShowConfirm(title, message, callback, window)
dialog.ShowCustom(title, dismiss, content, window)
dialog.ShowFileOpen(callback, window)
dialog.ShowFileSave(callback, window)
```

### Common Icons

```go
theme.HomeIcon()
theme.SettingsIcon()
theme.SearchIcon()
theme.MenuIcon()
theme.ContentAddIcon()
theme.ContentRemoveIcon()
theme.ContentCutIcon()
theme.ContentCopyIcon()
theme.ContentPasteIcon()
theme.DeleteIcon()
theme.MailComposeIcon()
theme.DocumentIcon()
theme.DocumentSaveIcon()
theme.FolderIcon()
theme.FolderOpenIcon()
```

### Common Colors

```go
theme.ColorNameBackground
theme.ColorNameForeground
theme.ColorNamePrimary
theme.ColorNameHover
theme.ColorNameFocus
theme.ColorNamePressed
theme.ColorNameScrollBar
theme.ColorNameShadow
```

### Common Sizes

```go
theme.SizeNamePadding
theme.SizeNameScrollBar
theme.SizeNameSeparator
theme.SizeNameText
theme.SizeNameHeadingText
theme.SizeNameSubHeadingText
theme.SizeNameCaptionText
theme.SizeNameInlineIcon
```

### Data Binding

```go
binding.NewBool()
binding.NewFloat()
binding.NewInt()
binding.NewString()
binding.NewStringList()
binding.NewBoolList()
binding.NewFloatList()
binding.NewIntList()
```

### Keyboard Shortcuts

```go
desktop.CustomShortcut{KeyName: fyne.KeyN, Modifier: fyne.KeyModifierControl}
desktop.CustomShortcut{KeyName: fyne.KeyS, Modifier: fyne.KeyModifierControl}
desktop.CustomShortcut{KeyName: fyne.KeyQ, Modifier: fyne.KeyModifierControl}
```

---

## Appendix: Conversion to PDF

To convert this markdown to PDF, use one of these tools:

**1. Using Pandoc (Recommended)**

```bash
pandoc fyne-framework-guide.md -o fyne-framework-guide.pdf \
    --pdf-engine=xelatex \
    --toc \
    --number-sections
```

**2. Using VS Code Extension**

- Install "Markdown PDF" extension
- Open this file
- Right-click → "Markdown PDF: Export (pdf)"

**3. Using Online Tools**

- https://www.markdowntopdf.com/
- https://www.markdown2pdf.com/

**4. Using Python (requires markdown and weasyprint)**

```bash
pip install markdown weasyprint
python -c "
import markdown
from weasyprint import HTML

with open('fyne-framework-guide.md') as f:
    html = markdown.markdown(f.read())
    HTML(string=html).write_pdf('fyne-framework-guide.pdf')
"
```

---

## Summary

**Remember These Key Points:**

1. **One Window, One Content** - Use containers to combine multiple widgets
2. **Always Refresh** - After changes, call `.Refresh()` or use data binding
3. **Containers Handle Layout** - Don't fight the layout system
4. **Pre-create or Lazy Load** - Choose based on your needs
5. **Centralize State** - Single source of truth
6. **Use Data Binding** - Automatic UI updates
7. **Theme Consistently** - Custom themes for branding
8. **Bundle Resources** - Better performance and distribution

**Your architecture is solid!** Keep building on these foundations.

---

**End of Guide**
