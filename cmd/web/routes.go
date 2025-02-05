package main

import (
	"net/http"

	"github.com/alexesp/Go_Reserva_casa/pkg/config"
	"github.com/alexesp/Go_Reserva_casa/pkg/handlers"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func routes(app *config.AppConfig) http.Handler {
	// mux := pat.New()

	// mux.Get("/", http.HandlerFunc(handlers.Repo.Home))
	// mux.Get("/about", http.HandlerFunc(handlers.Repo.About))

	mux := chi.NewRouter()

	mux.Use(middleware.Recoverer)
	//mux.Use(WriteToConsole)
	//mux.Use(NoSurf)

	mux.Use(SessionLoad)

	mux.Get("/", handlers.Repo.Home)
	mux.Get("/casas", handlers.Repo.Casas)
	mux.Get("/habitaciones", handlers.Repo.Habitaciones)
	mux.Get("/precios", handlers.Repo.Precios)
	mux.Get("/contacto", handlers.Repo.Contacto)
	mux.Get("/make-reservation", handlers.Repo.Reservation)	
	mux.Post("/make-reservation", handlers.Repo.PostReservation)
	mux.Get("/search-abailability", handlers.Repo.Search)
	mux.Get("/generals", handlers.Repo.Generals)
	mux.Post("/generals", handlers.Repo.PostGenerals)

	fileServer := http.FileServer(http.Dir("./static/"))
	mux.Handle("/static/*", http.StripPrefix("/static", fileServer))

	return mux
}
