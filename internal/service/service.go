package service

import (
	"GeoServiseAppDate/internal/models"
	"bytes"
	"encoding/json"
	"log"
	"net/http"
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
	req.Header.Add("Authorization", "Token 9a84b6e525fb548e7170b77175e9e15af84a30ac")
	req.Header.Add("X-Secret", "6ecfe8510311d14daf5de31de9a5af4ceeb5b0d5")

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
	req.Header.Add("Authorization", "Token 9a84b6e525fb548e7170b77175e9e15af84a30ac")

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
