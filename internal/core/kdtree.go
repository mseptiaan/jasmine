package core

import (
	"github.com/mseptiaan/jasmine/internal/pb"
	"os"
	"sort"
	"sync"
	"time"
)

type Node struct {
	Point *pb.RidersNearby
	Left  *Node
	Right *Node
	Axis  int
}

type KDTree struct {
	Name       string
	Root       *Node
	LastUpdate time.Time
	mu         sync.RWMutex
}

func NewKDTree(name string) *KDTree {
	return &KDTree{Name: name}
}

func (tree *KDTree) Insert(rider *pb.RidersNearby) {
	tree.mu.Lock()
	defer tree.mu.Unlock()
	tree.Root = insert(tree.Root, rider, 0)
}

func insert(node *Node, rider *pb.RidersNearby, depth int) *Node {
	if node == nil {
		return &Node{Point: rider, Axis: depth % 2}
	}

	if depth%2 == 0 {
		if rider.Latitude < node.Point.Latitude {
			node.Left = insert(node.Left, rider, depth+1)
		} else {
			node.Right = insert(node.Right, rider, depth+1)
		}
	} else {
		if rider.Longitude < node.Point.Longitude {
			node.Left = insert(node.Left, rider, depth+1)
		} else {
			node.Right = insert(node.Right, rider, depth+1)
		}
	}
	return node
}

func (tree *KDTree) Update(rider *pb.RidersNearby) {
	tree.mu.Lock()
	defer tree.mu.Unlock()
	tree.Root = update(tree.Root, rider, 0)
}

func update(node *Node, rider *pb.RidersNearby, depth int) *Node {
	if node == nil {
		return nil
	}

	if node.Point.RiderId == rider.RiderId {
		node.Point = rider
		return node
	}

	if depth%2 == 0 {
		if rider.Latitude < node.Point.Latitude {
			node.Left = update(node.Left, rider, depth+1)
		} else {
			node.Right = update(node.Right, rider, depth+1)
		}
	} else {
		if rider.Longitude < node.Point.Longitude {
			node.Left = update(node.Left, rider, depth+1)
		} else {
			node.Right = update(node.Right, rider, depth+1)
		}
	}
	return node
}

func (tree *KDTree) SearchNearby(lat, long, radius float64) []*pb.RidersNearby {
	tree.mu.RLock()
	defer tree.mu.RUnlock()
	var results []*pb.RidersNearby
	searchNearby(tree.Root, lat, long, radius, 0, &results)
	return results
}

func searchNearby(node *Node, lat, long, radius float64, depth int, results *[]*pb.RidersNearby) {
	if node == nil {
		return
	}

	distance := haversine(lat, long, node.Point.Latitude, node.Point.Longitude)
	if distance <= radius {
		node.Point.Distance = distance // Set the distance
		*results = append(*results, node.Point)
	}

	if depth%2 == 0 {
		if lat-radius < node.Point.Latitude {
			searchNearby(node.Left, lat, long, radius, depth+1, results)
		}
		if lat+radius >= node.Point.Latitude {
			searchNearby(node.Right, lat, long, radius, depth+1, results)
		}
	} else {
		if long-radius < node.Point.Longitude {
			searchNearby(node.Left, lat, long, radius, depth+1, results)
		}
		if long+radius >= node.Point.Longitude {
			searchNearby(node.Right, lat, long, radius, depth+1, results)
		}
	}
}

func haversine(lat1, lon1, lat2, lon2 float64) float64 {
	const R = 6371 // Earth radius in kilometers
	dLat := (lat2 - lat1) * (3.141592653589793 / 180.0)
	dLon := (lon2 - lon1) * (3.141592653589793 / 180.0)
	a := 0.5 - (0.5 * (1 - (dLat*dLat/2 + dLon*dLon/2)))
	return R * 2 * (1 - a)
}

func (tree *KDTree) GetAllRiders() []*pb.RidersNearby {
	tree.mu.RLock()
	defer tree.mu.RUnlock()
	var riders []*pb.RidersNearby
	collectRiders(tree.Root, &riders)
	return riders
}

func collectRiders(node *Node, riders *[]*pb.RidersNearby) {
	if node == nil {
		return
	}
	*riders = append(*riders, node.Point)
	collectRiders(node.Left, riders)
	collectRiders(node.Right, riders)
}

func (tree *KDTree) FindRiderByID(id string) *pb.RidersNearby {
	tree.mu.RLock()
	defer tree.mu.RUnlock()
	return findRiderByID(tree.Root, id)
}

func findRiderByID(node *Node, id string) *pb.RidersNearby {
	if node == nil {
		return nil
	}
	if node.Point.RiderId == id {
		return node.Point
	}
	leftResult := findRiderByID(node.Left, id)
	if leftResult != nil {
		return leftResult
	}
	return findRiderByID(node.Right, id)
}

func (tree *KDTree) Rebuild() {
	tree.mu.Lock()
	defer tree.mu.Unlock()
	var riders []*pb.RidersNearby
	collectRiders(tree.Root, &riders)
	tree.Root = nil
	for _, rider := range riders {
		tree.Root = insert(tree.Root, rider, 0)
	}
}

func (tree *KDTree) SearchNearbyWithLimit(lat, long float64, limit int) []*pb.RidersNearby {
	tree.mu.RLock()
	defer tree.mu.RUnlock()
	var results []*pb.RidersNearby
	searchNearbyWithLimit(tree.Root, lat, long, limit, 0, &results)

	var distanceResult []float64
	if os.Getenv("ENGINE") == "OSRM" {
		distanceResult, _, _ = OSRM(results, lat, long)
	}

	for i := range results {
		switch os.Getenv("ENGINE") {
		case "OSRM":
			results[i].Distance = distanceResult[i+1]
		case "HAVERSINE":
		default:
			results[i].Distance = Haversine(lat, long, results[i].Latitude, results[i].Longitude)
		}
	}

	sort.Slice(results, func(i, j int) bool {
		return results[i].Distance < results[j].Distance
	})
	if len(results) > limit {
		results = results[:limit]
	}
	return results
}

func searchNearbyWithLimit(node *Node, lat, long float64, limit int, depth int, results *[]*pb.RidersNearby) {
	if node == nil {
		return
	}
	*results = append(*results, node.Point)

	if depth%2 == 0 {
		if lat < node.Point.Latitude {
			searchNearbyWithLimit(node.Left, lat, long, limit, depth+1, results)
			searchNearbyWithLimit(node.Right, lat, long, limit, depth+1, results)
		} else {
			searchNearbyWithLimit(node.Right, lat, long, limit, depth+1, results)
			searchNearbyWithLimit(node.Left, lat, long, limit, depth+1, results)
		}
	} else {
		if long < node.Point.Longitude {
			searchNearbyWithLimit(node.Left, lat, long, limit, depth+1, results)
			searchNearbyWithLimit(node.Right, lat, long, limit, depth+1, results)
		} else {
			searchNearbyWithLimit(node.Right, lat, long, limit, depth+1, results)
			searchNearbyWithLimit(node.Left, lat, long, limit, depth+1, results)
		}
	}
}
