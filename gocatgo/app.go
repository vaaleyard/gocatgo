package gocatgo

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"crypto/sha256"
	"fmt"
	"io"
	"io/ioutil"
	"log"

	"github.com/vaaleyard/gocatgo/models"
	"gorm.io/driver/postgres"
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
	AESCipherkey   []byte
	BinaryFilename string
	Database       Database
}

func (app *App) initializeDB() error {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s",
		app.Database.Host, app.Database.User, app.Database.Password,
		app.Database.Name, app.Database.Port)
	var err error
	app.DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
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

func AESEncrypt(key []byte, message string) (encoded string, err error) {
	text := []byte(message)

	aesCipher, err := aes.NewCipher(key)
	if err != nil {
		return
	}

	gcm, err := cipher.NewGCM(aesCipher)
	if err != nil {
		return
	}

	nonce := make([]byte, gcm.NonceSize())

	if _, err = io.ReadFull(rand.Reader, nonce); err != nil {
		return
	}

	return fmt.Sprintf("%x", gcm.Seal(nonce, nonce, text, nil)), nil
}

func (app *App) GetSha256() [32]byte {
	data, err := ioutil.ReadFile(app.BinaryFilename)
	if err != nil {
		panic(err)
	}

	return sha256.Sum256(data)
}
