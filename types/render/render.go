package render

// RenderComponent implements the IRenderComponent interface
type RenderComponent struct {
	Name string
}

// NewRenderComponent creates a IRenderComponent object
func NewRenderComponent(Name string) *RenderComponent {
	return &RenderComponent{
		Name: Name,
	}
}
