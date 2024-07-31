package core

import (
	"encoding/json"
	"fmt"
	_ "github.com/joho/godotenv/autoload"
	"github.com/useinsider/go-pkg/insrequester"
	"io"
	"os"
	"strings"
	"time"
)

type OSRMTable struct {
	Code         string      `json:"code"`
	Distances    [][]float64 `json:"distances"`
	Destinations []struct {
		Hint     string    `json:"hint"`
		Distance float64   `json:"distance"`
		Name     string    `json:"name"`
		Location []float64 `json:"location"`
	} `json:"destinations"`
	Durations [][]float64 `json:"durations"`
	Sources   []struct {
		Hint     string    `json:"hint"`
		Distance float64   `json:"distance"`
		Name     string    `json:"name"`
		Location []float64 `json:"location"`
	} `json:"sources"`
}

func OSRM(drivers []*Point, Latitude, Longitude float64) (distances []float64, durations []float64, err error) {
	requester := insrequester.NewRequester()
	retryConfig := insrequester.RetryConfig{
		WaitBase: time.Second,
		Times:    3,
	}
	requester.WithRetry(retryConfig).Load()

	var latlon []string
	latlon = append(latlon, fmt.Sprintf("%f,%f", Longitude, Latitude))
	for _, p := range drivers {
		latlon = append(latlon, fmt.Sprintf("%f,%f", p.Longitude, p.Latitude))
	}
	res, _ := requester.Get(insrequester.RequestEntity{Endpoint: fmt.Sprintf("%s/table/v1/driving/%s?sources=0&annotations=distance,duration", os.Getenv("OSRM_URL"), strings.Join(latlon, ";"))})
	var result OSRMTable
	bodyBytes, err := io.ReadAll(res.Body)
	if err != nil {
		return distances, durations, err
	}
	err = json.Unmarshal(bodyBytes, &result)
	if err != nil {
		return distances, durations, err
	}
	distances = result.Distances[0]
	durations = result.Durations[0]
	return distances, durations, nil
}
