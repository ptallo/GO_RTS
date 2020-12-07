package components

import "go_rts/geometry"

type IPositionComponent interface {
	GetPosition() *geometry.Point
	SetDestination(geometry.Point)
	MoveTowardsDestination()
}

func NewPositionComponent(position geometry.Point, speed float64) *PositionComponent {
	return &PositionComponent{
		Position:    &position,
		Destination: &position,
		Speed:       speed,
	}
}

type PositionComponent struct {
	Position    *geometry.Point
	Destination *geometry.Point
	Speed       float64
}

func (p *PositionComponent) GetPosition() *geometry.Point {
	return p.Position
}

func (p *PositionComponent) SetDestination(dest geometry.Point) {
	p.Destination = &dest
}

func (p *PositionComponent) MoveTowardsDestination() {
	pathToDest := p.Position.To(*p.Destination)
	if pathToDest.DistanceFrom(geometry.NewPoint(0.0, 0.0)) < p.Speed {
		p.Position = p.Destination
	} else {
		stepToDest := pathToDest.Unit().Scale(p.Speed)
		p.Position.Translate(stepToDest)
	}
}
