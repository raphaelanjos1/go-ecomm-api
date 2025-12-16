package main

import (
	"log"
	"net/http"
	"time"

	"github.com/raphaelanjos1/go-ecomm-api/internal/routes"
)

type application struct {
	config config
}

type config struct {
	addr string
}

func (app *application) mount() http.Handler {
	router := routes.NewRouter()

	return router
}

func (app *application) run(h http.Handler) error {
	srv := &http.Server{
		ReadTimeout:  30 * time.Second,
		WriteTimeout: 30 * time.Second,
		Addr:         app.config.addr,
		Handler:      h,
	}

	log.Printf("server has started at addr %s", app.config.addr)

	return srv.ListenAndServe()
}
