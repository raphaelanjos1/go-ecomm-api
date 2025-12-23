package main

import (
	"log"
	"net/http"
	"time"

	"github.com/jackc/pgx/v5"
	"github.com/raphaelanjos1/go-ecomm-api/internal/routes"
)

type application struct {
	config config
	db     *pgx.Conn
}

type config struct {
	addr string
	db   dbConfig
}

type dbConfig struct {
	dsn string
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
