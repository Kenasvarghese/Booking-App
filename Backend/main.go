package main

import (
	"fmt"
	"net/http"

	"github.com/Kenasvarghese/Booking-App/Backend/config"
	"github.com/Kenasvarghese/Booking-App/Backend/database"
	propertiesHandler "github.com/Kenasvarghese/Booking-App/Backend/properties/handler"
	propertiesRepo "github.com/Kenasvarghese/Booking-App/Backend/properties/repo"
	propertiesUsecase "github.com/Kenasvarghese/Booking-App/Backend/properties/usecase"

	roomsHandler "github.com/Kenasvarghese/Booking-App/Backend/rooms/handler"
	roomsRepo "github.com/Kenasvarghese/Booking-App/Backend/rooms/repo"
	roomsUsecase "github.com/Kenasvarghese/Booking-App/Backend/rooms/usecase"

	userHandler "github.com/Kenasvarghese/Booking-App/Backend/users/handler"

	"github.com/gorilla/mux"
)

func main() {
	cfg := config.LoadConfig()
	db := database.LoadDB(cfg)
	r := mux.NewRouter()
	apiRouter := r.PathPrefix("/" + cfg.BasePath).Subrouter()
	privateRoutes(apiRouter, db)

	server := http.Server{
		Addr:    fmt.Sprintf(":%d", cfg.ServerPort),
		Handler: r,
	}
	println("server started")
	server.ListenAndServe()

}
func privateRoutes(r *mux.Router, db database.DB) {
	pRepo := propertiesRepo.NewPropertiesRepo(db)
	pUsecase := propertiesUsecase.NewPropertiesUsecaseHandler(pRepo)
	propertiesHandler.NewPropertiesHandler(r, pUsecase)

	rRepo := roomsRepo.NewRoomsRepo(db)
	rUsecase := roomsUsecase.NewRoomUsecase(rRepo)
	roomsHandler.NewRoomsHandler(r, rUsecase)
	userHandler.NewUserHandler(r)
}
