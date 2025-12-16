package main

import (
	"log/slog"
	"os"

	"github.com/joho/godotenv"
	"github.com/raphaelanjos1/go-ecomm-api/config"
)

func main() {
	godotenv.Load("../.env")
	port := config.GetString("PORT", ":8080")

	cfg := configuration{
		addr: port,
	}

	api := application{
		config: cfg,
	}

	if err := api.run(api.mount()); err != nil {
		slog.Error("server failed to start", "error", err)
		os.Exit(1)
	}
}
