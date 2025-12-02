package theme

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/widget"
)

// FixedHeightContainer is a container that enforces a fixed height
type FixedHeightContainer struct {
	widget.BaseWidget
	content fyne.CanvasObject
	height  float32
}

// NewFixedHeightContainer creates a container with a fixed height
func NewFixedHeightContainer(height float32, content fyne.CanvasObject) *FixedHeightContainer {
	container := &FixedHeightContainer{
		content: content,
		height:  height,
	}
	container.ExtendBaseWidget(container)
	return container
}

func (f *FixedHeightContainer) CreateRenderer() fyne.WidgetRenderer {
	return &fixedHeightRenderer{
		container: f,
		content:   f.content,
	}
}

type fixedHeightRenderer struct {
	container *FixedHeightContainer
	content   fyne.CanvasObject
}

func (r *fixedHeightRenderer) Layout(size fyne.Size) {
	// Force the height to our fixed value
	r.content.Resize(fyne.NewSize(size.Width, r.container.height))
	r.content.Move(fyne.NewPos(0, 0))
}

func (r *fixedHeightRenderer) MinSize() fyne.Size {
	// Return our fixed height, width can be flexible
	return fyne.NewSize(r.content.MinSize().Width, r.container.height)
}

func (r *fixedHeightRenderer) Refresh() {
	r.content.Refresh()
}

func (r *fixedHeightRenderer) Objects() []fyne.CanvasObject {
	return []fyne.CanvasObject{r.content}
}

func (r *fixedHeightRenderer) Destroy() {}
