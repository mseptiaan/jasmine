package core

import (
	"context"
	"github.com/mseptiaan/jasmine/internal/pb"
	"testing"
)

func BenchmarkPostStore(b *testing.B) {
	j := New(true)
	req := &pb.RequestStore{
		Bucket:    "benchmark_bucket",
		RiderId:   "rider_1",
		Latitude:  12.34,
		Longitude: 56.78,
	}
	for i := 0; i < b.N; i++ {
		j.PostStore(context.Background(), req)
	}
}

func BenchmarkGetData(b *testing.B) {
	j := New(true)
	req := &pb.RequestStore{
		Bucket:    "benchmark_bucket",
		RiderId:   "rider_1",
		Latitude:  12.34,
		Longitude: 56.78,
	}
	j.PostStore(context.Background(), req)
	getReq := &pb.RequestGet{Bucket: "benchmark_bucket"}
	for i := 0; i < b.N; i++ {
		j.GetData(context.Background(), getReq)
	}
}

func BenchmarkGetDataGeoJson(b *testing.B) {
	j := New(true)
	req := &pb.RequestStore{
		Bucket:    "benchmark_bucket",
		RiderId:   "rider_1",
		Latitude:  12.34,
		Longitude: 56.78,
	}
	j.PostStore(context.Background(), req)
	getReq := &pb.RequestGet{Bucket: "benchmark_bucket"}
	for i := 0; i < b.N; i++ {
		j.GetDataGeoJson(context.Background(), getReq)
	}
}

func BenchmarkPostNearby(b *testing.B) {
	j := New(true)
	req := &pb.RequestStore{
		Bucket:    "benchmark_bucket",
		RiderId:   "rider_1",
		Latitude:  12.34,
		Longitude: 56.78,
	}
	j.PostStore(context.Background(), req)
	nearbyReq := &pb.RequestNearby{
		Bucket:    "benchmark_bucket",
		Latitude:  12.34,
		Longitude: 56.78,
		Limit:     1,
	}
	for i := 0; i < b.N; i++ {
		j.PostNearby(context.Background(), nearbyReq)
	}
}
