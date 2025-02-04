package main

import (
	"fmt"
	"net/http"

	"github.com/Kenasvarghese/Booking-App/Backend/config"
	"github.com/Kenasvarghese/Booking-App/Backend/database"
	userHandler "github.com/Kenasvarghese/Booking-App/Backend/users/handler"
	"github.com/go-chi/chi/v5"
)

func main() {
	cfg := config.LoadConfig()
	db := database.LoadDB(cfg)
	r := chi.NewRouter()
	apiRouter := chi.NewRouter()
	apiRouter.Group(func(r chi.Router) {
		privateRoutes(r.(*chi.Mux), db)
	})
	println("/" + cfg.BasePath)
	r.Mount("/"+cfg.BasePath, apiRouter)
	server := &http.Server{
		Addr:    fmt.Sprintf(":%d", cfg.ServerPort),
		Handler: r,
	}
	println("server started")
	server.ListenAndServe()

}
func privateRoutes(r *chi.Mux, db database.DB) {
	userHandler.NewUserHandler(r)
}
