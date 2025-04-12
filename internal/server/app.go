package server

import (
	"crypto/sha256"
	"github.com/jackc/pgx/v5/pgxpool"
	"os"
)

type App struct {
	DB             *pgxpool.Pool
	Alphabet       string
	BinaryFilename string
}

func NewApp() *App {
	return &App{
		BinaryFilename: "gcg",
		Alphabet:       "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789",
	}
}

func (app *App) getBinarySha256() [32]byte {
	data, err := os.ReadFile(app.BinaryFilename)
	if err != nil {
		panic(err)
	}

	return sha256.Sum256(data)
}
