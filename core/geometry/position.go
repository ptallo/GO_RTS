package geometry

// NewPositionComponent is a shortcut to create a PositionComponent
func NewPositionComponent(rect Rectangle, speed float64) *PositionComponent {
	return &PositionComponent{
		Rectangle:          rect,
		CurrentDestination: rect.Point,
		GoalDestination:    rect.Point,
		NodesToVisit:       []Rectangle{},
		Speed:              speed,
	}
}

// PositionComponent is a struct which implements the PositionComponent
type PositionComponent struct {
	Rectangle          Rectangle
	CurrentDestination Point
	GoalDestination    Point
	NodesToVisit       []Rectangle
	Speed              float64
}

// Equals checks if two position components are equal
func (p PositionComponent) Equals(p2 PositionComponent) bool {
	return p.Rectangle.Equals(p2.Rectangle) && p.Speed == p2.Speed
}

// SetDestination sets the destination of the PositionComponent
func (p *PositionComponent) SetDestination(goalDest Point, mapRect Rectangle, collidables []PositionComponent) {
	if !mapRect.Contains(goalDest) {
		goalDest = getInMapDestination(goalDest, mapRect)
	}
	p.GoalDestination = goalDest

	graph := NewGraph(collidables, mapRect)
	p.NodesToVisit = graph.PathFrom(p.Rectangle.Point, goalDest)
	p.updateCurrentDestination()
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

func (p *PositionComponent) updateCurrentDestination() {
	if len(p.NodesToVisit) >= 3 {
		currentDestination := p.getPointFromNodes(p.NodesToVisit[0], p.NodesToVisit[1], p.NodesToVisit[2])
		p.CurrentDestination = currentDestination
		p.NodesToVisit = p.NodesToVisit[1:]
	} else {
		p.CurrentDestination = p.GoalDestination
	}
}

func (p *PositionComponent) getPointFromNodes(currentNode, nextNode, thirdNode Rectangle) Point {
	leftX := nextNode.Point.X
	middleX := p.Rectangle.Point.X
	rightX := nextNode.Point.X + nextNode.Width - p.Rectangle.Width

	topY := nextNode.Point.Y
	middleY := p.Rectangle.Point.Y
	bottomY := nextNode.Point.Y + nextNode.Height - p.Rectangle.Height

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

// MoveTowardsDestination defines how to move towards the destination
func (p *PositionComponent) MoveTowardsDestination(collidables []PositionComponent) {
	if p.Rectangle.Point.Equals(p.CurrentDestination) && !p.CurrentDestination.Equals(p.GoalDestination) {
		p.updateCurrentDestination()
	}

	p.Rectangle = p.Rectangle.Move(*p.getTranslationVector())
	avgPoint := NewPoint(0.0, 0.0)
	for _, c := range p.getCollisions(collidables) {
		vecAway := c.Rectangle.Center().To(p.Rectangle.Center()).Unit().Scale(p.Speed * 1.5)
		avgPoint = avgPoint.Move(vecAway)
	}
	p.Rectangle = p.Rectangle.Move(avgPoint)
}

func (p *PositionComponent) getTranslationVector() *Point {
	var returnPoint Point
	if p.Rectangle.Point.DistanceFrom(p.CurrentDestination) == 0.0 {
		returnPoint = NewPoint(0.0, 0.0)
	} else if p.Rectangle.Point.DistanceFrom(p.CurrentDestination) < p.Speed {
		returnPoint = p.Rectangle.Point.To(p.CurrentDestination)
	} else {
		returnPoint = p.Rectangle.Point.To(p.CurrentDestination).Unit().Scale(p.Speed)
	}
	return &returnPoint
}

func (p *PositionComponent) getCollisions(pcs []PositionComponent) []PositionComponent {
	collisions := make([]PositionComponent, 0)
	for _, pc := range pcs {
		if pc.Rectangle.Intersects(p.Rectangle) {
			collisions = append(collisions, pc)
		}
	}
	return collisions
}
