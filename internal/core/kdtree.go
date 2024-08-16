package core

import (
	"os"
	"sort"
)

type KDTree struct {
	Root *Node
}

type Node struct {
	Point *Point
	Left  *Node
	Right *Node
	Axis  int
}

type Point struct {
	ID        string
	Latitude  float64
	Longitude float64
	Distance  float64
}

func NewKDTree() *KDTree {
	return &KDTree{}
}

func (tree *KDTree) ListPoints() []*Point {
	var points []*Point
	listPoints(tree.Root, &points)
	return points
}

func listPoints(node *Node, points *[]*Point) {
	if node == nil {
		return
	}
	*points = append(*points, node.Point)
	listPoints(node.Left, points)
	listPoints(node.Right, points)
}

func (tree *KDTree) AddPoint(point *Point) {
	tree.Root = addPoint(tree.Root, point, 0)
}

func addPoint(node *Node, point *Point, depth int) *Node {
	if node == nil {
		return &Node{Point: point, Axis: depth % 2}
	}

	if (depth%2 == 0 && point.Longitude < node.Point.Longitude) || (depth%2 == 1 && point.Latitude < node.Point.Latitude) {
		node.Left = addPoint(node.Left, point, depth+1)
	} else {
		node.Right = addPoint(node.Right, point, depth+1)
	}

	return node
}

func (tree *KDTree) UpdatePointByID(id string, newPoint *Point) bool {
	updated := false
	tree.Root = updatePointByID(tree.Root, id, newPoint, 0, &updated)
	return updated
}

func updatePointByID(node *Node, id string, newPoint *Point, depth int, updated *bool) *Node {
	if node == nil {
		return nil
	}

	if node.Point.ID == id {
		*updated = true
		points := collectPoints(node)
		points[0] = newPoint // Replace the point with the new one
		return buildKDTree(points, depth)
	}

	if (depth%2 == 0 && newPoint.Longitude < node.Point.Longitude) || (depth%2 == 1 && newPoint.Latitude < node.Point.Latitude) {
		node.Left = updatePointByID(node.Left, id, newPoint, depth+1, updated)
	} else {
		node.Right = updatePointByID(node.Right, id, newPoint, depth+1, updated)
	}

	return node
}

func (tree *KDTree) KNearestNeighbors(center *Point, k int) []*Point {
	var result []*Point
	knnSearch(tree.Root, center, k, 0, &result)

	var distanceResult []float64
	if os.Getenv("ENGINE") == "OSRM" {
		distanceResult, _, _ = OSRM(result, center.Latitude, center.Longitude)
	}

	for i := range result {
		switch os.Getenv("ENGINE") {
		case "OSRM":
			result[i].Distance = distanceResult[i+1]
		case "HAVERSINE":
		default:
			result[i].Distance = Haversine(center.Latitude, center.Longitude, result[i].Latitude, result[i].Longitude)
		}
	}
	sort.Slice(result, func(i, j int) bool {
		return result[i].Distance < result[j].Distance
	})

	if len(result) > k {
		return result[:k]
	}
	return result
}

func knnSearch(node *Node, center *Point, k int, depth int, result *[]*Point) {
	if node == nil {
		return
	}
	*result = append(*result, node.Point)
	if (depth%2 == 0 && center.Longitude < node.Point.Longitude) || (depth%2 == 1 && center.Latitude < node.Point.Latitude) {
		knnSearch(node.Left, center, k, depth+1, result)
		knnSearch(node.Right, center, k, depth+1, result)
	} else {
		knnSearch(node.Right, center, k, depth+1, result)
		knnSearch(node.Left, center, k, depth+1, result)
	}
}

func collectPoints(node *Node) []*Point {
	if node == nil {
		return nil
	}
	points := []*Point{node.Point}
	points = append(points, collectPoints(node.Left)...)
	points = append(points, collectPoints(node.Right)...)
	return points
}

func buildKDTree(points []*Point, depth int) *Node {
	if len(points) == 0 {
		return nil
	}

	axis := depth % 2
	sort.Slice(points, func(i, j int) bool {
		if axis == 0 {
			return points[i].Longitude < points[j].Longitude
		}
		return points[i].Latitude < points[j].Latitude
	})

	median := len(points) / 2
	return &Node{
		Point: points[median],
		Left:  buildKDTree(points[:median], depth+1),
		Right: buildKDTree(points[median+1:], depth+1),
		Axis:  axis,
	}
}
