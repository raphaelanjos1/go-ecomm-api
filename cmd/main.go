package main

import (
	"log/slog"
	"os"

	"github.com/joho/godotenv"
	"github.com/raphaelanjos1/go-ecomm-api/utils"
)

func main() {
	godotenv.Load("../.env")
	port := utils.GetEnv("PORT", ":8080")

	cfg := config{
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
