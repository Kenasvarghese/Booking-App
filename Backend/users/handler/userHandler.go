package handler

import (
	"net/http"

	"github.com/go-chi/chi/v5"
)

type userHandler struct {
}

func NewUserHandler(r *chi.Mux) {
	handler := &userHandler{}
	r.Get("/users", handler.ListUsers)
}

func (h userHandler) ListUsers(w http.ResponseWriter, r *http.Request) {
	println("reached")
}
