package main

import (
	"net/http"
	"time"

	"github.com/raphaelanjos1/go-ecomm-api/internal/routes"
)

func main() {
	router := routes.NewRouter()
	srv := &http.Server{
		ReadTimeout:  30 * time.Second,
		WriteTimeout: 30 * time.Second,
		Addr:         ":8080",
		Handler:      router,
	}
	err := srv.ListenAndServe()
	if err != nil {
		panic(err)
	}
}
