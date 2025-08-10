package main

import (
	"fmt"
	"net/http"

	"github.com/Kenasvarghese/Booking-App/Backend/auth"
	"github.com/Kenasvarghese/Booking-App/Backend/config"
	"github.com/Kenasvarghese/Booking-App/Backend/database"

	propertiesHandler "github.com/Kenasvarghese/Booking-App/Backend/properties/handler"
	propertiesRepo "github.com/Kenasvarghese/Booking-App/Backend/properties/repo"
	propertiesUsecase "github.com/Kenasvarghese/Booking-App/Backend/properties/usecase"

	roomsHandler "github.com/Kenasvarghese/Booking-App/Backend/rooms/handler"
	roomsRepo "github.com/Kenasvarghese/Booking-App/Backend/rooms/repo"
	roomsUsecase "github.com/Kenasvarghese/Booking-App/Backend/rooms/usecase"

	bookingsHandler "github.com/Kenasvarghese/Booking-App/Backend/bookings/handler"
	bookingsRepo "github.com/Kenasvarghese/Booking-App/Backend/bookings/repo"
	bookingsUsecase "github.com/Kenasvarghese/Booking-App/Backend/bookings/usecase"

	ssoHandler "github.com/Kenasvarghese/Booking-App/Backend/sso/handler"

	userHandler "github.com/Kenasvarghese/Booking-App/Backend/users/handler"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

func main() {
	cfg := config.LoadConfig()
	db := database.LoadDB(cfg)
	r := mux.NewRouter()
	apiRouter := r.PathPrefix("/" + cfg.BasePath).Subrouter()
	authConfig := auth.NewSSOConfig(
		auth.WithClientID(cfg.ClientID),
		auth.WithClientSecret(cfg.ClientSecret),
		auth.WithRedirectURL(cfg.RedirectURL),
	)
	authProvider := auth.NewAuthProvider(authConfig)
	privateRoutes(apiRouter, db, authProvider)
	corsHandler := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Content-Type", "Authorization"},
		AllowCredentials: true,
	})

	server := http.Server{
		Addr:    fmt.Sprintf(":%d", cfg.ServerPort),
		Handler: corsHandler.Handler(r),
	}
	println("server started")
	server.ListenAndServe()

}
func privateRoutes(r *mux.Router, db database.DB, authProvider auth.AuthProvider) {
	pRepo := propertiesRepo.NewPropertiesRepo(db)
	pUsecase := propertiesUsecase.NewPropertiesUsecaseHandler(pRepo)
	propertiesHandler.NewPropertiesHandler(r, pUsecase)

	rRepo := roomsRepo.NewRoomsRepo(db)
	rUsecase := roomsUsecase.NewRoomUsecase(rRepo)
	roomsHandler.NewRoomsHandler(r, rUsecase)

	bookingsRepo := bookingsRepo.NewBookingRepo(db)
	bookingsUsecase := bookingsUsecase.NewBookingsUsecase(bookingsRepo)
	bookingsHandler.NewBookingHandler(r, bookingsUsecase)

	ssoHandler.NewSSOHandler(r, authProvider)
	userHandler.NewUserHandler(r)

}
