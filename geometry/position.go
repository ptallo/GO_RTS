package geometry

// IPositionComponent is an interface which describes how to handle position
type IPositionComponent interface {
	GetPosition() *Point
	GetRectangle() Rectangle
	SetDestination(Point, Rectangle, []IPositionComponent)
	MoveTowardsDestination([]IPositionComponent)
}

// NewPositionComponent is a shortcut to create a PositionComponent
func NewPositionComponent(rect Rectangle, speed float64) *PositionComponent {
	return &PositionComponent{
		Rectangle:    &rect,
		Destination:  rect.Point,
		Destinations: []Point{},
		Speed:        speed,
	}
}

// PositionComponent is a struct which implements the IPositionComponent
type PositionComponent struct {
	Rectangle    *Rectangle
	Destination  *Point
	Destinations []Point
	Speed        float64
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
func (p *PositionComponent) SetDestination(goalDest Point, mapRect Rectangle, collidables []IPositionComponent) {

	if !mapRect.Contains(goalDest) {
		goalDest = getInMapDestination(goalDest, mapRect)
	}

	graph := NewGraph(collidables, mapRect)
	if !graph.DoesGraphContainPoint(goalDest) {
		return
	}

	nodesToDestination := graph.PathFrom(*p.Rectangle.Point, goalDest)
	destinations := make([]Point, 0)
	for _, n := range nodesToDestination {
		if n.Contains(goalDest) {
			destinations = append(destinations, goalDest)
		} else {
			destinations = append(destinations, n.Center())
		}
	}

	if len(destinations) > 0 {
		p.Destination = &destinations[0]
		p.Destinations = destinations[1:]
	}
}

func getInMapDestination(goalDest Point, mapRect Rectangle) Point {
	if mapRect.Contains(goalDest) {
		return goalDest
	}

	var newX float64
	if goalDest.X < mapRect.Point.X {
		newX = mapRect.Point.X
	} else if goalDest.X > mapRect.Point.X+mapRect.Width {
		newX = mapRect.Point.X + mapRect.Width
	} else {
		newX = goalDest.X
	}

	var newY float64
	if goalDest.Y < mapRect.Point.Y {
		newY = mapRect.Point.Y
	} else if goalDest.Y > mapRect.Point.Y+mapRect.Height {
		newY = mapRect.Point.Y + mapRect.Height
	} else {
		newY = goalDest.Y
	}

	return NewPoint(newX, newY)
}

// MoveTowardsDestination defines how to move towards the destination
func (p *PositionComponent) MoveTowardsDestination(collidables []IPositionComponent) {
	newTranslationVector := *p.getTranslationVector()
	p.Rectangle.Point.Translate(newTranslationVector)

	if p.Rectangle.Point.Equals(*p.Destination) && len(p.Destinations) > 0 {
		p.Destination = &p.Destinations[0]
		p.Destinations = p.Destinations[1:]
	}

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
