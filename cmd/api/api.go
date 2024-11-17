package main

import (
	"log"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

type application struct {
	config config
}
type config struct {
	addr string
}

func (app *application) mount() http.Handler {
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello World!"))
	})
	r.Get("/health", app.healthCheckHandler)
	return r
}

func (app *application) run(mux http.Handler) error {
	srv := http.Server{
		Addr:         app.config.addr,
		Handler:      mux,
		WriteTimeout: 60 * time.Second,
		ReadTimeout:  30 * time.Second,
		IdleTimeout:  time.Minute,
	}
	log.Printf("Server has started at %v", app.config.addr)
	srv.ListenAndServe()
	return nil
}
