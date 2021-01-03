package geometry

import "sort"

// Graph is a structure containing an adjacency list
type Graph struct {
	AdjacencyList map[Rectangle][]Rectangle
}

// NewGraph creates a graph with nodes based on the components and mapRect
func NewGraph(components []IPositionComponent, mapRect Rectangle) Graph {
	xPoints, yPoints := generateXYDivisions(components, mapRect)
	rects := makeRectsFromXYDivisions(xPoints, yPoints)
	nonPcRects := filterCollidableRects(rects, components)
	adjacencyList := constructAdjacencyList(nonPcRects)
	return Graph{
		AdjacencyList: adjacencyList,
	}
}

func generateXYDivisions(components []IPositionComponent, mapRect Rectangle) ([]float64, []float64) {
	xPoints := []float64{mapRect.Point.X, mapRect.Point.X + mapRect.Width}
	yPoints := []float64{mapRect.Point.Y, mapRect.Point.Y + mapRect.Height}

	for _, c := range components {
		xPoints = append(xPoints, c.GetRectangle().Point.X)
		xPoints = append(xPoints, c.GetRectangle().Point.X+c.GetRectangle().Width)
		yPoints = append(yPoints, c.GetRectangle().Point.Y)
		yPoints = append(yPoints, c.GetRectangle().Point.Y+c.GetRectangle().Height)
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

func makeRectsFromXYDivisions(xPoints, yPoints []float64) []Rectangle {
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

// PathFrom finds a set of points which will get the user from start to the destination
func (g Graph) PathFrom(start, dest Point) []Rectangle {
	nodes := make([]Rectangle, 0)
	for k := range g.AdjacencyList {
		nodes = append(nodes, k)
	}

	if !g.Contains(dest) {
		return []Rectangle{}
	}

	startNode := getNodeContainingPoint(start, nodes)
	endNode := getNodeContainingPoint(dest, nodes)
	return g.aStarSearch(startNode, endNode)
}

func (g Graph) aStarSearch(startNode, endNode Rectangle) []Rectangle {
	queue := []Rectangle{startNode}
	visited := make([]Rectangle, 0)
	parent := make(map[Rectangle]Rectangle)

	for !queue[0].Equals(endNode) {
		currentNode := queue[0]
		visited = append(visited, currentNode)
		nodesToAdd := g.AdjacencyList[currentNode]
		for _, n := range nodesToAdd {
			if !isItemInList(n, visited) {
				queue = append(queue, n)
				parent[n] = currentNode
			}
		}
		queue = queue[1:]
	}

	path := []Rectangle{endNode}
	for !path[0].Equals(startNode) {
		parentNode := parent[path[0]]
		path = append([]Rectangle{parentNode}, path...)
	}
	return path
}

func isItemInList(rect Rectangle, rects []Rectangle) bool {
	for _, r := range rects {
		if r.Equals(rect) {
			return true
		}
	}
	return false
}

func getNodeContainingPoint(point Point, nodes []Rectangle) Rectangle {
	for _, n := range nodes {
		if n.Contains(point) {
			return n
		}
	}
	return Rectangle{}
}

// Contains returns true if the point is in the graph else false
func (g Graph) Contains(p Point) bool {
	for k := range g.AdjacencyList {
		if k.Contains(p) {
			return true
		}
	}
	return false
}
