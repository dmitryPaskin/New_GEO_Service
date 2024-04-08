package searchGEOHandlers

import (
	"GeoServiseAppDate/internal/controller/responder"
	"GeoServiseAppDate/internal/models"
	"GeoServiseAppDate/internal/service"
	"encoding/json"
	"net/http"
)

type Handler struct {
	s service.Service
	r responder.Responder
}

func New(service service.Service, responder responder.Responder) Handler {
	return Handler{
		s: service,
		r: responder,
	}
}

func (h *Handler) SearchAddressHandler(w http.ResponseWriter, r *http.Request) {
	var searchRequest models.SearchRequest
	if err := json.NewDecoder(r.Body).Decode(&searchRequest.Query); err != nil {
		h.r.ErrorBedRequest(w, err)
		return
	}

	address, err := h.s.Address(searchRequest)
	if err != nil {
		h.r.ErrorInternal(w, err)
		return
	}

	h.r.OutputJSON(w, responder.Response{
		Success: true,
		Message: "address get",
		Data:    address,
	})
}

func (h *Handler) GeocodeHandler(w http.ResponseWriter, r *http.Request) {
	var geocodeRequest models.GeocodeRequest
	if err := json.NewDecoder(r.Body).Decode(&geocodeRequest); err != nil {
		h.r.ErrorBedRequest(w, err)
		return
	}

	geocode, err := h.s.Geocode(geocodeRequest)
	if err != nil {
		h.r.ErrorInternal(w, err)
		return
	}
	h.r.OutputJSON(w, responder.Response{
		Success: true,
		Message: "address get",
		Data:    geocode,
	})

}
