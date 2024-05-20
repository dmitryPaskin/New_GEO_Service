package ClientsRPC

import (
	"GeoServiseAppDate/internal/models"
	"log"
	"net/rpc/jsonrpc"
)

type JSONRPC struct {
}

func (r *JSONRPC) Address(request models.SearchRequest) ([]*models.AddressSearch, error) {
	client, err := jsonrpc.Dial("tcp", "json-rpc:4321")
	if err != nil {
		return nil, err
	}
	log.Println("Dial Sucessful")
	var address []*models.AddressSearch
	if err = client.Call("Geocoder.SearchService", request, &address); err != nil {
		return nil, err
	}
	return address, nil
}

func (r *JSONRPC) Geocode(request models.GeocodeRequest) (*models.AddressGeo, error) {
	client, err := jsonrpc.Dial("tcp", "json-rpc:4321")
	if err != nil {
		return nil, err
	}
	log.Println("Dial Sucessful")
	var address *models.AddressGeo
	if err = client.Call("Geocoder.GeocodeAddressService", request, address); err != nil {
		return nil, err
	}
	return address, nil
}
