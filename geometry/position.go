package geometry

// IPositionComponent is an interface which describes how to handle position
type IPositionComponent interface {
	GetPosition() *Point
	GetRectangle() Rectangle
	SetDestination(Point, []IPositionComponent)
	MoveTowardsDestination([]IPositionComponent)
}

// NewPositionComponent is a shortcut to create a PositionComponent
func NewPositionComponent(rect Rectangle, speed float64) *PositionComponent {
	return &PositionComponent{
		Rectangle:   &rect,
		Destination: rect.Point,
		Speed:       speed,
	}
}

// PositionComponent is a struct which implements the IPositionComponent
type PositionComponent struct {
	Rectangle   *Rectangle
	Destination *Point
	Speed       float64
}

// GetPosition returns the current position of the PositionComponent
func (p *PositionComponent) GetPosition() *Point {
	return p.Rectangle.Point
}

// GetRectangle returns the rectangle describing this position component
func (p *PositionComponent) GetRectangle() Rectangle {
	return *p.Rectangle
}

// SetDestination sets the destination of the PositionComponent
func (p *PositionComponent) SetDestination(dest Point, tiles []IPositionComponent) {
	if isInTiles(dest, tiles) {
		p.Destination = &dest
	} else {
		newDest := getDestinationInTiles(dest, tiles)
		p.Destination = &newDest
	}
}

func getDestinationInTiles(goalDest Point, tiles []IPositionComponent) Point {
	minTileDistance := 99999999.0
	var minTilePoint Point
	for _, t := range tiles {
		for _, p := range t.GetRectangle().GetCorners() {
			dist := p.DistanceFrom(goalDest)
			if dist < minTileDistance {
				minTileDistance = dist
				minTilePoint = p
			}
		}
	}
	return minTilePoint
}

func isInTiles(p Point, tiles []IPositionComponent) bool {
	inTiles := false
	for _, t := range tiles {
		if t.GetRectangle().Contains(p) {
			inTiles = true
		}
	}
	return inTiles
}

// MoveTowardsDestination defines how to move towards the destination
func (p *PositionComponent) MoveTowardsDestination(collidables []IPositionComponent) {
	newTranslationVector := *p.getTranslationVector()
	p.Rectangle.Point.Translate(newTranslationVector)

	if willCollide(p, collidables) {
		p.Rectangle.Point.Translate(newTranslationVector.Inverse())
	}
}

func (p *PositionComponent) getTranslationVector() *Point {
	var returnPoint Point
	if p.Rectangle.Point.DistanceFrom(*p.Destination) == 0.0 {
		returnPoint = NewPoint(0.0, 0.0)
	} else if p.Rectangle.Point.DistanceFrom(*p.Destination) < p.Speed {
		returnPoint = p.Rectangle.Point.To(*p.Destination)
	} else {
		returnPoint = p.Rectangle.Point.To(*p.Destination).Unit().Scale(p.Speed)
	}
	return &returnPoint
}

func willCollide(pc IPositionComponent, others []IPositionComponent) bool {
	for _, o := range others {
		if o.GetRectangle().Intersects(pc.GetRectangle()) {
			return true
		}
	}
	return false
}
