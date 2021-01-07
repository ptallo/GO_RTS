package render

// RenderComponent implements the IRenderComponent interface
type RenderComponent struct {
	Name string
}

// NewRenderComponent creates a IRenderComponent object
func NewRenderComponent(Name string) RenderComponent {
	return RenderComponent{
		Name: Name,
	}
}

// Equals will return true if two render components are identical
func (r RenderComponent) Equals(r2 RenderComponent) bool {
	return r.Name == r2.Name
}
