package components

import "go_rts/geometry"

// IPositionComponent is an interface which describes how to handle position
type IPositionComponent interface {
	GetPosition() *geometry.Point
	SetDestination(geometry.Point)
	MoveTowardsDestination()
}

// NewPositionComponent is a shortcut to create a PositionComponent
func NewPositionComponent(position geometry.Point, speed float64) *PositionComponent {
	return &PositionComponent{
		Position:    &position,
		Destination: &position,
		Speed:       speed,
	}
}

// PositionComponent is a struct which implements the IPositionComponent
type PositionComponent struct {
	Position    *geometry.Point
	Destination *geometry.Point
	Speed       float64
}

// GetPosition returns the current position of the PositionComponent
func (p *PositionComponent) GetPosition() *geometry.Point {
	return p.Position
}

// SetDestination sets the destination of the PositionComponent
func (p *PositionComponent) SetDestination(dest geometry.Point) {
	p.Destination = &dest
}

// MoveTowardsDestination defines how to move towards the destination
func (p *PositionComponent) MoveTowardsDestination() {
	pathToDest := p.Position.To(*p.Destination)
	if pathToDest.DistanceFrom(geometry.NewPoint(0.0, 0.0)) < p.Speed {
		p.Position = p.Destination
	} else {
		stepToDest := pathToDest.Unit().Scale(p.Speed)
		p.Position.Translate(stepToDest)
	}
}
