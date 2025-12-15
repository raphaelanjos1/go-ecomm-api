package routes

import (
	"fmt"
	"net/http"
)

func NewRouter() http.Handler {
	mux := http.NewServeMux()
	mux.HandleFunc("/v1/health-check", healthCheckHandler)
	return mux
}

func healthCheckHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println(w, "health check OK")
}
