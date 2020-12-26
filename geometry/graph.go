package geometry

import "sort"

type Graph struct {
	AdjacencyList map[Rectangle][]Rectangle
}

// NewGraph creates a graph with nodes based on the components and mapRect
func NewGraph(components []IPositionComponent, mapRect Rectangle) Graph {
	rects := generateNodes(components, mapRect)
	nonPcRects := filterCollidableRects(rects, components)
	adjacencyList := constructAdjacencyList(nonPcRects)
	return Graph{
		AdjacencyList: adjacencyList,
	}
}

func generateNodes(components []IPositionComponent, mapRect Rectangle) []Rectangle {
	xPoints, yPoints := generateXYDivisions(components, mapRect)
	rects := make([]Rectangle, 0)
	for i := range xPoints[:len(xPoints)-1] {
		for j := range yPoints[:len(yPoints)-1] {
			rect := NewRectangleFromPoints(
				NewPoint(xPoints[i], yPoints[j]),
				NewPoint(xPoints[i+1], yPoints[j+1]),
			)
			rects = append(rects, rect)
		}
	}

	return rects
}

func generateXYDivisions(components []IPositionComponent, mapRect Rectangle) ([]float64, []float64) {
	xPoints := []float64{mapRect.Point.X, mapRect.Point.X + mapRect.Width}
	yPoints := []float64{mapRect.Point.Y, mapRect.Point.Y + mapRect.Height}

	for _, c := range components {
		xPoints = append(xPoints, c.GetPosition().X)
		xPoints = append(xPoints, c.GetPosition().X+c.GetRectangle().Width)
		yPoints = append(yPoints, c.GetPosition().Y)
		yPoints = append(yPoints, c.GetPosition().Y+c.GetRectangle().Height)
	}

	return removeDuplicates(xPoints), removeDuplicates(yPoints)
}

func removeDuplicates(slice []float64) []float64 {
	sort.Float64s(slice)
	for i := len(slice) - 1; i > 0; i-- {
		if slice[i] == slice[i-1] {
			slice = append(slice[:i], slice[i+1:]...)
		}
	}
	return slice
}

func filterCollidableRects(rects []Rectangle, components []IPositionComponent) []Rectangle {
	nonPcRects := make([]Rectangle, 0)
	for _, r := range rects {
		inPcs := false
		for _, pc := range components {
			if pc.GetRectangle().Equals(r) {
				inPcs = true
			}
		}

		if !inPcs {
			nonPcRects = append(nonPcRects, r)
		}
	}
	return nonPcRects
}

func constructAdjacencyList(nodes []Rectangle) map[Rectangle][]Rectangle {
	adjacencyList := make(map[Rectangle][]Rectangle)
	for _, r := range nodes {
		adjacencyList[r] = make([]Rectangle, 0)
	}

	for i, r1 := range nodes {
		for j, r2 := range nodes {
			if i != j && r1.IsAdjacentTo(r2) {
				adjacencyList[r1] = append(adjacencyList[r1], r2)
			}
		}
	}
	return adjacencyList
}
