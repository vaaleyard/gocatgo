package gocatgo

import (
	"log"

	"github.com/vaaleyard/gocatgo/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type App struct {
	DB *gorm.DB
}

func (app *App) Initialize() error {
	dsn := "host=host user=postgres password=pwd dbname=dbname port=5432 sslmode=disable TimeZone=America/Sao_Paulo"

	var err error
	app.DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return err
	}

	return nil
}

func (app *App) Run() error {
	err := app.Initialize()
	if err != nil {
		return err
	}

	log.Println("Connected to the database")

	err = app.DB.AutoMigrate(&models.Pastebin{})
	if err != nil {
		return err
	}
	log.Println("Database migration completed")

	return nil
}
