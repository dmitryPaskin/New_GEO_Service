package service

import (
	"GeoServiseAppDate/internal/models"
	"GeoServiseAppDate/internal/repository"
)

type serviceProxy struct {
	realService service
	repo        repository.Repository
}

func NewServiceProxy(realObject service, repo repository.Repository) Service {
	return &serviceProxy{
		realService: realObject,
		repo:        repo,
	}
}

func (sp *serviceProxy) Address(request models.SearchRequest) ([]*models.AddressSearch, error) {
	isCache, err := sp.repo.CheckCacheAddress(&request)
	if err != nil {
		return nil, err
	}

	if isCache {
		result, err := sp.repo.GetDataAddress(&request)
		if err != nil {
			return nil, err
		}
		return result, nil
	} else {
		result, err := sp.realService.Address(request)
		if err != nil {
			return nil, err
		}

		if err = sp.repo.AddDataAddressToDB(&request, result); err != nil {
			return nil, err
		}
		return result, nil
	}
}

func (sp *serviceProxy) Geocode(request models.GeocodeRequest) (*models.AddressGeo, error) {
	isCache, err := sp.repo.CheckCacheGEO(&request)
	if err != nil {
		return nil, err
	}

	if isCache {
		result, err := sp.repo.GetDataGEO(&request)
		if err != nil {
			return nil, err
		}
		return result, nil
	} else {
		result, err := sp.realService.Geocode(request)
		if err != nil {
			return nil, err
		}

		if err = sp.repo.AddDataGEOToDB(&request, result); err != nil {
			return nil, err
		}
		return result, nil
	}
}
