package core

import (
	"context"
	"encoding/json"
	"errors"
	"github.com/go-playground/validator/v10"
	_ "github.com/joho/godotenv/autoload"
	"github.com/mseptiaan/jasmine/internal/pb"
	"os"
	"strconv"
	"sync"
	"time"
)

type Jasmine struct {
	mu    sync.RWMutex
	cache map[string]*KDTree
	pb.UnimplementedJasmineEndpointServer
	validate *validator.Validate
}

func New(test bool) *Jasmine {
	j := &Jasmine{cache: make(map[string]*KDTree), validate: NewValidator()}
	if !test {
		j.restoreCache(os.Getenv("CACHE_FILE"))
		go j.startCacheStoreInterval()
	}
	return j
}

func (j *Jasmine) PostStore(ctx context.Context, req *pb.RequestStore) (*pb.ResponseStore, error) {
	if err := j.validate.Struct(req); err != nil {
		return nil, errors.New("validation error")
	}
	name := req.Bucket
	j.mu.RLock()
	tree, exists := j.cache[name]
	if !exists {
		tree = NewKDTree(name)
		j.cache[name] = tree
	}
	j.mu.RUnlock()

	j.mu.Lock()
	defer j.mu.Unlock()
	existingRider := tree.FindRiderByID(req.RiderId)
	if existingRider != nil {
		existingRider.Latitude = req.Latitude
		existingRider.Longitude = req.Longitude
		tree.Rebuild()
	} else {
		tree.Insert(&pb.RidersNearby{RiderId: req.RiderId, Latitude: req.Latitude, Longitude: req.Longitude})
	}
	j.cache[name].LastUpdate = time.Now()
	return &pb.ResponseStore{Status: "OK"}, nil
}

func (j *Jasmine) GetData(ctx context.Context, req *pb.RequestGet) (*pb.ResponseGet, error) {
	j.mu.RLock()
	tree, exists := j.cache[req.Bucket]
	j.mu.RUnlock()
	if !exists {
		return &pb.ResponseGet{Riders: nil}, nil
	}
	points := tree.GetAllRiders()
	riders := make([]*pb.Riders, len(points))
	for i, p := range points {
		riders[i] = &pb.Riders{RiderId: p.GetRiderId(), Latitude: p.GetLatitude(), Longitude: p.GetLongitude()}
	}
	return &pb.ResponseGet{Riders: riders}, nil
}

func (j *Jasmine) GetDataGeoJson(ctx context.Context, req *pb.RequestGet) (*pb.ResponseGetGeoJson, error) {
	//j.mu.RLock()
	//data, exists := j.cache[req.Bucket]
	//j.mu.RUnlock()
	//if !exists {
	//	return &pb.ResponseGetGeoJson{}, nil
	//}
	//
	//data.LastUpdate = time.Now()
	//points := data.Driver.ListPoints()
	//features := make([]*pb.ResponseGetGeoJson_Features, len(points))
	//for i, p := range points {
	//	features[i] = &pb.ResponseGetGeoJson_Features{
	//		Type: "Feature",
	//		Properties: &pb.ResponseGetGeoJson_Properties{
	//			Markercolor:  "#e81515",
	//			Markersize:   "medium",
	//			Markersymbol: "circle",
	//		},
	//		Geometry: &pb.ResponseGetGeoJson_Geometry{
	//			Coordinates: []float64{p.Longitude, p.Latitude},
	//			Type:        "Point",
	//		},
	//		Id: uint32(i),
	//	}
	//}
	return &pb.ResponseGetGeoJson{Features: nil}, nil
}

func (j *Jasmine) PostNearby(ctx context.Context, req *pb.RequestNearby) (*pb.ResponseNearby, error) {
	name := req.Bucket
	j.mu.RLock()
	tree, exists := j.cache[name]
	j.mu.RUnlock()
	if !exists {
		return &pb.ResponseNearby{Rider: nil}, nil
	}

	results := tree.SearchNearbyWithLimit(req.GetLatitude(), req.GetLongitude(), int(req.GetLimit()))
	riders := make([]*pb.RidersNearby, len(results))
	for i, p := range results {
		riders[i] = &pb.RidersNearby{RiderId: p.GetRiderId(), Latitude: p.GetLatitude(), Longitude: p.GetLongitude(), Distance: p.GetDistance()}
	}

	return &pb.ResponseNearby{Rider: riders}, nil
}

func (j *Jasmine) startCacheStoreInterval() {
	cacheDuration, err := strconv.Atoi(os.Getenv("DURATION_BACKUP_RESTORE_CACHE"))
	if err != nil {
		cacheDuration = 60
	}
	ticker := time.NewTicker(time.Duration(cacheDuration) * time.Second)
	defer ticker.Stop()

	for range ticker.C {
		errBackup := j.backupCache(os.Getenv("CACHE_FILE"))
		if err != nil {
			panic(errBackup)
		}
		j.removeOldCacheEntries()
	}
}

func (j *Jasmine) backupCache(filename string) error {
	j.mu.RLock()
	defer j.mu.RUnlock()

	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	return json.NewEncoder(file).Encode(j.cache)
}

func (j *Jasmine) restoreCache(filename string) error {
	j.mu.RLock()
	defer j.mu.RUnlock()

	file, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	return json.NewDecoder(file).Decode(&j.cache)
}

func (j *Jasmine) removeOldCacheEntries() {
	j.mu.Lock()
	defer j.mu.Unlock()
	purgeDuration, err := strconv.Atoi(os.Getenv("DURATION_PURGE_CACHE"))
	if err != nil {
		purgeDuration = 720
	}
	oneMonthAgo := time.Now().Add(-time.Duration(purgeDuration) * time.Hour)

	for key, data := range j.cache {
		if data.LastUpdate.Before(oneMonthAgo) {
			delete(j.cache, key)
		}
	}
}
