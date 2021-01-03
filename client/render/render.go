package render

// RenderComponent implements the IRenderComponent interface
type RenderComponent struct {
	name string
}

// NewRenderComponent creates a IRenderComponent object
func NewRenderComponent(name string) *RenderComponent {
	return &RenderComponent{
		name: name,
	}
}
