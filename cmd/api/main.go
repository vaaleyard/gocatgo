package main

import (
	"context"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/vaaleyard/gocatgo/internal/server"
	"log/slog"
	"os"
)

func main() {
	var err error

	app := server.NewApp()

	app.DB, err = pgxpool.New(context.Background(), os.Getenv("DATABASE_URL"))
	if err != nil {
		slog.Error("Unable to create connection pool: " + err.Error())
		os.Exit(1)
	}
	slog.Info("Connected to the database")
	defer app.DB.Close()

	err = server.Run(app)
	if err != nil {
		slog.Error("Error " + err.Error())
		os.Exit(1)
	}
}
