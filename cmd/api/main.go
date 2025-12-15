package main

import (
	"net/http"
	"time"
)

func main() {
	srv := &http.Server{
		ReadTimeout:  30 * time.Second,
		WriteTimeout: 30 * time.Second,
		Addr:         ":8080",
		Handler:      http.DefaultServeMux,
	}
	srv.ListenAndServe()
}
