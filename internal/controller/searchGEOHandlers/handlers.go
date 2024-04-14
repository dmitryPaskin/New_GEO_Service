package searchGEOHandlers

import (
	"GeoServiseAppDate/internal/controller/responder"
	"GeoServiseAppDate/internal/metrics"
	"GeoServiseAppDate/internal/models"
	"GeoServiseAppDate/internal/service"
	"encoding/json"
	"github.com/prometheus/client_golang/prometheus"
	"net/http"
	"time"
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

// @Summary Search for an address
// @ID search_address
// @Tags GEO_Data
// @Accept json
// @Produce json
// @Param request body models.SearchRequest true "Search request"
// @Success 200 {object} models.AddressSearch "get data"
// @Failure 400
// @Failure 500
// @Security ApiKeyAuth
// @Router /address/search [post]
func (h *Handler) SearchAddressHandler(w http.ResponseWriter, r *http.Request) {
	start := time.Now()
	defer func() {
		duration := time.Since(start).Seconds()
		metrics.RequestDuration.With(prometheus.Labels{
			"method": r.Method, "path": r.URL.Path}).
			Observe(duration)

		metrics.RequestCount.With(prometheus.Labels{
			"method": r.Method, "path": r.URL.Path}).Inc()
	}()
	var searchRequest models.SearchRequest
	if err := json.NewDecoder(r.Body).Decode(&searchRequest); err != nil {
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

// @Summary Search for an GEO
// @ID GEO_address
// @Tags GEO_Data
// @Accept json
// @Produce json
// @Param request body models.GeocodeRequest true "Geocode request"
// @Success 200 {object} models.AddressGeo "get data"
// @Failure 400
// @Failure 500
// @Security ApiKeyAuth
// @Router /address/geocode [post]
func (h *Handler) GeocodeHandler(w http.ResponseWriter, r *http.Request) {
	start := time.Now()
	defer func() {
		duration := time.Since(start).Seconds()
		metrics.RequestDuration.With(prometheus.Labels{
			"method": r.Method, "path": r.URL.Path}).
			Observe(duration)

		metrics.RequestCount.With(prometheus.Labels{
			"method": r.Method, "path": r.URL.Path}).Inc()
	}()
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
