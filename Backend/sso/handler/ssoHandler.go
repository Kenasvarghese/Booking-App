package sso

import (
	"net/http"

	"github.com/Kenasvarghese/Booking-App/Backend/auth"
	"github.com/Kenasvarghese/Booking-App/Backend/utils"
	"github.com/gorilla/mux"
)

type ssoHandler struct {
	authProvider auth.AuthProvider
}

const state string = "random-state"

func NewSSOHandler(r *mux.Router, authProvider auth.AuthProvider) {
	handler := ssoHandler{authProvider: authProvider}
	r.HandleFunc("/login", handler.SSOLogin).Methods("GET")
	r.HandleFunc("/auth/callback", handler.SSOCallback).Methods("GET")

}

// SSOLogin user login handler
func (h *ssoHandler) SSOLogin(w http.ResponseWriter, r *http.Request) {
	url := h.authProvider.GetURL(state)
	http.Redirect(w, r, url, http.StatusTemporaryRedirect)
}

// SSOCallback handler for auth callback
func (h *ssoHandler) SSOCallback(w http.ResponseWriter, r *http.Request) {
	if r.URL.Query().Get("state") != state {
		utils.ApiErrorResponse(w, http.StatusBadRequest, "state mismatch")
		return
	}
	code := r.URL.Query().Get("code")
	_, err := h.authProvider.GetUserEmail(code)
	if err != nil {
		utils.ApiErrorResponse(w, http.StatusInternalServerError, err.Error())
	}
	// To Do
	// add implementation for access token and refresh token
	// verify user db for the email

}
