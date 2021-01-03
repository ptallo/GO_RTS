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
		Rectangle:          &rect,
		CurrentDestination: rect.Point,
		GoalDestination:    rect.Point,
		NodesToVisit:       []Rectangle{},
		Speed:              speed,
	}
}

// PositionComponent is a struct which implements the IPositionComponent
type PositionComponent struct {
	Rectangle          *Rectangle
	CurrentDestination *Point
	GoalDestination    *Point
	NodesToVisit       []Rectangle
	Speed              float64
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
	p.GoalDestination = &goalDest

	graph := NewGraph(collidables, mapRect)
	p.NodesToVisit = graph.PathFrom(*p.Rectangle.Point, goalDest)
	p.UpdateCurrentDestination()
}

// UpdateCurrentDestination will change the current destination based on the p.NodesToVisit and p.GoalDestination
func (p *PositionComponent) UpdateCurrentDestination() {
	if len(p.NodesToVisit) >= 3 {
		currentDdestination := p.getPointFromNodes(p.NodesToVisit[0], p.NodesToVisit[1], p.NodesToVisit[2])
		p.CurrentDestination = &currentDdestination
		p.NodesToVisit = p.NodesToVisit[1:]
	} else {
		p.CurrentDestination = p.GoalDestination
	}
}

func (p *PositionComponent) getPointFromNodes(currentNode, nextNode, thirdNode Rectangle) Point {
	leftX := nextNode.Point.X
	middleX := p.GetPosition().X
	rightX := nextNode.Point.X + nextNode.Width - p.GetRectangle().Width

	topY := nextNode.Point.Y
	middleY := p.GetPosition().Y
	bottomY := nextNode.Point.Y + nextNode.Height - p.GetRectangle().Height

	if currentNode.IsTopAdjacent(nextNode) {
		if nextNode.IsTopAdjacent(thirdNode) {
			return NewPoint(middleX, topY)
		} else if nextNode.IsLeftAdjacent(thirdNode) {
			return NewPoint(rightX, topY)
		} else if nextNode.IsRightAdjacent(thirdNode) {
			return NewPoint(leftX, topY)
		}
	}

	if currentNode.IsBottomAdjacent(nextNode) {
		if nextNode.IsBottomAdjacent(thirdNode) {
			return NewPoint(middleX, bottomY)
		} else if nextNode.IsLeftAdjacent(thirdNode) {
			return NewPoint(rightX, bottomY)
		} else if nextNode.IsRightAdjacent(thirdNode) {
			return NewPoint(leftX, bottomY)
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
	if p.Rectangle.Point.Equals(*p.CurrentDestination) && !p.CurrentDestination.Equals(*p.GoalDestination) {
		p.UpdateCurrentDestination()
	}

	p.Rectangle.Point.Translate(*p.getTranslationVector())

	avgPoint := NewPoint(0.0, 0.0)
	for _, c := range p.Collisions(collidables) {
		vecAway := c.GetRectangle().Center().To(p.GetRectangle().Center()).Unit().Scale(p.Speed * 1.5)
		avgPoint.Translate(vecAway)
	}
	p.Rectangle.Point.Translate(avgPoint)
}

func (p *PositionComponent) getTranslationVector() *Point {
	var returnPoint Point
	if p.Rectangle.Point.DistanceFrom(*p.CurrentDestination) == 0.0 {
		returnPoint = NewPoint(0.0, 0.0)
	} else if p.Rectangle.Point.DistanceFrom(*p.CurrentDestination) < p.Speed {
		returnPoint = p.Rectangle.Point.To(*p.CurrentDestination)
	} else {
		returnPoint = p.Rectangle.Point.To(*p.CurrentDestination).Unit().Scale(p.Speed)
	}
	return &returnPoint
}

// Collisions returns all position components from pcs which collide with p
func (p *PositionComponent) Collisions(pcs []IPositionComponent) []IPositionComponent {
	collisions := make([]IPositionComponent, 0)
	for _, pc := range pcs {
		if pc.GetRectangle().Intersects(p.GetRectangle()) {
			collisions = append(collisions, pc)
		}
	}
	return collisions
}
