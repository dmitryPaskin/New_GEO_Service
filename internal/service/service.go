package service

import (
	"GeoServiseAppDate/internal/metrics"
	"GeoServiseAppDate/internal/models"
	"bytes"
	"encoding/json"
	"github.com/prometheus/client_golang/prometheus"
	"log"
	"net/http"
	"time"
)

const (
	urlAddress = "https://cleaner.dadata.ru/api/v1/clean/address"
	urlGeocode = "https://suggestions.dadata.ru/suggestions/api/4_1/rs/geolocate/address"
)

type Service interface {
	Address(request models.SearchRequest) ([]*models.AddressSearch, error)
	Geocode(request models.GeocodeRequest) (*models.AddressGeo, error)
}

type service struct {
	client *http.Client
}

func NewService(client *http.Client) service {
	return service{
		client: client,
	}
}

func (s *service) Address(request models.SearchRequest) ([]*models.AddressSearch, error) {
	start := time.Now()
	defer func() {
		duration := time.Since(start).Seconds()
		metrics.DBDuration.With(prometheus.Labels{
			"method": "POST", "path": urlAddress}).
			Observe(duration)
	}()

	var result []*models.AddressSearch

	requestBody, err := json.Marshal([]string{request.Query})
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", urlAddress, bytes.NewBuffer(requestBody))
	if err != nil {
		return nil, err
	}

	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Authorization", "Token e6b91900da8a4f3c5138bc921a882ee75d42922a")
	req.Header.Add("X-Secret", "943062a0ae098458484fa91f7947fd31c3f549df")

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer func() {
		if err := resp.Body.Close(); err != nil {
			log.Fatal(err)
		}
	}()

	if err = json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, err
	}

	return result, nil
}

func (s *service) Geocode(request models.GeocodeRequest) (*models.AddressGeo, error) {
	start := time.Now()
	defer func() {
		duration := time.Since(start).Seconds()
		metrics.DBDuration.With(prometheus.Labels{
			"method": "POST", "path": urlGeocode}).
			Observe(duration)
	}()

	var result *models.AddressGeo

	requestBody, err := json.Marshal(request)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", urlGeocode, bytes.NewBuffer(requestBody))
	if err != nil {
		return nil, err
	}

	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Authorization", "Token e6b91900da8a4f3c5138bc921a882ee75d42922a")

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer func() {
		if err := resp.Body.Close(); err != nil {
			log.Fatal(err)
		}
	}()

	if err = json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, err
	}

	return result, nil
}
