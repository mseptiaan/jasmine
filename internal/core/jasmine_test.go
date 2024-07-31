package core

import (
	"context"
	"os"
	"testing"
	"time"

	"github.com/mseptiaan/jasmine/internal/pb"
	"github.com/stretchr/testify/assert"
)

func TestPostStore_NewBucket(t *testing.T) {
	j := New(true)
	req := &pb.RequestStore{
		Bucket:    "new_bucket",
		RiderId:   "rider_1",
		Latitude:  12.34,
		Longitude: 56.78,
	}
	resp, err := j.PostStore(context.Background(), req)
	assert.NoError(t, err)
	assert.Equal(t, "OK", resp.Status)
	assert.Contains(t, j.cache, "new_bucket")
}

func TestPostStore_UpdateExistingBucket(t *testing.T) {
	j := New(true)
	req := &pb.RequestStore{
		Bucket:    "existing_bucket",
		RiderId:   "rider_1",
		Latitude:  12.34,
		Longitude: 56.78,
	}
	j.PostStore(context.Background(), req)
	req.Latitude = 23.45
	req.Longitude = 67.89
	resp, err := j.PostStore(context.Background(), req)
	assert.NoError(t, err)
	assert.Equal(t, "OK", resp.Status)
	assert.Equal(t, 23.45, j.cache["existing_bucket"].Driver.ListPoints()[0].Latitude)
}

func TestGetData_ExistingBucket(t *testing.T) {
	j := New(true)
	req := &pb.RequestStore{
		Bucket:    "existing_bucket",
		RiderId:   "rider_1",
		Latitude:  12.34,
		Longitude: 56.78,
	}
	j.PostStore(context.Background(), req)
	getReq := &pb.RequestGet{Bucket: "existing_bucket"}
	resp, err := j.GetData(context.Background(), getReq)
	assert.NoError(t, err)
	assert.Len(t, resp.Riders, 1)
	assert.Equal(t, "rider_1", resp.Riders[0].RiderId)
}

func TestGetData_NonExistingBucket(t *testing.T) {
	j := New(true)
	getReq := &pb.RequestGet{Bucket: "non_existing_bucket"}
	resp, err := j.GetData(context.Background(), getReq)
	assert.NoError(t, err)
	assert.Len(t, resp.Riders, 0)
}

func TestGetDataGeoJson_ExistingBucket(t *testing.T) {
	j := New(true)
	req := &pb.RequestStore{
		Bucket:    "existing_bucket",
		RiderId:   "rider_1",
		Latitude:  12.34,
		Longitude: 56.78,
	}
	j.PostStore(context.Background(), req)
	getReq := &pb.RequestGet{Bucket: "existing_bucket"}
	resp, err := j.GetDataGeoJson(context.Background(), getReq)
	assert.NoError(t, err)
	assert.Len(t, resp.Features, 1)
	assert.Equal(t, "Feature", resp.Features[0].Type)
}

func TestGetDataGeoJson_NonExistingBucket(t *testing.T) {
	j := New(true)
	getReq := &pb.RequestGet{Bucket: "non_existing_bucket"}
	resp, err := j.GetDataGeoJson(context.Background(), getReq)
	assert.NoError(t, err)
	assert.Len(t, resp.Features, 0)
}

func TestPostNearby_ExistingBucket(t *testing.T) {
	j := New(true)
	req := &pb.RequestStore{
		Bucket:    "existing_bucket",
		RiderId:   "rider_1",
		Latitude:  12.34,
		Longitude: 56.78,
	}
	j.PostStore(context.Background(), req)
	nearbyReq := &pb.RequestNearby{
		Bucket:    "existing_bucket",
		Latitude:  12.34,
		Longitude: 56.78,
		Limit:     1,
	}
	resp, err := j.PostNearby(context.Background(), nearbyReq)
	assert.NoError(t, err)
	assert.Len(t, resp.Rider, 1)
	assert.Equal(t, "rider_1", resp.Rider[0].RiderId)
}

func TestPostNearby_NonExistingBucket(t *testing.T) {
	j := New(true)
	nearbyReq := &pb.RequestNearby{
		Bucket:    "non_existing_bucket",
		Latitude:  12.34,
		Longitude: 56.78,
		Limit:     1,
	}
	resp, err := j.PostNearby(context.Background(), nearbyReq)
	assert.NoError(t, err)
	assert.Len(t, resp.Rider, 0)
}

func TestStoreCacheToFile(t *testing.T) {
	j := New(true)
	req := &pb.RequestStore{
		Bucket:    "test_bucket",
		RiderId:   "rider_1",
		Latitude:  12.34,
		Longitude: 56.78,
	}
	j.PostStore(context.Background(), req)
	filename := "test_cache.json"
	j.storeCacheToFile(filename)
	defer os.Remove(filename)
	_, err := os.Stat(filename)
	assert.NoError(t, err)
}

func TestLoadCacheFromFile(t *testing.T) {
	j := New(true)
	filename := "test_cache.json"
	file, _ := os.Create(filename)
	defer os.Remove(filename)
	defer file.Close()
	file.WriteString(`{"test_bucket":{"Driver":{},"LastUpdate":"2023-01-01T00:00:00Z"}}`)
	j.loadCacheFromFile(filename)
	assert.Contains(t, j.cache, "test_bucket")
}

func TestRemoveOldCacheEntries(t *testing.T) {
	j := New(true)
	j.cache["old_bucket"] = &CargoData{LastUpdate: time.Now().AddDate(0, -2, 0)}
	j.cache["new_bucket"] = &CargoData{LastUpdate: time.Now()}
	j.removeOldCacheEntries()
	assert.NotContains(t, j.cache, "old_bucket")
	assert.Contains(t, j.cache, "new_bucket")
}
