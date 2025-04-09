package gocatgo

import (
	"crypto/sha256"
	"fmt"
	"log"
	"os"

	"github.com/vaaleyard/gocatgo/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Database struct {
	Host     string
	User     string
	Password string
	Name     string
	Port     string
}

type App struct {
	DB             *gorm.DB
	Host           string
	Alphabet       string
	BinaryFilename string
	Database       Database
}

func (app *App) initializeDB() error {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		app.Database.User, app.Database.Password, app.Database.Host,
		app.Database.Port, app.Database.Name)

	var err error
	app.DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return err
	}

	return nil
}

func (app *App) Run() error {
	err := app.initializeDB()
	if err != nil {
		return err
	}
	log.Println("Connected to the database")

	err = app.DB.AutoMigrate(&models.Pastebin{})
	if err != nil {
		return err
	}
	log.Println("Database migration completed")

	app.Alphabet = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

	return nil
}

func (app *App) GetSha256() [32]byte {
	data, err := os.ReadFile(app.BinaryFilename)
	if err != nil {
		panic(err)
	}

	return sha256.Sum256(data)
}
