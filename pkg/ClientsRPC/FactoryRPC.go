package ClientsRPC

import (
	"GeoServiseAppDate/internal/models"
	"net/rpc"
)

type RPC struct {
}

func (r *RPC) Address(request models.SearchRequest) ([]*models.AddressSearch, error) {
	client, err := rpc.Dial("tcp", "rpc:1234")
	if err != nil {
		return nil, err
	}

	var address []*models.AddressSearch
	if err = client.Call("Geocoder.SearchService", request, &address); err != nil {
		return nil, err
	}
	return address, nil
}

func (r *RPC) Geocode(request models.GeocodeRequest) (*models.AddressGeo, error) {
	client, err := rpc.Dial("tcp", "rpc:1234")
	if err != nil {
		return nil, err
	}

	var address *models.AddressGeo
	if err = client.Call("Geocoder.GeocodeAddressService", request, address); err != nil {
		return nil, err
	}
	return address, nil
}
