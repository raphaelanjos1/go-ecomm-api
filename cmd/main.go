package main

import (
	"context"
	"log/slog"
	"os"

	"github.com/jackc/pgx/v5"
	"github.com/joho/godotenv"
	"github.com/raphaelanjos1/go-ecomm-api/utils"
)

func main() {
	ctx := context.Background()
	godotenv.Load("../.env")
	port := utils.GetEnv("PORT", ":8080")

	cfg := config{
		addr: port,
		db: dbConfig{
			dsn: utils.GetEnv("GOOSE_DBSTRING", "host=localhost user=postgres password=postgres dbname=ecom sslmode=disable"),
		},
	}

	conn, err := pgx.Connect(ctx, cfg.db.dsn)
	if err != nil {
		panic(err)
	}
	defer conn.Close(ctx)

	api := application{
		config: cfg,
		db:     conn,
	}

	if err := api.run(api.mount()); err != nil {
		slog.Error("server failed to start", "error", err)
		os.Exit(1)
	}
}
