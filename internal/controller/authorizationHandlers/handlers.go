package authorizationHandlers

import (
	"GeoServiseAppDate/internal/controller/responder"
	"GeoServiseAppDate/internal/metrics"
	"GeoServiseAppDate/internal/models"
	"GeoServiseAppDate/internal/service/authService"
	"encoding/json"
	"github.com/prometheus/client_golang/prometheus"
	"net/http"
	"time"
)

type HandlerAuth struct {
	s authService.AuthService
	r responder.Responder
}

func New(service authService.AuthService, responder responder.Responder) HandlerAuth {
	return HandlerAuth{
		s: service,
		r: responder,
	}
}

// @Summary Register a user
// @ID SingUp
// @Tags authorization
// @Accept json
// @Produce json
// @Param input body models.User true "User"
// @Success 201 "User registered successfully"
// @Failure 400 "Invalid request format"
// @Failure 500 "Response writer error on write"
// @Router /register [post]
func (h *HandlerAuth) SingUpHandler(w http.ResponseWriter, r *http.Request) {
	start := time.Now()
	defer func() {
		duration := time.Since(start).Seconds()
		metrics.RequestDuration.With(prometheus.Labels{
			"method": r.Method, "path": r.URL.Path}).
			Observe(duration)

		metrics.RequestCount.With(prometheus.Labels{
			"method": r.Method, "path": r.URL.Path}).Inc()
	}()

	var singUpUser models.User

	if err := json.NewDecoder(r.Body).Decode(&singUpUser); err != nil {
		h.r.ErrorBedRequest(w, err)
		return
	}

	if err := h.s.SaveUser(singUpUser); err != nil {
		h.r.ErrorInternal(w, err)
		return
	}

	h.r.StatusCreated(w)
}

// @Summary SingIn a user
// @ID SingIn
// @Tags authorization
// @Accept json
// @Produce json
// @Param input body models.User true "User"
// @Success 200 "JWT token"
// @Failure 400 "Invalid request format"
// @Failure 500 "Response writer error on write"
// @Router /login [post]
func (h *HandlerAuth) SingInHandler(w http.ResponseWriter, r *http.Request) {
	start := time.Now()
	defer func() {
		duration := time.Since(start).Seconds()
		metrics.RequestDuration.With(prometheus.Labels{
			"method": r.Method, "path": r.URL.Path}).
			Observe(duration)

		metrics.RequestCount.With(prometheus.Labels{
			"method": r.Method, "path": r.URL.Path}).Inc()
	}()

	var singInUser models.User

	if err := json.NewDecoder(r.Body).Decode(&singInUser); err != nil {
		h.r.ErrorBedRequest(w, err)
		return
	}

	Token, err := h.s.GetToken(singInUser)
	if err != nil {
		h.r.ErrorInternal(w, err)
		return
	}

	h.r.OutputJSON(w, responder.Response{
		Success: true,
		Message: "Bearer: " + Token,
		Data:    nil,
	})
}
