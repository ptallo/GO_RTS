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

	nodes := graph.PathFrom(*p.Rectangle.Point, goalDest)
	destinations := p.getDestinationPoints(nodes, goalDest)

	if len(destinations) > 0 {
		p.Destination = &destinations[0]
		p.Destinations = destinations[1:]
	}
}

func (p *PositionComponent) getDestinationPoints(nodes []Rectangle, goalDest Point) []Point {
	destinations := make([]Point, 0)
	for i := range nodes {
		if nodes[i].Contains(goalDest) {
			destinations = append(destinations, goalDest)
		} else if i > 0 {
			destinations = append(destinations, p.getPointFromNodes(nodes[i-1], nodes[i], nodes[i+1]))
		}
	}
	return destinations
}

func (p *PositionComponent) getPointFromNodes(currentNode, nextNode, thirdNode Rectangle) Point {
	leftX := nextNode.Point.X
	middleX := nextNode.Point.X + (nextNode.Width+p.GetRectangle().Width)/2
	rightX := nextNode.Point.X + nextNode.Width - p.GetRectangle().Width

	topY := nextNode.Point.Y
	middleY := nextNode.Point.Y + (nextNode.Height+p.GetRectangle().Height)/2
	bottomY := nextNode.Point.Y + nextNode.Height - p.GetRectangle().Height

	if currentNode.IsTopAdjacent(nextNode) {
		if nextNode.IsTopAdjacent(thirdNode) {
			return NewPoint(middleX, topY)
		} else if nextNode.IsLeftAdjacent(thirdNode) {
			return NewPoint(leftX, topY)
		} else if nextNode.IsRightAdjacent(thirdNode) {
			return NewPoint(rightX, topY)
		}
	}

	if currentNode.IsBottomAdjacent(nextNode) {
		if nextNode.IsBottomAdjacent(thirdNode) {
			return NewPoint(middleX, bottomY)
		} else if nextNode.IsLeftAdjacent(thirdNode) {
			return NewPoint(leftX, bottomY)
		} else if nextNode.IsRightAdjacent(thirdNode) {
			return NewPoint(rightX, bottomY)
		}
	}

	if currentNode.IsLeftAdjacent(nextNode) {
		if nextNode.IsLeftAdjacent(thirdNode) {
			return NewPoint(leftX, middleY)
		} else if nextNode.IsBottomAdjacent(thirdNode) {
			return NewPoint(leftX, topY)
		} else if nextNode.IsTopAdjacent(thirdNode) {
			return NewPoint(leftX, bottomY)
		}
	}

	if currentNode.IsRightAdjacent(nextNode) {
		if nextNode.IsRightAdjacent(thirdNode) {
			return NewPoint(rightX, middleY)
		} else if nextNode.IsBottomAdjacent(thirdNode) {
			return NewPoint(rightX, topY)
		} else if nextNode.IsTopAdjacent(thirdNode) {
			return NewPoint(rightX, bottomY)
		}
	}

	return Point{}
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

	if willCollide(p, collidables) {
		p.Rectangle.Point.Translate(newTranslationVector.Inverse())
	}

	if p.Rectangle.Point.Equals(*p.Destination) && len(p.Destinations) > 0 {
		p.Destination = &p.Destinations[0]
		p.Destinations = p.Destinations[1:]
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
