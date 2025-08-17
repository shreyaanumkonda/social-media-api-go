package main

import (
	"log"
	"net/http"
	"os"
	"social/internal/storage"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

type dbConfig struct {
	addr         string
	maxOpenConns int
	maxIdleConns int
	maxIdleTime  time.Duration
}
type config struct {
	addr string
	db   dbConfig
}

type application struct {
	config  config
	dbLayer storage.Storage
}

func (app *application) mount() http.Handler {
	r := chi.NewRouter()
	r.Use(middleware.Recoverer)
	r.Use(middleware.Logger)
	r.Route("/v1", func(r chi.Router) {
		r.Get("/healthcheck", app.healthcheckHandler)
	})
	return r
}
func (app *application) run(handler http.Handler) error {

	srv := &http.Server{
		Addr:         app.config.addr,
		Handler:      handler,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 30 * time.Second,
		IdleTimeout:  time.Minute,
		ErrorLog:     log.New(os.Stderr, "", log.LstdFlags),
	}
	log.Println("Starting server on", app.config.addr)
	return srv.ListenAndServe()

}
