package authorizationHandlers

import (
	"GeoServiseAppDate/internal/controller/responder"
	"GeoServiseAppDate/internal/models"
	"GeoServiseAppDate/internal/service/authService"
	"encoding/json"
	"net/http"
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

func (h *HandlerAuth) SingUpHandler(w http.ResponseWriter, r *http.Request) {
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

func (h *HandlerAuth) SingInHandler(w http.ResponseWriter, r *http.Request) {
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
